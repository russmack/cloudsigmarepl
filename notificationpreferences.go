package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
	"os"
	"strings"
)

type CommandGetNotifyPrefs struct {
	channels *replizer.Channels
}

type CommandSetNotifyPrefs struct {
	channels *replizer.Channels
}

func NewGetNotifyPrefs() *CommandGetNotifyPrefs {
	return &CommandGetNotifyPrefs{}
}

func NewSetNotifyPrefs() *CommandSetNotifyPrefs {
	return &CommandSetNotifyPrefs{}
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
	fmt.Println("Username:", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	g.channels.ResponseChan <- string(resp)
	return nil
}

func (m *CommandSetNotifyPrefs) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setNotifyPrefsContact
	cargo := cloudsigma.Preference{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsContact(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Contact:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Contact = s
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyPrefsMedium(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsMedium(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Medium:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Medium = s
	} else {
		fmt.Println("asserton not ok")
	}
	return m.setNotifyPrefsType(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsType(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Type:"
	s := <-m.channels.UserChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Type = s
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.setNotifyPrefsValue(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsValue(cargo interface{}) statemachiner.StateFn {
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
			// TODO: clean this.
			fmt.Println("Invalid input.")
			os.Exit(1)
		}
		c.Value = val
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
		os.Exit(1)
	}
	return m.setNotifyPrefsSendRequest(c)
}
func (m *CommandSetNotifyPrefs) setNotifyPrefsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationPreferences()
	c, ok := cargo.(cloudsigma.Preference)
	if !ok {
		// TODO: clean this.
		fmt.Println("assertion not ok")
		os.Exit(1)
	}
	args := o.NewSet(c)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		// TODO: clean this.
		fmt.Println("Error calling client.", err)
	}
	m.channels.ResponseChan <- string(resp)
	return nil
}
