package util

// Notation is prepare for listener & ast package
//
// Because both of them need the same structure but can't have circle import
type Notation struct {
	Leading string
	Content []string
}
