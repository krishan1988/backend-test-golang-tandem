// Package models contains entities specific to
// the backend-test-golang application domain.
package models

// FactFilter contains fields that are needed to filter out factors.
type FactFilter struct {
	Found string
	Types []string
	Page  int
	Limit int
}

// GetFound return found if value exist and ok is true , other wise ok is false.
func (f FactFilter) GetFound() (value, ok bool) {
	if f.Found == "true" {
		return true, true
	} else if f.Found == "false" {
		return false, true
	} else {
		return false, false
	}
}

// GetTypes return types if value exist and ok is true , other wise ok is false.
func (f FactFilter) GetTypes() (value []string, ok bool) {
	if len(f.Types) == 0 {
		return nil, false
	}
	return f.Types, true
}
