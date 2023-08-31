package packages

import (
	std "github.com/eshu0/poe-core/pkg/std"
)

type Method struct {
	Name        string
	Parameters  []Parameter
	Structure   Structure
	ReturnError std.Error
	Return      []Parameter
}
