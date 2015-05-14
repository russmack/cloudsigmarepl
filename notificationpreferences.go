package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
	"strings"
)

type CommandGetNotifyPrefs struct {
	channels *replizer.Channels
}

type CommandEditNotifyPrefs struct {
	channels *replizer.Channels
}

func NewGetNotifyPrefs() *CommandGetNotifyPrefs {
	return &CommandGetNotifyPrefs{}
}

func NewEditNotifyPrefs() *CommandEditNotifyPrefs {
	return &CommandEditNotifyPrefs{}
}

func (g *CommandGetNotifyPrefs) Start(channels *replizer.Channels) {
	g.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.getNotifyPrefs
	cargo := cloudsigma.Preference{}
	stateMachine.Start(cargo)
}

func (g *CommandGetNotifyPrefs) getNotifyPrefs(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationPreferences()
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

func (m *CommandEditNotifyPrefs) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.editNotifyPrefsContact
	cargo := cloudsigma.Preference{}
	stateMachine.Start(cargo)
}

func (m *CommandEditNotifyPrefs) editNotifyPrefsContact(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Contact:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Contact = s
	} else {
		m.channels.ResponseChan <- "Error asserting Preference."
		return nil
	}
	return m.editNotifyPrefsMedium(c)
}

func (m *CommandEditNotifyPrefs) editNotifyPrefsMedium(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Medium:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Medium = s
	} else {
		m.channels.ResponseChan <- "Error asserting Preference."
		return nil
	}
	return m.editNotifyPrefsType(c)
}

func (m *CommandEditNotifyPrefs) editNotifyPrefsType(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Type:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Type = s
	} else {
		m.channels.ResponseChan <- "Error asserting Preference."
		// TODO: failed assertion should probably return the startFn.
		return nil
	}
	return m.editNotifyPrefsValue(c)
}

func (m *CommandEditNotifyPrefs) editNotifyPrefsValue(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Value (true|false):"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		val := false
		switch strings.ToLower(s) {
		case "true":
			val = true
		case "false":
			val = false
		default:
			m.channels.ResponseChan <- "Invalid input."
			return m.editNotifyPrefsValue(c)
		}
		c.Value = val
	} else {
		m.channels.ResponseChan <- "Error asserting Preference."
		return nil
	}
	return m.editNotifyPrefsSendRequest(c)
}

func (m *CommandEditNotifyPrefs) editNotifyPrefsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationPreferences()
	c, ok := cargo.(cloudsigma.Preference)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Preference."
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
