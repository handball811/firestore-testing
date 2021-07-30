package common

import (
	"context"

	firestore "cloud.google.com/go/firestore"
)

func GetClient() *firestore.Client {
	if err := LoadEnv(); err != nil {
		panic(err)
	}
	ctx := context.Background()
	store, err := firestore.NewClient(ctx, "testing")
	if err != nil {
		panic(err)
	}
	return store
}
