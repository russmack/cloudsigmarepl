package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
	"os"
	"strings"
)

type CommandGetNotifyPrefs struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

type CommandSetNotifyPrefs struct {
	//Contact      string `json:"contact"`
	//Medium       string `json:"medium"`
	//Type         string `json:"type"`
	//Value        bool   `json:"value"`
	//Pref cloudsigma.Preference
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewGetNotifyPrefs() *CommandGetNotifyPrefs {
	return &CommandGetNotifyPrefs{}
}

func (g *CommandGetNotifyPrefs) Start(respChan chan string, promptChan chan string, userChan chan string) {
	g.responseChan = respChan
	g.promptChan = promptChan
	g.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.getNotifyPrefs
	//cargo := CommandGetNotifyPrefs{}
	cargo := cloudsigma.Preference{}
	stateMachine.Start(cargo)
}

func (g *CommandGetNotifyPrefs) getNotifyPrefs(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationPreferences()
	args := o.NewGet()
	fmt.Println("Username:", config.Login().Username)
	args.Username = config.Login().Username
	args.Password = config.Login().Password
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	g.responseChan <- string(resp)
	return nil
}

func (m *CommandSetNotifyPrefs) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setNotifyPrefsContact
	cargo := CommandSetNotifyPrefs{}
	stateMachine.Start(cargo)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsContact(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.promptChan <- "Contact:"
	s := <-m.userChan
	//c, ok := cargo.(CommandSetNotifyPrefs)
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Contact = s
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyPrefsMedium(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsMedium(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Medium:"
	s := <-m.userChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Medium = s
	} else {
		fmt.Println("asserton not ok")
	}
	return m.setNotifyPrefsType(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsType(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Type:"
	s := <-m.userChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		c.Type = s
	} else {
		fmt.Println("assertion not ok")
	}
	return m.setNotifyPrefsValue(c)
}

func (m *CommandSetNotifyPrefs) setNotifyPrefsValue(cargo interface{}) statemachiner.StateFn {
	m.promptChan <- "Value (true|false):"
	s := <-m.userChan
	c, ok := cargo.(cloudsigma.Preference)
	if ok {
		val := false
		switch strings.ToLower(s) {
		case "true":
			val = true
		case "false":
			val = false
		default:
			fmt.Println("Invalid input.")
			os.Exit(1)
		}
		c.Value = val
	} else {
		fmt.Println("assertion not ok")
		os.Exit(1)
	}
	return m.setNotifyPrefsSendRequest(c)
}
func (m *CommandSetNotifyPrefs) setNotifyPrefsSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewNotificationPreferences()
	c, ok := cargo.(cloudsigma.Preference)
	if !ok {
		fmt.Println("assertion not ok")
		os.Exit(1)
	}
	args := o.NewSet(c)
	// TODO: this needs to be entered by repl user.
	args.Location = "zrh"
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	m.responseChan <- string(resp)
	return nil
}
