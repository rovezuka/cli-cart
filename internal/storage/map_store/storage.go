package map_store

import (
	"CliCart/internal/domain/cart"
	storage_errors "CliCart/internal/storage/errors"
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

func (s *Storage) Read(userID uint64) (string, error) {
	return s.read(cart.UserID(userID))
}

func (s *Storage) read(userID cart.UserID) (string, error) {
	builder := strings.Builder{}
	if _, ok := s.store[userID]; !ok {
		return "", fmt.Errorf("user %d: %w", userID, storage_errors.ErrHasNoCart)
	}

	for _, ci := range s.store[userID] {
		builder.Write([]byte(fmt.Sprintf("sku: %d - cnt: %d\n", ci.SKU, ci.QTY)))
	}
	return builder.String(), nil
}

func (s *Storage) Add(userID, sku uint64, skuCount uint) error {
	return s.add(cart.UserID(userID), cart.SKU(sku), skuCount)
}

func (s *Storage) add(userID cart.UserID, sku cart.SKU, skuCount uint) error {

	if skuCount < 1 {
		return storage_errors.ErrInvalidSKUCount
	}

	if cartItem, ok := s.store[userID][sku]; ok {
		cartItem.QTY += skuCount
	}
	s.store[userID] = make(map[cart.SKU]*cart.CartItem)
	s.store[userID][sku] = &cart.CartItem{
		SKU: sku,
		QTY: skuCount,
	}
	return nil
}
