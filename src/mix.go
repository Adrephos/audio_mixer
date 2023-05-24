package src

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

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

func Play(path string, duration time.Duration) {
	// Start ffplay command
	cmd := exec.Command("mpv", "--volume=40", "--no-video", "--quiet", path)

	// Start the command
	fmt.Println("Playing", path)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting ffplay command:", err)
		return
	}

	durationFormat := fmt.Sprintf("%02d:%02d", int(duration.Minutes()), int(duration.Seconds()-float64(int(duration.Minutes())*60)))

	go func() {
		// Get the initial time.
		before := time.Now()
		for {
			// Get the current time.
			now := time.Now()

			// Calculate the time remaining in the song.
			remainingTime := now.Sub(before)
			remainingFormat := fmt.Sprintf("%02d:%02d", int(remainingTime.Minutes()), int(remainingTime.Seconds()-float64(int(remainingTime.Minutes())*60)))
			
			if int(remainingTime.Seconds()) >= int(duration.Seconds()) {
				break
			}

			// Print the time remaining in the song.
			fmt.Printf("\rTime elapsed: %v / %v", remainingFormat, durationFormat)

			// Sleep for 1 second.
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error:", err)
	}

	// Wait for user input to exit
	fmt.Print("\rSong finished, press Enter to exit.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
