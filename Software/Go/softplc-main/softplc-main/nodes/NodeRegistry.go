package nodes

import (
	"fmt"
)

type servicesStruct struct {
	FriendlyName string   `json:"friendlyName"`
	NameServices []string `json:"nameServices"`
}
type dataTypeNameStruct struct {
	DataType string `json:"dataType"`
	Name     string `json:"name"`
}
type subServicesStruct struct {
	FriendlyName string               `json:"friendlyName"`
	Primary      string               `json:"primary"`
	Secondary    []dataTypeNameStruct `json:"secondary"`
}
type nodeDescription struct {
	AccordionName string               `json:"accordion"`
	PrimaryType   string               `json:"primaryType"`
	Type_         string               `json:"type"`
	Display       string               `json:"display"`
	Label         string               `json:"label"`
	Stretchable   bool                 `json:"stretchable"`
	Services      []servicesStruct     `json:"services"`
	SubServices   []subServicesStruct  `json:"subServices"`
	Input         []dataTypeNameStruct `json:"inputHandle"`
	Output        []dataTypeNameStruct `json:"outputHandle"`
}

type FinalJson struct {
	Nodes []nodeDescription `json:"nodes"`
}

type (
	NodeCreator func() (Node, error)
)

var nodeCreators map[string]NodeCreator
var nodeTypeDescriptions map[string]nodeDescription

func init() {
	nodeCreators = make(map[string]NodeCreator)
	nodeTypeDescriptions = make(map[string]nodeDescription)
}

func RegisterNodeCreator(name string, creator NodeCreator, description nodeDescription) {
	nodeCreators[name] = creator
	nodeTypeDescriptions[name] = description
}

func NodeTypes() []string {
	types := make([]string, 0, len(nodeCreators))
	for k := range nodeCreators {
		types = append(types, k)
	}
	return types
}

func NodeDescription(type_ string) (nodeDescription, error) {
	desc, ok := nodeTypeDescriptions[type_]
	if !ok {
		return nodeDescription{}, fmt.Errorf("node type %s not found", type_)
	}
	return desc, nil
}

func CreateNode(type_ string) (Node, error) {
	creator, ok := nodeCreators[type_]
	if !ok {
		return nil, fmt.Errorf("node type %s not found", type_)
	}
	return creator()
}

func SystemDescription() (FinalJson, error) {
	var finalJson FinalJson
	for _, v := range nodeTypeDescriptions {
		finalJson.Nodes = append(finalJson.Nodes, v)
	}
	return finalJson, nil
}
