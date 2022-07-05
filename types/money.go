package types

type MoneyType string

const (
	MoneyType_CZK = "czk"
)

func (mt MoneyType) DecimalLength() uint {
	switch mt {
	case MoneyType_CZK:
		return 100

	default:
		return 1
	}
}

type Money struct {
	value    uint64    ``
	positive bool      ``
	_type    MoneyType ``
}

func (m Money) ToFloat64() float64 {
	var positive float64

	if m.positive {
		positive = 1
	} else {
		positive = -1
	}

	return float64(m.value) / float64(m._type.DecimalLength()) * positive
}

func (m Money) Add(value float64) Money {
	m.value += uint64(value * float64(m._type.DecimalLength()))
	return m
}

func (m Money) Subtract(value float64) Money {
	return m.Add(value * -1)
}
