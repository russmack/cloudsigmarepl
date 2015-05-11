package main

import (
	"fmt"
	"github.com/russmack/cloudsigma"
	"github.com/russmack/statemachiner"
)

type CommandLocations struct {
	responseChan chan string
	promptChan   chan string
	userChan     chan string
}

func NewLocations() *CommandLocations {
	return &CommandLocations{}
}

func (m *CommandLocations) Start(respChan chan string, promptChan chan string, userChan chan string) {
	m.responseChan = respChan
	m.promptChan = promptChan
	m.userChan = userChan
	stateMachine := &statemachiner.StateMachine{}
	stateMachine.StartState = m.getLocations
	cargo := CommandLocations{}
	stateMachine.Start(cargo)
}

func (m *CommandLocations) getLocations(cargo interface{}) statemachiner.StateFn {
	o := cloudsigma.NewLocations()
	args := o.NewGet()
	client := &cloudsigma.Client{}
	resp, err := client.Call(args)
	if err != nil {
		fmt.Println("Error calling client.", err)
	}
	m.responseChan <- string(resp)
	return nil
}
