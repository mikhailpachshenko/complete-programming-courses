package main

import (
	"module/src/github.com/headfirstgo/gadget"
)

func playList(device gadget.TypePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TypePlayer{}
	mixtape := []string{"first", "second", "third"}
	playList(player, mixtape)
}
