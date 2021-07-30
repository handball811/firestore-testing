package structs

import (
	"context"

	"cloud.google.com/go/firestore"
)

func InitCityFull(c *firestore.Client) {
	init := &InitCity{
		City: City{
			Name: "sasakuna",
			Tags: []string{
				"hello",
				"world",
				"haha",
			},
			Pop: 56,
		}}
	SetInit(&init.Init)
	c.Doc("users/userid0").Create(context.Background(), init)
}

func InitCityPartial(c *firestore.Client) {
	init := &InitCity{
		City: City{
			Name: "sasakuna",
		}}
	SetInit(&init.Init)
	c.Doc("users/userid0").Create(context.Background(), init)
}

func InitPCFull(c *firestore.Client) {
	name := "sasakuna"
	pop := int64(56)

	init := &InitPC{
		PC: PC{
			Name: &name,
			Tags: []string{
				"hello",
				"world",
				"haha",
			},
			Pop: &pop,
		}}
	SetInit(&init.Init)
	c.Doc("users/userid0").Create(context.Background(), init)
}

func InitPCPartial(c *firestore.Client) {
	name := "sasakuna"

	init := &InitPC{
		PC: PC{
			Name: &name,
		}}
	SetInit(&init.Init)
	c.Doc("users/userid0").Create(context.Background(), init)
}
