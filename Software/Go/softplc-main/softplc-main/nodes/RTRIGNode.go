// Package nodes This package contains the definition of the RtrigNode struct, which is a node that performs the NOT operation on 1 input signals.
package nodes

// RtrigNode struct for a one input Rtrig node
type RtrigNode struct {
	id        int
	nodeType  string
	input     []InputHandle
	output    []OutputHandle
	prevInput string
}

var rtrigDescription = nodeDescription{
	AccordionName: "Edge Detection",
	PrimaryType:   "LogicalNode",
	Type_:         "RtrigNode",
	Display:       "Rtrig Node",
	Label:         "Rtrig",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
	Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("RtrigNode", func() (Node, error) {
		return &RtrigNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, rtrigDescription)
}

func (n *RtrigNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	if *n.input[0].Input == n.prevInput {
		n.output[0].Output = "0"
	} else if *n.input[0].Input == "1" {
		n.output[0].Output = "1"
	}

	n.prevInput = *n.input[0].Input
}
func (n *RtrigNode) GetNodeType() string {
	return n.nodeType
}

func (n *RtrigNode) GetId() int {
	return n.id
}

func (n *RtrigNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *RtrigNode) GetInput() []InputHandle {
	return n.input
}

func (n *RtrigNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *RtrigNode) DestroyToBuildAgain() {

}
