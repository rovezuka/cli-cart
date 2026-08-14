package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/multierr"
)

const (
	expectedFlagsCnt = 3
)

var rootCommand cobra.Command = cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCommand.AddCommand(
		&cobra.Command{
			Use:     "add",
			Short:   "a",
			Example: "",
			Args:    cobra.ExactArgs(expectedFlagsCnt),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) != expectedFlagsCnt {
					return errors.New("wrong number of arguments")
				}

				uid, err1 := strconv.Atoi(args[0])
				sku, err2 := strconv.Atoi(args[1])
				sc, err3 := strconv.Atoi(args[2])

				if err1 != nil || err2 != nil || err3 != nil {
					return multierr.Combine(err1, err2, err3)
				}

				Add(uint64(uid), uint64(sku), uint(sc))

				return nil
			},
		},
		&cobra.Command{
			Use:     "view",
			Short:   "v",
			Example: "",
			Args:    cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				uid, err1 := strconv.Atoi(args[0])

				if err1 != nil {
					return fmt.Errorf("cannot prase uid: %s %w", args[0], err1)
				}

				fmt.Println(Read(uint64(uid)))

				return nil
			},
		},
	)
}

func main() {

	//userID := pflag.Uint64("user", 0, "userID for add item to cart")
	//skuID := pflag.Uint64("sku", 0, "skuID for add item to cart")
	//skuCount := pflag.Uint("count", 0, "count items")
	//
	//pflag.Parse()
	//
	//if pflag.NFlag() != expectedFlagsCnt {
	//	log.Fatalf("unexpected number of arguments: %d, want: %d", pflag.NFlag(), expectedFlagsCnt)
	//}
	//
	//Add(*userID, *skuID, *skuCount)
	//
	//fmt.Println(Read(*userID))

	fmt.Println("Welcome")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("cart>")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		rootCommand.SetArgs(parts)

		if err := rootCommand.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, "Command failed, Input Error: %w", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}
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
