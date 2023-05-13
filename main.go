package main

import "fmt"

import "github.com/nomadiq-sw/syncSQL/cmd/dbconn"

func main() {
	fmt.Println(dbconn.SayHello())
}
