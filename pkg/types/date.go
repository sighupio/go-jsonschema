package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var ErrDateNotJSONString = errors.New("cannot parse non-string value as a date")

type SerializableDate struct {
	time.Time
}

func (d *SerializableDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Format(time.DateOnly) + "\""), nil
}

func (d *SerializableDate) parse(date string) error {
	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return fmt.Errorf("unable to parse date from JSON: %w", err)
	}

	d.Time = parsedDate

	return nil
}

func (d *SerializableDate) UnmarshalJSON(data []byte) error {
	stringifiedData := string(data)

	if stringifiedData == "null" {
		return nil
	}

	if !strings.HasPrefix(stringifiedData, "\"") || !strings.HasSuffix(stringifiedData, "\"") {
		return ErrDateNotJSONString
	}

	dataWithoutQuotes := stringifiedData[1 : len(stringifiedData)-1]

	return d.parse(dataWithoutQuotes)
}

func (d *SerializableDate) UnmarshalYAML(value *yaml.Node) error {
	return d.parse(value.Value)
}
