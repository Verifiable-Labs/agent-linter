package engine

type Rule interface {
	ID() string
	Description() string
	DefaultSeverity() Severity
	Apply(in Inputs) []Finding
}
