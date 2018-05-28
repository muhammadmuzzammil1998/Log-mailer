package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// Repeat -
func Repeat(f func(), interval string) {
	f()
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatalln(err)
	}
	for range time.Tick(d) {
		f()
	}
}

// Check -
func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Ask -
func Ask(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	r, _, _ := reader.ReadLine()
	return string(r)
}
