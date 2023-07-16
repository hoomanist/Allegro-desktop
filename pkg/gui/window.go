package gui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/speaker"
	"github.com/hoomanist/Allegro/pkg/audio"
)

func MainWindow() {
	a := app.New()
	w := a.NewWindow("Allegro")
	w.Resize(fyne.NewSize(900, 600))
	ap := audio.MakeStreamer("/home/hooman/test.flac")
	go audio.PlayMusic(ap.Streamer, ap.Format)
	speaker.Lock()
	ap.Ctrl.Paused = true
	speaker.Unlock()
	play_button := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		speaker.Lock()
		ap.Ctrl.Paused = !ap.Ctrl.Paused
		fmt.Println(ap.Ctrl.Paused)
		speaker.Unlock()
	})
	slider := widget.NewSlider(0, float64(ap.Streamer.Len()))

	go func() {
		for ap.Streamer.Position() != ap.Streamer.Len() {
			if !ap.Ctrl.Paused {
				play_button.SetIcon(theme.MediaPauseIcon())
			} else {
				play_button.SetIcon(theme.MediaPlayIcon())
			}
			slider.SetValue(float64(ap.Streamer.Position()))
			time.Sleep(1 * time.Second)
		}
	}()
	content := container.New(layout.NewVBoxLayout(), container.New(layout.NewCenterLayout(), play_button), slider)
	w.SetContent(content)
	w.ShowAndRun()
}
