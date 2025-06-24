package nodes

type ConstantInputNode struct {
	id       int
	nodeType string
	output   []InputNodeHandle
}

func (c *ConstantInputNode) GetId() int {
	return c.id
}

func (c *ConstantInputNode) GetNodeType() string {
	return c.nodeType
}

func (c *ConstantInputNode) InitNode(id_ int, nodeType_ string, input_ []InputNodeHandle, parameterValueData_ []string) {
	c.id = id_
	c.nodeType = nodeType_
	c.output = input_
}

func (c *ConstantInputNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range c.output {
		if name.InputHandle.Name == outName {
			return &c.output[i]
		}
	}
	return nil
}

func init() {
	RegisterNodeCreator("constantInput", func() (Node, error) {
		return &ConstantInputNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, nodeDescription{
		AccordionName: "Constant",
		PrimaryType:   "inputNode",
		Type_:         "constantInput",
		Display:       "Constant value",
		Label:         "Constant value Input",
		Stretchable:   false,
		Services:      []servicesStruct{},
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{},
		Output:        []dataTypeNameStruct{{DataType: "value", Name: "Output"}},
	})
}
