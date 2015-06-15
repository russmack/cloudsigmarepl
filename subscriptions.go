package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListTransactions struct {
	channels *replizer.Channels
}

func NewListTransactions() *CommandListTransactions {
	return &CommandListTransactions{}
}

func (m *CommandListTransactions) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listTransactions
	cargo := CommandListTransactions{}
	stateMachine.Start(cargo)
}

func (m *CommandListTransactions) listTransactions(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewTransactions()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
