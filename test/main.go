package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

type test struct {
	test_a string
	test_b string
}

type test_struct struct {
	a string
	b *test
	c []*test
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// catchs system signal
	t := &test_struct{
		a: "a",
	}
	t.b = &test{
		test_a: "test_a",
		test_b: "test_b",
	}
	t.c = append(t.c, &test{
		test_a: "slice_test_a",
		test_b: "slice_test_b",
	})
	log.Println(t)
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
