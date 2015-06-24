package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListAcls struct {
	channels *replizer.Channels
}

func NewListAcls() *CommandListAcls {
	return &CommandListAcls{}
}

func (m *CommandListAcls) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listAcls
	cargo := CommandListAcls{}
	stateMachine.Start(cargo)
}

func (m *CommandListAcls) listAcls(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewAcls()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
