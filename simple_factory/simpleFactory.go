package simple_factory

type Operation interface {
	GetResult(a, b float64) float64
}

type OperationAdd struct{}

func (ope *OperationAdd) GetResult(a, b float64) float64 {
	return a + b
}

type OperationSub struct{}

func (ope *OperationSub) GetResult(a, b float64) float64 {
	return a - b
}

type OperationMul struct{}

func (ope *OperationMul) GetResult(a, b float64) float64 {
	return a * b
}

type OperationDiv struct{}

func (ope *OperationDiv) GetResult(a, b float64) float64 {
	return a / b
}

// 根据不同的参数返回不同的实例
func CreateOperate(operate string) Operation {
	var ope Operation
	switch operate {
	case "+":
		ope = &OperationAdd{}
	case "-":
		ope = &OperationSub{}
	case "*":
		ope = &OperationMul{}
	case "/":
		ope = &OperationDiv{}
	}
	return ope
}
