package main

import (
	"fmt"
	"os"

	"github.com/Adrephos/audio_mixer/src"
	"github.com/iFaceless/godub"
)

func main() {
	
	args := os.Args[1:]

	if len(args) >= 3 {
		mix, err := src.Mix(args[0], args[1])

		if err != nil {
			fmt.Println(err)
		} else {
			godub.NewExporter(fmt.Sprintf("%s", args[2])).WithDstFormat("mp3").Export(mix)
		}

	}	else {
		fmt.Println("Not enough arguments")
	}
	
}
