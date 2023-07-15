package audio

import (
	"os"
	"time"

	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
)

func PlayMusic(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("no such file")
	}
	streamer, format, err := flac.Decode(f)
	if err != nil {
		panic("can't decode")
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
	defer streamer.Close()
}
