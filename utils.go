package main

import (
	"log"
	"time"
)

func repeat(f func(), interval string) {
	f()
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatalln(err)
	}
	for range time.Tick(d) {
		f()
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
