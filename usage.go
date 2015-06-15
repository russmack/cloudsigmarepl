package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandUsage struct {
	channels *replizer.Channels
}

func NewUsage() *CommandUsage {
	return &CommandUsage{}
}

func (m *CommandUsage) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getUsage
	cargo := CommandUsage{}
	stateMachine.Start(cargo)
}

func (m *CommandUsage) getUsage(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewUsage()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
