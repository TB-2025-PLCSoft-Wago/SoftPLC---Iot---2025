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

		//Reset especially comunication
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

		// Convert to JSON to save in file
		jsonData, err := json.MarshalIndent(jsonBody, "", "  ")
		if err != nil {
			return err
		}
		// save in file
		err2 := os.WriteFile("graph_marcelin_tof.json", jsonData, 0644)
		if err2 != nil {
			return err2
		}
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
		return c.JSON(http.StatusOK, data)
		//return c.JSON(http.StatusOK, savedJson)
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
		// read file
		//jsonData, err2 := os.ReadFile("graph_marcelin_tof.json")
		/*if err2 != nil {
			return err2
		}
		// Désérialiser le JSON
		var data interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			return err
		}*/
		return c.JSON(http.StatusOK, server.DebugGraphJson)
		//return c.JSON(http.StatusOK, savedJson)
	})
	e.GET("/debugStop", func(c echo.Context) error {
		// read file
		/*
			jsonData, err2 := os.ReadFile("graph_marcelin_tof.json")
			if err2 != nil {
				return err2
			}
			// Désérialiser le JSON
			var data interface{}
			if err := json.Unmarshal(jsonData, &data); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, data)
		*/
		server.RemoveLabelsFromDebugGraph()
		return c.JSON(http.StatusOK, server.DebugGraphJson)
	})
	e.POST("/json-save-toDebug", func(c echo.Context) error {
		var jsonBody interface{}
		if err := c.Bind(&jsonBody); err != nil {
			return err
		}
		server.DebugGraphJson = jsonBody

		// Convert to JSON to save in file
		//jsonData, err := json.MarshalIndent(jsonBody, "", "  ")
		/*if err != nil {
			return err
		}*/
		return c.HTML(http.StatusOK, "Graph saved to debug")
	})

	e.Logger.Fatal(e.Start(":8889"))
}
