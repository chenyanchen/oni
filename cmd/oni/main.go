package main

import "log"

func main() {
	if err := New().Execute(); err != nil {
		log.Fatalln(err)
	}
}
