package debug

import (
	"fmt"

	firestore "cloud.google.com/go/firestore"
)

type ActionSet struct {
	Name   string
	Action func(*firestore.Client)
}

type ActionPipline struct {
	name       string
	client     *firestore.Client
	actionSets []ActionSet
}

func NewActionPipeline(
	name string,
	client *firestore.Client,
	sets ...ActionSet,
) *ActionPipline {
	return &ActionPipline{
		name:       name,
		client:     client,
		actionSets: sets,
	}
}

func (p *ActionPipline) Run() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(p.name)
	fmt.Println("")
	fmt.Println("")
	for _, set := range p.actionSets {
		set.Action(p.client)
		fmt.Println("Name: ", set.Name)
		ShowAll(p.client)
	}
}
