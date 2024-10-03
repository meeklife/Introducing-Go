# Concurrency

Large programs are often made up of many smaller subprograms. For example, a web server handles requests made from web browsers and serves up HTML web pages in response. Each request is handled like a small program.

It would be ideal for programs like these to be able to run their smaller components at the same time (in the case of the web server, to handle multiple requests). Making progress on more than one task simultaneously is known as concurrency. Go has rich support for concurrency using goroutines and channels.

## Goroutines

A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine, we use the keyword `go` followed by a function invocation:

```go
func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
    }
}
func main() {
    go f(0)
    var input string
    fmt.Scanln(&input)
}
```

**_NB: This program consists of two goroutines. The first goroutine is implicit and is the main function itself. The second goroutine is created when we call go f(0). Nor‐ mally, when we invoke a function, our program will execute all the statements in a function and then return to the next line following the invocation. With a goroutine, we return immediately to the next line and don’t wait for the function to complete. This is why the call to the Scanln function has been included; without it, the program would exit before being given the opportunity to print all the numbers._**

Goroutines are lightweight and we can easily create thousands of them.

## Channels

Channels provide a way for two goroutines to communicate with each other and synchronize their execution and communicate by passing values.

A new channel value can be made using the built-in function make.

```go
ic := make(chan int)
```

To send a value on a channel, use <- as a binary operator. To receive a value on a channel, use it as a unary operator.

```go
ic <- 3   // Send 3 on the channel.
n := <-ic // Receive a string from the channel.
```

Here is an example program using channels:
**_This program will print ping forever (hit Enter to stop it). A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel (in this case, we are passing strings). The left arrow operator (<-) is used to send and receive messages on the channel. c <- "ping" means send "ping". msg := <- c means receive a message and store it in msg._**

```go
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
    var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var inputt string
	fmt.Scanln(&inputt)
}
```

Using a channel like this synchronizes the two goroutines. When pinger attempts to send a message on the channel, it will wait until printer is ready to receive the message (this is known as blocking).

### Channel Direction

We can specify a direction on a channel type, thus restricting it to either sending or receiving. For example, pinger’s function signature can be changed to this:

```go
func pinger(c chan<- string) // can only be used to send strings
```

Now pinger is only allowed to send to c. Attempting to receive from c will result in a
compile-time error. Similarly, we can change printer to this:

```go
func printer(c <-chan string) // can only be used to receive strings
```

A channel that doesn’t have these restrictions is known as bidirectional. A bidirectional channel can be passed to a function that takes send-only or receive-only channels, but the reverse is not true.

### Select

The select statement waits for multiple send or receive opera­tions simul­taneously. it works like a switch but for channels:

- The statement blocks as a whole until one of the operations becomes unblocked.
- If several cases can proceed, a single one of them will be chosen at random.

Example:
**_This program prints “from 1” every 2 seconds and “from 2” every 3 seconds. select picks the first channel that is ready and receives from it (or sends to it). If more than one of the channels are ready, then it randomly picks which one to receive from. If none of the channels are ready, the statement blocks until one becomes available._**

```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "from 1"
            time.sleep(time.second * 2)
        }
    }()

    go func() {
        for {
            c2 <- "from 2"
            time.sleep(time.second * 3)
        }
    }()

    go func() {
        for {
            select{
                case msg1 := <- c1:
                    fmt.Println(msg1)
                case msg2 := <- c2:
                    fmt.Println(msg2)

            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}
```

The default case is always able to proceed and runs if all other cases are blocked.

```go
// never blocks
select {
    case x := <-ch:
        fmt.Println("Received", x)
    default:
        fmt.Println("Nothing available")
}
```

### Buffered Channels

By default channels are unbuffered, meaning that they will only accept sends (chan <-). Buffered channels accept a limited number of values without a corresponding receiver for those values.

To create a buffered channel, we pass a second argument to make function when creating a channel

```go
c := make(chan int, 2)

c <- "buffered"
c <- "channel"
```

This creates a buffered channel with a capacity of 2.
