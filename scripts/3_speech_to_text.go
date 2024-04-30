// https://pkg.go.dev/github.com/sashabaranov/go-openai
// https://github.com/sashabaranov/go-openai

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	openai "github.com/sashabaranov/go-openai"
)

func SpeechToText(
	date string,
) {
	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	c := openai.NewClient(OPENAI_API_KEY)
	ctx := context.Background()

	audioChunks, _ := os.ReadDir(PATH_MAIN + PATH_AUDIO_CHUNKS + date + "/")
	chunkCount := len(audioChunks)
	chunk := 0

	for chunk < chunkCount {
		clips, _ := os.ReadDir(PATH_MAIN + PATH_AUDIO_CHUNKS + date + "/" + strconv.Itoa(chunk) + "/")
		clipCount := len(clips)
		clip := 0

		newpath := filepath.Join(PATH_MAIN+PATH_TEXT_CHUNKS+date, strconv.Itoa(chunk))
		os.MkdirAll(newpath, os.ModePerm)

		for clip < clipCount {
			TranscribeClip(c, ctx, date, chunk, clip)
			clip += 1
		}
		chunk += 1
	}
}

func TranscribeClip(
	c *openai.Client,
	ctx context.Context,
	date string,
	chunk int,
	clip int,
) {
	filePathAudio := PATH_MAIN + PATH_AUDIO_CHUNKS + date + "/" + strconv.Itoa(chunk) + "/" + strconv.Itoa(clip) + FORMAT_AUDIO
	filePathText := PATH_MAIN + PATH_TEXT_CHUNKS + date + "/" + strconv.Itoa(chunk) + "/" + strconv.Itoa(clip) + FORMAT_TEXT

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePathAudio,
	}

	resp, errRequest := c.CreateTranscription(ctx, req)

	if errRequest != nil {
		fmt.Printf("Transcription error: %v\n", errRequest)
		return
	}

	f, _ := os.Create(filePathText)

	_, err := f.WriteString(resp.Text)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
