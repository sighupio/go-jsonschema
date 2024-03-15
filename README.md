**go-jsonschema is a tool to generate Go data types from [JSON Schema](http://json-schema.org/) definitions.**

This tool generates Go data types and structs that corresponds to definitions in the schema, along with unmarshalling code that validates the input JSON according to the schema's validation rules.

## Badges
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/sighupio/go-jsonschema?style=flat)](https://github.com/sighupio/go-jsonschema/releases/latest)
[![GitHub Workflow Status (event)](https://img.shields.io/github/actions/workflow/status/sighupio/go-jsonschema/development.yaml?style=flat)](https://github.com/sighupio/go-jsonschema/actions?workflow=development)
[![License](https://img.shields.io/github/license/sighupio/go-jsonschema?style=flat)](/LICENSE.md)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sighupio/go-jsonschema?style=flat)](https://tip.golang.org/doc/go1.21)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/sighupio/go-jsonschema?style=flat)](https://github.com/sighupio/go-jsonschema)
[![GitHub repo file count (file type)](https://img.shields.io/github/directory-file-count/sighupio/go-jsonschema?style=flat)](https://github.com/sighupio/go-jsonschema)
[![GitHub all releases](https://img.shields.io/github/downloads/sighupio/go-jsonschema/total?style=flat)](https://github.com/sighupio/go-jsonschema)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/sighupio/go-jsonschema?style=flat)](https://github.com/sighupio/go-jsonschema/commits)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=flat)](https://conventionalcommits.org)
[![Codecov](https://img.shields.io/codecov/c/gh/sighupio/go-jsonschema?style=flat&token=lPWlXd3MVK)](https://codecov.io/gh/sighupio/go-jsonschema)
[![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability/sighupio/go-jsonschema?style=flat)](https://codeclimate.com/github/sighupio/go-jsonschema)
[![Go Report Card](https://goreportcard.com/badge/github.com/sighupio/go-jsonschema)](https://goreportcard.com/report/github.com/sighupio/go-jsonschema)

## Installing

* **Download**: Get a release [here](https://github.com/sighupio/go-jsonschema/releases).

* **Install from source**: To install with Go 1.16+:

```shell
$ go get github.com/sighupio/go-jsonschema/...
$ go install github.com/sighupio/go-jsonschema@latest
```

* **Install with Brew**: To install with [Homebrew](https://brew.sh):

```shell
$ brew tap sighupio/go-jsonschema
$ brew install go-jsonschema
```

## Contributing

This project makes use of [go workspaces](https://go.dev/ref/mod#workspaces) in order to ease testing of the generated code during development while keeping the codebase as tidy and maintainable as possible.
It is bit of an unusual choice, but it allows to not only test the code-generation logic, but also the generated code itself.

## Usage

At its most basic:

```shell
$ go-jsonschema -p main schema.json
```

This will write a Go source file to standard output, declared under the package `main`.

You can generate code for multiple schemas in the same invocation, optionally writing to different files inside different packages:

```shell
$ go-jsonschema \
  --schema-package=https://example.com/schema1=github.com/myuser/myproject \
   --schema-output=https://example.com/schema1=schema1.go \
  --schema-package=https://example.com/schema2=github.com/myuser/myproject/stuff \
   --schema-output=https://example.com/schema2=stuff/schema2.go \
  schema1.json schema2.json
```

This will create `schema1.go` (declared as `package myproject`) and `stuff/schema2.go` (declared as `package stuff`). If `schema1.json` refers to `schema2.json` or vice versa, the two Go files will import the other package that it depends on. Note the flag format:

```
--schema-package=https://example.com/schema1=github.com/myuser/myproject \
                 ^                           ^
                 |                           |
                 schema $id                  full import URL
```

### Special types

In a few cases, special types are used to help with serializing/deserializing
data frrom JSON. Namely a custom types is provided for the following semantic
types:

* `SerializableDate`
* `SerializableTime`

These types are needed because there is no native type provided by Go which
properly handles them.

## Status

While not finished, go-jsonschema can be used today. Aside from some minor features, only specific validations remain to be fully implemented.

## Release

When releasing a new version, a problem may occur when running `make release`, to fix this, run the following command:

```shell
$ docker buildx create --bootstrap --use
```

### Validation

- Core ([RFC draft](http://json-schema.org/latest/json-schema-core.html))
  - [x] Data model (§4.2.1)
    - [x] `null`
    - [x] `boolean`
    - [x] `object`
    - [x] `array`
    - [x] `number`
      - [ ] Option to use `json.Number`
    - [x] `string`
  - [ ] Location identifiers (§8.2.3)
    - [x] References against top-level names: `#/$defs/someName`
    - [ ] References against nested names: `#/$defs/someName/$defs/someOtherName`
    - [x] References against top-level names in external files: `myschema.json#/$defs/someName`
    - [ ] References against nested names: `myschema.json#/$defs/someName/$defs/someOtherName`
  - [x] Comments (§9)
- Validation ([RFC draft](http://json-schema.org/latest/json-schema-validation.html))
  - [ ] Schema annotations (§10)
    - [x] `description`
    - [x] `default` (only for struct fields)
    - [ ] `readOnly`
    - [ ] `writeOnly`
    - [ ] ~~`title`~~ (N/A)
    - [ ] ~~`examples`~~ (N/A)
  - [ ] General validation (§6.1)
    - [x] `enum`
    - [x] `type` (single)
    - [x] `type` (multiple; **note**: partial support, limited validation)
    - [ ] `const`
  - [ ] Numeric validation (§6.2)
    - [ ] `multipleOf`
    - [ ] `maximum`
    - [ ] `exclusiveMaximum`
    - [ ] `minimum`
    - [ ] `exclusiveMinimum`
  - [ ] String validation (§6.3)
    - [X] `maxLength`
    - [X] `minLength`
    - [ ] `pattern`
  - [ ] Array validation (§6.4)
    - [X] `items`
    - [x] `maxItems`
    - [x] `minItems`
    - [ ] `uniqueItems`
    - [ ] `additionalItems`
    - [ ] `contains`
  - [ ] Object validation (§6.5)
    - [x] `required`
    - [x] `properties`
    - [ ] `patternProperties`
    - [ ] `dependencies`
    - [ ] `propertyNames`
    - [ ] `maxProperties`
    - [ ] `minProperties`
  - [ ] Conditional subschemas (§6.6)
    - [ ] `if`
    - [ ] `then`
    - [ ] `else`
  - [ ] Boolean subschemas (§6.7)
    - [ ] `allOf`
    - [ ] `anyOf`
    - [ ] `oneOf`
    - [ ] `not`
  - [ ] Semantic formats (§7.3)
    - [x] Dates and times
    - [ ] Email addresses
    - [ ] Hostnames
    - [ ] IP addresses
    - [ ] Resource identifiers
    - [ ] URI-template
    - [ ] JSON pointers
    - [ ] Regex

## License

MIT license. See `LICENSE` file.
