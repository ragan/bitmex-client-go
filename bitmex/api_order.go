package bitmex

import (
	"net/http"
	"net/url"
	"strconv"
)

type OrderForm struct {
	Symbol   string
	OrderQty float64
	Price    float64
}

func (o OrderForm) toValues() url.Values {
	return map[string][]string{
		"symbol":   {o.Symbol},
		"orderQty": {strconv.FormatFloat(o.OrderQty, 'f', 4, 32)},
		"price":    {strconv.FormatFloat(o.Price, 'f', 4, 32)},
	}
}

func DeleteOrderAll(c Client) (o *Order, err *Error) {
	o = &Order{}
	err = c.do(http.MethodDelete, "/oder/all", "", nil, o)
	if err != nil {
		o = nil
	}
	return
}

func PostOrder(c Client, of *OrderForm) (o *Order, errData *Error) {
	errData = c.do(http.MethodPost, "/order", of.toValues().Encode(), nil, &o)
	if errData != nil {
		o = nil
	}
	return
}
