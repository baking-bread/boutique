package value_kind

type ValueKind int64

const (
	Undefined ValueKind = iota
	Input
	Options
	Regex
)
