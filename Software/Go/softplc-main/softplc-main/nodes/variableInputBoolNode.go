package nodes

import (
	"SoftPLC/variable"
)

type VariableInputBoolNode struct {
	id                 int
	nodeType           string
	output             []InputNodeHandle
	parameterValueData []string //select Appliance name + value +
}

/*
type Body struct {
	Di   []bool    `json:"di"`
	Do   []bool    `json:"do"`
	Ai   []float64 `json:"ai"`
	Ao   []float64 `json:"ao"`
	Temp []float64 `json:"temp"`
}*/

func (n *VariableInputBoolNode) GetNodeType() string {
	return n.nodeType
}

func (n *VariableInputBoolNode) GetId() int {
	return n.id
}

func init() {
	/*
		resp, _ := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var input Body
		json.Unmarshal(body, &input)
	*/
	var nameServices []string
	nameServices = []string{"bool", "number", "string"}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var variableInputDescription = nodeDescription{
		AccordionName:     "Variables",
		PrimaryType:       "inputNode",
		Type_:             "variableInputBool",
		Display:           "Input bool ",
		Label:             "variable Input bool",
		Stretchable:       false,
		Services:          services,
		SubServices:       []subServicesStruct{},
		Input:             []dataTypeNameStruct{},
		Output:            []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
		ParameterNameData: []string{"name", "default value"},
	}
	RegisterNodeCreator("variableInputBool", func() (Node, error) {
		return &VariableInputBoolNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, variableInputDescription)
}

func (n *VariableInputBoolNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle, parameterValueData_ []string) {
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
	variable.AddInput(n.parameterValueData[0], n.parameterValueData[1])

}

func (n *VariableInputBoolNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
