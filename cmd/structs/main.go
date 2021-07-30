package main

import (
	"github.com/handball811/firestore-testing/cmd/common"
	. "github.com/handball811/firestore-testing/internal/structs"
)

// 構造体を利用した
func main() {
	client := common.GetClient()
	Do0(client)
	Do1(client)
	Do2(client)
	Do3(client)
}
