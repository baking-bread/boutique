package value_type

type ValueType int64

const (
	Undefined ValueType = iota
	Text
	Bool
	Number
)
