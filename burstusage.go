package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandBurstUsage struct {
	channels *replizer.Channels
}

func NewBurstUsage() *CommandBurstUsage {
	return &CommandBurstUsage{}
}

func (m *CommandBurstUsage) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getBurstUsage
	cargo := CommandBurstUsage{}
	stateMachine.Start(cargo)
}

func (m *CommandBurstUsage) getBurstUsage(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewBurstUsage()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
