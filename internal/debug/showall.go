package debug

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/handball811/firestore-testing/internal/debug/format"
)

func ShowAll(c *firestore.Client) {
	ctx := context.Background()
	root := c.Collections(ctx)
	f := format.ColIter{
		Target: root,
		Prefix: "*",
	}
	for {
		s, err := f.Next()
		if err != nil {
			fmt.Println("(end)")
			return
		}
		fmt.Println(s)
	}
}

/*
イメージ


users/
	userid0/
		name: (string -> Sasakuna)
		age: (int64 -> 25)
		evaluation/
			shopid0/
				comment: (string -> great)
			shopid1/
shops/
	shopid0/
		name: (string -> SasaShop)
		tags: ([]interface{} -> )
			0: (string -> hello)
			1: (map[string]interface{} -> )
				123: (map[string]interface{} -> )
					456: (string -> good)
					what: ([]interface{} -> )
						0: (int64 -> 56)
			2: (string -> go)
*/
