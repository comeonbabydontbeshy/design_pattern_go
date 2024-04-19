package builder

type BigClass struct {
	name   string
	age    int
	sex    string
	weight float64
	height float64
	width  float64
	fieldA string
	fieldB string
	fieldC string
}

type Option func(*BigClass)

func WithNameOption(name string) Option {
	return func(obj *BigClass) {
		obj.name = name
	}
}

func WithAgeOption(age int) Option {
	return func(obj *BigClass) {
		obj.age = age
	}
}

func WithSexOption(sex string) Option {
	return func(obj *BigClass) {
		obj.sex = sex
	}
}

func WithWeightOption(weight float64) Option {
	return func(obj *BigClass) {
		obj.weight = weight
	}
}

func WithHeightOption(height float64) Option {
	return func(obj *BigClass) {
		obj.height = height
	}
}

func WithWidthOption(width float64) Option {
	return func(obj *BigClass) {
		obj.width = width
	}
}

func WithFieldAOption(fieldA string) Option {
	return func(obj *BigClass) {
		obj.fieldA = fieldA
	}
}

func WithFieldBOption(fieldB string) Option {
	return func(obj *BigClass) {
		obj.fieldB = fieldB
	}
}

func WithFieldCOption(fieldC string) Option {
	return func(obj *BigClass) {
		obj.fieldC = fieldC
	}
}

func NewBigClass(opts ...Option) *BigClass {
	obj := &BigClass{}
	for _, opt := range opts {
		opt(obj)
	}
	if obj.name == "" {
		obj.name = "小明"
	}
	if obj.age == 0 {
		obj.age = 20
	}
	return obj
}
