package gui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/speaker"
	"github.com/hoomanist/Allegro/pkg/audio"
)

func MainWindow() {
	a := app.New()
	pwd, _ := os.Getwd()
	icon, err := fyne.LoadResourceFromPath(strings.Join([]string{pwd, "clef.png"}, "/"))
	if err != nil {
		panic("no icon")
	}
	app.SetMetadata(fyne.AppMetadata{
		Name:    "Allegro",
		Version: "pre-alpha",
		Icon:    icon,
	})
	w := a.NewWindow("Allegro")
	w.Resize(fyne.NewSize(900, 600))
	ap := audio.MakeStreamer(strings.Join([]string{pwd, "test.flac"}, "/"))
	ap.Ctrl.Paused = true
	skip_next := widget.NewToolbarAction(theme.MediaSkipNextIcon(), func() {
		//TODO
		fmt.Println("next")
	})
	skip_previous := widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), func() {
		fmt.Println("previous")
	})

	play_button := widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
		ap.PausePlay()
	})
	speaker.Lock()
	Position := NewToolbarLabel("")
	speaker.Unlock()
	slider := NewToolbarSlider(0, float64(ap.Streamer.Len()))
	go func() {
		for ap.Streamer.Position() != ap.Streamer.Len() {
			if !ap.Ctrl.Paused {
				play_button.Icon = theme.MediaPauseIcon()
			} else {
				play_button.Icon = theme.MediaPlayIcon()
			}
			slider.Value = float64(ap.Streamer.Position())
			Position.Text = ap.Position()
			time.Sleep(time.Second / 20)

		}
		fmt.Println("wifj")
		speaker.Lock()
		ap.Ctrl.Paused = true
		speaker.Unlock()
		play_button.Icon = theme.MediaPlayIcon()
		slider.Value = slider.Min
	}()
	toolbar := widget.NewToolbar(
		skip_previous,
		play_button,
		skip_next,
		widget.NewToolbarSeparator(),
		slider,
		widget.NewToolbarSeparator(),
		Position,
	)
	go func() {
		for {
			toolbar.Refresh()
			time.Sleep(time.Second)
		}
	}()
	content := container.NewBorder(nil, toolbar, nil, nil, nil)
	w.SetContent(content)
	w.ShowAndRun()
}
