// Code generated by github.com/sighupio/go-jsonschema, DO NOT EDIT.

package schema

import other "github.com/sighupio/go-jsonschema/tests/data/crossPackage/other"

type Schema struct {
	// DefInOtherSchema corresponds to the JSON schema field "defInOtherSchema".
	DefInOtherSchema *other.Thing `json:"defInOtherSchema,omitempty" yaml:"defInOtherSchema,omitempty" mapstructure:"defInOtherSchema,omitempty"`

	// DefInSameSchema corresponds to the JSON schema field "defInSameSchema".
	DefInSameSchema *Thing `json:"defInSameSchema,omitempty" yaml:"defInSameSchema,omitempty" mapstructure:"defInSameSchema,omitempty"`
}

type Thing struct {
	// S corresponds to the JSON schema field "s".
	S *string `json:"s,omitempty" yaml:"s,omitempty" mapstructure:"s,omitempty"`
}
