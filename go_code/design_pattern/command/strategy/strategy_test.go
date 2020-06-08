package strategy

import "testing"

func TestPayByCash(t *testing.T) {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
}

func TestPayByBank(t *testing.T) {
	payment := NewPayment("Ada", "", 123, &Bank{})
	payment.Pay()
}
