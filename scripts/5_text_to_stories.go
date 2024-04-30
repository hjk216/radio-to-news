// https://pkg.go.dev/github.com/sashabaranov/go-openai
// https://github.com/sashabaranov/go-openai

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	openai "github.com/sashabaranov/go-openai"
)

func TextToStories(date string) {
	pathParent := PATH_MAIN + PATH_TEXT_FULL + date + "/"
	textChunks, _ := os.ReadDir(pathParent)
	chunkCount := len(textChunks)
	chunk := 0

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	c := openai.NewClient(OPENAI_API_KEY)
	ctx := context.Background()

	prompt := `
		You are given radio transcripts. Organize the data into paragraphs by incident, write in past tense, with each
		incident separated by a newline character.
		Include as much information as possible but do not make anything up.
		Decipher police codes for additional context.
	`

	for chunk < chunkCount {
		filePathText := pathParent + strconv.Itoa(chunk) + FORMAT_TEXT
		data, _ := os.ReadFile(filePathText)

		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0125,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: string(data),
				},
			},
		}

		resp, errRequest := c.CreateChatCompletion(ctx, req)

		if errRequest != nil {
			fmt.Printf("Completion error: %v\n", errRequest)
			return
		}

		f, _ := os.Create(PATH_MAIN + PATH_STORIES + date + "/" + strconv.Itoa(chunk) + FORMAT_TEXT)

		_, err := f.WriteString(resp.Choices[0].Message.Content)

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		chunk += 1
	}
}
