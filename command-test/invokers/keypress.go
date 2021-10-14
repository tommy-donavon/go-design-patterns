package invokers

import "command-ui/interfaces"

type KeyPressInvoker struct {
	command interfaces.ICommand
}

func NewKeyPressInvoker(command interfaces.ICommand) *KeyPressInvoker {
	return &KeyPressInvoker{command}
}

func (kp *KeyPressInvoker) ExecuteCommand() {
	kp.command.Execute()
}

func (kp *KeyPressInvoker) UndoCommand() {
	kp.command.Undo()
}
