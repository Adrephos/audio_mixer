package src

import (
	"log"
	"github.com/Vernacular-ai/godub"
)

func Mix(path1 string, path2 string) *godub.AudioSegment {
	segment, err := godub.NewLoader().Load(path1)

	if err != nil {
		log.Fatal(err)
	}

	segment2, err := godub.NewLoader().Load(path2)

	if err != nil {
		log.Fatal(err)
	}

	overlaidSeg, err := segment.Overlay(segment2, &godub.OverlayConfig{LoopToEnd: false})

	if err != nil {
		log.Fatal(err)
	}

	return overlaidSeg
}
