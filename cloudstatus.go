package main

import (
	"fmt"
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
	args := o.NewGet()
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
