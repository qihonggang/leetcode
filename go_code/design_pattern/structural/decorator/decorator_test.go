package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpADDDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}
