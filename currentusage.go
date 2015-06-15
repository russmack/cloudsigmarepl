package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandCurrentUsage struct {
	channels *replizer.Channels
}

func NewCurrentUsage() *CommandCurrentUsage {
	return &CommandCurrentUsage{}
}

func (m *CommandCurrentUsage) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getCurrentUsage
	cargo := CommandCurrentUsage{}
	stateMachine.Start(cargo)
}

func (m *CommandCurrentUsage) getCurrentUsage(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewCurrentUsage()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
