package builder

import "testing"

func TestBuilder4(t *testing.T) {
	obj := NewBigClass(WithAgeOption(15), WithFieldBOption("fieldB"), WithHeightOption(180))
	t.Logf("bigclass: +%v", obj)
}
