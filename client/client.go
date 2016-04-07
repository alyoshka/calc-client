package client

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/alyoshka/calc-server/gen-go/calculate"
)

// Client - client for rpc calculator server
type Client struct {
	calculator *calculate.CalculatorClient
}

func NewClient(addr string) (*Client, error) {
	protocolFactory := thrift.NewTJSONProtocolFactory()
	transportFactory := thrift.NewTTransportFactory()
	socket, err := thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}
	if socket == nil {
		return nil, fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}
	transport := transportFactory.GetTransport(socket)
	if transport == nil {
		return nil, fmt.Errorf("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}

	err = transport.Open()
	if err != nil {
		return nil, err
	}

	return &Client{
		calculator: calculate.NewCalculatorClientFactory(transport, protocolFactory),
	}, nil
}

// Handle - send single command to server
func (c *Client) Handle(input string) (result float64, err error) {
	if c.calculator == nil {
		err = errors.New("Client is not initialized")
		return
	}
	command := strings.Split(strings.Trim(input, "\n"), " ")
	if len(command) != 3 {
		err = errors.New("Not enough arguments")
		return
	}
	arg1, err := strconv.ParseFloat(command[0], 64)
	if err != nil {
		err = errors.New("Failed to parse first argument")
		return
	}
	arg2, err := strconv.ParseFloat(command[2], 64)
	if err != nil {
		err = errors.New("Failed to parse second argument")
		return
	}
	var operation calculate.Operation
	switch command[1] {
	case "+":
		operation = calculate.Operation_ADD
	case "-":
		operation = calculate.Operation_SUB
	case "/":
		operation = calculate.Operation_DIVIDE
	case "*":
		operation = calculate.Operation_MULTIPLY
	default:
		err = errors.New("Unknown operation")
		return
	}
	return c.calculator.Calculate(operation, arg1, arg2)
}

// Interact - start interactive session
func (c *Client) Interact() {
	fmt.Printf("Please enter the command to server\nIf you want to leave, type \"quit\"\n")
	consolereader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := consolereader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read string, error: %s", err)
			return
		}
		input = strings.Trim(input, "\n")
		if input == "quit" {
			fmt.Println("Bye")
			break
		}
		result, err := c.Handle(input)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println(result)
		}
	}
}

// Close - close connection to server
func (c *Client) Close() {
	c.calculator.Transport.Close()
}
