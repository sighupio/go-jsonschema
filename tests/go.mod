module github.com/sighupio/go-jsonschema/tests

go 1.21

replace (
	github.com/sighupio/go-jsonschema => ../
	github.com/sighupio/go-jsonschema/tests/helpers/other => ./helpers/other
)

require (
	github.com/magiconair/properties v1.8.7
	github.com/sighupio/go-jsonschema v0.0.0-00010101000000-000000000000
	github.com/sighupio/go-jsonschema/tests/helpers/other v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/goccy/go-yaml v1.11.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sanity-io/litter v1.5.5 // indirect
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
