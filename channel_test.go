package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// Channel is use for sending and receiving data from 1 goroutines to another different goroutine
// create new channel can used make() method

// IMPORTANT! make sure to CLOSE the channel after using it to avoid memory leaks and make the go garbage collectors remove it from memory
// Sending data to CHANNEL: channelName <- data
// Receive data from CHANNEL: data <- channelName
// Use data from channel to params: fmt.Println(<- channel)

// CAREFULL!
// sending goroutine data to channel without goroutines receive can caused hang or block the proccess until there is 1 goroutine to receive
// receive goroutine data from channel without goroutine sender can cause deadlock! because all goroutine is sleep

func TestCreateChannel(t *testing.T) {
	// create channel
	channel := make(chan string)
	// approach 1: close channel using defer to make sure channel closed at the end of func
	defer close(channel) 

	go func(){
		time.Sleep(2 * time.Second)

		channel <- "Dipa Galatian"
		fmt.Println("Done sending data to channel")

		channel <- "as;ldkfja;sldk"
	}()

	data2 := <- channel
	fmt.Println("Received data2 from channel:", data2)

	// why the receiver is not goroutines? NO, basically all proccess in go is treats as MAIN goroutines
	// so sending from goroutine anonymous func to data with channel is correct way
	data := <- channel
	fmt.Println("Received data from channel:", data)

	time.Sleep(5 * time.Second)


	// close channel at the end (approach 2)
	// close(channel)
}

// Channel as parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Secret data user"
}

func TestChannelAsParams(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println("Received secret from channel:", data)

	time.Sleep(5 * time.Second)
}

// Channel IN
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "super super secret data"
}

// Channel Out
func OnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println("Received super secret from channel:", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// send data to channel (using func params)
	go OnlyIn(channel)

	// receive data from channel (using func params)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel
// by default, channel is unbuffered, means data must sent and received at the same time
// buffered channel means data can be sent to channel without waiting for receiver to receive it
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func(){
		channel <- "dipa"
		channel <- "galatian"
		channel <- "super secret data"
	}()

	go func(){
		time.Sleep(2 * time.Second)

		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	// Send to channel without goroutine
	// channel <- "dipa"
	// channel <- "galatian"

	time.Sleep(2 * time.Second)
	fmt.Println("Channel done")

	// Receive from channel without goroutine
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	
}