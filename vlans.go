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
type CommandGetVlan struct {
	channels *replizer.Channels
}

type VlanCargo struct {
	Uuid string
	Body cloudsigma.VlanRequest
}

func NewListVlans() *CommandListVlans {
	return &CommandListVlans{}
}
func NewListVlansDetailed() *CommandListVlansDetailed {
	return &CommandListVlansDetailed{}
}
func NewGetVlan() *CommandGetVlan {
	return &CommandGetVlan{}
}

func (m *CommandListVlans) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listVlans
	cargo := VlanCargo{}
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
	cargo := VlanCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandListVlansDetailed) listVlansDetailed(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewVlans()
	args := o.NewListDetailed()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandGetVlan) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getVlanUuid
	cargo := VlanCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandGetVlan) getVlanUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(VlanCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Vlan in getVlanUuid."
		return m.getVlanUuid(c)
	}
	return m.getVlanSendRequest(c)
}

func (m *CommandGetVlan) getVlanSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewVlans()
	c, ok := cargo.(VlanCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Vlan in getVlanSendRequest."
		return nil
	}
	args := o.NewGet(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
