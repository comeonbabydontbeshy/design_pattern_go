package decorate

import (
	"fmt"
	"testing"
)

func TestBirdFactory(t *testing.T) {
	var beverage Beverage
	beverage = &Coffee{description: "Coffee"}
	beverage = Mocha{beverage: beverage, description: "Mocha"}
	beverage = Milk{beverage: beverage, description: "Milk"}
	fmt.Printf("description : %s , cost : %d", beverage.getDescription(), beverage.cost())
	if beverage.cost() != 18 {
		t.Errorf("cost calculate wrong")
	}
}
