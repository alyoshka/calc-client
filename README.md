#calc-server
RPC calculator server  
Calculates simple arithmetic tasks.

##Requirements
You need Golang 1.5 or newer [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/code.html)  
Then install Godep:  
`go get github.com/tools/godep`

##Installation
Download calc-server:  
`go get github.com/alyoshka/calc-client`  
`cd $GOPATH/src/github.com/alyoshka/calc-client`  
Install  
`godep go install`  

##Usage
`calc-client [-host host:port] [-command "arg1 operation arg2"]`  
Client starts interactive session if no command argument passed.

##Commands
client accepts commands of the following form  
`arg1 operation arg2`  
Where arg1 and arg2 - numbers, operation is one of supported mathematical signs (+, -, * or /)
