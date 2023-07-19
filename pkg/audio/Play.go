package audio

import (
	"os"
	"strings"
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
func (ap *AudioPanel) PausePlay() {
	if ap.Streamer.Position() == 0 && ap.Ctrl.Paused {
		go ap.PlayMusic()
	}
	speaker.Lock()
	ap.Ctrl.Paused = !ap.Ctrl.Paused
	speaker.Unlock()
}
func (ap *AudioPanel) Position() string {
	pos := ap.Format.SampleRate.D(ap.Streamer.Position()).Truncate(time.Second)
	len := ap.Format.SampleRate.D(ap.Streamer.Len()).Truncate(time.Second)
	converter := func(t time.Duration) string {
		if t.Truncate(time.Minute) == 0 {
			t_sec, _, _ := strings.Cut(t.String(), "s")
			return strings.Join([]string{"0", t_sec}, ":")
		}
		t_min, t_sec, _ := strings.Cut(t.String(), "m")
		t_sec, _, _ = strings.Cut(t_sec, "s")
		return strings.Join([]string{t_min, t_sec}, ":")
	}
	return strings.Join([]string{converter(pos), converter(len)}, "/")
}
