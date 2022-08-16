package decorate

import "fmt"

type Beverage interface {
	getDescription() string
	cost() int
}

type Coffee struct {
	description string
}

func (this Coffee) getDescription() string {
	return this.description
}

func (this Coffee) cost() int {
	return 10
}

// 辅料装饰类是整个装饰器模式的精华，既有继承也有组合
type Mocha struct {
	beverage    Beverage
	description string
}

func (this Mocha) getDescription() string {
	return fmt.Sprintf("%s,%s", this.beverage.getDescription(), this.description)
}

func (this Mocha) cost() int {
	return 3 + this.beverage.cost()
}

type Milk struct {
	beverage    Beverage
	description string
}

func (this Milk) getDescription() string {
	return fmt.Sprintf("%s,%s", this.beverage.getDescription(), this.description)
}

func (this Milk) cost() int {
	return 5 + this.beverage.cost()
}
