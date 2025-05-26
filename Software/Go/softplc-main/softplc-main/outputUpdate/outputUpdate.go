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

/*
Pour chaque Output, on parcourt toutes les connexions de sortie (OutputList)
associées à ce nœud pour détecter si une mise à jour est nécessaire.
La première boucle sert à itérer sur tous les nœuds de sortie du graphe de processus.
La seconde boucle parcourt chaque élément de la liste de sortie d'un nœud (c’est-à-dire les connexions
vers des services ou équipements externes).

Si nl’élément n’est pas un équipement (FriedlyName vide), on parcourt ensuite tous les états d'entrée/sortie
actuels (InputsOutputsState) pour retrouver l’état correspondant à ce service.

Si la valeur actuelle est différente de la valeur attendue (au-delà de la tolérance),
on construit une requête HTTP PUT pour mettre à jour la sortie concernée via l’API REST.
Selon le type de donnée (booléen ou flottant), on adapte la structure de la donnée et l'URL à utiliser.

Enfin, si l’élément est un équipement (FriendlyName non vide), on affiche un message car
la gestion des appliances n’est pas encore implémentée.
*/

func UpdateOutput() {

	for _, output := range processGraph.OutputNodes {
		for _, nodeOutputList := range output.GetOutputList() {
			if nodeOutputList.FriendlyName == "" { // Output isn't an appliance
				for _, ccIOState := range inputUpdate.InputsOutputsState {
					if ccIOState.Service == nodeOutputList.Service {
						ccIOState_, _ := strconv.ParseFloat(ccIOState.Value, 64)
						outputHandleInput_, _ := strconv.ParseFloat(*nodeOutputList.OutputHandle.Input, 64)

						if math.Abs(ccIOState_-outputHandleInput_) > tolerance {
							var data *strings.Reader
							var url string
							if nodeOutputList.OutputHandle.DataType == "bool" {
								if *nodeOutputList.OutputHandle.Input == "1" {
									data = strings.NewReader("true")
								} else {
									data = strings.NewReader("false")
								}
								if match, _ := regexp.MatchString(`^DO\d+$`, nodeOutputList.Service); match {
									doNb := strings.Trim(nodeOutputList.Service, "DO")
									doNbInt, _ := strconv.Atoi(doNb)
									doNb = strconv.Itoa(doNbInt - 1)
									url = "http://192.168.37.134:8888/api/v1/hal/do/" + doNb
								}
							} else {
								if match, _ := regexp.MatchString(`^AO\d+$`, nodeOutputList.Service); match {
									doNb := strings.Trim(nodeOutputList.Service, "AO")
									doNbInt, _ := strconv.Atoi(doNb)
									doNb = strconv.Itoa(doNbInt - 1)
									url = "http://192.168.37.134:8888/api/v1/hal/ao/" + doNb
									data = strings.NewReader(*nodeOutputList.OutputHandle.Input) //data = strings.NewReader(strconv.FormatFloat(*nodeOutputList.OutputHandle.Input, 'f', -1, 64))
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
