package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

const (
	PATH_MAIN           string = "content/"
	PATH_AUDIO_ORIGINAL string = "audio_original/"
	PATH_AUDIO_CHUNKS   string = "audio_chunks/"
	PATH_TEXT_CHUNKS    string = "text_chunks/"
	PATH_TEXT_FULL      string = "text_full/"
	PATH_STORIES        string = "stories/"
	FORMAT_AUDIO        string = ".mp3"
	FORMAT_TEXT         string = ".txt"
)

func main() {
	fmt.Println("Started")

	const date = "2024-04-14"

	if !SerializeDate(date) {
		return
	}

	godotenv.Load()

	AudioDownload()

	AudioSplit(date)

	SpeechToText(date)

	CombineText(date)

	TextToStories(date)

	fmt.Println("Finished")
}
