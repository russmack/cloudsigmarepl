package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListTags struct {
	channels *replizer.Channels
}
type CommandGetTag struct {
	channels *replizer.Channels
}

type TagCargo struct {
	Uuid string
	//Body cloudsigma.TagRequest
}

func NewListTags() *CommandListTags {
	return &CommandListTags{}
}
func NewGetTag() *CommandGetTag {
	return &CommandGetTag{}
}

func (m *CommandListTags) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listTags
	cargo := CommandListTags{}
	stateMachine.Start(cargo)
}

func (m *CommandListTags) listTags(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewTags()
	args := o.NewList()
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

func (m *CommandGetTag) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getTagUuid
	cargo := TagCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandGetTag) getTagUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(TagCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Tag in getTagUuid."
		return m.getTagUuid(c)
	}
	return m.getTagSendRequest(c)
}

func (m *CommandGetTag) getTagSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewTags()
	c, ok := cargo.(TagCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Tag in getTagSendRequest."
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
