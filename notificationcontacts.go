package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandListNotifyContacts struct {
	channels *replizer.Channels
}

type CommandCreateNotifyContacts struct {
	channels *replizer.Channels
}

type CommandEditNotifyContacts struct {
	channels *replizer.Channels
}

type CommandDeleteNotifyContacts struct {
	channels *replizer.Channels
}

type ContactCargo struct {
	Uuid string
	Body cloudsigma.ContactRequest
}

func NewListNotifyContacts() *CommandListNotifyContacts {
	return &CommandListNotifyContacts{}
}

func NewCreateNotifyContacts() *CommandCreateNotifyContacts {
	return &CommandCreateNotifyContacts{}
}

func NewEditNotifyContacts() *CommandEditNotifyContacts {
	return &CommandEditNotifyContacts{}
}

func NewDeleteNotifyContacts() *CommandDeleteNotifyContacts {
	return &CommandDeleteNotifyContacts{}
}

func (g *CommandListNotifyContacts) Start(channels *replizer.Channels) {
	g.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.listNotifyContacts
	cargo := CommandListNotifyContacts{}
	stateMachine.Start(cargo)
}

func (g *CommandListNotifyContacts) listNotifyContacts(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
	args := o.NewList()
	g.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(g.channels, args)
	return nil
}

func (m *CommandCreateNotifyContacts) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createNotifyContactsEmail
	cargo := ContactCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandEditNotifyContacts) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.editNotifyContactsUuid
	cargo := ContactCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandDeleteNotifyContacts) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.deleteNotifyContactsUuid
	cargo := ContactCargo{}
	stateMachine.Start(cargo)
}

// Create contact state functions.

func (m *CommandCreateNotifyContacts) createNotifyContactsEmail(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Email:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Body.Email = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.createNotifyContactsEmail(c)
	}
	return m.createNotifyContactsName(c)
}

func (m *CommandCreateNotifyContacts) createNotifyContactsName(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Name:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Body.Name = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.createNotifyContactsName(c)
	}
	return m.createNotifyContactsPhone(c)
}

func (m *CommandCreateNotifyContacts) createNotifyContactsPhone(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Phone (must begin with +):"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Body.Phone = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.createNotifyContactsPhone(c)
	}
	return m.createNotifyContactsSendRequest(c)
}

func (m *CommandCreateNotifyContacts) createNotifyContactsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
	c, ok := cargo.(ContactCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Contact."
		return nil
	}
	contacts := []cloudsigma.ContactRequest{c.Body}
	args := o.NewCreate(contacts)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	err := sendRequest(m.channels, args)
	if err != nil {
		m.channels.MessageChan <- "Ensure phone begins with +"
	}
	return nil
}

// Edit contact state functions.

func (m *CommandEditNotifyContacts) editNotifyContactsUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Contact uuid:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Uuid = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.editNotifyContactsUuid(c)
	}
	return m.editNotifyContactsEmail(c)
}

func (m *CommandEditNotifyContacts) editNotifyContactsEmail(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Email:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Body.Email = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.editNotifyContactsEmail(c)
	}
	return m.editNotifyContactsName(c)
}

func (m *CommandEditNotifyContacts) editNotifyContactsName(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Name:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Body.Name = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.editNotifyContactsName(c)
	}
	return m.editNotifyContactsPhone(c)
}

func (m *CommandEditNotifyContacts) editNotifyContactsPhone(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Phone:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Body.Phone = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.editNotifyContactsPhone(c)
	}
	return m.editNotifyContactsSendRequest(c)
}

func (m *CommandEditNotifyContacts) editNotifyContactsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
	c, ok := cargo.(ContactCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Contact."
		return nil
	}
	args := o.NewEdit(c.Uuid, c.Body)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
	return nil
}

// Delete contact state functions.

func (m *CommandDeleteNotifyContacts) deleteNotifyContactsUuid(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Contact uuid:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ContactCargo)
	if ok {
		c.Uuid = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.deleteNotifyContactsUuid(c)
	}
	return m.deleteNotifyContactsSendRequest(c)
}

func (m *CommandDeleteNotifyContacts) deleteNotifyContactsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
	c, ok := cargo.(ContactCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Contact."
		return nil
	}
	args := o.NewDelete(c.Uuid)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
	return nil
}
