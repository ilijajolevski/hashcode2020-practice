package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lineNumber int
	var maxPizzaSlices = 0
	var PizzaTypes = 0

	fptr := flag.String("fpath", "./a_example.in", "file path to read from")
	//fptr := flag.String("fpath", "./b_small.in", "file path to read from")
	//fptr := flag.String("fpath", "./c_medium.in", "file path to read from")
	//fptr := flag.String("fpath", "./d_quite_big.in", "file path to read from")
	//fptr := flag.String("fpath", "./e_also_big.in", "file path to read from")

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

		if lineNumber == 0 {
			fmt.Println(s.Text())
			words := strings.Fields(s.Text())

			maxPizzaSlices, err = strconv.Atoi(words[0])
			if err != nil {
				// handle error
			}
			PizzaTypes, err = strconv.Atoi(words[1])
			if err != nil {
				// handle error
			}

			fmt.Println("Max Pizza Slices: " + strconv.Itoa(maxPizzaSlices))
			fmt.Println("Pizza Types: " + strconv.Itoa(PizzaTypes))
		}

		if lineNumber == 1 {
			//fmt.Println(s.Text())
			//words := strings.Fields(s.Text())
			sliceData := strings.Split(s.Text(), " ")
			fmt.Println(sliceData)
		}
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
