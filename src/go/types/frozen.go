package types

type Frozen struct {
	base Type
	aType
}

func (f *Frozen) Underlying() Type {
	return f
}

func (f *Frozen) Under() Type {
	return f
}

func (f *Frozen) String() string {
	return TypeString(f, nil)
}

func (f *Frozen) Elem() Type { return f.base }
