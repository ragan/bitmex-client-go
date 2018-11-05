package bitmex

func GetMargin(c Client) (m *Margin, err *Error) {
	m = &Margin{}
	err = c.do("GET", "/user/margin", "currency=XBt", nil, m)
	if err != nil {
		return nil, err
	}
	return
}
