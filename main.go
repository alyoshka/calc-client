package main

import (
	"flag"
	"fmt"

	"github.com/alyoshka/calc-client/client"
)

var host = flag.String("host", "localhost:8080", "Host of the calculator rpc server")
var command = flag.String("command", "", "Command to server")

func main() {
	flag.Parse()
	client, err := client.NewClient(*host)
	if err != nil {
		fmt.Println("error running client:", err)
		return
	}
	defer client.Close()
	if *command != "" {
		result, err := client.Handle(*command)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println(result)
		}
	} else {
		client.Interact()
	}
}
