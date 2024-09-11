package golanggoroutine

import (
	"fmt"
	"strconv"
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

func OnlyIn(channel chan<- string) {
	// data := <- channel  -- will be invalid
	channel <- "Izz Luthfi El Shirazy"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
	// channel <- "Izz Luthfi El Shirazy" -- will be invalid
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	// Create channel with 3 buffer
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Izz"
		channel <- "Luthfi"
		channel <- "El"
		channel <- "Shirazy" // -- wait buffer
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		// fmt.Println(<-channel)
	}()

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i+1)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
