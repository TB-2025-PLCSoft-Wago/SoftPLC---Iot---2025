package echo

import (
	"SoftPLC/nodes"
	"SoftPLC/outputUpdate"
	"SoftPLC/processGraph"
	"SoftPLC/server"
	"SoftPLC/serverResponse"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

var savedJson interface{}

func EchoServer() {
	e := echo.New()

	e.Use(middleware.CORS())

	//After build
	e.POST("/json-graph", func(c echo.Context) error {
		serverResponse.ResponseProcessGraph = "Graph received"

		//Reset especially communication
		processGraph.Mutex.Lock()
		for _, v := range processGraph.LogicalNode {
			for _, n := range v {
				if logicalNode, ok := n.(nodes.LogicalNodeInterface); ok {
					logicalNode.DestroyToBuildAgain()
				}
			}
		}
		server.ResetAll()
		outputUpdate.UpdateOutput()
		processGraph.Mutex.Unlock()
		//nil
		var graph processGraph.Graph
		if err := c.Bind(&graph); err != nil {
			return err
		}
		processGraph.OutputNodes = nil
		processGraph.LogicalNode = nil
		processGraph.InputNodes = nil
		processGraph.ConstValue = nil
		processGraph.CreateQueue(graph)

		if serverResponse.ResponseProcessGraph != "Graph received" {
			processGraph.OutputNodes = nil
			processGraph.LogicalNode = nil
			processGraph.InputNodes = nil
			processGraph.ConstValue = nil
		}

		fmt.Println("OutputNodes: ")
		for _, v := range processGraph.OutputNodes {
			fmt.Println(v)
		}
		fmt.Println("LogicalNode: ")
		for _, v := range processGraph.LogicalNode {
			for _, n := range v {
				fmt.Println(n)
			}
		}
		fmt.Println("InputNodes: ")
		for _, v := range processGraph.InputNodes {
			fmt.Println(v)
		}
		fmt.Println("Const: ")
		for _, v := range processGraph.ConstValue {
			fmt.Println(v)
		}
		return c.HTML(http.StatusOK, serverResponse.ResponseProcessGraph)
	})

	e.POST("/json-save", func(c echo.Context) error {
		var jsonBody interface{}
		if err := c.Bind(&jsonBody); err != nil {
			return err
		}
		savedJson = jsonBody
		//TODO : Uncomment to save in file
		// Convert to JSON to save in file
		/*jsonData, err := json.MarshalIndent(jsonBody, "", "  ")
		if err != nil {
			return err
		}
		// save in file
		err2 := os.WriteFile("graph_marcelin_tof.json", jsonData, 0644)
		if err2 != nil {
			return err2
		}*/
		return c.HTML(http.StatusOK, "Graph saved")
	})

	e.GET("/get-saved-json", func(c echo.Context) error {
		// read file
		jsonData, err2 := os.ReadFile("graph_marcelin_tof.json")
		if err2 != nil {
			return err2
		}
		// Désérialiser le JSON
		var data interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			return err
		}
		//return c.JSON(http.StatusOK, data) //return from file
		return c.JSON(http.StatusOK, savedJson) //return from variable
	})

	e.GET("/get-description", func(c echo.Context) error {
		answer, err := nodes.SystemDescription()
		fmt.Println(answer)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, answer)
	})

	e.GET("/debug", func(c echo.Context) error {
		server.IsActiveDebug = true
		return c.JSON(http.StatusOK, server.DebugGraphJson)
	})
	e.GET("/debugStop", func(c echo.Context) error {
		server.RemoveLabelsFromDebugGraph()
		server.IsActiveDebug = false
		return c.JSON(http.StatusOK, server.DebugGraphJsonCopy)
	})
	e.POST("/json-save-toDebug", func(c echo.Context) error {
		var jsonBody interface{}
		if err := c.Bind(&jsonBody); err != nil {
			return err
		}
		server.DebugGraphJson = jsonBody
		server.DebugGraphJsonCopy = server.CopyInterface(server.DebugGraphJson)
		return c.HTML(http.StatusOK, "Graph saved to debug")
	})

	e.POST("/new-function", func(c echo.Context) error {
		var jsonBody interface{}
		if err := c.Bind(&jsonBody); err != nil {
			return err
		}
		/*graphData, ok := jsonBody.(map[string]interface{})["data"]
		if !ok {
			return c.HTML(http.StatusBadRequest, "Graph function wrong format")
		}
		name := jsonBody.(map[string]interface{})["name"].(string)
		CreateFunctionQueue(name, graphData)*/
		nodes.FunctionGraphJson = jsonBody
		nodes.FunctionAddToList()
		return c.HTML(http.StatusOK, "Graph function saved")
	})

	e.Logger.Fatal(e.Start(":8889"))
}

func CreateFunctionQueue(name string, graphFunc interface{}) error {
	/*if len(name) >= 5 {
		name = name[:len(name)-5]
	}*/
	graphNodes, ok := graphFunc.(map[string]interface{})["nodes"]
	if !ok {
		return fmt.Errorf("graph function wrong format")
	}
	nodeInterfaces, ok := graphNodes.([]interface{})
	if !ok {
		return fmt.Errorf("invalid node list format")
	}

	var nodesJson []processGraph.NodeJson
	for _, node := range nodeInterfaces {
		rawMap, ok := node.(map[string]interface{})
		if !ok {
			continue
		}

		convertedNode, err := convertToNodeJson(rawMap)
		if err != nil {
			log.Println("node conversion error:", err)
			continue
		}

		nodesJson = append(nodesJson, convertedNode)
	}

	graphEdges := graphFunc.(map[string]interface{})["edges"]
	graphComplet := map[string]interface{}{
		"nodes": nodesJson, // graphData2
		"edges": graphEdges,
	}

	var graph processGraph.Graph

	jsonBytes, err := json.Marshal(graphComplet)
	if err != nil {
		return fmt.Errorf("cannot marshal graphFunc: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, &graph); err != nil {
		fmt.Println("JSON Function parse error:", err)
		return err
	}

	processGraph.CreateQueueFunction(graph, name)
	fmt.Println("function input :", processGraph.ListProcessFunction[name].InputNodes)
	fp := nodes.FunctionNodeListProcess[name]
	if processGraph.ListProcessFunction[name].InputNodes != nil {
		fp.InputNodes = processGraph.ListProcessFunction[name].InputNodes
	}
	if processGraph.ListProcessFunction[name].LogicalNode != nil {
		fp.LogicalNode = processGraph.ListProcessFunction[name].LogicalNode
	}
	if processGraph.ListProcessFunction[name].OutputNodes != nil {
		fp.OutputNodes = processGraph.ListProcessFunction[name].OutputNodes
	}
	nodes.FunctionNodeListProcess[name] = fp

	return nil

}

func convertToNodeJson(fullNode map[string]interface{}) (processGraph.NodeJson, error) {
	data, ok := fullNode["data"].(map[string]interface{})
	if !ok {
		return processGraph.NodeJson{}, fmt.Errorf("missing or invalid 'data' field")
	}
	id := toString(fullNode["id"])
	typ := toString(data["type"])

	// Extraire proprement les string
	friendlyName := toString(data["selectedFriendlyNameData"])
	service := toString(data["selectedServiceData"])
	subService := toString(data["selectedSubServiceData"])
	value := toString(data["valueData"])

	// Convertir parameterValueData []interface{} -> []string
	rawParamValues, _ := data["parameterValueData"].([]interface{})
	var paramValues []string
	for _, val := range rawParamValues {
		strVal, ok := val.(string)
		if ok {
			paramValues = append(paramValues, strVal)
		}
	}

	return processGraph.NodeJson{
		Id:   id,
		Type: typ,
		Data: processGraph.Data_{
			FriendlyName:       friendlyName,
			Service:            service,
			SubService:         subService,
			Value:              value,
			ParameterValueData: paramValues,
		},
	}, nil
}

func toString(val interface{}) string {
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}
