package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var lineNumber int

	fptr := flag.String("fpath", "./a_example.in", "file path to read from")

	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	//s.Scan()
	for s.Scan() {
		fmt.Println(s.Text())
		//read line by line

		lineNumber++
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("read lines from file: " + strconv.Itoa(lineNumber))

	//end reading the input file
	fmt.Println("read the file completely")

}
