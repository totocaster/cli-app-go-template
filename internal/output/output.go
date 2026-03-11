package output

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSON writes one JSON document followed by a newline.
func JSON(w io.Writer, v any) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}

// Line writes a formatted line and appends a trailing newline.
func Line(w io.Writer, format string, args ...any) error {
	_, err := fmt.Fprintf(w, format+"\n", args...)
	return err
}

// Raw writes text exactly as provided.
func Raw(w io.Writer, text string) error {
	_, err := io.WriteString(w, text)
	return err
}
