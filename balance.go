package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandBalance struct {
	channels *replizer.Channels
}

func NewBalance() *CommandBalance {
	return &CommandBalance{}
}

func (m *CommandBalance) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getBalance
	cargo := CommandBalance{}
	stateMachine.Start(cargo)
}

func (m *CommandBalance) getBalance(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewBalance()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
