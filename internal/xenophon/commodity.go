package xenophon

type commodity int

const (
	Oats commodity = iota + 1
	Peas
	Beans
	Barley
)

func (c commodity) String() string {
	switch c {
	case Oats:
		return "Oats"
	case Peas:
		return "Peas"
	case Beans:
		return "Beans"
	case Barley:
		return "Barley"
	default:
		return "not found: " + string(c)
	}
}
