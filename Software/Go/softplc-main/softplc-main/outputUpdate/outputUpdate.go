package outputUpdate

import (
	"SoftPLC/inputUpdate"
	"SoftPLC/processGraph"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const tolerance = 0.01

func UpdateOutput() {
	for _, output := range processGraph.OutputNodes {
		for _, nodeOutputList := range output.GetOutputList() {
			if nodeOutputList.FriendlyName == "" { // Output isn't an appliance
				for _, ccIOState := range inputUpdate.InputsOutputsState {
					if ccIOState.Service == nodeOutputList.Service {
						if math.Abs(ccIOState.Value-*nodeOutputList.OutputHandle.Input) > tolerance {
							var data *strings.Reader
							var url string
							if nodeOutputList.OutputHandle.DataType == "bool" {
								if *nodeOutputList.OutputHandle.Input == 1 {
									data = strings.NewReader("true")
								} else {
									data = strings.NewReader("false")
								}
								if match, _ := regexp.MatchString(`^DO\d+$`, nodeOutputList.Service); match {
									doNb := strings.Trim(nodeOutputList.Service, "DO")
									doNbInt, _ := strconv.Atoi(doNb)
									doNb = strconv.Itoa(doNbInt - 1)
									url = "http://192.168.1.175:8888/api/v1/hal/do/" + doNb
								}
							} else {
								if match, _ := regexp.MatchString(`^AO\d+$`, nodeOutputList.Service); match {
									doNb := strings.Trim(nodeOutputList.Service, "AO")
									doNbInt, _ := strconv.Atoi(doNb)
									doNb = strconv.Itoa(doNbInt - 1)
									url = "http://192.168.1.175:8888/api/v1/hal/ao/" + doNb
									data = strings.NewReader(strconv.FormatFloat(*nodeOutputList.OutputHandle.Input, 'f', -1, 64))
								}
							}
							req, err := http.NewRequest(http.MethodPut, url, data)
							if err != nil {
								fmt.Println(err)
							}
							req.Header.Set("Content-Type", "application/json")

							client := &http.Client{}
							resp, err := client.Do(req)
							if err != nil {
								fmt.Println(err)
							}
							if resp.StatusCode != 204 {
								fmt.Println("Error while updating output on " + nodeOutputList.Service)
							}
							resp.Body.Close()
						}
					}
				}
			} else { // Output is an appliance
				//TODO
				fmt.Println("pas encore fait pour les appliance")
			}
		}
	}
}
