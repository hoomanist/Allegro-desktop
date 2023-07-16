package audio

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
)

type AudioPanel struct {
	Ctrl     *beep.Ctrl
	Format   beep.Format
	Streamer beep.StreamSeekCloser
}

func MakeStreamer(path string) *AudioPanel {
	f, err := os.Open(path)
	if err != nil {
		panic("no such file")
	}
	streamer, format, err := flac.Decode(f)
	if err != nil {
		panic("can't decode")
	}
	ctrl := &beep.Ctrl{Streamer: streamer}
	return &AudioPanel{Ctrl: ctrl, Streamer: streamer, Format: format}
}
func PlayMusic(streamer beep.StreamCloser, format beep.Format) {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
