package nodes

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	_URL          = 0
	_USER         = 1
	_PASSWORD     = 2
	_HEADER_START = 3
)

type HttpNode struct {
	id                 int
	nodeType           string
	input              []InputHandle
	output             []OutputHandle
	parameterValueData []string
	client             *http.Client
	outputFlag         bool
	lastResponse       string
}

var httpDescription = nodeDescription{
	AccordionName: "Communication",
	PrimaryType:   "LogicalNode",
	Type_:         "ConfigurableNodeHttp",
	Display:       "HTTP Client Node",
	Label:         "HTTP Client",
	Stretchable:   false,
	Services:      []servicesStruct{},
	SubServices:   []subServicesStruct{},
	Input: []dataTypeNameStruct{
		{DataType: "bool", Name: "xSend"},
		{DataType: "value", Name: "url path"},
		{DataType: "value", Name: "method"}, // "GET" or "POST"
		{DataType: "value", Name: "body"},   // Only for POST
	},
	Output: []dataTypeNameStruct{
		{DataType: "bool", Name: "xDone"},
		{DataType: "value", Name: "response"},
	},
	ParameterNameData: []string{"url", "user", "password", "header Key 1", "header Value 1", "header Key 2", "header Value 2"},
}

func init() {
	RegisterNodeCreator("ConfigurableNodeHttp", func() (Node, error) {
		return &HttpNode{
			id:       -1,
			nodeType: "",
			input:    nil,
			output:   nil,
		}, nil
	}, httpDescription)
}

func (n *HttpNode) InitNode(id_ int, nodeType_ string, input_ []InputHandle, output_ []OutputHandle, parameterValueData_ []string) {
	n.id = id_
	n.nodeType = nodeType_
	n.input = input_
	n.output = output_
	n.parameterValueData = parameterValueData_
	n.client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	n.outputFlag = false
	n.lastResponse = ""
}

func (n *HttpNode) ProcessLogic() {
	if n.input == nil {
		n.output[0].Output = "0"
		return
	}
	if n.input[0].Input == nil {
		n.output[0].Output = "0"
		n.output[1].Output = ""
		return
	}

	if *n.input[0].Input == "1" {
		go func() {
			var url string
			if n.input[1].Input != nil && *n.input[1].Input != "empty" && *n.input[1].Input != "null" {
				urlPath := strings.Split(*n.input[1].Input, " ,, ")
				url = n.parameterValueData[_URL] + urlPath[0]
			} else {
				return
			}
			var method string
			if n.input[2].Input != nil {
				method = strings.ToUpper(*n.input[2].Input)
			}

			payload := ""
			if n.input[3].Input != nil && *n.input[3].Input != "empty" && *n.input[3].Input != "null" {
				payload = *n.input[3].Input
			}

			var req *http.Request
			var err error
			switch method {
			case "POST":
				req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(payload)))
			case "PUT":
				req, err = http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(payload)))
			case "PATCH":
				req, err = http.NewRequest(http.MethodPatch, url, bytes.NewBuffer([]byte(payload)))
			case "DELETE":
				req, err = http.NewRequest(http.MethodDelete, url, nil) // Or with body if needed
			case "HEAD":
				req, err = http.NewRequest(http.MethodHead, url, nil)
			case "OPTIONS":
				req, err = http.NewRequest(http.MethodOptions, url, nil)
			default:
				req, err = http.NewRequest(http.MethodGet, url, nil)
			}

			if err != nil {
				n.lastResponse = fmt.Sprintf("request error: %v", err)
				n.outputFlag = true
				return
			}

			// Optional headers
			for i := _HEADER_START; i+1 < len(n.parameterValueData); i += 2 {
				key := strings.TrimSpace(n.parameterValueData[i])
				value := strings.TrimSpace(n.parameterValueData[i+1])
				if key != "" && value != "" {
					req.Header.Set(key, value)
				}
			}

			// Optional authenticate
			if n.parameterValueData[_USER] != "" && n.parameterValueData[_PASSWORD] != "" {
				req.SetBasicAuth(n.parameterValueData[_USER], n.parameterValueData[_PASSWORD])
			}

			if n.client == nil {
				n.lastResponse = "http client is nil"
				n.outputFlag = true
				return
			}
			resp, err := n.client.Do(req)
			if err != nil {
				n.lastResponse = fmt.Sprintf("http error: %v", err)
				n.outputFlag = true
				return
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			n.lastResponse = string(body)
			n.outputFlag = true
		}()
	}

	if n.outputFlag {
		n.output[0].Output = "1"
		n.output[1].Output = n.lastResponse
		n.outputFlag = false
	} else {
		n.output[0].Output = "0"
	}
}

func (n *HttpNode) GetNodeType() string {
	return n.nodeType
}

func (n *HttpNode) GetId() int {
	return n.id
}

func (n *HttpNode) GetOutput(outName string) *OutputHandle {
	for i, name := range n.output {
		if name.Name == outName {
			return &n.output[i]
		}
	}
	return nil
}

func (n *HttpNode) GetInput() []InputHandle {
	return n.input
}

func (n *HttpNode) DestroyToBuildAgain() {
	n.client = nil
	n.outputFlag = false
	n.lastResponse = ""
}
