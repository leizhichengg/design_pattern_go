package simple_factory

import "testing"

func TestCreateOperateAdd(t *testing.T) {
	opeAdd := CreateOperate("+")
	result := opeAdd.GetResult(1, 2)
	if result != 3 {
		t.Errorf("The result should be %f, not %f\n", 3.0, result)
	}
}

func TestCreateOperateMul(t *testing.T) {
	opeAdd := CreateOperate("*")
	result := opeAdd.GetResult(1, 2)
	if result != 2 {
		t.Errorf("The result should be %f, not %f\n", 2.0, result)
	}
}