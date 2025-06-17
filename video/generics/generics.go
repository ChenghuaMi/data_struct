package main

import (
	"fmt"
	"sync"
)

// 通过 泛型实现 配置文件的读取

type ConfigManager[T any] interface {
	Get() T
	Update(data T) error
	Watch() ConfigManager[T]
	OnChange(func(T)) (cancel func())
	InitData(data T) ConfigManager[T]
}
type LocalConfigManager[T any] struct {
	data      T
	mu        sync.RWMutex
	isWatched bool
	fns       []func(T)
}

func (l *LocalConfigManager[T]) Get() T {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.data
}
func (l *LocalConfigManager[T]) Update(data T) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = data
	if l.isWatched {
		for _, fn := range l.fns {
			fn(data)
		}
	}
	return nil
}

// 监听

func (l *LocalConfigManager[T]) Watch() ConfigManager[T] {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.isWatched = true
	return l
}
func (l *LocalConfigManager[T]) OnChange(fn func(T)) (cancel func()) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fns = append(l.fns, fn)
	if l.isWatched {
		return func() {
			for i, f := range l.fns {
				if &f == &fn {
					l.fns = append(l.fns[:i], l.fns[i+1:]...)
					break
				}
			}
		}
	}
	return nil
}

func (l *LocalConfigManager[T]) InitData(data T) ConfigManager[T] {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = data
	return l
}
func Local[T any]() ConfigManager[T] {
	return &LocalConfigManager[T]{
		fns: make([]func(T), 0),
	}
}

type config struct {
	Name string
}

func (c *config) Get() string {
	return c.Name
}

func main() {

	c := Local[config]()
	c = c.Watch().InitData(config{Name: "init data"})
	fmt.Println(c.Get())

	c.Update(config{Name: "updata data..."})
	c.OnChange(func(cf config) {
		cf.Name = "cccc"
	})
	fmt.Println(c.Get())
}
