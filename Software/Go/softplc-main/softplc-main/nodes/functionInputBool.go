package nodes

import "SoftPLC/function"

type FunctionInputBoolNode struct {
	id                 int
	nodeType           string
	output             []InputNodeHandle
	parameterValueData []string //select Appliance name + value +
	functionName       string
}

func (n *FunctionInputBoolNode) GiveFunctionName(name string) {
	n.functionName = name
}
func (n *FunctionInputBoolNode) GetNodeType() string {
	return n.nodeType
}

func (n *FunctionInputBoolNode) GetId() int {
	return n.id
}

func init() {
	var nameServices []string
	nameServices = []string{"bool", "number", "string"}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var functionInputDescription = nodeDescription{
		AccordionName:     "Functions",
		PrimaryType:       "inputNode",
		Type_:             "functionInputBool",
		Display:           " Input bool ",
		Label:             "function Input bool",
		Stretchable:       false,
		Services:          services,
		SubServices:       []subServicesStruct{},
		Input:             []dataTypeNameStruct{},
		Output:            []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
		ParameterNameData: []string{"name", "default value"},
	}
	RegisterNodeCreator("functionInputBool", func() (Node, error) {
		return &FunctionInputBoolNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, functionInputDescription)
}

func (n *FunctionInputBoolNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = output_
	n.parameterValueData = parameterValueData_

	if len(n.parameterValueData) <= 1 {
		n.parameterValueData = append(n.parameterValueData, "")
	}
	if n.parameterValueData[1] == "" {
		n.parameterValueData[1] = "0"
	}
	n.output[0].FriendlyName = n.parameterValueData[0]
	function.AddInput(n.parameterValueData[0], n.parameterValueData[1], n.functionName)
}

func (n *FunctionInputBoolNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
func (n *FunctionInputBoolNode) Clone() *FunctionInputBoolNode {
	clonedOutput := make([]InputNodeHandle, len(n.output))
	copy(clonedOutput, n.output)

	clonedParamValues := make([]string, len(n.parameterValueData))
	copy(clonedParamValues, n.parameterValueData)

	return &FunctionInputBoolNode{
		id:                 n.id,
		nodeType:           n.nodeType,
		output:             clonedOutput,
		parameterValueData: clonedParamValues,
	}
}
