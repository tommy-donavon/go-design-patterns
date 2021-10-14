package components

import (
	"command-ui/commands"
	"command-ui/interfaces"
	"command-ui/invokers"
	"command-ui/receivers"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type App struct {
	app.Compo
	btnPresses []invokers.KeyPressInvoker
	receiver   interfaces.IReceiver
}

func (b *App) Render() app.UI {
	box := receivers.NewBox()
	div := app.
		Div().Body(
		box.Canvas,
	).TabIndex(0).Style("width", "100%").Style("height", "100%").OnKeyDown(b.press)
	b.receiver = box
	return div
}

func (b *App) press(ctx app.Context, e app.Event) {
	if b.btnPresses == nil {
		b.btnPresses = []invokers.KeyPressInvoker{}
	}
	switch e.Value.Get("key").String() {
	case "ArrowUp":
		kp := *invokers.NewKeyPressInvoker(commands.NewCommandUp(b.receiver))
		kp.ExecuteCommand()
		b.btnPresses = append(b.btnPresses, kp)
	case "ArrowDown":
		kp := *invokers.NewKeyPressInvoker(commands.NewCommandDown(b.receiver))
		kp.ExecuteCommand()
		b.btnPresses = append(b.btnPresses, kp)
	case "ArrowRight":
		kp := *invokers.NewKeyPressInvoker(commands.NewCommandRight(b.receiver))
		kp.ExecuteCommand()
		b.btnPresses = append(b.btnPresses, kp)
	case "ArrowLeft":
		kp := *invokers.NewKeyPressInvoker(commands.NewCommandLeft(b.receiver))
		kp.ExecuteCommand()
		b.btnPresses = append(b.btnPresses, kp)
	case "c":
		kp := *invokers.NewKeyPressInvoker(commands.NewCommandColorChange(b.receiver))
		kp.ExecuteCommand()
		b.btnPresses = append(b.btnPresses, kp)
	case "Escape":
		if len(b.btnPresses) > 0 {
			kp := b.btnPresses[len(b.btnPresses)-1]
			kp.UndoCommand()
			b.btnPresses = b.btnPresses[:len(b.btnPresses)-1]
		}
	}
}
