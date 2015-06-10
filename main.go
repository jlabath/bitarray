package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

func main() {
	limit := flag.Int("limit", 1000, "upper limit at which we abort the sieve")
	flag.Parse()
	if *limit < 1 {
		log.Fatalf("limit is invalid or pointless: %d", *limit)
	}
	primes := eratosthenes(*limit)
	//print only last few
	offset := len(primes) - 50
	if offset < 0 {
		offset = 0
	}
	for _, v := range primes[offset:] {
		fmt.Println(v)
	}
}

//eratosthenes calculates the Sieve of Eratosthenes up to n
func eratosthenes(n int) []int {
	//allocs
	primes := make([]int, 0, 100)
	ary := make([]bool, n-1) //ary contains 2 .. n
	//init to true
	for i := range ary {
		ary[i] = true
	}
	//calc
	sqn := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqn; i++ {
		if ary[i-2] {
			step := 0
			for j := i * i; j <= n; j = (i * i) + (step * i) {
				ary[j-2] = false
				step++

			}
		}
	}
	for i := range ary {
		if ary[i] {
			primes = append(primes, i+2)
		}
	}
	return primes
}
