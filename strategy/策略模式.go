package main

import "fmt"

type pay interface {
	Pay()
}
type aliPay struct {
}

func (a *aliPay) Pay() {
	fmt.Println("alipay")
}

type wxPay struct {
}

func (wx *wxPay) Pay() {
	fmt.Println("wx pay")
}

type strategy struct {
	pay pay
}

func (s *strategy) setPay(pay pay) {
	s.pay = pay
}
func (s *strategy) Pay() {
	s.pay.Pay()
}
func main() {
	ali := &aliPay{}
	wx := &wxPay{}
	s := &strategy{}
	s.setPay(ali)
	s.Pay()
	s.setPay(wx)
	s.Pay()
}