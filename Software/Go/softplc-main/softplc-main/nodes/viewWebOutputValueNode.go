package nodes

import (
	"SoftPLC/server"
	"strconv"
)

type ViewWebOutputValueNode struct {
	id       int
	nodeType string
	input    []OutputNodeHandle
}

func (d *ViewWebOutputValueNode) GetId() int {
	return d.id
}

func (d *ViewWebOutputValueNode) GetNodeType() string {
	return d.nodeType
}

func (d *ViewWebOutputValueNode) InitNode(id_ int, nodeType_ string, output_ []OutputNodeHandle) {
	d.id = id_
	d.nodeType = nodeType_
	d.input = output_
}

func init() {
	/*
		resp, _ := http.Get("http://192.168.37.134:8888/api/v1/hal/io")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var input Body
		json.Unmarshal(body, &input)*/
	var nameServices []string
	//for i := range input.Do {
	for i := 9; i <= 16; i++ {
		nameServices = append(nameServices, "DO"+strconv.Itoa(i-8))
	}
	services := []servicesStruct{{FriendlyName: "", NameServices: nameServices}}
	var digitalOutputDescription = nodeDescription{
		AccordionName: "View web",
		PrimaryType:   "outputNode",
		Type_:         "viewWebOutputValue",
		Display:       "Output value",
		Label:         "view Output Value",
		Stretchable:   false,
		Services:      services,
		SubServices:   []subServicesStruct{},
		Input:         []dataTypeNameStruct{{DataType: "value", Name: "Input"}},
		Output:        []dataTypeNameStruct{},
	}
	RegisterNodeCreator("viewWebOutputValue", func() (Node, error) {
		return &ViewWebOutputValueNode{
			id:       -1,
			nodeType: "",
			input:    nil,
		}, nil
	}, digitalOutputDescription)
}

func (d *ViewWebOutputValueNode) GetOutput(outName string) *OutputNodeHandle {
	for i, name := range d.input {
		if name.OutputHandle.Name == outName {
			if d.input[i].FriendlyName == "default" {
				d.input[i].FriendlyName = server.AddOutputToAppliance(d.input[i].Service, d.input[i].SubService, "value", "")
			}
			return &d.input[i]
		}
	}
	return nil
}

func (d *ViewWebOutputValueNode) GetOutputList() []OutputNodeHandle {
	return d.input
}
