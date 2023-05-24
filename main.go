package main

import (
	"fmt"
	"os"

	"github.com/Adrephos/audio_mixer/src"
	"github.com/iFaceless/godub"
)

func main() {
	
	args := os.Args[1:]
	var outputPath string

	if len(args) >= 3 {
		index := 0
		if args[0] == "-p" {
			index = 1
		}
		fmt.Println("Mxing songs ....")
		mix, err := src.Mix(args[index], args[index + 1])

		if err != nil {
			fmt.Println(err)
		} else {
			outputPath = fmt.Sprintf("%s", args[index + 2])
			godub.NewExporter(outputPath).WithDstFormat("mp3").Export(mix)
			fmt.Println("Done")
		}

		if index == 1 {
			src.Play(outputPath, mix.Duration())
		}

	}	else {
		fmt.Println("Not enough arguments")
	}
	
}
