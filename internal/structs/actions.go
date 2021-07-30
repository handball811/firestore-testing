package structs

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
)

func InitCityFull(c *firestore.Client) bool {
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
	return true
}

func InitCityPartial(c *firestore.Client) bool {
	init := &InitCity{
		City: City{
			Name: "sasakuna",
		}}
	SetInit(&init.Init)
	c.Doc("users/userid0").Create(context.Background(), init)
	return true
}

func InitPCFull(c *firestore.Client) bool {
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
	return true
}

func InitPCPartial(c *firestore.Client) bool {
	name := "sasakuna"

	init := &InitPC{
		PC: PC{
			Name: &name,
		}}
	SetInit(&init.Init)
	c.Doc("users/userid0").Create(context.Background(), init)
	return true
}

func DataToCity(c *firestore.Client) bool {
	city := new(City)
	snap, _ := c.Doc("users/userid0").Get(context.Background())
	snap.DataTo(city)
	fmt.Println("City: ", city)
	return false
}

func DataToPC(c *firestore.Client) bool {
	city := new(PC)
	snap, _ := c.Doc("users/userid0").Get(context.Background())
	snap.DataTo(city)
	fmt.Printf("City: %s\n", cmp.Diff(nil, city))
	return false
}
