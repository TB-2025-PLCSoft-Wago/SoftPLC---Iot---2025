// Package nodes This package contains the definition of the RFtrigNode struct, which is a node that performs the NOT operation on 1 input signals.
package nodes

// RFtrigNode struct for a one input RFtrig node
type RFtrigNode struct {
	id        int
	nodeType  string
	input     []InputHandle
	output    []OutputHandle
	prevInput string
}

var rftrigDescription = nodeDescription{
	AccordionName: "Edge Detection",
	PrimaryType:   "LogicalNode",
	Type_:         "RFtrigNode",
	Display:       "RF_trig Node",
	Label:         "RF_trig",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
	Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("RFtrigNode", func() (Node, error) {
		return &RFtrigNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, rftrigDescription)
}

func (n *RFtrigNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	if *n.input[0].Input == n.prevInput {
		n.output[0].Output = "0"
	} else {
		n.output[0].Output = "1"
	}

	n.prevInput = *n.input[0].Input

}
func (n *RFtrigNode) GetNodeType() string {
	return n.nodeType
}

func (n *RFtrigNode) GetId() int {
	return n.id
}

func (n *RFtrigNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *RFtrigNode) GetInput() []InputHandle {
	return n.input
}

func (n *RFtrigNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
