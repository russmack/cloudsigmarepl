package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListAcls struct {
	channels *replizer.Channels
}
type CommandGetAcl struct {
	channels *replizer.Channels
}

type AclCargo struct {
	Uuid string
	//Body cloudsigma.AclRequest
}

func NewListAcls() *CommandListAcls {
	return &CommandListAcls{}
}
func NewGetAcl() *CommandGetAcl {
	return &CommandGetAcl{}
}

func (m *CommandListAcls) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listAcls
	cargo := CommandListAcls{}
	stateMachine.Start(cargo)
}

func (m *CommandListAcls) listAcls(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewAcls()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandGetAcl) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getAclUuid
	cargo := AclCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandGetAcl) getAclUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(AclCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Acl in getAclUuid."
		return m.getAclUuid(c)
	}
	return m.getAclSendRequest(c)
}

func (m *CommandGetAcl) getAclSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewAcls()
	c, ok := cargo.(AclCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Acl in getAclSendRequest."
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
