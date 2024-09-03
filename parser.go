package validate

import (
	"encoding/json"
	"io"
)

// Parse parses a JSON expression into the provided struct instance
// and validates it.
// Any fields whose type is [Required] are checked that they have a value assigned to them from the incoming JSON.
//
// Parse returns a multi-error wrapping all the errors that have occurred in the course of the verification.
func Parse(r io.Reader, obj interface{}) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}

	return Struct(obj)
}
