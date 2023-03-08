package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	Discount   float64
	FinalPrice float64
	Quantity   int
}

func NewOrder(id string, price float64, tax float64, discount float64, quantity int) (*Order, error) {
	order := Order{
		ID:       id,
		Price:    price,
		Tax:      tax,
		Discount: discount,
		Quantity: quantity,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) SetPrice(price float64) {
	o.Price = price
}

func (o Order) GetTotal() float64 {
	return ((o.Price + o.Tax) * float64(o.Quantity) * o.Discount) / 100
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

	if o.Discount >= 100.0 {
		return errors.New("Order discount is invalid")
	}

	if o.Quantity <= 0 {
		return errors.New("Order quantity is invalid")
	}

	return nil
}
