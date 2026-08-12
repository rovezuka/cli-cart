package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	expectedFlagsCnt = 3
)

var rootCommand cobra.Command = cobra.Command{
	Use:   "cli",
	Short: "CLI tool for Cart",
	Long:  "CLI tool for Cart",
}

func main() {

	userID := pflag.Uint64("user", 0, "userID for add item to cart")
	skuID := pflag.Uint64("sku", 0, "skuID for add item to cart")
	skuCount := pflag.Uint("count", 0, "count items")

	pflag.Parse()

	if pflag.NFlag() != expectedFlagsCnt {
		log.Fatalf("unexpected number of arguments: %d, want: %d", pflag.NFlag(), expectedFlagsCnt)
	}

	Add(*userID, *skuID, *skuCount)

	fmt.Println(Read(*userID))
}

func Read(userID uint64) string {
	builder := strings.Builder{}
	for _, ci := range store[userID] {
		builder.Write([]byte(fmt.Sprintf("sku: %d - cnt: %d\n", ci.SKU, ci.QTY)))
	}
	return builder.String()
}

type CartItem struct {
	SKU uint64
	QTY uint
}

//type CartItem map[uint64]uint

var store = make(map[uint64]map[uint64]*CartItem)

func Add(userID, sku uint64, skuCount uint) {
	if cartItem, ok := store[userID][sku]; ok {
		cartItem.QTY += skuCount
	}
	store[userID] = make(map[uint64]*CartItem)
	store[userID][sku] = &CartItem{
		SKU: sku,
		QTY: skuCount,
	}
	//fmt.Printf("user: %d, sku: %d, sku_cnt: %d\n", userID, skuID, skuCount)
}
