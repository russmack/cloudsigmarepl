package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandCurrentUsage struct {
	channels *replizer.Channels
}

func NewCurrentUsage() *CommandCurrentUsage {
	return &CommandCurrentUsage{}
}

func (m *CommandCurrentUsage) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getCurrentUsage
	cargo := CommandCurrentUsage{}
	stateMachine.Start(cargo)
}

func (m *CommandCurrentUsage) getCurrentUsage(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewCurrentUsage()
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