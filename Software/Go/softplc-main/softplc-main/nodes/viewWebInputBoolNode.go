package nodes

import (
	"SoftPLC/server"
	"SoftPLC/serverResponse"
)

type ViewWebInputBoolNode struct {
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

func (n *ViewWebInputBoolNode) GetNodeType() string {
	return n.nodeType
}

func (n *ViewWebInputBoolNode) GetId() int {
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
	var viewWebInputDescription = nodeDescription{
		AccordionName:     "View web",
		PrimaryType:       "inputNode",
		Type_:             "viewWebInputBool",
		Display:           "Input bool",
		Label:             "view Input bool",
		Stretchable:       false,
		Services:          services,
		SubServices:       []subServicesStruct{},
		Input:             []dataTypeNameStruct{},
		Output:            []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
		ParameterNameData: []string{"appliance name", "signal name"},
	}
	RegisterNodeCreator("viewWebInputBool", func() (Node, error) {
		return &ViewWebInputBoolNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, viewWebInputDescription)
}

func (n *ViewWebInputBoolNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = output_
	n.parameterValueData = parameterValueData_
	if n.parameterValueData[1] == "" {
		serverResponse.ResponseProcessGraph = "view Input Bool - signal name : empty"
	}
	if n.parameterValueData[0] == "" {
		serverResponse.ResponseProcessGraph = "view Input Bool - appliance name : empty"
	}

	n.output[0].FriendlyName = server.AddInputToAppliance(n.parameterValueData[0], n.parameterValueData[1], "bool") //save the irCode // to find where use in processGraph.go : if strconv.Itoa(inputLink.IRCode) == inputHandle.FriendlyName {

}

func (n *ViewWebInputBoolNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
