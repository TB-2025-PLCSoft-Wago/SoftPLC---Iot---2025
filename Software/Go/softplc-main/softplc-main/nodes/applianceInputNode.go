package nodes

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApplianceInputNode struct {
	id       int
	nodeType string
	output   []InputNodeHandle
}

func (a *ApplianceInputNode) GetId() int {
	return a.id
}

func (a *ApplianceInputNode) GetNodeType() string {
	return a.nodeType
}

func (a *ApplianceInputNode) InitNode(id_ int, nodeType_ string, input_ []InputNodeHandle) {
	a.id = id_
	a.nodeType = nodeType_
	a.output = input_
}

func (a *ApplianceInputNode) GetOutput(outName string) *InputNodeHandle {
	for i, name := range a.output {
		if name.InputHandle.Name == outName {
			return &a.output[i]
		}
	}
	return nil
}

func init() {
	var services []servicesStruct
	var subServices []subServicesStruct
	var result map[string]interface{}
	var tabResult []map[string]interface{}
	resp, err := http.Get("http://192.168.1.175:8888/api/v1/appliance/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &tabResult)
	if err != nil {
		panic(err)
	}
	var friendlyNameList []string
	var idList []string
	var servicesList [][]string
	for _, v := range tabResult {
		var servicesOfThisId []string
		friendlyNameList = append(friendlyNameList, v["friendlyName"].(string))
		idList = append(idList, v["id"].(string))
		for _, s := range v["services"].([]interface{}) {
			servicesOfThisId = append(servicesOfThisId, s.(string))
		}
		servicesList = append(servicesList, servicesOfThisId)
	}
	for i, actualId := range idList {
		var updateServicesList []string
		for _, v := range servicesList[i] {
			var subServ []dataTypeNameStruct
			res, err := http.Get("http://192.168.1.175:8888/api/v1/appliance/" + actualId + "/" + v)
			if err != nil {
				panic(err)
			}
			body, err = io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(body, &result)
			if err != nil {
				panic(err)
			}
			res.Body.Close()
			if len(result["dataPoints"].(map[string]interface{})) != 0 {
				updateServicesList = append(updateServicesList, v)
				for key, value := range result["dataPoints"].(map[string]interface{}) {
					if value.(map[string]interface{})["type"].(string) == "NumericReading" {
						subServ = append(subServ, dataTypeNameStruct{
							DataType: "value",
							Name:     key,
						})
					} else {
						subServ = append(subServ, dataTypeNameStruct{
							DataType: "bool",
							Name:     key,
						})
					}
				}
				subServices = append(subServices, subServicesStruct{
					FriendlyName: friendlyNameList[i],
					Primary:      v,
					Secondary:    subServ,
				})
			}
		}
		services = append(services, servicesStruct{
			FriendlyName: friendlyNameList[i],
			NameServices: updateServicesList,
		})
	}

	var applianceNodeDescription = nodeDescription{
		AccordionName: "Input",
		PrimaryType:   "inputNode",
		Type_:         "appliancesInput",
		Display:       "Appliance Input",
		Label:         "Input",
		Stretchable:   false,
		Services:      services,
		SubServices:   subServices,
		Input:         []dataTypeNameStruct{},
		Output:        []dataTypeNameStruct{{DataType: "variable", Name: "Output"}},
	}
	RegisterNodeCreator("appliancesInput", func() (Node, error) {
		return &ApplianceInputNode{
			id:       -1,
			nodeType: "",
			output:   nil,
		}, nil
	}, applianceNodeDescription)
}
