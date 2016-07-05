package max

import (
	"testing"
)

func Test_max(t *testing.T) {
	m := NewMax()
	for i := 0; i < 1000; i++ {
		m.Get()
		println(i)
		go func(j int) {
			m.Set()
			println("close", j)
		}(i)
	}
}
