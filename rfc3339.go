package rfc3339

import "time"

const layout = time.RFC3339

// Parse parses a formatted string and returns the time value it represents.
func Parse(s string) (time.Time, error) {
	return time.Parse(layout, s)
}

// MustParse must version of Parse().
func MustParse(s string) time.Time {
	t, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	return t
}

// Format returns a textual representation.
func Format(t time.Time) string {
	return t.Format(layout)
}
