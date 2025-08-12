package function

import "fmt"

type OutputState struct {
	Name         string
	Value        string
	FunctionName string
}

var OutputsStateFunction []OutputState

type InputState struct {
	Name         string
	Value        string
	FunctionName string
}

var InputsStateFunction []InputState

func AddInput(name string, value string, functionName string) {
	InputsStateFunction = append(InputsStateFunction, InputState{name, value, functionName})
}

func AddOutput(name string, value string, functionName string) {
	OutputsStateFunction = append(OutputsStateFunction, OutputState{name, value, functionName})
}

func UpdateFunctionOutput(name string, value string, functionName string) {
	for i := range OutputsStateFunction {
		if OutputsStateFunction[i].Name == name && OutputsStateFunction[i].FunctionName == functionName {
			OutputsStateFunction[i].Value = value
			//return
		}
	}
}

func GetFunctionOutput(name string, functionName string) (value string) {
	for i := range OutputsStateFunction {
		if OutputsStateFunction[i].Name == name && OutputsStateFunction[i].FunctionName == functionName {
			return OutputsStateFunction[i].Value
		}
	}
	fmt.Println("Function output : no value found")
	return "Function output : no value found"
}

func UpdateFunctionInput(name string, value string, functionName string) {
	for i := range InputsStateFunction {
		if InputsStateFunction[i].Name == name && InputsStateFunction[i].FunctionName == functionName {
			InputsStateFunction[i].Value = value
			//return
		}
	}
}

func GetFunctionInput(name string, functionName string) (value string) {
	for i := range InputsStateFunction {
		if InputsStateFunction[i].Name == name && InputsStateFunction[i].FunctionName == functionName {
			return InputsStateFunction[i].Value
		}
	}
	fmt.Println("Function input : no value found")
	return "Function input : no value found"
}
