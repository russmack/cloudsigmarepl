package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandVlans struct {
	channels *replizer.Channels
}

func NewVlans() *CommandVlans {
	return &CommandVlans{}
}

func (m *CommandVlans) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getVlans
	cargo := CommandVlans{}
	stateMachine.Start(cargo)
}

func (m *CommandVlans) getVlans(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewVlans()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		m.channels.ResponseChan <- fmt.Sprintf("Error calling client. %s", err)
		return nil
	}
	m.channels.ResponseChan <- string(resp)
	return nil
}
