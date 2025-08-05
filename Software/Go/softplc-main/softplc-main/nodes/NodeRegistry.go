package nodes

import (
	"SoftPLC/function"
	"SoftPLC/inputUpdate"
	"SoftPLC/variable"
	"fmt"
	"sort"
	"strings"
)

var LogicalNode [][]LogicalNodeInterface

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
	AccordionName     string               `json:"accordion"`
	PrimaryType       string               `json:"primaryType"`
	Type_             string               `json:"type"`
	Display           string               `json:"display"`
	Label             string               `json:"label"`
	Stretchable       bool                 `json:"stretchable"`
	Services          []servicesStruct     `json:"services"`
	SubServices       []subServicesStruct  `json:"subServices"`
	Input             []dataTypeNameStruct `json:"inputHandle"`
	Output            []dataTypeNameStruct `json:"outputHandle"`
	ParameterNameData []string             `json:"parameterNameData"`
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
	fmt.Println("Init NodeRegistry")
	nodeCreators = make(map[string]NodeCreator)
	nodeTypeDescriptions = make(map[string]nodeDescription)
	FunctionNodeListProcess = make(map[string]functionProcess)
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

// Use by processGraph
func NodeDescription(type_ string) (nodeDescription, error) {
	desc, ok := nodeTypeDescriptions[type_]
	if !ok {
		return nodeDescription{}, fmt.Errorf("node type %s not found", type_)
	}
	return desc, nil
}

// Use to append node to processGraph slice (OutputNodes, InputNodes, LogicalNode
func CreateNode(type_ string) (Node, error) {
	creator, ok := nodeCreators[type_]
	if !ok {
		return nil, fmt.Errorf("node type %s not found", type_)
	}
	return creator()
}

func SystemDescription() (FinalJson, error) {
	var finalJson FinalJson
	keys := make([]string, 0, len(nodeTypeDescriptions))
	for k := range nodeTypeDescriptions {
		keys = append(keys, k) //Prepare the keys to be sorted
	}
	sort.Strings(keys)
	for _, v := range keys {
		finalJson.Nodes = append(finalJson.Nodes, nodeTypeDescriptions[v])
	}
	return finalJson, nil
}

/*** Functions Nodes ***/

type functionProcess struct {
	//nameFunction string
	LogicalNode [][]LogicalNodeInterface
	OutputNodes []OutputNodeInterface
	InputNodes  []InputNodeInterface
	//ConstValue  []processGraph.Const
}

var FunctionNodeListProcess map[string]functionProcess

func ProcessFunction(name string) {
	if len(FunctionNodeListProcess[name].OutputNodes) != 0 {
		inputUpdate.UpdateInputs() // TO DO : DELETE
		for _, v := range FunctionNodeListProcess[name].LogicalNode {
			variable.UpdateVariableInputs()
			for _, n := range v {
				if logicalNode, ok := n.(LogicalNodeInterface); ok {
					logicalNode.ProcessLogic()
				}
			}
			//outputUpdate.UpdateVariables()
			updateVariables(name)
		}
		/*
			outputUpdate.UpdateOutput()
			server.UpdateOutputValueView()
			if server.IsActiveDebug {
				server.DebugMode()
			}
		*/

		//variable.UpdateVariableInputs()
		updateFunctionOutputs(name)
	}
}
func updateVariables(name string) {
	for _, output := range FunctionNodeListProcess[name].OutputNodes {
		for _, nodeOutputList := range output.GetOutputList() {
			/*** Variable output ***/
			if strings.Contains(output.GetNodeType(), "variable") {
				for _, ccIOState := range variable.OutputsStateVariable {
					if ccIOState.Name == nodeOutputList.FriendlyName {
						if ccIOState.Value != *nodeOutputList.OutputHandle.Input {
							//ccIOState.Value = *nodeOutputList.OutputHandle.Input
							variable.UpdateVariableOutput(ccIOState.Name, *nodeOutputList.OutputHandle.Input)
						}
					}
				}
				continue
			}
		}
	}
}
func updateFunctionOutputs(name string) {
	for _, output := range FunctionNodeListProcess[name].OutputNodes {
		for _, nodeOutputList := range output.GetOutputList() {
			/*** Function output ***/
			if strings.Contains(output.GetNodeType(), "function") {
				for _, ccIOState := range function.OutputsStateFunction {
					if ccIOState.Name == nodeOutputList.FriendlyName {
						if ccIOState.Value != *nodeOutputList.OutputHandle.Input {
							//ccIOState.Value = *nodeOutputList.OutputHandle.Input
							function.UpdateFunctionOutput(ccIOState.Name, *nodeOutputList.OutputHandle.Input)
							fmt.Println("updateFunctionOutputs")
						}
					}
				}
				continue
			}
		}
	}
}
