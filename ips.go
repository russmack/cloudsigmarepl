package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListIps struct {
	channels *replizer.Channels
}

func NewListIps() *CommandListIps {
	return &CommandListIps{}
}

func (m *CommandListIps) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listIps
	cargo := CommandListIps{}
	stateMachine.Start(cargo)
}

func (m *CommandListIps) listIps(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewIps()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
