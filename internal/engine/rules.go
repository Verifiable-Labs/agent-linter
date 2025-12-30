package engine

type Rule interface {
	ID() string
	DefaultSeverity() Severity
	Apply(in Inputs) []Finding
}
