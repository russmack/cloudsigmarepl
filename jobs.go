package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListJobs struct {
	channels *replizer.Channels
}

func NewListJobs() *CommandListJobs {
	return &CommandListJobs{}
}

func (m *CommandListJobs) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listJobs
	cargo := CommandListJobs{}
	stateMachine.Start(cargo)
}

func (m *CommandListJobs) listJobs(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewJobs()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
