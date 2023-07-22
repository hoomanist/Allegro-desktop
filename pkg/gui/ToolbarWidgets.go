package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ToolbarSlider struct {
	Min   float64
	Max   float64
	Value float64
}
type ToolbarLabel struct {
	Text string
}
type ExtendedSlider struct {
	widget.Slider
}

func (t *ToolbarLabel) ToolbarObject() fyne.CanvasObject {
	l := widget.NewLabel(t.Text)
	return l
}

func NewToolbarLabel(Text string) *ToolbarLabel {
	return &ToolbarLabel{Text: Text}
}

func (t *ToolbarSlider) ToolbarObject() fyne.CanvasObject {
	s := &ExtendedSlider{}
	s.Max = t.Max
	s.Min = t.Min
	s.ExtendBaseWidget(s)
	s.SetValue(t.Value)
	return s
}

func (t *ExtendedSlider) MinSize() fyne.Size {
	return fyne.NewSize(700, 30)
}
func NewToolbarSlider(Min float64, Max float64) *ToolbarSlider {
	return &ToolbarSlider{Min: Min, Max: Max}
}
