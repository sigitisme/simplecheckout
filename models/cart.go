package models

// Cart ...
type Cart struct {
	Contents []*Content `json:"contents"`
}

// Content ...
type Content struct {
	SKU      string
	Name     string
	Quantity int
}
