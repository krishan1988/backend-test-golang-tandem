// Package models contains entities specific to
// the backend-test-golang application domain.
package models

import (
	"fmt"
	"strconv"
	"strings"
)

const defaultPageSize = 10

// FactFilterBuilder responsible for creating FactFilter
type FactFilterBuilder struct {
	errors     []error
	factFilter FactFilter
}

// NewFactFilterBuilder create a new instance of a FactFilterBuilder
// with default values and returns its' pointer.
func NewFactFilterBuilder() *FactFilterBuilder {
	return &FactFilterBuilder{
		errors: make([]error, 0),
		factFilter: FactFilter{
			Types: make([]string, 0),
		},
	}
}

// SetPage set page number to the FactFilter.
func (b *FactFilterBuilder) SetPage(page string) *FactFilterBuilder {
	if page == "" {
		b.handleError("request should have page filter")
		return b
	}
	val, err := strconv.Atoi(page)
	if err != nil {
		b.handleError("invalid page filter. error: %s", err)
		return b
	}

	if val < 1 {
		b.handleError("page should be grater than '0'")
	}

	b.factFilter.Page = val
	return b
}

// SetLimit set limit (number of records in the page) to the FactFilter.
func (b *FactFilterBuilder) SetLimit(limit string) *FactFilterBuilder {
	if limit == "" {
		b.factFilter.Limit = defaultPageSize
		return b
	}

	val, err := strconv.Atoi(limit)
	if err != nil {
		b.handleError("invalid limit filter. error: %s", err)
		return b
	}

	if val < 10 {
		b.handleError("limit should be grater than '10'")
		return b
	}

	b.factFilter.Limit = val
	return b
}

// SetTypes set types to  the FactFilter.
func (b *FactFilterBuilder) SetTypes(types []string) *FactFilterBuilder {
	b.factFilter.Types = types
	return b
}

// SetFound set found to the FactFilter by converting given string
// to boolean. For invalid string, holds the error in the receiver instance.
func (b *FactFilterBuilder) SetFound(found string) *FactFilterBuilder {
	if found == "" {
		return b
	}
	_, err := strconv.ParseBool(found)
	if err != nil {
		return b.handleError("found should be boolean. error: %s", err)
	}
	b.factFilter.Found = found
	return b
}

// Get returns the created FactFilter if there are no errors
// in the building step, otherwise return and error.
func (b *FactFilterBuilder) Get() (FactFilter, error) {
	if len(b.errors) > 0 {
		args := make([]any, 0)
		for _, err := range b.errors {
			args = append(args, err)
		}
		return FactFilter{}, fmt.Errorf("%s"+strings.Repeat(",%s", len(args)-1), args...)
	}
	return b.factFilter, nil
}

func (b *FactFilterBuilder) handleError(format string, errs ...error) *FactFilterBuilder {
	args := make([]any, 0)
	for _, err := range errs {
		args = append(args, err)
	}
	b.errors = append(b.errors, fmt.Errorf(format, args...))
	return b
}
