package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ToolbarSlider struct {
	Min    float64
	Max    float64
	Slider *widget.Slider
}

func (t *ToolbarSlider) ToolbarObject() fyne.CanvasObject {
	s := widget.NewSlider(t.Min, t.Max)
	t.Slider = s
	return s
}
