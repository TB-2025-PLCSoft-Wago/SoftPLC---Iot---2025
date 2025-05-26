// This node is a TOF node.
// timer with a turn on delay
// Input: IN, PT
// Output: Q
// on IN falling edge a timer is started with PT time on IN rising edge the timer is reset
// Q is set to 1 and stay when the timer is not elapsed and IN is 0
// Q is set to 0 if IN is 0 or if the timer is not elapsed
package nodes

import (
	"fmt"
	"strconv"
	"time"
)

type TOFNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
	timer    time.Timer
	elapsed  bool
	oldIN    float64
	fired    bool
}

var tofDescription = nodeDescription{
	AccordionName: "Timer",
	PrimaryType:   "LogicalNode",
	Type_:         "TOFNode",
	Display:       "Tof Node",
	Label:         "TOF",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "Input"},
		{DataType: "value", Name: "Time [ms]"},
	},
	Output: []dataTypeNameStruct{{DataType: "bool", Name: "Output"}},
}

func init() {
	RegisterNodeCreator("TOFNode", func() (Node, error) {
		return &TOFNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, tofDescription)
}

func (t *TOFNode) GetId() int {
	return t.id
}

func (t *TOFNode) GetNodeType() string {
	return t.nodeType
}

func (t *TOFNode) ProcessLogic() {
	if t.fired {
		select {
		case <-t.timer.C:
			t.elapsed = true
		default:
		}
	}

	indexPT := -1
	indexIN := -1
	indexOut := -1

	for i, in := range t.input {
		if in.Name == "Time [ms]" {
			indexPT = i
		}
		if in.Name == "Input" {
			indexIN = i
		}
	}
	for i, out := range t.output {
		if out.Name == "Output" {
			indexOut = i
		}
	}

	if t.input[indexIN].Input == nil {
		fmt.Println("TOFNode: processLogic: Error: Input has no link")
		return
	}
	if t.input[indexPT].Input == nil {
		fmt.Println("TOFNode: processLogic: Error: Time has no link")
		return
	}

	if indexPT != -1 && indexIN != -1 && indexOut != -1 {
		if *t.input[indexIN].Input == "0" && t.oldIN == 1 {
			if !t.fired {
				//fmt.Println(*t.input[indexPT].Input)
				valueStr := *t.input[indexPT].Input
				valueFloat, _ := strconv.ParseFloat(valueStr, 64)
				fmt.Println(time.Duration(valueFloat) * time.Millisecond)
				t.timer = *time.NewTimer(time.Duration(valueFloat) * time.Millisecond)
				t.fired = true
			}
			t.elapsed = false
			t.oldIN = 0
		}

		if *t.input[indexIN].Input == "1" && t.fired == true {
			t.elapsed = false
			t.oldIN = 1
			t.fired = false
			t.timer.Stop()
		}
		if *t.input[indexIN].Input == "1" {
			t.oldIN = 1
		}

		if *t.input[indexIN].Input == "1" || (!t.elapsed && t.fired == true) {
			t.output[indexOut].Output = "1"
		} else {
			t.output[indexOut].Output = "0"
		}
	} else {
		fmt.Println("TOFNode: processLogic: Error: Input or output not found")
	}

}

func (t *TOFNode) GetOutput(outName string) *OutputHandle {
	for i, name := range t.output {
		if name.Name == outName {
			return &t.output[i]
		}
	}
	return nil
}

func (t *TOFNode) GetInput() []InputHandle {
	return t.input
}

func (t *TOFNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	t.id = id_
	t.nodeType = nodeType_
	t.input = input_
	t.output = output_
}
