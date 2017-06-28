package pantheon

import (
	"bytes"
	"encoding/json"
)

// jsonFloat is an int64 which unmarshals from JSON
// as either unquoted or quoted (with any amount
// of internal leading/trailing whitespace).
// it is based off of https://play.golang.org/p/KNPxDL1yqL
type jsonInt64 int64

func (f jsonInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(f))
	// OR if you want "string" output, something like:
	//return json.Marshal(strconv.FormatFloat(int64(f), 'g', -1, 64))
}

func (f *jsonInt64) UnmarshalJSON(data []byte) error {
	var v int64

	// Option C:
	// Both the above might allow for weird things like only
	// a leading quote or trailing quote or multiple
	// leading/trailing quotes.
	//
	// To be more strict one could do something like this:

	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		// Remove single set of matching quotes
		data = data[1 : len(data)-1]
	}
	// And, optionally if you also then want to remove any whitespace:
	data = bytes.TrimSpace(data)

	err := json.Unmarshal(data, &v)
	*f = jsonInt64(v)
	return err
}
