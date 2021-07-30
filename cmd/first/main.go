package main

import (
	"context"

	"github.com/handball811/firestore-testing/cmd/common"
	"github.com/handball811/firestore-testing/internal/debug"
)

func main() {
	client := common.GetClient()

	ctx := context.Background()
	userCol := client.Collection("user")
	if _, err := userCol.Doc("hello0").Set(ctx, map[string]interface{}{
		"name": "sasakuna",
		"age":  6,
		"tags": []string{"go", "firestore"},
	}); err != nil {
		panic(err)
	}

	debug.ShowAll(client)
}
