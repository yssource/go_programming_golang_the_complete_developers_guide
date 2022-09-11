package ui

import (
	"fyne.io/fyne/v2"
	"github.com/exvimmer/go_programming_language/pixl/apptype"
	"github.com/exvimmer/go_programming_language/pixl/pxcanvas"
	"github.com/exvimmer/go_programming_language/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
