package src

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	yt "github.com/kkdai/youtube/v2"
)

func Download(url string, output string) error {
	var client yt.Client
	video, err := client.GetVideo(url)

	if err != nil {
		return err
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])

	if err != nil {
		return err
	}

	file, err := os.Create(output)

	if err != nil {
		return err
	}

	_, err = io.Copy(file, stream)

	if err != nil {
		return err
	}
	
	defer file.Close()

	cmd := exec.Command("ffmpeg", "-i", output, "-vn", "-acodec", "libmp3lame", fmt.Sprintf("%v.mp3", output))
	cmd.Run()
	cmd = exec.Command("rm", output)
	cmd.Run()

	return nil
}
