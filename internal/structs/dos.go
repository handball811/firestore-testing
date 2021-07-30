package structs

import (
	"cloud.google.com/go/firestore"
	"github.com/handball811/firestore-testing/internal/debug"
)

func Do0(c *firestore.Client) {
	// とりあえずリフレッシュ
	debug.DeleteAll(c)
	debug.NewActionPipeline(
		"Do0",
		c,
		[]debug.ActionSet{
			{
				Name:   "City Full",
				Action: InitCityFull,
			},
		}...,
	).Run()
}

func Do1(c *firestore.Client) {
	// とりあえずリフレッシュ
	debug.DeleteAll(c)
	debug.NewActionPipeline(
		"Do1",
		c,
		[]debug.ActionSet{
			{
				Name:   "City Partial",
				Action: InitCityPartial,
			},
		}...,
	).Run()
}

func Do2(c *firestore.Client) {
	// とりあえずリフレッシュ
	debug.DeleteAll(c)
	debug.NewActionPipeline(
		"Do2",
		c,
		[]debug.ActionSet{
			{
				Name:   "PC Full",
				Action: InitPCFull,
			},
		}...,
	).Run()
}

func Do3(c *firestore.Client) {
	// とりあえずリフレッシュ
	debug.DeleteAll(c)
	debug.NewActionPipeline(
		"Do3",
		c,
		[]debug.ActionSet{
			{
				Name:   "PC Partial",
				Action: InitPCPartial,
			},
		}...,
	).Run()
}
