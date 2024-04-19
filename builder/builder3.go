package builder

import "errors"

type Food struct {
	// 种类
	typ string
	// 名称
	name string
	// 重量
	weight float64
	// 品牌
	brand string
	// 价格
	cost float64
}

func NewFood(typ string, name string, weight float64, brand string, cost float64) *Food {
	return &Food{
		typ:    typ,
		name:   name,
		weight: weight,
		brand:  brand,
		cost:   cost,
	}
}

type FoodBuilder struct {
	Food
}

func NewFoodBuilder() *FoodBuilder {
	return &FoodBuilder{}
}

func (f *FoodBuilder) Type(typ string) *FoodBuilder {
	f.typ = typ
	return f
}

func (f *FoodBuilder) Name(name string) *FoodBuilder {
	f.name = name
	return f
}

func (f *FoodBuilder) Weight(weight float64) *FoodBuilder {
	f.weight = weight
	return f
}

func (f *FoodBuilder) Brand(brand string) *FoodBuilder {
	f.brand = brand
	return f
}

func (f *FoodBuilder) Cost(cost float64) *FoodBuilder {
	f.cost = cost
	return f
}

func (f *FoodBuilder) Build() (*Food, error) {
	if f.typ == "" {
		return nil, errors.New("miss type info")
	}
	if f.name == "" {
		return nil, errors.New("miss name info")
	}

	return &Food{
		typ:    f.typ,
		name:   f.name,
		brand:  f.brand,
		weight: f.weight,
		cost:   f.cost,
	}, nil
}
