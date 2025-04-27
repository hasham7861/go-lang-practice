package main


import "fmt"

// channels are a way to send data between main thread and go routine(s)
func main() {

	// to make a channel, simply use make function with channel value type
	messages := make(chan string)

	// to send a value to declared channel on this code, use <- operator
	go func () { messages <- "ping"}()

	// we can read from the channel using the same operator
	msg := <-messages
	fmt.Println(msg)
}