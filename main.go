package main

import (
	"fmt"
	"os"

	"github.com/Adrephos/audio_mixer/src"
)

func main() {
	
	args := os.Args[1:]
	var outputPath string

	if len(args) >= 3 {
		index := 0
		if args[0] == "-p" {
			index = 1
		}
		outputPath = args[index + 2]
		mix, err := src.Mix(args[index], args[index + 1], outputPath)

		if err != nil {
			fmt.Println(err)
		}

		if index == 1 {
			src.Play(outputPath, mix.Duration())
		}

	}	else if len(args) >= 1 {
		if args[0] == "-y" {
			_, err := src.MixYoutubeAudio()

			if err != nil {
				fmt.Println(err)
			}
		}

	}else {
		fmt.Println("Not enough arguments")
	}

}
