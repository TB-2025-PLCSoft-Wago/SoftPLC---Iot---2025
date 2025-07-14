package main

import (
	"SoftPLC/echo"
	"SoftPLC/inputUpdate"
	"SoftPLC/nodes"
	"SoftPLC/outputUpdate"
	"SoftPLC/processGraph"
	"SoftPLC/server"
	"SoftPLC/variable"
	"time"
)

var start2 time.Time

func main() {
	//nodes.CreateNode("TONNode")
	//fmt.Println(nodes.NodeDescription("digitalOutput"))
	//fmt.Println(nodes.SystemDescription())

	inputUpdate.InitInputs()
	timer := time.NewTimer(1000 * time.Millisecond)
	<-timer.C
	go echo.EchoServer()
	go server.CreateWebSocket()
	ticker := time.NewTicker(10 * time.Millisecond)
	ticker2 := time.NewTicker(600 * time.Second) // monitoring Lists
	go func() {
		for {
			select {
			case <-ticker2.C:

				inputUpdate.CreateMonitoringLists() // Recreate
			}
		}
	}()
	for {
		select {
		case <-ticker.C:
			if len(processGraph.OutputNodes) != 0 {
				//start := time.Now()
				//fmt.Printf("total time before restart is  %s\n", time.Since(start2))
				inputUpdate.UpdateInputs()
				processGraph.Mutex.Lock()
				for _, v := range processGraph.LogicalNode {
					for _, n := range v {
						if logicalNode, ok := n.(nodes.LogicalNodeInterface); ok {
							logicalNode.ProcessLogic()

						}
					}
				}
				outputUpdate.UpdateOutput()
				server.UpdateOutputValueView()
				variable.UpdateVariableInputs()
				if server.IsActiveDebug {
					server.DebugMode()
				}

				processGraph.Mutex.Unlock()
				//fmt.Printf("total time value is  %s\n", time.Since(start))
				start2 = time.Now()

			}
		}
	}

	/*in1 := 1.0
	in2 := 2.0
	in3 := 3.0
	in4 := 4.0
	in5 := 5.0
	in6 := 6.0
	var ln []tg
	ln = append(ln, tg{eddegIn: []*float64{&in1, &in2, &in3}})
	for _, n := range ln[0].eddegIn {
		fmt.Println(*n)
	}
	pn := ln[0].getInput()
	pn[0] = &in4
	pn[1] = &in5
	pn[2] = &in6
	for _, n := range ln[0].eddegIn {
		fmt.Println(*n)
	}

	/*var i1 float64 = 1.0
	var i2 float64 = 0.0

	var s []*float64
	var m map[*float64]string

	s = append(s, &i1)
	s = append(s, &i2)

	m = make(map[*float64]string)
	m[s[0]] = "cst"
	m[s[1]] = "prout"

	fmt.Println(m[s[0]], m[s[1]])
	fmt.Println(*s[0], *s[1])

	i1 = 10000.0
	i2 = 20000.0

	fmt.Println(m[s[0]], m[s[1]])
	fmt.Println(*s[0], *s[1])*/

	/*var in1 float64 = 0.0
	var in2 float64 = 0.0
	var in3 float64 = 1.0

	var ln []nodes.LogicalNodeInterface

	ln = append(ln, &nodes.AndNode{id: 1, Type: "AND", Output: 0, Input1: &in1, Input2: &in2})
	ln = append(ln, &nodes.OrNode{id: 2, Type: "OR", Output: 0, Input1: &in3, Input2: &in2})

	var out nodes.OutputNode = nodes.OutputNode{Value: ln[1].GetOutput(), Destination: "Q1"}

	for _, n := range ln {
		n.ProcessLogic()
		fmt.Println(*n.GetOutput())
	}

	fmt.Println(*out.Value)

	in2 = 1.0

	for _, n := range ln {
		n.ProcessLogic()
		fmt.Println(*n.GetOutput())
	}

	fmt.Println(*out.Value)*/

}
