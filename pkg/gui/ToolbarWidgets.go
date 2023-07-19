package gui

import (
	"fmt"

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

func (t *ToolbarLabel) ToolbarObject() fyne.CanvasObject {
	l := widget.NewLabel(t.Text)
	return l
}

func NewToolbarLabel(Text string) *ToolbarLabel {
	return &ToolbarLabel{Text: Text}
}

func (t *ToolbarSlider) ToolbarObject() fyne.CanvasObject {
	s := widget.NewSlider(t.Min, t.Max)
	fmt.Println()
	s.SetValue(t.Value)
	return s
}
func (t *ToolbarSlider) MinSize() fyne.Size {
	return fyne.NewSize(750, 50)
}
func NewToolbarSlider(Min float64, Max float64) *ToolbarSlider {
	return &ToolbarSlider{Min: Min, Max: Max}
}
