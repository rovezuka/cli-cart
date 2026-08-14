package map_store

import (
	"CliCart/internal/domain/cart"
	"fmt"
	"strings"
)

type Storage struct {
	store map[cart.UserID]map[cart.SKU]*cart.CartItem
}

func NewMapStorage() *Storage {
	return &Storage{
		store: make(map[cart.UserID]map[cart.SKU]*cart.CartItem),
	}
}

func (s *Storage) Read(userID cart.UserID) string {
	builder := strings.Builder{}
	for _, ci := range s.store[userID] {
		builder.Write([]byte(fmt.Sprintf("sku: %d - cnt: %d\n", ci.SKU, ci.QTY)))
	}
	return builder.String()
}

func (s *Storage) Add(userID cart.UserID, sku cart.SKU, skuCount uint) {

	if cartItem, ok := s.store[userID][sku]; ok {
		cartItem.QTY += skuCount
	}
	s.store[userID] = make(map[cart.SKU]*cart.CartItem)
	s.store[userID][sku] = &cart.CartItem{
		SKU: sku,
		QTY: skuCount,
	}
	//fmt.Printf("user: %d, sku: %d, sku_cnt: %d\n", userID, skuID, skuCount)
}
