package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
	"strconv"
)

type CommandListServers struct {
	channels *replizer.Channels
}
type CommandStartServer struct {
	channels *replizer.Channels
}
type CommandStopServer struct {
	channels *replizer.Channels
}
type CommandCreateServer struct {
	channels *replizer.Channels
}

type ServerCargo struct {
	Uuid string
	Body cloudsigma.ServerRequest
}

func NewListServers() *CommandListServers {
	return &CommandListServers{}
}
func NewStartServer() *CommandStartServer {
	return &CommandStartServer{}
}
func NewStopServer() *CommandStopServer {
	return &CommandStopServer{}
}
func NewCreateServer() *CommandCreateServer {
	return &CommandCreateServer{}
}

// Start is the start state of the CommandListServers state machine.
func (m *CommandListServers) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listServers
	cargo := ServerCargo{}
	stateMachine.Start(cargo)
}

func (g *CommandListServers) listServers(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	args := o.NewList()
	g.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(g.channels, args)
	return nil
}

// Start is the start state of the CommandStartServer state machine.
func (m *CommandStartServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.startServerUuid
	cargo := ServerCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandStartServer) startServerUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(ServerCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Server."
		return m.startServerUuid(c)
	}
	return m.startServerSendRequest(c)
}

func (m *CommandStartServer) startServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(ServerCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Server."
		return nil
	}

	args := o.NewStart(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

// Start is the start state of the CommandStopServer state machine.
func (m *CommandStopServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.stopServerUuid
	cargo := ServerCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandStopServer) stopServerUuid(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Uuid:"
	n := <-m.channels.UserChan
	c, ok := cargo.(ServerCargo)
	if ok {
		c.Uuid = n
	} else {
		m.channels.ResponseChan <- "Error asserting Server."
		return m.stopServerUuid(c)
	}
	return m.stopServerSendRequest(c)
}

func (m *CommandStopServer) stopServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(ServerCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Server."
		return nil
	}

	args := o.NewStop(c.Uuid)
	m.channels.MessageChan <- fmt.Sprintf("Using username: %s", session.Username)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}

// Start is the start state of the CommandCreateServer state machine.
func (m *CommandCreateServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createServerName
	cargo := ServerCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateServer) createServerName(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Name:"
	n := <-m.channels.UserChan
	c, ok := cargo.(ServerCargo)
	if ok {
		c.Body.Name = n
	} else {
		m.channels.ResponseChan <- "Error asserting Server."
		return m.createServerName(c)
	}
	return m.createServerCpu(c)
}

func (m *CommandCreateServer) createServerCpu(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "CPU:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ServerCargo)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			// TODO: clean this.
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Body.Cpu = n
		}
	} else {
		m.channels.ResponseChan <- "Error asserting Server."
		return m.createServerCpu(c)
	}
	return m.createServerMemory(c)
}

func (m *CommandCreateServer) createServerMemory(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Memory:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ServerCargo)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			// TODO: clean this.
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Body.Memory = n
		}
	} else {
		m.channels.ResponseChan <- "Error asserting Server."
		return m.createServerMemory(c)
	}
	return m.createServerVncPassword(c)
}

func (m *CommandCreateServer) createServerVncPassword(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "VNC password:"
	s := <-m.channels.UserChan
	c, ok := cargo.(ServerCargo)
	if ok {
		c.Body.VncPassword = s
	} else {
		m.channels.ResponseChan <- "Error asserting Server."
		return m.createServerVncPassword(c)
	}
	return m.createServerSendRequest(c)
}

func (m *CommandCreateServer) createServerSendRequest(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	c, ok := cargo.(ServerCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asserting Server."
		return nil
	}
	newServers := []cloudsigma.ServerRequest{
		cloudsigma.ServerRequest{c.Body.Name, c.Body.Cpu, c.Body.Memory, c.Body.VncPassword},
	}
	args := o.NewCreate(newServers)
	args.Username = session.Username
	args.Password = session.Password
	args.Location = session.Location
	_ = sendRequest(m.channels, args)
	return nil
}
