package channels

import (
	"fmt"
	"time"
)

func ChannelsImpl() {
	fmt.Println("This is the channels implementation.")

	ch := make(chan int)

	// go routine to send data to the channel
	go func() {
		for i := range 6 {
			ch <- i
			fmt.Println("Sending:", i)
		}
		close(ch)
		fmt.Println("Sending data to channel...")
	}()

	time.Sleep(1 * time.Second) // wait for a moment to ensure data is sent
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
