package nodes

import (
	"SoftPLC/function"
	"strconv"
)

type FunctionOutputBoolNode struct {
	id           int
	nodeType     string
	input        []OutputNodeHandle
	functionName string
}

func (d *FunctionOutputBoolNode) GiveFunctionName(name string) {
	d.functionName = name
}

func (d *FunctionOutputBoolNode) GetId() int {
	return d.id
}

func (d *FunctionOutputBoolNode) GetNodeType() string {
	return d.nodeType
}

func (d *FunctionOutputBoolNode) InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle) {
	d.id = id_
	d.nodeType = nodeType_
	d.input = output_
}

func init() {
	var nameServices []string
	for i := 9; i <= 16; i++ {
		nameServices = append(nameServices, "DO"+strconv.Itoa(i-8))
	}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var digitalOutputDescription = nodeDescription{
		AccordionName: "Functions",
		PrimaryType:   "outputNode",
		Type_:         "functionOutputBool",
		Display:       " Output Bool ",
		Label:         "function Output Bool",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
		Output:        []dataTypeNameStruct{},
	}
	RegisterNodeCreator("functionOutputBool", func() (Node, error) {
		return &FunctionOutputBoolNode{
			id:       -1,
			nodeType: "",
			input:    nil,
		}, nil
	}, digitalOutputDescription)
}

func (d *FunctionOutputBoolNode) GetOutput(outName string) *OutputNodeHandle {
	for i, name := range d.input {
		if name.OutputHandle.Name == outName {
			if d.input[i].FriendlyName == "default" {
				d.input[i].FriendlyName = d.input[i].Service //save Name
				function.AddOutput(d.input[i].Service, "0", d.functionName)
			}
			return &d.input[i]
		}
	}
	return nil
}

func (d *FunctionOutputBoolNode) GetOutputList() []OutputNodeHandle {
	return d.input
}
