package main

import (
"fmt"
"net/http"
"flag"
"runtime"
)

func getPage(page string, done chan int) {
	fmt.Printf("reading %s\n", page)
	_, err := http.Get(page)
	if err != nil {
		fmt.Printf("error reading the page \nPage=%s \nError=%v\n", page, err)
	}
	done <- 1
}

var (

flagThread = flag.Int("T", 1, "Number of parllel request")
flagPage   = flag.String("P", "http://www.google.com", "The webpage you want to read ")

)

func main () {

	flag.Parse()

	fin := make (chan int)

	runtime.GOMAXPROCS(*flagThread)

	fmt.Printf("flags Parllel=%d Page=%s\n", *flagThread, *flagPage)

	for i:=0; i<*flagThread; i++ {

		go getPage(*flagPage, fin)
	}
	

	for i:=0; i<*flagThread; i++ {

		<-fin
	}
}
