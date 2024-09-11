package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Izz Luthfi El Shirazy"
		fmt.Println("Selesai mengirim data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)

}

func GiveMeRespone(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Izz Luthfi El Shirazy"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespone(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}
