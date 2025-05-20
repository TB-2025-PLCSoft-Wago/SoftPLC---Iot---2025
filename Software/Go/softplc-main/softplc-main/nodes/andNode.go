// Package nodes This package contains the definition of the AndNode struct, which is a node that performs the AND operation on two input signals.
package nodes

// AndNode struct for a two input AND node
type AndNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var andDescription = nodeDescription{
	AccordionName: "Logical gate",
	PrimaryType:   "LogicalNode",
	Type_:         "AndNode",
	Display:       "And Node",
	Label:         "&",
	Stretchable:   true,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input:         []dataTypeNameStruct{{DataType: "bool", Name: "Input"}},
	Output:        []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("AndNode", func() (Node, error) {
		return &AndNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, andDescription)
}

func (n *AndNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = 0
		return
	}
	for _, in := range n.input {
		if *in.Input == 0 {
			n.output[0].Output = 0
			return
		}
	}
	n.output[0].Output = 1
}
func (n *AndNode) GetNodeType() string {
	return n.nodeType
}

func (n *AndNode) GetId() int {
	return n.id
}

func (n *AndNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *AndNode) GetInput() []InputHandle {
	return n.input
}

func (n *AndNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
