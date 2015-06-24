package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListVlans struct {
	channels *replizer.Channels
}
type CommandListVlansDetailed struct {
	channels *replizer.Channels
}

func NewListVlans() *CommandListVlans {
	return &CommandListVlans{}
}
func NewListVlansDetailed() *CommandListVlansDetailed {
	return &CommandListVlansDetailed{}
}

func (m *CommandListVlans) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listVlans
	cargo := CommandListVlans{}
	stateMachine.Start(cargo)
}

func (m *CommandListVlans) listVlans(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewVlans()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandListVlansDetailed) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listVlansDetailed
	cargo := CommandListVlansDetailed{}
	stateMachine.Start(cargo)
}

func (m *CommandListVlansDetailed) listVlansDetailed(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewVlansDetailed()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
