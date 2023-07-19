package gui

import (
	"fmt"
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
	icon, err := fyne.LoadResourceFromPath("/home/hooman/code/allegro/clef.png")
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
	ap := audio.MakeStreamer("/home/hooman/test.flac")
	ap.Ctrl.Paused = true
	skip_next := widget.NewToolbarAction(theme.MediaSkipNextIcon(), func() {
		//TODO
		fmt.Println("next")
	})
	skip_previous := widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), func() {
		fmt.Println("previous")
	})

	play_button := widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
		if ap.Streamer.Position() == 0 && ap.Ctrl.Paused {
			go ap.PlayMusic()
		}
		speaker.Lock()
		ap.Ctrl.Paused = !ap.Ctrl.Paused
		speaker.Unlock()
	})
	speaker.Lock()
	Length := widget.NewLabel(ap.Format.SampleRate.D(ap.Streamer.Len()).Truncate(time.Second).String())
	Position := widget.NewLabel("0s")
	speaker.Unlock()
	slider := widget.NewSlider(0, float64(ap.Streamer.Len()))
	toolbar_slider := ToolbarSlider{Slider: slider, Min: slider.Min, Max: slider.Max}
	go func() {
		for ap.Streamer.Position() != ap.Streamer.Len() {
			if !ap.Ctrl.Paused {
				play_button.Icon = theme.MediaPauseIcon()
			} else {
				play_button.Icon = theme.MediaPlayIcon()
			}
			toolbar_slider.Slider.SetValue(float64(ap.Streamer.Position()))
			Position.SetText(ap.Format.SampleRate.D(ap.Streamer.Position()).Truncate(time.Second).String())
			time.Sleep(time.Second / 20)

		}
		speaker.Lock()
		ap.Ctrl.Paused = true
		speaker.Unlock()
		play_button.SetIcon(theme.MediaPlayIcon())
	}()
	toolbar := widget.NewToolbar(
		skip_previous,
		play_button,
		skip_next,
		widget.NewToolbarSeparator(),
		&toolbar_slider,
	)
	go func() {
		for {
			//toolbar.Refresh()
			time.Sleep(time.Second)
		}
	}()
	fmt.Println(Position, Length)
	content := container.NewBorder(toolbar, nil, nil, nil, nil)
	w.SetContent(content)
	w.ShowAndRun()
}
