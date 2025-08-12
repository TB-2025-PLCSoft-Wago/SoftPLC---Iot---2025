package nodes

import (
	"SoftPLC/function"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type functionNodeDescription struct {
	name        string //Node Source
	description nodeDescription
	graph       interface{}
}
type functionInOut struct {
	name  string //exemple : xEnable
	value string
}

var FunctionGraphJson interface{}
var FunctionNodeList []functionNodeDescription
var functionInputBool []functionInOut
var functionOutputBool []functionInOut

//FunctionNodeList[0].description.Input =  Function input
//FunctionNodeList[0].description.Input =  Function output

func FunctionAddToList() {
	name := FunctionGraphJson.(map[string]interface{})["name"].(string)
	//drop .json
	if len(name) >= 5 {
		name = name[:len(name)-5]
	}
	newFuncNode := functionNodeDescription{
		name: name,
		description: nodeDescription{
			AccordionName:     "Functions",
			PrimaryType:       "LogicalNode",
			Display:           name,
			Type_:             name + "ConfigurableNodeFunction",
			Label:             name,
			Stretchable:       false,
			Services:          []servicesStruct{},
			SubServices:       []subServicesStruct{},
			ParameterNameData: nil,

			Input:  []dataTypeNameStruct{},
			Output: []dataTypeNameStruct{},
		},
	}

	graphData, ok := FunctionGraphJson.(map[string]interface{})["data"]
	if !ok {
		log.Println("Format invalide pour FunctionGraphJson")
		return
	}
	/*
		for n := range graphData.(map[string]interface{})["node"]{
			if strings.Contains(n.getType, "functionInputBool"){
				newFuncNode.description.Input = append( newFuncNode.description.Input,{DataType: "bool", Name: n.parameterValueData[0]})
			}
		}*/
	/*for n := range graphData.(map[string]interface{})["node"] {
		d := n.(map[string]interface{})["data"]
		if strings.Contains(d.(map[string]interface{})["type"], "functionInput") {
			newFuncNode.description.Input = append(newFuncNode.description.Input,
			{
				DataType: d.(map[string]interface{})["outputHandle"], Name: d.(map[string]interface{})["parameterValueData"][0]
			})
		}
	}*/
	nodesRaw, ok := graphData.(map[string]interface{})["nodes"]
	if !ok {
		log.Println("Pas de noeuds dans le graphe")
		return
	}

	nodesArray, ok := nodesRaw.([]interface{})
	if !ok {
		log.Println("Les noeuds ne sont pas un tableau")
		return
	}

	for _, node := range nodesArray {
		nodeMap, ok := node.(map[string]interface{})
		if !ok {
			continue
		}
		data, ok := nodeMap["data"].(map[string]interface{})
		if !ok {
			continue
		}

		nodeType, _ := data["type"].(string)
		if strings.Contains(nodeType, "functionInputBool") {
			paramList, _ := data["parameterValueData"].([]interface{})
			if len(paramList) > 0 {
				paramName, _ := paramList[0].(string)
				newFuncNode.description.Input = append(newFuncNode.description.Input, dataTypeNameStruct{
					DataType: "bool",
					Name:     paramName,
				})
			}
		}
		if strings.Contains(nodeType, "functionOutputBool") {
			paramName, _ := data["selectedServiceData"].(string)
			newFuncNode.description.Output = append(newFuncNode.description.Output, dataTypeNameStruct{
				DataType: "bool",
				Name:     paramName,
			})

		}

		if strings.Contains(nodeType, "functionInputValue") {
			paramList, _ := data["parameterValueData"].([]interface{})
			if len(paramList) > 0 {
				paramName, _ := paramList[0].(string)
				newFuncNode.description.Input = append(newFuncNode.description.Input, dataTypeNameStruct{
					DataType: "value",
					Name:     paramName,
				})
			}
		}
		if strings.Contains(nodeType, "functionOutputValue") {
			paramName, _ := data["selectedServiceData"].(string)
			newFuncNode.description.Output = append(newFuncNode.description.Output, dataTypeNameStruct{
				DataType: "value",
				Name:     paramName,
			})

		}
	}

	newFuncNode.graph = graphData

	/*err := CreateFunctionQueue(name, graphData)
	if err != nil {
		serverResponse.ResponseProcessGraph = "function " + name + " : error Graph"
		return
	}*/
	FunctionNodeList = append(FunctionNodeList, newFuncNode)
	update()
}

// FunctionNode struct for a one input Function node
type FunctionNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
	isInit   bool
}

func init() {
	for i := range FunctionNodeList {
		RegisterNodeCreator(FunctionNodeList[i].description.Type_, func() (Node, error) {
			return &FunctionNode{
				id:       -1,
				nodeType: "",
				input:    nil,
				output:   nil,
			}, nil
		}, FunctionNodeList[i].description)
	}
}

func update() {
	for i := range FunctionNodeList {
		RegisterNodeCreator(FunctionNodeList[i].description.Type_, func() (Node, error) {
			return &FunctionNode{
				id:       -1,
				nodeType: "",
				input:    nil,
				output:   nil,
			}, nil
		}, FunctionNodeList[i].description)
	}
}

func (n *FunctionNode) ProcessLogic() {

	name := strings.ReplaceAll(n.nodeType, "ConfigurableNodeFunction", "")
	if !n.isInit {
		for i := range FunctionNodeList {
			if name == FunctionNodeList[i].name {

				/*graphData, ok := FunctionNodeList[i].graph.(map[string]interface{})["data"]
				if !ok {
					fmt.Println("Function error graphData : ", name)
					return
				}*/
				// call echo.CreateFunctionQueue
				tempName := name + strconv.Itoa(n.id)
				if err := CreatorEcho(tempName, FunctionNodeList[i].graph); err != nil {
					fmt.Println("Queue creation failed:", err)
					return
				}
				break
			}
		}

		n.isInit = true
	}
	inputIndex := 0
	for _, input := range FunctionNodeListProcess[name+strconv.Itoa(n.id)].InputNodes {
		nodeType := input.GetNodeType()
		if strings.Contains(nodeType, "function") {
			if len(n.input) > inputIndex {
				function.UpdateFunctionInput(input.GetOutput("Output").FriendlyName, *n.input[inputIndex].Input, name+strconv.Itoa(n.id)) //TO DO : edge.TargetHandle
				inputIndex++
			} else {
				fmt.Println("nodeFunctions input error index : ", inputIndex)
			}

		}
	}

	ProcessFunction(name + strconv.Itoa(n.id)) //name + "ConfigurableNodeFunction"
	inputIndex = 0
	for _, out := range FunctionNodeListProcess[name+strconv.Itoa(n.id)].OutputNodes {
		nodeType := out.GetNodeType()
		if strings.Contains(nodeType, "function") {
			//TO DO : edge.TargetHandle
			n.output[inputIndex].Output = function.GetFunctionOutput(out.GetOutput("Input").Service, name+strconv.Itoa(n.id))
			inputIndex++
		}
	}
}
func (n *FunctionNode) GetNodeType() string {
	return n.nodeType
}

func (n *FunctionNode) GetId() int {
	return n.id
}

func (n *FunctionNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *FunctionNode) GetInput() []InputHandle {
	return n.input
}

func (n *FunctionNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
}
func (n *FunctionNode) DestroyToBuildAgain() {

}
