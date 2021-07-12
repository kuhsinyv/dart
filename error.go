package dart

import "errors"

var (
	ErrFetch         = errors.New("fetch error")
	ErrEmptyPatterns = errors.New("empty patterns")
)
