package main

import (
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandCloudStatus struct {
	channels *replizer.Channels
}

func NewCloudStatus() *CommandCloudStatus {
	return &CommandCloudStatus{}
}

func (m *CommandCloudStatus) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getCloudStatus
	cargo := CommandCloudStatus{}
	stateMachine.Start(cargo)
}

func (m *CommandCloudStatus) getCloudStatus(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewCloudStatus()
	args := o.NewList()
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
