package cart

type CartItem struct {
	SKU SKU  `json:"sku"`
	QTY uint `json:"quantity,omitempty"`
}

type (
	UserID uint64
	SKU    uint64
)
