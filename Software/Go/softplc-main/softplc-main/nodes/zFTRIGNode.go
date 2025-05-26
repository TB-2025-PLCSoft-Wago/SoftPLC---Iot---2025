// Package nodes This package contains the definition of the FtrigNode struct, which is a node that performs the NOT operation on 1 input signals.
package nodes

import "fmt"

// FtrigNode struct for a one input Ftrig node
type FtrigNode struct {
	id        int
	nodeType  string
	input     []InputHandle
	output    []OutputHandle
	prevInput float64
}

var ftrigDescription = nodeDescription{
	AccordionName: "Edge Detection",
	PrimaryType:   "LogicalNode",
	Type_:         "FtrigNode",
	Display:       "Ftrig Node",
	Label:         "Ftrig",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
	Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	fmt.Println("Registering FTRIGNode")
	RegisterNodeCreator("FtrigNode", func() (Node, error) {
		return &FtrigNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, ftrigDescription)
}

func (n *FtrigNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = 0
		return
	}
	if *n.input[0].Input == n.prevInput {
		n.output[0].Output = 0
	} else if *n.input[0].Input == 0 {
		n.output[0].Output = 1
	}

	n.prevInput = *n.input[0].Input
}
func (n *FtrigNode) GetNodeType() string {
	return n.nodeType
}

func (n *FtrigNode) GetId() int {
	return n.id
}

func (n *FtrigNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *FtrigNode) GetInput() []InputHandle {
	return n.input
}

func (n *FtrigNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
