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
			{
				Name:   "Data To City",
				Action: DataToCity,
			},
			{
				Name:   "Data To PC",
				Action: DataToPC,
			},
		}...,
	).Run()
}

/*

Do0



Name:  City Full
*  users/
*      userid0/
*        tags: ([]interface {} -> [hello world haha])
*        population: (int64 -> 56)
*        createdAt: (time.Time -> 2021-07-30 11:57:13.986231 +0000 UTC)
*        name: (string -> sasakuna)
*        updatedAt: (time.Time -> 2021-07-30 11:57:13.986231 +0000 UTC)
(end)

Name:  Data To City
City:  &{sasakuna [hello world haha] 56}

Name:  Data To PC
City:   interface{}(
+       &structs.PC{Name: &"sasakuna", Tags: []string{"hello", "world", "haha"}, Pop: &56},
  )

*/

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
			{
				Name:   "Data To City",
				Action: DataToCity,
			},
			{
				Name:   "Data To PC",
				Action: DataToPC,
			},
		}...,
	).Run()
}

/*

Do1



Name:  City Partial
*  users/
*      userid0/
*        population: (int64 -> 0)
*        tags: (Nil -> <nil>)
*        createdAt: (time.Time -> 2021-07-30 11:57:14.002157 +0000 UTC)
*        name: (string -> sasakuna)
*        updatedAt: (time.Time -> 2021-07-30 11:57:14.002157 +0000 UTC)
(end)

Name:  Data To City
City:  &{sasakuna [] 0}

Name:  Data To PC
City:   interface{}(
+       &structs.PC{Name: &"sasakuna", Pop: &0},
  )

*/

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
			{
				Name:   "Data To City",
				Action: DataToCity,
			},
			{
				Name:   "Data To PC",
				Action: DataToPC,
			},
		}...,
	).Run()
}

/*

Do2



Name:  PC Full
*  users/
*      userid0/
*        name: (string -> sasakuna)
*        updatedAt: (time.Time -> 2021-07-30 11:57:14.016314 +0000 UTC)
*        tags: ([]interface {} -> [hello world haha])
*        population: (int64 -> 56)
*        createdAt: (time.Time -> 2021-07-30 11:57:14.016314 +0000 UTC)
(end)

Name:  Data To City
City:  &{sasakuna [hello world haha] 56}

Name:  Data To PC
City:   interface{}(
+       &structs.PC{Name: &"sasakuna", Tags: []string{"hello", "world", "haha"}, Pop: &56},
  )

*/

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
			{
				Name:   "Data To City",
				Action: DataToCity,
			},
			{
				Name:   "Data To PC",
				Action: DataToPC,
			},
		}...,
	).Run()
}

/*

Do3



Name:  PC Partial
*  users/
*      userid0/
*        createdAt: (time.Time -> 2021-07-30 11:57:14.030879 +0000 UTC)
*        name: (string -> sasakuna)
*        updatedAt: (time.Time -> 2021-07-30 11:57:14.030879 +0000 UTC)
*        tags: (Nil -> <nil>)
*        population: (Nil -> <nil>)
(end)

Name:  Data To City
City:  &{sasakuna [] 0}

Name:  Data To PC
City:   interface{}(
+       &structs.PC{Name: &"sasakuna"},
  )

*/
