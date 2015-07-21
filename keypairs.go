package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListKeypairs struct {
	channels *replizer.Channels
}

func NewListKeypairs() *CommandListKeypairs {
	return &CommandListKeypairs{}
}

func (m *CommandListKeypairs) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listKeypairs
	cargo := CommandListKeypairs{}
	stateMachine.Start(cargo)
}

func (m *CommandListKeypairs) listKeypairs(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewKeypairs()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
