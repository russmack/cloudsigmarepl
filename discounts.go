package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListDiscounts struct {
	channels *replizer.Channels
}

func NewListDiscounts() *CommandListDiscounts {
	return &CommandListDiscounts{}
}

func (m *CommandListDiscounts) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listDiscounts
	cargo := CommandListDiscounts{}
	stateMachine.Start(cargo)
}

func (m *CommandListDiscounts) listDiscounts(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDiscounts()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
