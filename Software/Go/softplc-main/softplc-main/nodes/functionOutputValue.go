package nodes

import (
	"SoftPLC/function"
	"strconv"
)

type FunctionOutputValueNode struct {
	id           int
	nodeType     string
	input        []OutputNodeHandle
	functionName string
}

func (d *FunctionOutputValueNode) GiveFunctionName(name string) {
	d.functionName = name
}

func (d *FunctionOutputValueNode) GetId() int {
	return d.id
}

func (d *FunctionOutputValueNode) GetNodeType() string {
	return d.nodeType
}

func (d *FunctionOutputValueNode) InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle) {
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
		AccordionName: "Functions",
		PrimaryType:   "outputNode",
		Type_:         "functionOutputValue",
		Display:       " Output Value ",
		Label:         "function Output Value",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{{DataType: "value", Name: "Input"}},
		Output:        []dataTypeNameStruct{},
	}
	RegisterNodeCreator("functionOutputValue", func() (Node, error) {
		return &FunctionOutputValueNode{
			id:       -1,
			nodeType: "",
			input:    nil,
		}, nil
	}, digitalOutputDescription)
}

func (d *FunctionOutputValueNode) GetOutput(outName string) *OutputNodeHandle {
	for i, name := range d.input {
		if name.OutputHandle.Name == outName {
			if d.input[i].FriendlyName == "default" {
				d.input[i].FriendlyName = d.input[i].Service //save Name
				function.AddOutput(d.input[i].Service, "", d.functionName)
			}
			return &d.input[i]
		}
	}
	return nil
}

func (d *FunctionOutputValueNode) GetOutputList() []OutputNodeHandle {
	return d.input
}
