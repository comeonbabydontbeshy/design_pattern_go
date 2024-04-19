package builder

import "testing"

func TestBuilder3(t *testing.T) {
	fb := NewFoodBuilder()
	food1, err := fb.Type("apple").Cost(12.12).Brand("hongfushi").Name("woshihongfushi").Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("food: %+v", food1)
	}
}
