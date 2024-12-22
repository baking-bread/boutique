package domain

import (
	"github.com/baking-bread/boutique/internal/core/domain/value_kind"
	"github.com/baking-bread/boutique/internal/core/domain/value_type"
)

type TemplateValue struct {
	Name        string
	Description string
	Type        value_type.ValueType
	Kind        value_kind.ValueKind
	Options     []any
}
