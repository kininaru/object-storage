package main

import "fmt"

func MainLoop() {
	var input string
	var err error
	for {
		_, err = fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}
	}
}
