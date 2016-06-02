package channelp

import (
	"fmt"
	"runtime"
	"sync"
)

func ChanBasic() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		<-c
	}()
	c <- false
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func Invoke() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func GoSync(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}

func InvokeSync() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go GoSync(&wg, i)
	}
	wg.Wait()
}

func SelectChannel() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1")
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2")
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)

	for i := 0; i < 2; i++ {
		<-o
	}
}

func Chat(c chan string) {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Chat :Hi ,#%d", i)
		i++
	}
}
