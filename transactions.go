package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListSubscriptions struct {
	channels *replizer.Channels
}

func NewListSubscriptions() *CommandListSubscriptions {
	return &CommandListSubscriptions{}
}

func (m *CommandListSubscriptions) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listSubscriptions
	cargo := CommandListSubscriptions{}
	stateMachine.Start(cargo)
}

func (m *CommandListSubscriptions) listSubscriptions(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewSubscriptions()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
