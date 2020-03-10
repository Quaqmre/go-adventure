package token

type Token interface {
	Literal() string
	Type() string
}

type Number struct {
	Value string
}

func (n Number) Literal() string {
	return n.Value
}

func (n Number) Type() string {
	return "Number"
}

type Operator struct {
	Value string
}

func (o Operator) Literal() string {
	return o.Value
}

func (o Operator) Type() string {
	return "Operator"
}
