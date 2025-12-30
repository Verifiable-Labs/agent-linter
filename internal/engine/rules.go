package engine

type Rule interface {
	ID() string
	Apply(in Inputs) []Finding
}
