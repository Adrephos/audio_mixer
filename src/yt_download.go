package src

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/iFaceless/godub"
	yt "github.com/kkdai/youtube/v2"
)

func download(wg *sync.WaitGroup, url string, outputName *string, errOut *error) {
	// Close wg at the end
	defer wg.Done()

	// Create YouTube client and get video with URL
	var client yt.Client
	video, err := client.GetVideo(url)

	if err != nil {
		*errOut = err
		return
	}

	// Get the video data with Audio
	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])

	if err != nil {
		*errOut = err
		return
	}

	output := video.ID

	// Create file to store video
	file, err := os.Create(output)

	if err != nil {
		*errOut = err
		return
	}

	// Store video
	_, err = io.Copy(file, stream)

	if err != nil {
		*errOut = err
		return
	}
	
	defer file.Close()

	// Convert video to audio
	cmd := exec.Command("ffmpeg", "-i", output, "-vn", "-acodec", "libmp3lame", fmt.Sprintf("./songs/%v.mp3", output))
	cmd.Run()
	// Delete videos
	cmd = exec.Command("rm", output)
	cmd.Run()

	*outputName = output
	*errOut = err
}

func MixYoutubeAudio() (*godub.AudioSegment, error)  {
	var wg sync.WaitGroup
	wg.Add(2)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("First video URL: ")
	scanner.Scan()
	first := scanner.Text()

	fmt.Print("Second video URL: ")
	scanner.Scan()
	second := scanner.Text()

	fmt.Print("Output file name: ")
	scanner.Scan()
	output := scanner.Text()

	// File names of downloaded audio files
	var firstFile string
	var secondFile string
	var err error

	fmt.Println("Downloading videos...")
	go download(&wg, first, &firstFile, &err)
	go download(&wg, second, &secondFile, &err)

	wg.Wait()

	if err != nil {
		return nil, err
	}

	// Mix downloaded audio files
	firstFile = "./songs/" + firstFile + ".mp3"
	secondFile = "./songs/" + secondFile + ".mp3"
	result, err := Mix(firstFile, secondFile, output)

	// Delete downloaded audio files
	cmd := exec.Command("rm", firstFile, secondFile)
	cmd.Run()

	if err != nil {
		return nil, err
	}

	// Play song
	fmt.Print("Play song? (y/n): ")
	scanner.Scan()
	ans := scanner.Text()
	if ans == "y" || ans == "Y" {
		Play(output, result.Duration())
	}

	return result, nil
}
