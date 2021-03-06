package main

import (
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
)

type CommandGetConfigLocation struct {
	channels *replizer.Channels
}

type CommandSetConfigLocation struct {
	channels *replizer.Channels
}

func ListConfigLocation() *CommandGetConfigLocation {
	return &CommandGetConfigLocation{}
}

func EditConfigLocation() *CommandSetConfigLocation {
	return &CommandSetConfigLocation{}
}

func (g *CommandGetConfigLocation) Start(channels *replizer.Channels) {
	g.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = g.getConfigLocation
	cargo := ""
	stateMachine.Start(cargo)
}

func (g *CommandGetConfigLocation) getConfigLocation(cargo interface{}) statemachiner.StateFn {
	g.channels.ResponseChan <- session.Location
	return nil
}

func (m *CommandSetConfigLocation) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.setConfigLocation
	cargo := CommandSetConfigLocation{}
	stateMachine.Start(cargo)
}

func (m *CommandSetConfigLocation) setConfigLocation(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Enter service location:"
	s := <-m.channels.UserChan
	if s != "zrh" && s != "hnl" && s != "wdc" && s != "sjs" && s != "localhost" {
		m.channels.MessageChan <- "Unknown location.  Options: zrh | hnl | mia | sjc | wdc"
		return m.setConfigLocation(cargo)
	}
	session.Location = s
	m.channels.ResponseChan <- "{ \"location\": \"" + s + "\" }"
	return nil
}
