package nodes

import (
	"SoftPLC/variable"
	"strconv"
)

type VariableOutputValueNode struct {
	id       int
	nodeType string
	input    []OutputNodeHandle
}

func (d *VariableOutputValueNode) GetId() int {
	return d.id
}

func (d *VariableOutputValueNode) GetNodeType() string {
	return d.nodeType
}

func (d *VariableOutputValueNode) InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle) {
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
		AccordionName: "Variables",
		PrimaryType:   "outputNode",
		Type_:         "variableOutputValue",
		Display:       "Output Value ",
		Label:         "variable Output Value",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{{DataType: "value", Name: "Input"}},
		Output:        []dataTypeNameStruct{},
	}
	RegisterNodeCreator("variableOutputValue", func() (Node, error) {
		return &VariableOutputValueNode{
			id:       -1,
			nodeType: "",
			input:    nil,
		}, nil
	}, digitalOutputDescription)
}

func (d *VariableOutputValueNode) GetOutput(outName string) *OutputNodeHandle {
	for i, name := range d.input {
		if name.OutputHandle.Name == outName {
			if d.input[i].FriendlyName == "default" {
				d.input[i].FriendlyName = d.input[i].Service //save Name
				variable.AddOutput(d.input[i].Service, "")
			}
			return &d.input[i]
		}
	}
	return nil
}

func (d *VariableOutputValueNode) GetOutputList() []OutputNodeHandle {
	return d.input
}
