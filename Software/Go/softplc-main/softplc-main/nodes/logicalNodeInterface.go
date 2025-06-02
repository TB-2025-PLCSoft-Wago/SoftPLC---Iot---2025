package nodes

// LogicalNodeInterface is an interface for logical nodes
type LogicalNodeInterface interface {

	// Node interface
	GetId() int
	GetNodeType() string
	//LogicalNode methodes
	InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string)
	ProcessLogic()
	GetOutput(outName string) *OutputHandle
	GetInput() []InputHandle
	DestroyToBuildAgain()
}
