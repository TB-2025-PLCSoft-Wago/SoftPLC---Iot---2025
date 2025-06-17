package processGraph

import (
	"SoftPLC/inputUpdate"
	"SoftPLC/nodes"
	"SoftPLC/server"
	"SoftPLC/serverResponse"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var Mutex sync.Mutex

// nodeJson is a struct that represents a node in the JSON file.
type nodeJson struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Data data_  `json:"data"`
}

// data_ is a struct that represents the data of a node in the JSON file.
type data_ struct {
	FriendlyName       string   `json:"friendlyName"`
	Service            string   `json:"service"`
	SubService         string   `json:"subService"`
	Value              string   `json:"value"`
	ParameterValueData []string `json:"parameterValueData"`
}

// edge is a struct that represents an edge in the JSON file an edge is the link between two nodes.
type edge struct {
	Source       string `json:"source"`
	Target       string `json:"target"`
	SourceHandle string `json:"sourceHandle"`
	TargetHandle string `json:"targetHandle"`
}

// Graph is a struct that represents the JSON file.
type Graph struct {
	NodesJson []nodeJson `json:"nodes"`
	Edges     []edge     `json:"edges"`
}

type Const struct {
	Id    int
	Value float64
}

// LogicalNode is a slice of slice of nodes.Node tha contains all the nodes in the JSON file except the output/input nodes.
var LogicalNode [][]nodes.LogicalNodeInterface

// OutputNodes is a slice of nodes.OutputNode that contains all the output nodes in the JSON file.
var OutputNodes []nodes.OutputNodeInterface

// InputNodes is a slice of nodes.InputNode that contains all the input nodes in the JSON file.
var InputNodes []nodes.InputNodeInterface

var ConstValue []Const

//CreateQueue is a function that creates the process queues from the JSON file and order the nodes to assure that every node is processed in the right order.
//It also creates the input and output nodes and link them to the right source/destination

func CreateQueue(g Graph) {
	Mutex.Lock()
	//find the output nodes and create the process queues
	for _, nodeJson := range g.NodesJson {
		if strings.Contains(nodeJson.Type, "Output") {
			for _, out := range OutputNodes {
				for _, outHandle := range out.GetOutputList() {
					//Prevent the ability to put multiple times the same output
					if outHandle.Service == nodeJson.Data.Service && outHandle.SubService == nodeJson.Data.SubService {
						if outHandle.SubService != "" {
							serverResponse.ResponseProcessGraph = "Multiple use of the same output, service : " + outHandle.Service + " and SubService : " + outHandle.SubService
						} else {
							serverResponse.ResponseProcessGraph = "Multiple use of the same output : " + outHandle.Service
						}
						Mutex.Unlock()
						return
					}
				}
			}

			var queue []nodes.LogicalNodeInterface
			findPreviousNode(&queue, nodeJson, g)    //find the node link ahead of the output node
			LogicalNode = append(LogicalNode, queue) //Add a new queue to the process queues
			//keepOnlyOneLogicalNode()
			//create the output node
			outputToAdd, err := nodes.CreateNode(nodeJson.Type)
			if err != nil {
				fmt.Println(err)
			}
			outputNodeToAddInterface, ok := outputToAdd.(nodes.OutputNodeInterface)
			if !ok {
				fmt.Println("Error this node type isn't outputNodeInterface: ", outputToAdd)
			} else {
				var nodeInput nodes.InputHandle
				id, _ := strconv.Atoi(nodeJson.Id)
				description, _ := nodes.NodeDescription(nodeJson.Type)
				if description.Input[0].DataType == "variable" {
					for _, subSer := range description.SubServices {
						if subSer.FriendlyName == nodeJson.Data.FriendlyName && subSer.Primary == nodeJson.Data.Service {
							for _, sec := range subSer.Secondary {
								if sec.Name == nodeJson.Data.SubService {
									nodeInput = nodes.InputHandle{Input: nil, Name: description.Input[0].Name, DataType: sec.DataType}
								}
							}
						}
					}
				} else {
					nodeInput = nodes.InputHandle{
						Input:    nil,
						Name:     description.Input[0].Name,
						DataType: description.Input[0].DataType,
					}
				}
				nodeInputHandle := nodes.OutputNodeHandle{
					FriendlyName: nodeJson.Data.FriendlyName,
					Service:      nodeJson.Data.Service,
					SubService:   nodeJson.Data.SubService,
					OutputHandle: nodeInput,
				}
				var tabOutputNodeHandle []nodes.OutputNodeHandle
				tabOutputNodeHandle = append(tabOutputNodeHandle, nodeInputHandle)
				outputNodeToAddInterface.InitNode(id, nodeJson.Type, tabOutputNodeHandle)
				OutputNodes = append(OutputNodes, outputNodeToAddInterface) //Add the output node to the OutputNodes slice
			}
		}
	}
	linkNodes(g) //Link the nodes to the right source/destination
	Mutex.Unlock()
}

// findPreviousNode is a recursive function that find the nodes ahead of the output node and add them to the right slice
func findPreviousNode(queue *[]nodes.LogicalNodeInterface, nodeJson nodeJson, g Graph) {
	for _, edge := range g.Edges {
		if edge.Target == nodeJson.Id { //find the edge which target the actual node
			nextNodeJson := findNodeById(edge.Source, g)       //find the source node
			if !strings.Contains(nextNodeJson.Type, "Input") { //Check if the node ahead is not an input node
				nextNodeJsonId, _ := strconv.Atoi(nextNodeJson.Id)
				for i, v := range *queue { //Check if the node is already in the queue	if it is remove it
					if logicalNode, ok := v.(nodes.LogicalNodeInterface); ok {
						if logicalNode.GetId() == nextNodeJsonId {
							*queue = append((*queue)[:i], (*queue)[i+1:]...)
							break
						}
					} else {
						fmt.Println("Error this node type isn't logicalNodeInterface: ", v)
					}
				}
				// create the node and add it to the queue
				nodeToAdd := createNode(nextNodeJson, g)
				*queue = append([]nodes.LogicalNodeInterface{nodeToAdd}, *queue...)
				findPreviousNode(queue, nextNodeJson, g)
			} else { //if the node ahead is an input node
				isIn := false
				for _, input := range InputNodes {
					nodeJsonId, _ := strconv.Atoi(nextNodeJson.Id)
					if input.GetId() == nodeJsonId {
						isIn = true
						break
					}
				}
				if !isIn {
					inNodeToAdd, err := nodes.CreateNode(nextNodeJson.Type)
					if err != nil {
						fmt.Println(err)
					}
					inputNodeToAdd, ok := inNodeToAdd.(nodes.InputNodeInterface)
					if !ok {
						fmt.Println("Error this node type isn't inputNodeInterface: ", inputNodeToAdd)
					} else {
						inId, _ := strconv.Atoi(nextNodeJson.Id)
						var inputHandle nodes.InputHandle
						description, _ := nodes.NodeDescription(nextNodeJson.Type)
						if strings.Contains(nextNodeJson.Type, "constant") {
							val, _ := strconv.ParseFloat(nextNodeJson.Data.Value, 64)
							ConstValue = append(ConstValue, Const{Id: inId, Value: val})
							for i := range ConstValue {
								if ConstValue[i].Id == inId {
									valFloat := ConstValue[i].Value
									valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
									valStrCopy := valStr
									inputHandle = nodes.InputHandle{Input: &valStrCopy, Name: description.Output[0].Name, DataType: description.Output[0].DataType}
									break
								}
							}
						} else {
							if description.Output[0].DataType == "variable" {
								for _, subSer := range description.SubServices {
									if subSer.FriendlyName == nextNodeJson.Data.FriendlyName && subSer.Primary == nextNodeJson.Data.Service {
										for _, sec := range subSer.Secondary {
											if sec.Name == nextNodeJson.Data.SubService {
												inputHandle = nodes.InputHandle{Input: nil, Name: description.Output[0].Name, DataType: sec.DataType}
											}
										}
									}
								}
							} else {
								inputHandle = nodes.InputHandle{Input: nil, Name: description.Output[0].Name, DataType: description.Output[0].DataType}
							}
						}
						inputNodeHandle := nodes.InputNodeHandle{
							FriendlyName: nextNodeJson.Data.FriendlyName,
							Service:      nextNodeJson.Data.Service,
							SubService:   nextNodeJson.Data.SubService,
							InputHandle:  inputHandle,
						}
						var tabInputNodeHandle []nodes.InputNodeHandle
						tabInputNodeHandle = append(tabInputNodeHandle, inputNodeHandle)
						inputNodeToAdd.InitNode(inId, nextNodeJson.Type, tabInputNodeHandle, nextNodeJson.Data.ParameterValueData)
						InputNodes = append(InputNodes, inputNodeToAdd)
					}
				}
			}
		}
	}
}

// findNodeById is a function that find a node in the JSON file by its id
func findNodeById(source string, g Graph) nodeJson {
	for _, nodeJson := range g.NodesJson {
		if nodeJson.Id == source {
			return nodeJson
		}
	}
	return nodeJson{}
}

func getNbInputs(id string, g Graph) int {
	nb := 0
	var inputs []string
	for _, edge := range g.Edges {
		if edge.Target == id {
			isIn := false
			for _, input := range inputs {
				if input == edge.TargetHandle {
					isIn = true
					break
				}
			}
			if !isIn {
				nb++
				inputs = append(inputs, edge.TargetHandle)
			}
		}
	}
	return nb
}

// createNode is a function that create a node from a nodeJson struct
func createNode(nodeJson nodeJson, g Graph) nodes.LogicalNodeInterface {
	NodeJsonId, _ := strconv.Atoi(nodeJson.Id)
	nodeToAdd, err := nodes.CreateNode(nodeJson.Type)
	if err != nil {
		fmt.Println(err)
	}
	logicalNodeToAdd, ok := nodeToAdd.(nodes.LogicalNodeInterface)
	if !ok {
		fmt.Println("Error this nodeToAdd type isn't logicalNodeInterface: ", nodeToAdd)
	} else {
		description, _ := nodes.NodeDescription(nodeJson.Type)
		var input []nodes.InputHandle
		if description.Stretchable && !(description.Type_ == "StringToBoolNode") { //StringToBoolNode is Stretchable with only one input
			nbInputs := getNbInputs(nodeJson.Id, g)
			for i := 0; i < nbInputs; i++ {
				input = append(input, nodes.InputHandle{Input: nil, Name: description.Input[0].Name + strconv.Itoa(i), DataType: description.Input[0].DataType})
			}
		} else {
			for _, in := range description.Input {
				input = append(input, nodes.InputHandle{Input: nil, Name: in.Name, DataType: in.DataType})
			}
		}
		var output []nodes.OutputHandle
		for i, out := range description.Output {
			output = append(output, nodes.OutputHandle{Output: strconv.Itoa(i), Name: out.Name, DataType: out.DataType})
			fmt.Println(i)
		}
		logicalNodeToAdd.InitNode(NodeJsonId, nodeJson.Type, input, output, nodeJson.Data.ParameterValueData)
	}
	return nodeToAdd.(nodes.LogicalNodeInterface)
}

// findEdgeById is a function that find the edges that target a node in the JSON file
func findLinkedNode(targetId, targetHandle string, g Graph) nodeJson {
	for _, edge := range g.Edges {
		if edge.Target == targetId && edge.TargetHandle == targetHandle {
			for _, node := range g.NodesJson {
				if node.Id == edge.Source {
					return node
				}
			}
		}
	}
	return nodeJson{}
}

func findLinkedEdgeName(targetId, targetHandle string, g Graph) string {
	for _, edge := range g.Edges {
		if edge.Target == targetId && edge.TargetHandle == targetHandle {
			return edge.SourceHandle
		}
	}
	return ""
}

func linkNodes(g Graph) {
	/*** Verify InputNodes link ***/
	for i := range InputNodes {
		ableToConnect := false
		for _, edge := range g.Edges {
			if edge.Source == strconv.Itoa(InputNodes[i].GetId()) {
				for j := range inputUpdate.InputsOutputsState {
					inputHandle := InputNodes[i].GetOutput(edge.SourceHandle)
					inputLink := inputUpdate.InputsOutputsState[j]
					if inputLink.Service == inputHandle.Service && inputLink.SubService == inputHandle.SubService && inputLink.FriendlyName == inputHandle.FriendlyName {
						InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.Input = &inputUpdate.InputsOutputsState[j].Value
						ableToConnect = true
						break
					}
				}
			}
		}
		if strings.Contains(InputNodes[i].GetNodeType(), "constant") { //if contains constant
			ableToConnect = true
		}

		if strings.Contains(InputNodes[i].GetNodeType(), "viewWeb") {
			for _, edge := range g.Edges {
				if edge.Source == strconv.Itoa(InputNodes[i].GetId()) {
					for j := range server.InputsStateWeb {
						inputHandle := InputNodes[i].GetOutput(edge.SourceHandle)
						inputLink := server.InputsStateWeb[j]
						if strconv.Itoa(inputLink.IRCode) == inputHandle.FriendlyName {
							InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.Input = &server.InputsStateWeb[j].Value
							ableToConnect = true
							break
						}
					}
				}
			}
		}
		if !ableToConnect {
			serverResponse.ResponseProcessGraph = "Input node not connected" + strconv.Itoa(InputNodes[i].GetId())
			fmt.Println("Input node not connected" + strconv.Itoa(InputNodes[i].GetId()))
			break
		}
	}

	/*** Verify LogicalNode link ***/
	for i := range LogicalNode {
		for j := range LogicalNode[i] {
			actualNode := LogicalNode[i][j]
			actualNodeInputs := actualNode.GetInput()
			for k, in := range actualNodeInputs {
				srcNodeJson := findLinkedNode(strconv.Itoa(actualNode.GetId()), in.Name, g)
				switch {
				case strings.Contains(srcNodeJson.Type, "Input"):
					for l := range InputNodes {
						nodeJsonId, _ := strconv.Atoi(srcNodeJson.Id)
						if InputNodes[l].GetId() == nodeJsonId {
							targetNodeHandle := InputNodes[l].GetOutput(findLinkedEdgeName(strconv.Itoa(actualNode.GetId()), in.Name, g))
							//Check if the data type of the input and the output are the same
							if targetNodeHandle.InputHandle.DataType == actualNodeInputs[k].DataType {
								actualNodeInputs[k].Input = targetNodeHandle.InputHandle.Input
							} else {
								serverResponse.ResponseProcessGraph = "Data type mismatch"
								fmt.Println("Data type mismatch on two logical nodes " + strconv.Itoa(actualNode.GetId()) + " and " + strconv.Itoa(nodeJsonId))
							}
						}
					}
				default:
					for l := range LogicalNode[i] {
						nodeJsonId, _ := strconv.Atoi(srcNodeJson.Id)
						if LogicalNode[i][l].GetId() == nodeJsonId {
							targetNodeHandle := LogicalNode[i][l].GetOutput(findLinkedEdgeName(strconv.Itoa(actualNode.GetId()), in.Name, g))
							if targetNodeHandle.DataType == actualNodeInputs[k].DataType {
								fmt.Println("logical link : " + strconv.Itoa(actualNode.GetId()) + " (" + actualNode.GetNodeType() + ")" + " and " + strconv.Itoa(nodeJsonId) + " (" + srcNodeJson.Type + ")")
								actualNodeInputs[k].Input = &targetNodeHandle.Output
							} else {
								serverResponse.ResponseProcessGraph = "Data type mismatch"
								fmt.Println("Data type mismatch on two logical nodes " + strconv.Itoa(actualNode.GetId()) + " and " + strconv.Itoa(nodeJsonId))
							}
						}
					}
				}
			}
		}
	}

	/*** Verify Output link ***/
	for i := range OutputNodes {
		isLinked := false
		for _, edge := range g.Edges {
			if edge.Target == strconv.Itoa(OutputNodes[i].GetId()) {
				srcHandleName := edge.SourceHandle
				srcNodeJson := findLinkedNode(strconv.Itoa(OutputNodes[i].GetId()), edge.TargetHandle, g)
				if strings.Contains(srcNodeJson.Type, "Input") {
					for j := range InputNodes {
						src, _ := strconv.Atoi(edge.Source)
						if InputNodes[j].GetId() == src {
							dataTypeScr := InputNodes[j].GetOutput(srcHandleName).InputHandle.DataType
							dataTypeTarget := OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.DataType
							if dataTypeScr == dataTypeTarget {
								isLinked = true
								OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.Input = InputNodes[j].GetOutput(srcHandleName).InputHandle.Input //we pass the input address to the output
								break
							}
						}
					}
				} else {
					for j := range LogicalNode {
						for k := range LogicalNode[j] {
							src, _ := strconv.Atoi(edge.Source)
							if LogicalNode[j][k].GetId() == src {
								dataTypeScr := LogicalNode[j][k].GetOutput(srcHandleName).DataType
								dataTypeTarget := OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.DataType
								if dataTypeScr == dataTypeTarget {
									isLinked = true
									OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.Input = &LogicalNode[j][k].GetOutput(srcHandleName).Output
									break
								}
							}
						}
					}
				}
			}
			if isLinked {
				break
			}
		}
		if !isLinked {
			serverResponse.ResponseProcessGraph = "Data type mismatch, output node not connected"
			fmt.Println("Data type mismatch on two nodes while linking an output to a input")
			break
		}
	}
}

// function to find a node by ID in LogicalNode[i]
func getLogicalNodeById(nodes []nodes.LogicalNodeInterface, id int) nodes.LogicalNodeInterface {
	for _, n := range nodes {
		if n.GetId() == id {
			return n
		}
	}
	return nil
}

// function minimum of size for LogicalNode[i]
func keepOnlyOneLogicalNode() {
	lastQueueIndex := len(LogicalNode) - 1
	if lastQueueIndex == 0 {
		return
	}
	for n2 := range LogicalNode[lastQueueIndex] {
		for q1 := range LogicalNode {
			for n1 := range LogicalNode[q1] {
				if n1 < (len(LogicalNode[lastQueueIndex]) - 1) {
					if LogicalNode[lastQueueIndex][n2].GetId() == LogicalNode[q1][n1].GetId() {
						fmt.Println("remove id : " + strconv.Itoa(n2))
						LogicalNode[lastQueueIndex] = removeAtIndex(LogicalNode[lastQueueIndex], n2)

					}
				}
			}
		}
	}

}
func removeAtIndex(s []nodes.LogicalNodeInterface, i int) []nodes.LogicalNodeInterface {
	return append(s[:i], s[i+1:]...)
}
