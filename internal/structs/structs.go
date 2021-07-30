package structs

import "time"

type City struct {
	Name string   `firestore:"name"`
	Tags []string `firestore:"tags"`
	Pop  int64    `firestore:"population"`
}

type PC struct {
	Name *string  `firestore:"name"`
	Tags []string `firestore:"tags"`
	Pop  *int64   `firestore:"population"`
}

type Init struct {
	CreatedAt time.Time `firestore:"createdAt"`
	UpdatedAt time.Time `firestore:"updatedAt"`
}

func SetInit(dst *Init) {
	dst.CreatedAt = time.Now()
	dst.UpdatedAt = time.Now()
}

type Update struct {
	UpdatedAt time.Time `firestore:"updatedAt"`
}

func SetUpdate(dst *Update) {
	dst.UpdatedAt = time.Now()
}

type InitCity struct {
	Init
	City
}

type UpdateCity struct {
	Update
	City
}

type InitPC struct {
	Init
	PC
}

type UpdatePC struct {
	Update
	PC
}
