package nodes

import (
	"strconv"
)

type AnalogueInputNode struct {
	id       int
	nodeType string
	output   []InputNodeHandle
}

func (n *AnalogueInputNode) GetNodeType() string {
	return n.nodeType
}

func (n *AnalogueInputNode) GetId() int {
	return n.id
}

func init() {
	/*
		resp, _ := http.Get("http://192.168.37.134:8888/api/v1/hal/io") // marcelin http://192.168.1.175:8888/api/v1/hal/io
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var input Body
		json.Unmarshal(body, &input)
	*/
	var nameServices []string
	//for i := range input.Ai {
	for i := 31; i <= 34; i++ {
		nameServices = append(nameServices, "AI"+strconv.Itoa(i-30))
	}
	//for i := range input.Temp {
	/*
		for i := 29; i <= 30; i++ {
			nameServices = append(nameServices, "TEMP"+strconv.Itoa(i-28))
		}*/
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var digitalInputDescription = nodeDescription{
		AccordionName: "Input",
		PrimaryType:   "inputNode",
		Type_:         "analogueInput",
		Display:       "Analogue Input",
		Label:         "Input analogue",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{},
		Output:        []dataTypeNameStruct{{DataType: "value", Name: "Output"}},
	}
	RegisterNodeCreator("analogueInput", func() (Node, error) {
		return &AnalogueInputNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, digitalInputDescription)
}

func (n *AnalogueInputNode) InitNode(id_ int, nodeType_ string, input_ []InputNodeHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = input_
}

func (n *AnalogueInputNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
