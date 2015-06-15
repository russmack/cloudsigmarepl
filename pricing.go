package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListPricing struct {
	channels *replizer.Channels
}

func NewListPricing() *CommandListPricing {
	return &CommandListPricing{}
}

func (m *CommandListPricing) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listPricing
	cargo := CommandListPricing{}
	stateMachine.Start(cargo)
}

func (m *CommandListPricing) listPricing(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewPricing()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
