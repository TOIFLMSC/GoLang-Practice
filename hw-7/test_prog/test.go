package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	vars := os.Args[1:]
	for _, envvar := range vars {
		path, ok := os.LookupEnv(envvar)
		if ok != true {
			log.Println("Error with getting value of the environment variable")
			return
		} else {
			fmt.Println(path)
		}
	}
}
