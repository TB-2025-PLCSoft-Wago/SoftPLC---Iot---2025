package nodes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type DigitalInputNode struct {
	id       int
	nodeType string
	output   []InputNodeHandle
}

type Body struct {
	Di   []bool    `json:"di"`
	Do   []bool    `json:"do"`
	Ai   []float64 `json:"ai"`
	Ao   []float64 `json:"ao"`
	Temp []float64 `json:"temp"`
}

func (n *DigitalInputNode) GetNodeType() string {
	return n.nodeType
}

func (n *DigitalInputNode) GetId() int {
	return n.id
}

func init() {
	resp, _ := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var input Body
	json.Unmarshal(body, &input)
	var nameServices []string
	for i := range input.Di {
		nameServices = append(nameServices, "DI"+strconv.Itoa(i+1))
	}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var digitalInputDescription = nodeDescription{
		AccordionName: "Input",
		PrimaryType:   "inputNode",
		Type_:         "digitalInput",
		Display:       "Digital Input",
		Label:         "Input",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{},
		Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
	}
	RegisterNodeCreator("digitalInput", func() (Node, error) {
		return &DigitalInputNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, digitalInputDescription)
}

func (n *DigitalInputNode) InitNode(id_ int, nodeType_ string, output_ []InputNodeHandle) {
	n.id = id_
	n.nodeType = nodeType_
	n.output = output_
}

func (n *DigitalInputNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range n.output {
		if name.InputHandle.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}
