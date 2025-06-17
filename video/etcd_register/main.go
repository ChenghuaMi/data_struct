package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

// etcd 服务注册与发现
type etcdClient struct {
	client *clientv3.Client
	ctx    context.Context
	ttl    int64
	lease  *clientv3.LeaseGrantResponse
}

// 生成契约

func (c *etcdClient) Grant() (*clientv3.LeaseGrantResponse, error) {
	lease, err := c.client.Grant(c.ctx, c.ttl)
	c.lease = lease
	return lease, err
}
func (c *etcdClient) Put(key string, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return c.client.Put(c.ctx, key, val, opts...)
}
func (c *etcdClient) KeepAlive() {
	res, err := c.client.KeepAlive(c.ctx, c.lease.ID)
	if err != nil {
		panic(fmt.Sprintf("keepalive error:%s", err))
	}
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case rr, ok := <-res:
				if !ok {
					return
				}
				fmt.Printf("id = %d,ttl=%d \n", rr.ID, rr.TTL)

			}
		}
	}()
}
func (c *etcdClient) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	res, err := c.client.Get(c.ctx, key, opts...)
	return res, err
}
func (c *etcdClient) Watch(key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	watchChan := c.client.Watch(c.ctx, key, opts...)
	return watchChan

}

func NewEtcdClient(ctx context.Context, ttl int64) *etcdClient {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2381", "localhost:2382"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	return &etcdClient{
		client: c,
		ctx:    ctx,
		ttl:    ttl,
	}
}
func main() {
	var mp sync.Map
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	e := NewEtcdClient(ctx, 5)
	_, err := e.Grant()
	if err != nil {
		panic(fmt.Sprintf("lease error:%s", err))
	}
	e.Put("test1", "test11")
	e.Put("test2", "test22")
	e.Put("test3", "test33")
	e.KeepAlive()
	res, err := e.Get("test", clientv3.WithPrefix())
	if err != nil {
		panic(fmt.Sprintf("get error:%s", err))
	}
	for _, v := range res.Kvs {
		mp.Store(string(v.Key), string(v.Value))
	}
	fmt.Printf("get map:%#v", mp)
	go func() {
		watch_chan := e.Watch("test", clientv3.WithPrefix())
		for rs := range watch_chan {
			for _, v := range rs.Events {
				fmt.Println("watch:", v.Type, string(v.Kv.Key), string(v.Kv.Value))
				switch v.Type {
				case clientv3.EventTypePut:
					mp.Store(string(v.Kv.Key), string(v.Kv.Value))

				case clientv3.EventTypeDelete:
					mp.Delete(string(v.Kv.Key))
				}
			}
		}
	}()

	time.Sleep(time.Second * 70)
	mp.Range(func(key, value any) bool {
		fmt.Println("key:", key)
		fmt.Println("val:", value)
		return true
	})
	fmt.Println("main....")
}
