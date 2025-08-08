package processGraph

import (
	"SoftPLC/function"
	"SoftPLC/inputUpdate"
	"SoftPLC/nodes"
	"SoftPLC/server"
	"SoftPLC/serverResponse"
	"SoftPLC/variable"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var MutexFunction sync.Mutex

type functionProcess2 struct {
	//nameFunction string
	LogicalNode [][]nodes.LogicalNodeInterface
	OutputNodes []nodes.OutputNodeInterface
	InputNodes  []nodes.InputNodeInterface
	ConstValue  []Const
}
type functionOutputNodes struct {
	OutputNodes []nodes.OutputNodeInterface
}

var ListProcessFunction map[string]functionProcess2
var ListProcessFunctionLogicalNode map[string][][]nodes.LogicalNodeInterface
var ListProcessFunctionOutputNodes map[string]functionOutputNodes
var ListProcessFunctionInputNodes map[string][][]nodes.InputNodeInterface

func init() {
	ListProcessFunction = make(map[string]functionProcess2)
}
func CreateQueueFunction(g Graph, name string) {
	ListProcessFunction[name] = functionProcess2{}
	fp := ListProcessFunction[name]
	MutexFunction.Lock()
	//find the output nodes and create the process queues
	for _, nodeJson := range g.NodesJson {
		if strings.Contains(nodeJson.Type, "Output") {
			for _, out := range fp.OutputNodes {
				for _, outHandle := range out.GetOutputList() {
					//Prevent the ability to put multiple times the same output
					if outHandle.Service == nodeJson.Data.Service && outHandle.SubService == nodeJson.Data.SubService {
						if outHandle.SubService != "" {
							serverResponse.ResponseProcessGraph = "Multiple use of the same output, service : " + outHandle.Service + " and SubService : " + outHandle.SubService
						} else {
							serverResponse.ResponseProcessGraph = "Multiple use of the same output : " + outHandle.Service
						}
						MutexFunction.Unlock()
						return
					}
				}
			}

			var queue []nodes.LogicalNodeInterface
			findPreviousNodeFunction(&queue, nodeJson, g, name) //find the node link ahead of the output node
			fp = ListProcessFunction[name]
			fp.LogicalNode = append(fp.LogicalNode, queue) //Add a new queue to the process queues
			ListProcessFunction[name] = fp
			//keepOnlyOneLogicalNodeFunction ()
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
				} else if description.Input[0].DataType == "function" {
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
				fp = ListProcessFunction[name]
				fp.OutputNodes = append(fp.OutputNodes, outputNodeToAddInterface) //Add the output node to the fp.OutputNodes slice
				ListProcessFunction[name] = fp

			}
		}
	}
	ListProcessFunction[name] = fp
	linkNodesFunction(g, name) //Link the nodes to the right source/destination
	MutexFunction.Unlock()
}

// findPreviousNodeFunction is a recursive function that find the nodes ahead of the output node and add them to the right slice
func findPreviousNodeFunction(queue *[]nodes.LogicalNodeInterface, nodeJson NodeJson, g Graph, name string) {
	fp := ListProcessFunction[name]
	for _, edge := range g.Edges {
		if edge.Target == nodeJson.Id { //find the edge which target the actual node
			nextNodeJson := findNodeByIdFunction(edge.Source, g) //find the source node
			if !strings.Contains(nextNodeJson.Type, "Input") {   //Check if the node ahead is not an input node
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
				nodeToAdd := createNodeFunction(nextNodeJson, g)
				if nodeToAdd == nil {
					// we stop everything or we ignore this node
					fmt.Println("Node unknown or invalid, addition to the queue stopped.")
					return
				}
				*queue = append([]nodes.LogicalNodeInterface{nodeToAdd}, *queue...)
				findPreviousNodeFunction(queue, nextNodeJson, g, name)
				//TO DO : maybe add ListProcessFunction[name] = fp
			} else { //if the node ahead is an input node
				isIn := false
				for _, input := range fp.InputNodes {
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
							//val, _ := strconv.ParseFloat(nextNodeJson.Data.Value, 64)
							//fp.ConstValue = append(fp.ConstValue, Const{Id: inId, Value: val})
							//val, _ := strconv.ParseFloat(nextNodeJson.Data.Value, 64)
							fp = ListProcessFunction[name]
							fp.ConstValue = append(fp.ConstValue, Const{Id: inId, Value: nextNodeJson.Data.Value})
							ListProcessFunction[name] = fp
							for i := range fp.ConstValue {
								if fp.ConstValue[i].Id == inId {
									//valFloat := fp.ConstValue[i].Value
									//valStr := strconv.FormatFloat(valFloat, 'f', -1, 64)
									//valStrCopy := valStr
									valStrCopy := fp.ConstValue[i].Value
									inputHandle = nodes.InputHandle{Input: &valStrCopy, Name: description.Output[0].Name, DataType: description.Output[0].DataType}
									server.AddDebugState(edge.Source, &fp.ConstValue[i].Value, edge.SourceHandle, description.Output[0].DataType)
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
							} else if description.Output[0].DataType == "function" {
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
						fp = ListProcessFunction[name]
						fp.InputNodes = append(fp.InputNodes, inputNodeToAdd)
						ListProcessFunction[name] = fp
					}
				}
			}
		}
	}
	//ListProcessFunction[name] = fp
}

// findNodeByIdFunction is a function that find a node in the JSON file by its id
func findNodeByIdFunction(source string, g Graph) NodeJson {
	for _, nodeJson := range g.NodesJson {
		if nodeJson.Id == source {
			return nodeJson
		}
	}
	return NodeJson{}
}

func getNbInputsFunction(id string, g Graph) int {
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

// createNodeFunction is a function that create a node from a NodeJson struct
func createNodeFunction(nodeJson NodeJson, g Graph) nodes.LogicalNodeInterface {
	NodeJsonId, _ := strconv.Atoi(nodeJson.Id)
	nodeToAdd, err := nodes.CreateNode(nodeJson.Type)
	if err != nil {
		fmt.Println("createNode", err)
		serverResponse.ResponseProcessGraph = "Error this node to add type isn't logicalNodeInterface : " + nodeJson.Type
		return nil
	}
	logicalNodeToAdd, ok := nodeToAdd.(nodes.LogicalNodeInterface)
	if !ok {
		fmt.Println("Error this nodeToAdd type isn't logicalNodeInterface: ", nodeToAdd)
		serverResponse.ResponseProcessGraph = "Error this node to add type isn't logicalNodeInterface : " + nodeJson.Type
	} else {
		description, _ := nodes.NodeDescription(nodeJson.Type)
		var input []nodes.InputHandle
		if description.Stretchable && !(description.Type_ == "StringToBoolNode") { //StringToBoolNode is Stretchable with only one input
			nbInputs := getNbInputsFunction(nodeJson.Id, g)
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
		}
		logicalNodeToAdd.InitNode(NodeJsonId, nodeJson.Type, input, output, nodeJson.Data.ParameterValueData)
	}
	return nodeToAdd.(nodes.LogicalNodeInterface)
}

// findEdgeById is a function that find the edges that target a node in the JSON file
func findLinkedNodeFunction(targetId, targetHandle string, g Graph) NodeJson {
	for _, edge := range g.Edges {
		if edge.Target == targetId && edge.TargetHandle == targetHandle {
			for _, node := range g.NodesJson {
				if node.Id == edge.Source {
					return node
				}
			}
		}
	}
	return NodeJson{}
}

func findLinkedEdgeNameFunction(targetId, targetHandle string, g Graph) string {
	for _, edge := range g.Edges {
		if edge.Target == targetId && edge.TargetHandle == targetHandle {
			return edge.SourceHandle
		}
	}
	return ""
}

func linkNodesFunction(g Graph, name string) {
	fp := ListProcessFunction[name]
	/*** Verify fp.InputNodes link ***/
	for i := range fp.InputNodes {
		ableToConnect := false
		for _, edge := range g.Edges {
			if edge.Source == strconv.Itoa(fp.InputNodes[i].GetId()) {
				for j := range inputUpdate.InputsOutputsState {
					inputHandle := fp.InputNodes[i].GetOutput(edge.SourceHandle)
					inputLink := inputUpdate.InputsOutputsState[j]
					if inputLink.Service == inputHandle.Service && inputLink.SubService == inputHandle.SubService && inputLink.FriendlyName == inputHandle.FriendlyName {
						fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.Input = &inputUpdate.InputsOutputsState[j].Value
						server.AddDebugState(edge.Source, &inputUpdate.InputsOutputsState[j].Value, edge.SourceHandle, fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.DataType)
						ableToConnect = true
						break
					}
				}
			}
		}
		if strings.Contains(fp.InputNodes[i].GetNodeType(), "constant") { //if contains constant
			ableToConnect = true
		}

		if strings.Contains(fp.InputNodes[i].GetNodeType(), "viewWeb") {
			for _, edge := range g.Edges {
				if edge.Source == strconv.Itoa(fp.InputNodes[i].GetId()) {
					for j := range server.InputsStateWeb {
						inputHandle := fp.InputNodes[i].GetOutput(edge.SourceHandle)
						inputLink := server.InputsStateWeb[j]
						if strconv.Itoa(inputLink.IRCode) == inputHandle.FriendlyName {
							fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.Input = &server.InputsStateWeb[j].Value
							server.AddDebugState(edge.Source, &server.InputsStateWeb[j].Value, edge.SourceHandle, fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.DataType)
							ableToConnect = true
							break
						}
					}
				}
			}
		}
		if strings.Contains(fp.InputNodes[i].GetNodeType(), "variable") {
			for _, edge := range g.Edges {
				if edge.Source == strconv.Itoa(fp.InputNodes[i].GetId()) {
					for j := range variable.InputsStateVariable {
						inputHandle := fp.InputNodes[i].GetOutput(edge.SourceHandle)
						inputLink := variable.InputsStateVariable[j]
						if inputLink.Name == inputHandle.FriendlyName {
							fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.Input = &variable.InputsStateVariable[j].Value
							server.AddDebugState(edge.Source, &variable.InputsStateVariable[j].Value, edge.SourceHandle, fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.DataType)
							ableToConnect = true
							break
						}
					}
				}
			}
		}

		if strings.Contains(fp.InputNodes[i].GetNodeType(), "function") {
			for _, edge := range g.Edges {
				if edge.Source == strconv.Itoa(fp.InputNodes[i].GetId()) {
					for j := range function.InputsStateFunction {
						inputHandle := fp.InputNodes[i].GetOutput(edge.SourceHandle)
						inputLink := function.InputsStateFunction[j]
						if inputLink.Name == inputHandle.FriendlyName {
							fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.Input = &function.InputsStateFunction[j].Value
							server.AddDebugState(edge.Source, &function.InputsStateFunction[j].Value, edge.SourceHandle, fp.InputNodes[i].GetOutput(edge.SourceHandle).InputHandle.DataType)
							ableToConnect = true
							break
						}
					}
				}
			}
		}
		if !ableToConnect {
			serverResponse.ResponseProcessGraph = "Input node not connected : " + fp.InputNodes[i].GetNodeType() //+ strconv.Itoa(fp.InputNodes[i].GetId())
			fmt.Println("Input node not connected" + strconv.Itoa(fp.InputNodes[i].GetId()))
			break
		}
	}

	/*** Verify fp.LogicalNode link ***/
	for i := range fp.LogicalNode {
		for j := range fp.LogicalNode[i] {
			actualNode := fp.LogicalNode[i][j]
			actualNodeInputs := actualNode.GetInput()
			for k, in := range actualNodeInputs {
				srcNodeJson := findLinkedNodeFunction(strconv.Itoa(actualNode.GetId()), in.Name, g)
				switch {
				case strings.Contains(srcNodeJson.Type, "Input"):
					for l := range fp.InputNodes {
						nodeJsonId, _ := strconv.Atoi(srcNodeJson.Id)
						if fp.InputNodes[l].GetId() == nodeJsonId {
							targetNodeHandle := fp.InputNodes[l].GetOutput(findLinkedEdgeNameFunction(strconv.Itoa(actualNode.GetId()), in.Name, g))
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
					for l := range fp.LogicalNode[i] {
						nodeJsonId, _ := strconv.Atoi(srcNodeJson.Id)
						if fp.LogicalNode[i][l].GetId() == nodeJsonId {
							targetNodeHandle := fp.LogicalNode[i][l].GetOutput(findLinkedEdgeNameFunction(strconv.Itoa(actualNode.GetId()), in.Name, g))
							if targetNodeHandle.DataType == actualNodeInputs[k].DataType {
								fmt.Println("logical link : " + strconv.Itoa(actualNode.GetId()) + " (" + actualNode.GetNodeType() + ")" + " and " + strconv.Itoa(nodeJsonId) + " (" + srcNodeJson.Type + ")")
								actualNodeInputs[k].Input = &targetNodeHandle.Output                                                                       // we pass the output address of the first logical node to the input of the second
								server.AddDebugState(strconv.Itoa(nodeJsonId), &targetNodeHandle.Output, targetNodeHandle.Name, targetNodeHandle.DataType) //value 1 out
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
	for i := range fp.OutputNodes {
		isLinked := false
		for _, edge := range g.Edges {
			if edge.Target == strconv.Itoa(fp.OutputNodes[i].GetId()) {
				srcHandleName := edge.SourceHandle
				srcNodeJson := findLinkedNodeFunction(strconv.Itoa(fp.OutputNodes[i].GetId()), edge.TargetHandle, g)
				if strings.Contains(srcNodeJson.Type, "Input") {
					for j := range fp.InputNodes {
						src, _ := strconv.Atoi(edge.Source)
						if fp.InputNodes[j].GetId() == src {
							dataTypeScr := fp.InputNodes[j].GetOutput(srcHandleName).InputHandle.DataType
							dataTypeTarget := fp.OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.DataType
							if dataTypeScr == dataTypeTarget {
								isLinked = true
								fp.OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.Input = fp.InputNodes[j].GetOutput(srcHandleName).InputHandle.Input //we pass the input address to the output
								server.AddDebugState(edge.Source, fp.InputNodes[j].GetOutput(srcHandleName).InputHandle.Input, srcHandleName, dataTypeTarget)
								break
							}
						}
					}
				} else {
					for j := range fp.LogicalNode {
						for k := range fp.LogicalNode[j] {
							src, _ := strconv.Atoi(edge.Source)
							if fp.LogicalNode[j][k].GetId() == src {
								dataTypeScr := fp.LogicalNode[j][k].GetOutput(srcHandleName).DataType
								dataTypeTarget := fp.OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.DataType
								if dataTypeScr == dataTypeTarget {
									isLinked = true
									fp.OutputNodes[i].GetOutput(edge.TargetHandle).OutputHandle.Input = &fp.LogicalNode[j][k].GetOutput(srcHandleName).Output //we pass the input address of logical to the output
									server.AddDebugState(edge.Source, &fp.LogicalNode[j][k].GetOutput(srcHandleName).Output, srcHandleName, dataTypeTarget)
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
			serverResponse.ResponseProcessGraph = "Data type mismatch, output node " + fp.OutputNodes[i].GetNodeType() + " of id " + strconv.Itoa(fp.OutputNodes[i].GetId()) + " not connected"
			fmt.Println("Data type mismatch on two nodes while linking an output to a input : " + fp.OutputNodes[i].GetNodeType() + ", id :" + strconv.Itoa(fp.OutputNodes[i].GetId()))
			break
		}
	}
	ListProcessFunction[name] = fp
}
