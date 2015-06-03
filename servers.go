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

type CommandCreateServer struct {
	channels *replizer.Channels
}

type CreateServerCargo struct {
	Uuid string
	Body cloudsigma.ServerRequest
}

func NewListServers() *CommandListServers {
	return &CommandListServers{}
}

func NewCreateServer() *CommandCreateServer {
	return &CommandCreateServer{}
}

func (m *CommandListServers) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.listServers
	cargo := CommandListServers{}
	stateMachine.Start(cargo)
}

func (g *CommandListServers) listServers(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewServers()
	args := o.NewList()
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

func (m *CommandCreateServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createServerName
	cargo := CreateServerCargo{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateServer) createServerName(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Name:"
	n := <-m.channels.UserChan
	c, ok := cargo.(CreateServerCargo)
	if ok {
		c.Body.Name = n
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.createServerCpu(c)
}

func (m *CommandCreateServer) createServerCpu(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "CPU:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CreateServerCargo)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			// TODO: clean this.
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Body.Cpu = n
		}
	} else {
		// TODO: clean this.
		fmt.Println("asserton not ok")
	}
	return m.createServerMemory(c)
}

func (m *CommandCreateServer) createServerMemory(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "Memory:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CreateServerCargo)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			// TODO: clean this.
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Body.Memory = n
		}
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.createServerVncPassword(c)
}

func (m *CommandCreateServer) createServerVncPassword(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "VNC password:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CreateServerCargo)
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
	c, ok := cargo.(CreateServerCargo)
	if !ok {
		m.channels.ResponseChan <- "Error asswerting Server."
		return nil
	}
	newServers := []cloudsigma.ServerRequest{
		cloudsigma.ServerRequest{c.Body.Name, c.Body.Cpu, c.Body.Memory, c.Body.VncPassword},
	}
	args := o.NewCreate(newServers)
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
