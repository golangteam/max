package max

import (
	"sync"
)

type Max struct {
	max    int
	locker *sync.Cond
	curr   int
}

func NewMax() *Max {
	return &Max{
		max:    100,
		curr:   100,
		locker: sync.NewCond(&sync.Mutex{}),
	}
}

func (this *Max) SetMax(max int) {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	this.max = max
	this.curr = max
}
func (this *Max) Set() {
	this.locker.L.Lock()
	defer this.locker.L.Unlock()
	this.curr += 1
	this.locker.Signal()
}
func (this *Max) Get() {

	this.locker.L.Lock()
	for this.curr <= 0 {
		this.locker.Wait()
	}
	this.curr -= 1
	this.locker.L.Unlock()
}
