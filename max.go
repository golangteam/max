package max

import (
	"fmt"
	"sync"
)

//Max
type Max struct {
	max    int
	locker *sync.Cond
	curr   int
	wait   sync.WaitGroup
	// 是否允许所有等待
	EnableWait bool
}

//新建一个Max
func NewMax() *Max {
	return &Max{
		max:    100,
		curr:   0,
		locker: sync.NewCond(&sync.Mutex{}),
	}
}

//重新设置一下max值
//
//  max int类型，最大值
func (this *Max) SetMax(max int) {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	this.max = max
}

//返回一个请求
func (this *Max) Set() {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	this.curr -= 1
	this.locker.Broadcast()
	if this.EnableWait {
		this.wait.Done()
	}
}
func (this *Max) Info() string {
	return fmt.Sprintf("max:%d,curr:%d", this.max, this.curr)
}

//获取一个请求，如果超过最大请求数，会阻塞
func (this *Max) Get() {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	for this.curr >= this.max {
		this.locker.Wait()
	}
	this.curr += 1
	if this.EnableWait {
		this.wait.Add(1)
	}
}

//完成一个非阻塞请求
func (this *Max) Done() {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	this.curr -= 1

	if this.EnableWait {
		this.wait.Done()
	}
}

//请求一个非阻塞请求
func (this *Max) Add() bool {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	this.curr += 1
	if this.EnableWait {
		this.wait.Add(1)
	}
	return this.max < this.curr
}

//等待所有请求完成
func (this *Max) Wait() {
	if this.EnableWait {
		this.wait.Wait()
	}
}
