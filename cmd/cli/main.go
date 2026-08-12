package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	expectedFlagsCnt = 3
)

func main() {

	userID := flag.Uint64("user", 0, "userID for add item to cart")
	skuID := flag.Uint64("sku", 0, "skuID for add item to cart")
	skuCount := flag.Uint64("count", 0, "count items")

	flag.Parse()

	if flag.NFlag() != expectedFlagsCnt {
		log.Fatalf("unexpected number of arguments: %d, want: %d", flag.NFlag(), expectedFlagsCnt)
	}

	fmt.Printf("user: %d, sku: %d, sku_cnt: %d\n", *userID, *skuID, *skuCount)
}
