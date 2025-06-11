package nodes

type InputNodeInterface interface {
	GetId() int
	GetNodeType() string

	InitNode(id_ int, nodeType_ string, input_ []InputNodeHandle, parameterValueData_ []string)
	GetOutput(outName string) *InputNodeHandle
}

type InputNodeHandle struct {
	FriendlyName string
	Service      string
	SubService   string
	InputHandle  InputHandle
}
