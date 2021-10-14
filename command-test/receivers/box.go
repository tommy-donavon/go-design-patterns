package receivers

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Box struct {
	Canvas app.HTMLDiv
}

var (
	leftPosition      int      = 50
	topPosition       int      = 50
	possibleColors    []string = []string{"red", "yellow", "green", "blue"}
	currentColorIndex int      = 0
)

func NewBox() *Box {
	return &Box{
		Canvas: app.Div().
			ID("box").
			Class("red").
			Style("width", "50px").
			Style("height", "50px").
			Style("left", fmt.Sprintf("%d%%", leftPosition)).
			Style("top", fmt.Sprintf("%d%%", topPosition)).
			Style("position", "absolute"),
	}
}

func (b *Box) Up() {
	if topPosition > 0 {
		topPosition -= 1
		elm := app.Window().GetElementByID("box")
		elm.Set("style.top", fmt.Sprintf("%d%%", topPosition))
	}
}
func (b *Box) Down() {
	if topPosition < 100 {
		topPosition += 1
		elm := app.Window().GetElementByID("box")
		elm.Set("style.top", fmt.Sprintf("%d%%", topPosition))
	}
}
func (b *Box) Left() {
	if leftPosition > 0 {
		leftPosition -= 1
		elm := app.Window().GetElementByID("box")
		elm.Set("style.left", fmt.Sprintf("%d%%", leftPosition))
	}
}
func (b *Box) Right() {
	if leftPosition < 100 {
		leftPosition += 1
		elm := app.Window().GetElementByID("box")
		elm.Set("style.left", fmt.Sprintf("%d%%", leftPosition))
	}
}

func (b *Box) NextColor() {
	if currentColorIndex < len(possibleColors)-1 {
		currentColorIndex++
		elm := app.Window().GetElementByID("box")
		elm.Set("className", possibleColors[currentColorIndex])
	}

}
func (b *Box) PreviousColor() {
	if currentColorIndex > 0 {
		currentColorIndex--
		elm := app.Window().GetElementByID("box")
		elm.Set("className", possibleColors[currentColorIndex])
	}

}
