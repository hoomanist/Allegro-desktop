package main

import (
	"fmt"
	"os"

	"github.com/hoomanist/Allegro/pkg/gui"
)

func main() {
	// Check if the user provided the audio file path as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: Allegro <audio_file_path>")
		return
	}

	audioFilePath := os.Args[1]
	gui.MainWindow(audioFilePath)
}
