package variable

type OutputState struct {
	Name  string
	Value string
}

var OutputsStateVariable []OutputState

type InputState struct {
	Name  string
	Value string
}

var InputsStateVariable []InputState

func AddInput(name string, value string) {
	InputsStateVariable = append(InputsStateVariable, InputState{name, value})
}
func AddOutput(name string, value string) {
	OutputsStateVariable = append(OutputsStateVariable, OutputState{name, value})
}

func UpdateVariableOutput(name string, value string) {
	//output
	for i := range OutputsStateVariable {
		if OutputsStateVariable[i].Name == name {
			OutputsStateVariable[i].Value = value
		}
	}
}
func UpdateVariableInputs() {
	//input
	for _, outputTemp := range OutputsStateVariable {
		for i, inputTemp := range InputsStateVariable {
			if inputTemp.Name == outputTemp.Name {
				if inputTemp.Value != outputTemp.Value {
					//fmt.Println("value : ", outputTemp.Value, " Name : ", outputTemp.Name)
					InputsStateVariable[i].Value = outputTemp.Value
				}
			}
		}
	}
}
