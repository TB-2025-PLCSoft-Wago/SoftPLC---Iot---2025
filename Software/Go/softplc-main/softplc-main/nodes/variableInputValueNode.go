package nodes

import (
	"SoftPLC/variable"
)

type VariableInputValueNode struct {
	id                 int
	nodeType           string
	output             []InputNodeHandle
	parameterValueData []string //select Appliance name + value +
}

func (n *VariableInputValueNode) GetNodeType() string {
	return n.nodeType
}

func (n *VariableInputValueNode) GetId() int {
	return n.id
}

func init() {
	var nameServices []string
	nameServices = []string{"value", "number", "string"}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var variableInputDescription = nodeDescription{
		AccordionName:     "Variables",
		PrimaryType:       "inputNode",
		Type_:             "variableInputValue",
		Display:           "Input value ",
		Label:             "variable Input value",
		Stretchable:       false,
		Services:          services,
		SubServices:       []subServicesStruct{},
		Input:             []dataTypeNameStruct{},
		Output:            []dataTypeNameStruct{{DataType: "value", Name: "Output"}},
		ParameterNameData: []string{"name", "default value"},
	}
	RegisterNodeCreator("variableInputValue", func() (Node, error) {
		return &VariableInputValueNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, variableInputDescription)
}

func (n *VariableInputValueNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = output_
	n.parameterValueData = parameterValueData_

	if len(n.parameterValueData) <= 1 {
		n.parameterValueData = append(n.parameterValueData, "")
	}
	n.output[0].FriendlyName = n.parameterValueData[0]
	variable.AddInput(n.parameterValueData[0], n.parameterValueData[1])

}

func (n *VariableInputValueNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *VariableInputValueNode) Clone() *VariableInputValueNode {
	clonedOutput := make([]InputNodeHandle, len(n.output))
	copy(clonedOutput, n.output)

	clonedParamValues := make([]string, len(n.parameterValueData))
	copy(clonedParamValues, n.parameterValueData)

	return &VariableInputValueNode{
		id:                 n.id,
		nodeType:           n.nodeType,
		output:             clonedOutput,
		parameterValueData: clonedParamValues,
	}
}
