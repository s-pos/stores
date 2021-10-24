package main

import "github.com/s-pos/go-utils/config"

func init() {
	serviceName := "stores"

	config.Load(serviceName)
}

func main() {
	forever := make(chan bool)
	<-forever
}
