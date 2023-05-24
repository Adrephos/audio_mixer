package src

import (
	"github.com/iFaceless/godub"
)

func Mix(path1 string, path2 string) (*godub.AudioSegment, error) {
	segment, err := godub.NewLoader().Load(path1)

	if err != nil {
		return nil, err
	}

	segment2, err := godub.NewLoader().Load(path2)

	if err != nil {
		return nil, err
	}

	var overlaidSeg *godub.AudioSegment

	if segment.Duration() > segment2.Duration() {
		overlaidSeg, err = segment.Overlay(segment2, &godub.OverlayConfig{LoopToEnd: false})
	} else {
		overlaidSeg, err = segment2.Overlay(segment, &godub.OverlayConfig{LoopToEnd: false})
	}
	

	if err != nil {
		return nil, err
	}

	return overlaidSeg, err
}
