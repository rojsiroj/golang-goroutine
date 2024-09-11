package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(id string) {
	fmt.Println("Hello World ", id)
}

func DisplayNumber(number int) {
	fmt.Println("Number: ", number)
}

func TestCreateGoRoutine(t *testing.T) {
	// go routine, async
	go RunHelloWorld("3")

	// not using go routine, sync
	RunHelloWorld("1")

	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func BenchmarkDisplayNumberNotUsingGoRoutine(b *testing.B) {
	for i := 1; i < b.N; i++ {
		DisplayNumber(i)
	}
}
func BenchmarkDisplayNumberManyGoRoutine(b *testing.B) {
	for i := 1; i < b.N; i++ {
		go DisplayNumber(i)
	}
}
