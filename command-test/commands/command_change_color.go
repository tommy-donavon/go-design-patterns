package commands

import "command-ui/interfaces"

type commandColorChange struct {
	receiver interfaces.IReceiver
}

func NewCommandColorChange(receiver interfaces.IReceiver) *commandColorChange {
	return &commandColorChange{receiver}
}

func (cu *commandColorChange) Execute() {
	cu.receiver.NextColor()
}

func (cu *commandColorChange) Undo() {
	cu.receiver.PreviousColor()
}
