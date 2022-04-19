package factory

type IAnimal interface {
	legCnt() int32
}

type Bird struct{}

func (b Bird) legCnt() int32 {
	return 2
}

type Tiger struct{}

func (t Tiger) legCnt() int32 {
	return 4
}

type IAnimalFactory interface {
	createAnimal() IAnimal
}

type BirdFactory struct{}

func (bf BirdFactory) createAnimal() IAnimal {
	return &Bird{}
}

type TigerFactory struct{}

func (tf TigerFactory) createAnimal() IAnimal {
	return &Tiger{}
}
