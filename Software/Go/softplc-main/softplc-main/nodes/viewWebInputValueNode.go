package nodes

import "SoftPLC/server"

type ViewWebInputValueNode struct {
	id                 int
	nodeType           string
	output             []InputNodeHandle
	parameterValueData []string //select Appliance name + value +
}

func (n *ViewWebInputValueNode) GetNodeType() string {
	return n.nodeType
}

func (n *ViewWebInputValueNode) GetId() int {
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
	nameServices = []string{"value", "number", "string"}

	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var viewWebInputDescription = nodeDescription{
		AccordionName:     "View web",
		PrimaryType:       "inputNode",
		Type_:             "viewWebInputValue",
		Display:           "Input value",
		Label:             "view Input value",
		Stretchable:       false,
		Services:          services,
		SubServices:       []subServicesStruct{},
		Input:             []dataTypeNameStruct{},
		Output:            []dataTypeNameStruct{{DataType: "value", Name: "Output"}},
		ParameterNameData: []string{"appliance name", "signal name"},
	}
	RegisterNodeCreator("viewWebInputValue", func() (Node, error) {
		return &ViewWebInputValueNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, viewWebInputDescription)
}

func (n *ViewWebInputValueNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.output[0].FriendlyName = server.AddInputToAppliance(n.parameterValueData[0], n.parameterValueData[1], "string") //save the irCode

}

func (n *ViewWebInputValueNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
