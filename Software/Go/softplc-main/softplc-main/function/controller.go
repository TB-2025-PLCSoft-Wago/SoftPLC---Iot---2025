package function

import "fmt"

type OutputState struct {
	Name  string
	Value string
}

var OutputsStateFunction []OutputState

type InputState struct {
	Name  string
	Value string
}

var InputsStateFunction []InputState

func AddInput(name string, value string) {
	InputsStateFunction = append(InputsStateFunction, InputState{name, value})
}
func AddOutput(name string, value string) {
	OutputsStateFunction = append(OutputsStateFunction, OutputState{name, value})
}

func UpdateFunctionOutput(name string, value string) {
	//output
	for i := range OutputsStateFunction {
		if OutputsStateFunction[i].Name == name {
			OutputsStateFunction[i].Value = value
		}
	}
}
func GetFunctionOutput(name string) (value string) {
	//output
	for i := range OutputsStateFunction {
		if OutputsStateFunction[i].Name == name {
			return OutputsStateFunction[i].Value
		}
	}
	fmt.Println("Function output : no value find ")
	return "Function output : no value find"
}
func UpdateFunctionInputs(name string, value string) {
	//input
	for i := range InputsStateFunction {
		if InputsStateFunction[i].Name == name {
			InputsStateFunction[i].Value = value
		}
	}
}
