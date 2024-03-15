// Code generated by github.com/sighupio/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type MaxItems struct {
	// MyNestedArray corresponds to the JSON schema field "myNestedArray".
	MyNestedArray [][]interface{} `json:"myNestedArray,omitempty" yaml:"myNestedArray,omitempty" mapstructure:"myNestedArray,omitempty"`

	// MyStringArray corresponds to the JSON schema field "myStringArray".
	MyStringArray []string `json:"myStringArray,omitempty" yaml:"myStringArray,omitempty" mapstructure:"myStringArray,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MaxItems) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain MaxItems
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.MyNestedArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myNestedArray", 5)
	}
	for i1 := range plain.MyNestedArray {
		if len(plain.MyNestedArray[i1]) > 5 {
			return fmt.Errorf("field %s length: must be <= %d", fmt.Sprintf("myNestedArray[%d]", i1), 5)
		}
	}
	if len(plain.MyStringArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myStringArray", 5)
	}
	*j = MaxItems(plain)
	return nil
}
