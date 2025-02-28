package echo

import (
	"SoftPLC/nodes"
	"SoftPLC/processGraph"
	"SoftPLC/serverResponse"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var savedJson interface{}

func EchoServer() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.POST("/json-graph", func(c echo.Context) error {
		serverResponse.ResponseProcessGraph = "Graph received"
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
		return c.HTML(http.StatusOK, "Graph saved")
	})

	e.GET("/get-saved-json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, savedJson)
	})

	e.GET("/get-description", func(c echo.Context) error {
		answer, err := nodes.SystemDescription()
		fmt.Println(answer)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, answer)
	})

	e.Logger.Fatal(e.Start(":8889"))
}
