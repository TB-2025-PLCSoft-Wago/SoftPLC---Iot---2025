package nodes

import (
	"encoding/json"
	"io"
	"net/http"
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
	resp, _ := http.Get("http://192.168.1.175:8888/api/v1/hal/io")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var input Body
	json.Unmarshal(body, &input)
	var nameServices []string
	for i := range input.Ai {
		nameServices = append(nameServices, "AI"+strconv.Itoa(i+1))
	}
	for i := range input.Temp {
		nameServices = append(nameServices, "TEMP"+strconv.Itoa(i+1))
	}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var digitalInputDescription = nodeDescription{
		AccordionName: "Input",
		PrimaryType:   "inputNode",
		Type_:         "analogueInput",
		Display:       "Analogue Input",
		Label:         "Input",
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

func (n *AnalogueInputNode) InitNode(id_ int, nodeType_ string, input_ []InputNodeHandle) {
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
