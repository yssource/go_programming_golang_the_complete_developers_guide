package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/exvimmer/go_programming_language/pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	ClickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh() // part of widget.BaseWidget
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		ClickHandler: clickHandler,
	}
	swatch.ExtendBaseWidget(swatch)

	return swatch
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		parent:  s,
		objects: objects,
	}
}
