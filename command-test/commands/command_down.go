package commands

import "command-ui/interfaces"

type commandDown struct {
	receiver interfaces.IReceiver
}

func NewCommandDown(receiver interfaces.IReceiver) *commandDown {
	return &commandDown{receiver}
}

func (cu *commandDown) Execute() {
	cu.receiver.Down()
}

func (cu *commandDown) Undo() {
	cu.receiver.Up()
}
