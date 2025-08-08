package nodes

import "strings"

// DeleteAndShowFirstElemNode : Provide single string from multiple
type DeleteAndShowFirstElemNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
}

var deleteAndShowFirstElemDescription = nodeDescription{
	AccordionName: "Handling value",
	PrimaryType:   "LogicalNode",
	Type_:         "DeleteAndShowFirstElemNode",
	Display:       "DeleteAndShowFirstElem",
	Label:         "D+S1",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "value", Name: "arrayIn"},
		{DataType: "value", Name: "splitter"},
	},
	Output: []dataTypeNameStruct{{DataType: "value", Name: "arrayOut"}, {DataType: "value", Name: "1er Elem"}},
}

func init() {
	RegisterNodeCreator("DeleteAndShowFirstElemNode", func() (Node, error) {
		return &DeleteAndShowFirstElemNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, deleteAndShowFirstElemDescription)
}

func (n *DeleteAndShowFirstElemNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = ""
		return
	}
	if n.input[0].Input == nil {
		n.output[0].Output = ""
		n.output[1].Output = ""
		return
	}
	var splitter string
	if n.input[1].Input == nil {
		splitter = " ,, "
	} else {
		splitter = *n.input[1].Input
	}
	inputOrder := strings.Split(*n.input[0].Input, splitter)
	n.output[1].Output = inputOrder[0]

	if len(inputOrder) > 1 {
		inputOrder = inputOrder[1:]
	} else {
		inputOrder = []string{}
	}

	n.output[0].Output = strings.Join(inputOrder, splitter)
}

func (n *DeleteAndShowFirstElemNode) GetNodeType() string {
	return n.nodeType
}

func (n *DeleteAndShowFirstElemNode) GetId() int {
	return n.id
}

func (n *DeleteAndShowFirstElemNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *DeleteAndShowFirstElemNode) GetInput() []InputHandle {
	return n.input
}

func (n *DeleteAndShowFirstElemNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *DeleteAndShowFirstElemNode) DestroyToBuildAgain() {

}
