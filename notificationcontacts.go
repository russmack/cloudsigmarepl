package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandGetNotifyContacts struct {
	channels *replizer.Channels
}

type CommandSetNotifyContacts struct {
	channels *replizer.Channels
}

func NewGetNotifyContacts() *CommandGetNotifyContacts {
	return &CommandGetNotifyContacts{}
}

func (g *CommandGetNotifyContacts) Start(channels *replizer.Channels) {
	g.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.getNotifyContacts
	cargo := CommandGetNotifyContacts{}
	stateMachine.Start(cargo)
}

func (g *CommandGetNotifyContacts) getNotifyContacts(cargs interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
	args := o.NewGet()
	g.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		g.channels.ResponseChan <- fmt.Sprintf("Error calling client. %s", err)
		return nil
	}
	g.channels.ResponseChan <- string(resp)
	return nil
}

func (m *CommandSetNotifyContacts) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setNotifyContactsEmail
	cargo := cloudsigma.Contact{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyContacts) setNotifyContactsEmail(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Email:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Contact)
	if ok {
		c.Email = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.setNotifyContactsEmail(c)
	}
	return m.setNotifyContactsName(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsName(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Name:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Contact)
	if ok {
		c.Name = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.setNotifyContactsName(c)
	}
	return m.setNotifyContactsPhone(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsPhone(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Phone:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Contact)
	if ok {
		c.Phone = s
	} else {
		m.channels.ResponseChan <- "Error asserting Contact."
		return m.setNotifyContactsPhone(c)
	}
	return m.setNotifyContactsSendRequest(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationContacts()
	c, ok := cargo.(cloudsigma.Contact)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Contact."
		return nil
	}
	args := o.NewSet(c)
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
