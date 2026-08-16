package json_store

import (
	"CliCart/internal/domain/cart"
	storage_errors "CliCart/internal/storage/errors"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Storage struct {
	path  string
	store map[cart.UserID]map[cart.SKU]*cart.CartItem
}

func NewStorage(path string) (*Storage, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := &Storage{
		path:  path,
		store: make(map[cart.UserID]map[cart.SKU]*cart.CartItem),
	}
	s.readDataFromFile(f)

	return s, nil
}

func (s *Storage) Add(userID, sku uint64, skuCount uint) error {
	return s.add(cart.UserID(userID), cart.SKU(sku), skuCount)
}

func (s *Storage) add(userID cart.UserID, sku cart.SKU, skuCount uint) (err error) {

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

	return s.writeDataToFile()
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

func (s Storage) readDataFromFile(file *os.File) (err error) {
	//var file *os.File
	//
	//file, err = os.OpenFile(j.path, os.O_RDWR, 0666)
	//if err != nil {
	//	return err
	//}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s.store)
	if err != nil {
		return fmt.Errorf("cannot unmarshal json %s: %w", s.path, err)
	}

	return nil

}

func (s Storage) writeDataToFile() (err error) {
	var file *os.File

	file, err = os.OpenFile(s.path, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(s.store)
	if err != nil {
		return fmt.Errorf("cannot write file %s: %w", s.path, err)
	}

	return nil
}
