package test

import (
	"testing"
	"time"
)

func TestUnix(t *testing.T) {
	pre := time.Now().Unix()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		pre = time.Now().Unix()
		if i != 0 && time.Now().Unix()-pre != 1 {
			t.Errorf("not really 1 second")
		}
	}
}
