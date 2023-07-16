package audio

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
)

type AudioPanel struct {
	Streamer  beep.StreamSeeker
	Ctrl      *beep.Ctrl
	Resampler *beep.Resampler
	Volume    *effects.Volume
	Format    beep.Format
}

func NewAudioPanel(streamer beep.StreamSeeker, format beep.Format) *AudioPanel {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return &AudioPanel{streamer, ctrl, resampler, volume, format}

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
	return NewAudioPanel(streamer, format)
}
func (ap *AudioPanel) PlayMusic() {
	speaker.Init(ap.Format.SampleRate, ap.Format.SampleRate.N(time.Second/10))
	speaker.Play(ap.Volume)
	select {}
}
