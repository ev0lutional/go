package main

import (
	"fmt"
	"os"
)

func help() {
	var msg string

	fmt.Print("\033[H\033[2J") // clear Linux screen

	msg = `Algorithmic music
	
http://rcarduino.blogspot.com/2012/09/algorithmic-music-on-arduino.html
The goal is to send to STDOUT a pulse code modulation (PCM format)
Usage : ./algorithmicmusic | aplay

TODO:
fmt.Println("- pipe directly the value to aplay
fmt.Println("- remove the infinite loop in music()

`
	fmt.Printf("%s", msg)
}

func music() {
	var t int
	var c int

	for t = 0; ; t++ {
		c = (t>>7|t|t>>6)*10 + 4*(t&t>>13|t>>6)
		fmt.Printf("%c", c)
	}
}

func main() {
	if len(os.Args[1:]) == 1 && os.Args[1] == "--help" {
		help()
	} else {
		music()
	}
}
