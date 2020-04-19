package transition

import (
	"fmt"
	"reflect"
)

type checkUnit struct {
	checkFunction         interface{}
	parameters            []interface{}
	expectResult          bool
	errorToUnexpectResult error
}
func NewCheckUnit(checkFunction interface{},parameters []interface{},expectResult bool,errorToUnexpectResult error) checkUnit{
	return checkUnit{
		checkFunction,
		parameters,
		expectResult,
		errorToUnexpectResult,
	}
}
// 校验
func (cku checkUnit) check() error {
	if result, err := check(cku.checkFunction, cku.parameters...); err != nil {
		return err
	} else {
		if result != cku.expectResult {
			return cku.errorToUnexpectResult
		}
	}
	return nil
}

// checkFunction 为要使用的函数
// parameters 为其参数
func check(checkFunction interface{}, parameters ...interface{}) (bool, error) {

	fType := reflect.TypeOf(checkFunction)
	fValue := reflect.ValueOf(checkFunction)
	if len(parameters) != fType.NumIn() {
		return false, fmt.Errorf("the count(%d) of parameters is not equal to the count(%d) of checkType", len(parameters), fType.NumIn())
	}

	for i := 0; i < fType.NumIn(); i++ {
		if fType.In(i) != reflect.TypeOf(parameters[i]) {
			return false, fmt.Errorf("parameter(%d) can't fit function's field(%d)", i, i)
		}
	}
	in := make([]reflect.Value, len(parameters))
	for i := 0; i < len(parameters); i++ {
		in[i] = reflect.ValueOf(parameters[i])
	}
	out := fValue.Call(in)
	if len(out) != 2 {
		return false, fmt.Errorf("the length of return to the function(%v) is wrong", checkFunction)
	}

	boolInterface := out[0].Interface()
	errInterface := out[1].Interface()

	if boolInterface == nil {
		return false, fmt.Errorf("the function(%v) is not a valid check function (0)", checkFunction)
	}
	if _, ok := boolInterface.(bool); !ok {
		return false, fmt.Errorf("the function(%v) is not a valid check function (1)", checkFunction)
	}

	if errInterface == nil {
		return boolInterface.(bool), nil
	}

	if _, ok := errInterface.(error); !ok {
		return false, fmt.Errorf("the function(%v) is not a valid check function (2)", checkFunction)
	}

	return boolInterface.(bool), errInterface.(error)

}
