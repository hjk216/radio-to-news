package main

import (
	"fmt"
	"os"
	"strconv"
)

func CombineText(date string) {
	pathParent := PATH_MAIN + PATH_TEXT_CHUNKS + date + "/"
	textChunks, _ := os.ReadDir(pathParent)
	chunkCount := len(textChunks)
	chunk := 0

	for chunk < chunkCount {
		pathChunk := pathParent + strconv.Itoa(chunk) + "/"
		clips, _ := os.ReadDir(pathChunk)
		clipCount := len(clips)
		clip := 0

		fullText := ""

		for clip < clipCount {
			filePathText := pathChunk + strconv.Itoa(clip) + FORMAT_TEXT
			data, err := os.ReadFile(filePathText)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			fullText = fullText + " " + string(data)
			clip += 1
		}

		f, _ := os.Create(PATH_MAIN + PATH_TEXT_FULL + date + "/" + strconv.Itoa(chunk) + FORMAT_TEXT)

		_, err := f.WriteString(fullText)

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		chunk += 1
	}
}
