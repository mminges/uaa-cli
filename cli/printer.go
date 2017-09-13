package cli

import (
	"encoding/json"
)

type Printer interface {
	Print(interface{}) error
}

type JsonPrinter struct {
	Log Logger
}

func NewJsonPrinter(log Logger) JsonPrinter {
	return JsonPrinter{log}
}
func (jp JsonPrinter) Print(obj interface{}) error {
	j, err := json.MarshalIndent(&obj, "", "  ")
	if err != nil {
		return err
	}
	jp.Log.Robots(string(j))
	return nil
}

type TestPrinter struct {
	CallData map[string]interface{}
}

func NewTestPrinter() TestPrinter {
	tp := TestPrinter{}
	tp.CallData = make(map[string]interface{})
	return tp
}
func (tp TestPrinter) Print(obj interface{}) error {
	tp.CallData["Print"] = obj
	return nil
}
