package factory

import (
	"reflect"
	"testing"
)

func TestBirdFactory(t *testing.T) {
	tests := []struct {
		name string
		want IAnimal
	}{
		{
			name: "bird",
			want: &Bird{},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			var bf IAnimalFactory = &BirdFactory{}
			if got := bf.createAnimal(); !reflect.DeepEqual(got, v.want) {
				t.Errorf("createAnimal() = %v, want %v", got, v.want)
			}
		})
	}
}

func TestTigerFactory(t *testing.T) {
	tests := []struct {
		name string
		want IAnimal
	}{
		{
			name: "tiger",
			want: &Tiger{},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			var bf IAnimalFactory = &TigerFactory{}
			if got := bf.createAnimal(); !reflect.DeepEqual(got, v.want) {
				t.Errorf("createAnimal() = %v, want %v", got, v.want)
			}
		})
	}
}
