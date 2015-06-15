package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandProfile struct {
	channels *replizer.Channels
}

func NewProfile() *CommandProfile {
	return &CommandProfile{}
}

func (m *CommandProfile) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getProfile
	cargo := CommandProfile{}
	stateMachine.Start(cargo)
}

func (m *CommandProfile) getProfile(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewProfile()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
