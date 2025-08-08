package nodes

import (
	"SoftPLC/function"
)

type FunctionInputValueNode struct {
	id                 int
	nodeType           string
	output             []InputNodeHandle
	parameterValueData []string //select Appliance name + value +
}

func (n *FunctionInputValueNode) GetNodeType() string {
	return n.nodeType
}

func (n *FunctionInputValueNode) GetId() int {
	return n.id
}

func init() {
	var nameServices []string
	nameServices = []string{"value", "number", "string"}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var functionInputDescription = nodeDescription{
		AccordionName:     "Functions",
		PrimaryType:       "inputNode",
		Type_:             "functionInputValue",
		Display:           " Input value ",
		Label:             "function Input value",
		Stretchable:       false,
		Services:          services,
		SubServices:       []subServicesStruct{},
		Input:             []dataTypeNameStruct{},
		Output:            []dataTypeNameStruct{{DataType: "value", Name: "Output"}},
		ParameterNameData: []string{"name", "default value"},
	}
	RegisterNodeCreator("functionInputValue", func() (Node, error) {
		return &FunctionInputValueNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, functionInputDescription)
}

func (n *FunctionInputValueNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = output_
	n.parameterValueData = parameterValueData_

	if len(n.parameterValueData) <= 1 {
		n.parameterValueData = append(n.parameterValueData, "")
	}
	n.output[0].FriendlyName = n.parameterValueData[0]
	function.AddInput(n.parameterValueData[0], n.parameterValueData[1])

}

func (n *FunctionInputValueNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *FunctionInputValueNode) Clone() *FunctionInputValueNode {
	clonedOutput := make([]InputNodeHandle, len(n.output))
	copy(clonedOutput, n.output)

	clonedParamValues := make([]string, len(n.parameterValueData))
	copy(clonedParamValues, n.parameterValueData)

	return &FunctionInputValueNode{
		id:                 n.id,
		nodeType:           n.nodeType,
		output:             clonedOutput,
		parameterValueData: clonedParamValues,
	}
}
