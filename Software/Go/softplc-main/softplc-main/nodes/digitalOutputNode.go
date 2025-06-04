package nodes

import (
	"strconv"
)

type DigitalOutputNode struct {
	id       int
	nodeType string
	input    []OutputNodeHandle
}

func (d *DigitalOutputNode) GetId() int {
	return d.id
}

func (d *DigitalOutputNode) GetNodeType() string {
	return d.nodeType
}

func (d *DigitalOutputNode) InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle) {
	d.id = id_
	d.nodeType = nodeType_
	d.input = output_
}

func init() {
	/*
		resp, _ := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var input Body
		json.Unmarshal(body, &input)*/
	var nameServices []string
	//for i := range input.Do {
	for i := 9; i <= 16; i++ {
		nameServices = append(nameServices, "DO"+strconv.Itoa(i-8))
	}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var digitalOutputDescription = nodeDescription{
		AccordionName: "Output",
		PrimaryType:   "outputNode",
		Type_:         "digitalOutput",
		Display:       "Digital output",
		Label:         "Output",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
		Output:        []dataTypeNameStruct{},
	}
	RegisterNodeCreator("digitalOutput", func() (Node, error) {
		return &DigitalOutputNode{
			id:       -1,
			nodeType: "",
			input:    nil,
		}, nil
	}, digitalOutputDescription)
}

func (d *DigitalOutputNode) GetOutput(outName string) *OutputNodeHandle {
	for i, name := range d.input {
		if name.OutputHandle.Name == outName {
			return &d.input[i]
		}
	}
	return nil
}

func (d *DigitalOutputNode) GetOutputList() []OutputNodeHandle {
	return d.input
}
