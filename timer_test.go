package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// A timer is useful when you need to perform a task once after a specific duration has elapsed.
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestTimerAfter(t *testing.T) {
	// only pass channel, not timer object
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	fmt.Println(time.Now())

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute after 5 second")
		fmt.Println(time.Now())
		group.Done()
	})

	group.Wait()
	fmt.Println(time.Now())
}
