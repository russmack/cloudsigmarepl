package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandLocations struct {
	channels *replizer.Channels
}

func NewLocations() *CommandLocations {
	return &CommandLocations{}
}

func (m *CommandLocations) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getLocations
	cargo := CommandLocations{}
	stateMachine.Start(cargo)
}

func (m *CommandLocations) getLocations(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewLocations()
	args := o.NewList()
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
