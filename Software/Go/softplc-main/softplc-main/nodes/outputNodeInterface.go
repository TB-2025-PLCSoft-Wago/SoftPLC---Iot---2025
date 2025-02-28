package nodes

type OutputNodeInterface interface {
	GetId() int
	GetNodeType() string

	InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle)
	GetOutput(outName string) *OutputNodeHandle
	GetOutputList() []OutputNodeHandle
}

type OutputNodeHandle struct {
	FriendlyName string
	Service      string
	SubService   string
	OutputHandle InputHandle
}
