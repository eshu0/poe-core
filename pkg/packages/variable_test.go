package packages

import (
	"testing"
)

const (
	vartest = "wills var"
)

func TestVariableAddVariable(t *testing.T) {
	str := NewStructure(strcname1)

	v := NewRefVariable(vartest, "int")
	_, err := str.VariableUpdate(v)
	if err != nil {
		t.Fatalf(`TestStructureAddVariable() Variable AddNew: %s error :%s`, str.Name, err.Error())
	}

	_, err = str.VariableRead(vartest)
	if err != nil {
		t.Fatalf(`TestStructureAddVariable() Member Get: %s error :%s`, strcname, err.Error())
	}
}
