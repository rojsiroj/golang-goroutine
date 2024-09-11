package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(sender *UserBalance, receiver *UserBalance, amount int) {
	sender.Lock()
	fmt.Println("Lock", sender.Name)
	sender.Change(-amount)

	time.Sleep(1 * time.Second)

	receiver.Lock()
	fmt.Println("Lock", receiver.Name)
	receiver.Change(amount)

	time.Sleep(1 * time.Second)

	sender.Unlock()
	receiver.Unlock()
}

func TestDeadLock(t *testing.T) {
	sender := UserBalance{Name: "Siroj", Balance: 100000}
	receiver := UserBalance{Name: "Izz", Balance: 100000}

	// Deadlock, goroutine saling menunggu
	go Transfer(&sender, &receiver, 10000)
	go Transfer(&receiver, &sender, 20000)

	time.Sleep(3 * time.Second)

	fmt.Println("User ", sender.Name, ", Balance ", sender.Balance)
	fmt.Println("User ", receiver.Name, ", Balance ", receiver.Balance)
}
