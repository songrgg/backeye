package common

type ContextKey int

const (
	ResponseKey         = ContextKey(0)
	AssertionResultsKey = ContextKey(1)
	VariablesKey        = ContextKey(2)
)
