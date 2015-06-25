package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListIps struct {
	channels *replizer.Channels
}
type CommandListIpsDetailed struct {
	channels *replizer.Channels
}
type CommandGetIp struct {
	channels *replizer.Channels
}

type IpsCargo struct {
	Uuid string
	Body cloudsigma.IpRequest
}

func NewListIps() *CommandListIps {
	return &CommandListIps{}
}
func NewListIpsDetailed() *CommandListIpsDetailed {
	return &CommandListIpsDetailed{}
}
func NewGetIp() *CommandGetIp {
	return &CommandGetIp{}
}

func (m *CommandListIps) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listIps
	cargo := CommandListIps{}
	stateMachine.Start(cargo)
}

func (m *CommandListIps) listIps(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewIps()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandListIpsDetailed) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listIpsDetailed
	cargo := IpsCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandListIpsDetailed) listIpsDetailed(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewIps()
	args := o.NewListDetailed()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandGetIp) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getIpUuid
	cargo := IpsCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandGetIp) getIpUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(IpsCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Ip in getIpUuid."
		return m.getIpUuid(c)
	}
	return m.getIpSendRequest(c)
}

func (m *CommandGetIp) getIpSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewIps()
	c, ok := cargo.(IpsCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Ip in getIpSendRequest."
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
