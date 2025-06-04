package nodes

import (
	"strconv"
)

type AnalogueOutputNode struct {
	id       int
	nodeType string
	input    []OutputNodeHandle
}

func (a *AnalogueOutputNode) GetId() int {
	return a.id
}

func (a *AnalogueOutputNode) GetNodeType() string {
	return a.nodeType
}

func (a *AnalogueOutputNode) InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle) {
	a.id = id_
	a.nodeType = nodeType_
	a.input = output_
}

func (a *AnalogueOutputNode) GetOutput(outName string) *OutputNodeHandle {
	for i, name := range a.input {
		if name.OutputHandle.Name == outName {
			return &a.input[i]
		}
	}
	return nil
}

func (a *AnalogueOutputNode) GetOutputList() []OutputNodeHandle {
	return a.input
}

func init() {
	/*
		resp, _ := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var input Body
		json.Unmarshal(body, &input)
	*/
	var nameServices []string
	//for i := range input.Ao {
	for i := 19; i <= 20; i++ {
		nameServices = append(nameServices, "AO"+strconv.Itoa(i-18))
	}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}

	var analogueOutputDescription = nodeDescription{
		AccordionName: "Output",
		PrimaryType:   "outputNode",
		Type_:         "analogueOutput",
		Display:       "Analogue output",
		Label:         "Output",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{{DataType: "value", Name: "Input"}},
		Output:        []dataTypeNameStruct{},
	}
	RegisterNodeCreator("analogueOutput", func() (Node, error) {
		return &AnalogueOutputNode{
			id:       -1,
			nodeType: "",
			input:    nil,
		}, nil
	}, analogueOutputDescription)
}
