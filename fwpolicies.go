package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListFwPolicies struct {
	channels *replizer.Channels
}

func NewListFwPolicies() *CommandListFwPolicies {
	return &CommandListFwPolicies{}
}

func (m *CommandListFwPolicies) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listFwPolicies
	cargo := CommandListFwPolicies{}
	stateMachine.Start(cargo)
}

func (m *CommandListFwPolicies) listFwPolicies(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewFwPolicies()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
