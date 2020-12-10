package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := fib(os.Stdout, 7); err != nil {
		log.Fatal(err)
	}
}

func fib(writer io.Writer, n int) error {
	current := 0
	next := 1
	sign := 1
	if n < 0 {
		sign = -1
		n *= -1
	}
	nextSign := sign
	for i := 0; i < n; i++ {
		sign *= nextSign
		current, next = next, current+next
		if _, err := fmt.Fprintln(writer, current*sign); err != nil {
			return err
		}
	}
	return nil
}
