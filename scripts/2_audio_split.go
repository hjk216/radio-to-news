// https://pkg.go.dev/github.com/Vernacular-ai/godub#section-readme

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Vernacular-ai/godub"
)

func AudioSplit(
	date string,
) {
	filePathOriginal := PATH_MAIN + PATH_AUDIO_ORIGINAL + date + "/"
	filePathChunks := PATH_MAIN + PATH_AUDIO_CHUNKS + date + "/"

	audio, err := os.ReadDir(PATH_MAIN + PATH_AUDIO_ORIGINAL + date + "/")
	audioCount := len(audio)

	if err != nil || audioCount == 0 {
		fmt.Println("Error: Audio or directory not found.")
		return
	}
	i := 0
	for i < audioCount {
		segment, _ := godub.NewLoader().Load(filePathOriginal + audio[i].Name())
		filePath := filePathChunks + strconv.Itoa(i) + "/"

		SplitOriginalChunk(segment, filePathOriginal, filePath)

		i += 1
	}
}

func SplitOriginalChunk(
	segment *godub.AudioSegment,
	filePathOriginal string,
	filePathChunks string,
) {
	segment_duration := segment.Duration().Seconds()

	clipCount := 0
	clipStart := 0 * time.Microsecond

	for (float32(clipStart) / 1000) < float32(segment_duration) {
		slicedSegment, err := segment.Slice(clipStart, clipStart+(30*time.Microsecond))

		if err != nil {
			fmt.Println("Error:", err)
		}

		filePath := filePathChunks + strconv.Itoa(clipCount) + FORMAT_AUDIO

		godub.NewExporter(filePath).WithDstFormat("mp3").Export(slicedSegment)

		clipCount += 1
		clipStart += (30 * time.Microsecond)
	}
}
