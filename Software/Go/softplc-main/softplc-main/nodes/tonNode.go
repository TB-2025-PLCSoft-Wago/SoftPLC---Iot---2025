// This node is a TON node.
// timer with a turn on delay
// Input: IN, PT
// Output: Q
// on IN rising edge a timer is started with PT time on IN falling edge the timer is reset
// Q is set to 1 when the timer elapsed and IN is 1
// Q is set to 0 if IN is 0 or if the timer is not elapsed
package nodes

import (
	"fmt"
	"time"
)

type TONNode struct {
	id       int
	nodeType string
	input    []InputHandle
	output   []OutputHandle
	timer    time.Timer
	elapsed  bool
	oldIN    float64
	fired    bool
}

var tonDescription = nodeDescription{
	AccordionName: "Timer",
	PrimaryType:   "LogicalNode",
	Type_:         "TONNode",
	Display:       "Ton Node",
	Label:         "TON",
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
	RegisterNodeCreator("TONNode", func() (Node, error) {
		return &TONNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, tonDescription)
}

func (t *TONNode) GetId() int {
	return t.id
}

func (t *TONNode) GetNodeType() string {
	return t.nodeType
}

func (t *TONNode) ProcessLogic() {
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

	if indexPT != -1 && indexIN != -1 && indexOut != -1 {
		if *t.input[indexIN].Input == 1 && t.oldIN == 0 {
			if !t.fired {
				//fmt.Println(*t.input[indexPT].Input)
				fmt.Println(time.Duration(*t.input[indexPT].Input) * time.Millisecond)
				t.timer = *time.NewTimer(time.Duration(*t.input[indexPT].Input) * time.Millisecond)
				t.fired = true
			}
			t.elapsed = false
			t.oldIN = 1
		}

		if *t.input[indexIN].Input == 0 && t.oldIN == 1 {
			t.elapsed = false
			t.oldIN = 0
			t.fired = false
			t.timer.Stop()
		}

		if *t.input[indexIN].Input == 1 && t.elapsed {
			t.output[indexOut].Output = 1
		} else {
			t.output[indexOut].Output = 0
		}
	} else {
		fmt.Println("TONNode: processLogic: Error: Input or output not found")
	}

}

func (t *TONNode) GetOutput(outName string) *OutputHandle {
	for i, name := range t.output {
		if name.Name == outName {
			return &t.output[i]
		}
	}
	return nil
}

func (t *TONNode) GetInput() []InputHandle {
	return t.input
}

func (t *TONNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle) {
	t.id = id_
	t.nodeType = nodeType_
	t.input = input_
	t.output = output_
}
