package main

import (
	"fmt"
	"os"

	"github.com/Adrephos/audio_mixer/src"
	"github.com/Vernacular-ai/godub"
)

func main() {
	
	args := os.Args[1:]

	if len(args) >= 3 {
		mix := src.Mix(args[0], args[1])

		godub.NewExporter(fmt.Sprintf("./songs/%s.mp3", args[2])).WithDstFormat("mp3").Export(mix)
	}	else {
		fmt.Println("Not enough arguments")
	}
	
}
