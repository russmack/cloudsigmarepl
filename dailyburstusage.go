package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandDailyBurstUsage struct {
	channels *replizer.Channels
}

func NewDailyBurstUsage() *CommandDailyBurstUsage {
	return &CommandDailyBurstUsage{}
}

func (m *CommandDailyBurstUsage) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getDailyBurstUsage
	cargo := CommandDailyBurstUsage{}
	stateMachine.Start(cargo)
}

func (m *CommandDailyBurstUsage) getDailyBurstUsage(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewDailyBurstUsage()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
