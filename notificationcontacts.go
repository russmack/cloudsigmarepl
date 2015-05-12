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
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
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
	fmt.Println("Username:", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	g.channels.ResponseChan <- string(resp)
	return nil
}

func (m *CommandSetNotifyContacts) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setNotifyContactsEmail
	cargo := CommandSetNotifyContacts{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyContacts) setNotifyContactsEmail(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Email:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Email = s
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.setNotifyContactsName(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsName(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Name:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Name = s
	} else {
		// TODO: clean this.
		fmt.Println("asserton not ok")
	}
	return m.setNotifyContactsPhone(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsPhone(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Phone:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CommandSetNotifyContacts)
	if ok {
		c.Phone = s
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.setNotifyContactsValue(c)
}

func (m *CommandSetNotifyContacts) setNotifyContactsValue(cargo interface{}) statemachiner.StateFn {
	//o := cloudsigma.NewNotificationPreferences()
	//args := o.NewGet()
	//client := &cloudsigma.Client{}
	//resp, err := client.Call(args)
	//if err != nil {
	//	fmt.Println("Error calling client.", err)
	//}
	m.channels.ResponseChan <- "Not yet implemented"
	return nil
}
