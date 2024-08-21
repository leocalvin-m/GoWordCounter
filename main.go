package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	
	countLine := flag.Bool("l",false,"Count Lines")
	flag.Parse()
	args := flag.Args()
	var input io.Reader
	if(len(args) > 0){
		file, err := os.Open(args[0])
		if err != nil {
            fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}

		defer file.Close()
		input = file
	}else {
		input = os.Stdin
	}

	scanner := bufio.NewScanner(input)

	if !*countLine {
		scanner.Split(bufio.ScanWords)
	}

	count := 0
	for scanner.Scan() {

		count ++
	}


	if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }


	log.Printf("Count: %d\n",count)
	









}