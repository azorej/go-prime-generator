# go-prime-generator
Prime numbers generator on Golang with timeout functionality

# install
go get github.com/Azorej/go-prime-generator

#how to

```go

package main

import (
	"fmt"
	primes "github.com/Azorej/go-prime-generator"
	"time"
)

func main() {
	//generate primes less than 10000000, timeout 2 seconds
	primes := generate(10000000, time.Second*2)

	//primes count
	fmt.Println(len(primes))
}

func generate(maxN int, timeout time.Duration) []int {
	ch := make(chan int, 100)
	enough := primes.NewChanSignal()

	go primes.Generate(maxN, ch, enough)

	timer := time.NewTimer(timeout)
	res := make([]int, 0)

	for {
		select {
		case v, more := <-ch:
			if !more {
				return res
			}
			res = append(res, v)

		case <-timer.C:
			enough <- primes.NewSignal()
			println("Timeout :(")
			return res
		}
	}
}


