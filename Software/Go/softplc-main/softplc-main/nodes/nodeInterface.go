package nodes

// Node is an interface that all nodes must implement
type Node interface {
	GetNodeType() string
	GetId() int
}

type InputHandle struct {
	Input    *string
	Name     string
	DataType string
}

type OutputHandle struct {
	Output   string
	Name     string
	DataType string
}
