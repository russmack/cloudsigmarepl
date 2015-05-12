package main

import (
	"fmt"
	"github.com/russmack/replizer"
	"github.com/russmack/statemachiner"
	"strconv"
)

type CommandCreateServer struct {
	Name        string `json:"name"`
	Cpu         int    `json:"cpu"`
	Memory      int    `json:"mem"`
	VncPassword string `json:"vnc_password"`
	channels    *replizer.Channels
}

func NewCreateServer() *CommandCreateServer {
	return &CommandCreateServer{}
}

func (m *CommandCreateServer) Start(channels *replizer.Channels) {
	m.channels = channels
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.createServerName
	cargo := CommandCreateServer{}
	stateMachine.Start(cargo)
}

func (m *CommandCreateServer) createServerName(cargo interface{}) statemachiner.StateFn {
	// The state machine will not progress beyond this point until the repl
	// pops from the promptChan.
	m.channels.PromptChan <- "Name:"
	n := <-m.channels.UserChan
	c, ok := cargo.(CommandCreateServer)
	if ok {
		c.Name = n
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.createServerCpu(c)
}

func (m *CommandCreateServer) createServerCpu(cargo interface{}) statemachiner.StateFn {
	m.channels.PromptChan <- "CPU:"
	s := <-m.channels.UserChan
	c, ok := cargo.(CommandCreateServer)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			// TODO: clean this.
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Cpu = n
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
	c, ok := cargo.(CommandCreateServer)
	if ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			// TODO: clean this.
			fmt.Println("this should be a request to re-enter info")
		} else {
			c.Memory = n
		}
	} else {
		// TODO: clean this.
		fmt.Println("assertion not ok")
	}
	return m.createServerVncPassword(c)
}

func (m *CommandCreateServer) createServerVncPassword(cargo interface{}) statemachiner.StateFn {
	//o := cloudsigma.NewCreateServer()
	//args := o.NewGet()
	//client := &cloudsigma.Client{}
	//resp, err := client.Call(args)
	//if err != nil {
	//	fmt.Println("Error calling client.", err)
	//}
	m.channels.ResponseChan <- "Not yet implemented"
	return nil
}
