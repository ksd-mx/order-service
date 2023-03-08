package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
	Quantity   int
}

func NewOrder(id string, price float64, tax float64, quantity int) *Order {
	return &Order{
		ID:       id,
		Price:    price,
		Tax:      tax,
		Quantity: quantity,
	}
}

func (o *Order) SetPrice(price float64) {
	o.Price = price
}

func (o Order) GetTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func (o Order) Validate() error {
	if o.ID == "" {
		return errors.New("Order ID cannot be empty")
	}

	if o.Price <= 0 {
		return errors.New("Order price is invalid")
	}

	if o.Tax <= 0 {
		return errors.New("Order tax is invalid")
	}

	if o.Quantity <= 0 {
		return errors.New("Order quantity is invalid")
	}

	return nil
}
