package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/jlabath/bitarray"
)

func main() {
	limit := flag.Int("limit", 100000000, "upper limit at which we abort the sieve")
	flag.Parse()
	if *limit < 1 {
		log.Fatalf("limit is invalid or pointless: %d", *limit)
	}
	primes := eratosthenes(*limit)
	//print only last few
	offset := len(primes) - 50
	if offset < 0 {
		offset = 0
	} else {
		fmt.Println("Printing last 50 primes found")
	}
	for _, v := range primes[offset:] {
		fmt.Println(v)
	}
}

// eratosthenes calculates the Sieve of Eratosthenes up to n
// http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func eratosthenes(n int) []int {
	//allocs
	primes := make([]int, 0, 100)
	ary := bitarray.New(n - 1) //ary contains 2 .. n
	//init to true
	ary.Fill(1)
	//calc
	sqn := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqn; i++ {
		if ary.IsSet(i - 2) {
			step := 0
			for j := i * i; j <= n; j = (i * i) + (step * i) {
				ary.Unset(j - 2)
				step++

			}
		}
	}
	for i := 0; i < ary.Length(); i++ {
		if ary.IsSet(i) {
			primes = append(primes, i+2)
		}
	}
	return primes
}
