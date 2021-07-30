package main

import (
	"context"

	"github.com/handball811/firestore-testing/cmd/common"
	"github.com/handball811/firestore-testing/internal/debug"
)

func main() {
	client := common.GetClient()
	// とりあえずリフレッシュ
	debug.DeleteAll(client)

	// データの追加
	ctx := context.Background()
	userCol := client.Collection("user")
	userCol.Doc("hello0").Set(ctx, map[string]interface{}{
		"name": "sasakuna",
		"age":  6,
		"tags": []string{"go", "firestore"},
		"wow": map[string]interface{}{
			"hello": map[string]interface{}{
				"world": 123,
			},
		},
		"docref": userCol.Doc("hello0"),
	})

	debug.ShowAll(client)
}

/*
Result
*  user/
*      hello0/
*        age: (int64 -> 6)
*        tags: ([]interface {} -> [go firestore])
*        docref: (*firestore.DocumentRef -> &{0xc000696000 projects/testing/databases/(default)/documents/user/hello0 user/hello0 hello0})
*        name: (string -> sasakuna)
*        wow: (map[string]interface {} -> map[hello:map[world:123]])
(end)
*/
