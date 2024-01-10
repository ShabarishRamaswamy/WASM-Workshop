package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	stringsToBeHashed := []string{"hashThisString", "hashThatThing", "1123544884451484454145451154544", "Channels are a typed conduit through which you can send and receive values with the channel operator.", "1", "95403d97179341f258a6e903b53bcdba6024e00f0c9eeea86cecf81b15f75f93"}

	c := make(chan string)

	for _, str := range stringsToBeHashed {
		go hashString(str, c)
	}

	fmt.Println(<-c, <-c, <-c, <-c, <-c)

	fmt.Println(time.Since(startTime))
}

func hashString(str string, c chan string) {
	h := sha256.New()

	h.Write([]byte(str))
	bs := h.Sum(nil)

	c <- fmt.Sprintf("%x\n", bs)
}
