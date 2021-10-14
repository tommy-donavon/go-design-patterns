package commands

import "command-ui/interfaces"

type commandUp struct {
	receiver interfaces.IReceiver
}

func NewCommandUp(receiver interfaces.IReceiver) *commandUp {
	return &commandUp{receiver}
}

func (cu *commandUp) Execute() {
	cu.receiver.Up()
}

func (cu *commandUp) Undo() {
	cu.receiver.Down()
}
