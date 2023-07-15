package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/hoomanist/Allegro/pkg/audio"
)

func MainWindow() {
	a := app.New()
	w := a.NewWindow("Allegro")
	w.Resize(fyne.NewSize(900, 600))
	play_button := widget.NewButtonWithIcon("play", theme.MediaPlayIcon(), func() {
		audio.PlayMusic("/home/hooman/test.flac")
	})
	w.SetContent(play_button)
	w.ShowAndRun()
}
