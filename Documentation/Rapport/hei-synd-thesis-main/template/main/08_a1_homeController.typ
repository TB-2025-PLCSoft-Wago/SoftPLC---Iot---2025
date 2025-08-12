/*
#pagebreak()
= Exemples codes : vue programme + JSON

== Exemple WebSocket <sec:exempleUtililationWebSocket-vs-vue>
#figure(
  image("/resources/img/19_OutputPatchWDA.png", width: 100%),
  caption: [
    *WebSocket* : vue programmation
  ],
)
#label("fig:toto-vs-vue")
#figure(
    align(left,
    ```json
      {
        "data": {
          "id": "0-0-io-channels-9-dovalue",
          "type": "parameters",
          "attributes": {
            "value": true
          }
        }
      }
    ```
    ),
    caption: [*WebSocket* : graphique format JSON],
  
  )
  #label("fig:toto-vs-vue")
  */

  /*

          ```
    ),
    
  )
  #figure(
   
  align(left,
    ```json

  */
  #pagebreak()
= Application: Home Controller <sec:homeController-vs-vue>

== Reçevoir commande boutons - HTTP serveur <sec:HTTPServeur_boutons-vs-vue>
#figure(
  image("/resources/img/83_ProofOfConcept_Boutton_Config.png", width: 100%),
  caption: [
    *Gestion boutons* : vue programmation - configuration du bloc _HTTP serveur_
  ],
)
#label("fig:83_ProofOfConcept_Boutton_Config-vs-vue")
#figure(
  image("/resources/img/83_ProofOfConcept_Boutton_complet.png", width: 100%),
  caption: [
    *Gestion boutons* : vue programmation - complet
  ],
)
#label("fig:83_ProofOfConcept_Boutton_Config-vs-vue")
== Gestion enclenchement du chauffage <sec:enclenchementChauffage-vs-vue>
L'idée est de faire du tout ou rien sur un relais _myStrom_. Cependant, on offre également la possibilitée d'activer ou désactiver manuellement via l'interface. Sur l'input _topicToReceive_ du bloc _MQTT_, on a la constante "shellies/shelly-s1-Heater-Home/sensor/temperature" qu'on ne voit pas en complet sur le capture d'écran.
#figure(
  image("/resources/img/80_Proof_of_concept_Heater_Interface.png", width: 100%),
  caption: [
    *Gestion du chauffage* : vue user
  ],
)
#label("fig:80_Proof_of_concept_Heater_Interface-vs-vue")
#figure(
  image("/resources/img/80_Proof_of_concept_Heater_config_mqtt.png", width: 100%),
  caption: [
    *Gestion du chauffage* : vue programmation - configuration du bloc _MQTT_
  ],
)
#label("fig:80_Proof_of_concept_Heater_config_mqtt-vs-vue")
#figure(
  image("/resources/img/80_Proof_of_concept_Heater_config_httpClient.png", width: 100%),
  caption: [
    *Gestion du chauffage* : vue programmation - configuration du bloc _HTTP Client_
  ],
)
#label("fig:80_Proof_of_concept_Heater_config_mqtt-vs-vue")
#set page(
  flipped: true,
)
#figure(
  image("/resources/img/80_Proof_of_concept_Heater_Complet.png", width: 120%),
  caption: [
    *Gestion du chauffage* : vue programmation - vue en complet
  ],
)
#label("fig:80_Proof_of_concept_Heater_Complet-vs-vue")

#figure(
  image("/resources/img/80_Proof_of_concept_Heater_MQTT_part1.png", width: 100%),
  caption: [
    *Gestion du chauffage* : vue programmation - MQTT attente température
  ],
)
#label("fig:80_Proof_of_concept_Heater_MQTT_part1-vs-vue")

#figure(
  image("/resources/img/80_Proof_of_concept_Heater_MQTT_part2_GestionLogiqueTemperature.png", width: 100%),
  caption: [
    *Gestion du chauffage* : vue programmation - Gestion automatique selon température (seuil)
  ],
)
#label("fig:80_Proof_of_concept_Heater_MQTT_part2_GestionLogiqueTemperature-vs-vue")

#figure(
  image("/resources/img/80_Proof_of_concept_Heater_SendToMyStrom.png", width: 100%),
  caption: [
    *Gestion du chauffage* : vue programmation - envoi au relais de _My Strom_ (*on* ou *off*)
  ],
)
#label("fig:80_Proof_of_concept_Heater_MQTT_part2_GestionLogiqueTemperature-vs-vue")

#set page(
  flipped: false,
)
== Gestion porte du garage <sec:porteGarage-vs-vue>

#figure(
  image("/resources/img/82_ProofOfConcept_Garage_interface.png", width: 100%),
  caption: [
    *porte du garage* : vue user
  ],
)
#label("fig:82_ProofOfConcept_Garage_interface-vs-vue")
=== Modbus - porte du garage
#figure(
  image("/resources/img/82_ProofOfConcept_Garage_Config_Modbus.png", width: 100%),
  caption: [
    *porte du garage* : vue programmation - Modbus - configuration
  ],
)
#label("fig:82_ProofOfConcept_Garage_Config_Modbus-vs-vue")
#set page(
  flipped: true,
)
#figure(
  image("/resources/img/82_ProofOfConcept_Garage_Read_Modbus.png", width: 100%),
  caption: [
    *garage* : vue programmation - Modbus - lecture
  ],
)
#label("fig:82_ProofOfConcept_Garage_Read_Modbus-vs-vue")

#figure(
  image("/resources/img/82_ProofOfConcept_Garage_write_Modbus.png", width: 100%),
  caption: [
    *garage* : vue programmation - Modbus - écriture
  ],
)
#label("fig:82_ProofOfConcept_Garage_Read_Modbus-vs-vue")
#set page(
  flipped: false,
)
=== logique - porte du garage
#figure(
  image("/resources/img/82_ProofOfConcept_Log_Garage_Complet.png", width: 100%),
  caption: [
    *porte du garage* : vue programmation - logique - vue en complet
  ],
)
#label("fig:82_ProofOfConcept_Log_Garage_Complet-vs-vue")

#set page(
  flipped: false,
)
#figure(
  image("/resources/img/82_ProofOfConcept_Log_Garage_dontMove_LastGo.png", width: 120%),
  caption: [
    *porte du garage* : vue programmation - logique - écriture des variables "door don't move" et "Garage door last go up"
  ],
)
#label("fig:82_ProofOfConcept_Log_Garage_dontMove_LastGo-vs-vue")

#figure(
  image("/resources/img/82_ProofOfConcept_Log_Garage_dontMove_LastGo.png", width: 120%),
  caption: [
    *porte du garage* : vue programmation - logique - écriture des variables "Garage door open" et "Garage door close"
  ],
)
#label("fig:82_ProofOfConcept_Log_Garage_dontMove_LastGo-vs-vue")







#pagebreak()
#set page(
  flipped: false,
)
== Gestion de la lampe de couleur <sec:lampe-vs-vue>
L'idée de ce programme est de permettre à l'utilisateur de changer la couleur et luminosité de la lampe. C'est aussi lui qui est responsable d'allumer la lampe en vert quand elle est ouverte au complet sinon en rouge.
#figure(
  image("/resources/img/81_Proof_of_concept_LightColor_interfacepng.png", width: 80%),
  caption: [
    *Gestion de la lampe de couleur _Shelly_* : vue user
  ],
)
#label("fig:81_Proof_of_concept_LightColor_interfacepng-vs-vue")

#figure(
  image("/resources/img/81_Proof_of_concept_LightColor_configurationHTTPClient.png", width: 100%),
  caption: [
    *Gestion de la lampe de couleur _Shelly_* : vue programmation - configuration du bloc _HTTP Client_
  ],
)
#label("fig:80_Proof_of_concept_Heater_config_mqtt-vs-vue")

#figure(
  image("/resources/img/81_Proof_of_concept_LightColor_Complet.png", width: 120%),
  caption: [
    *Gestion de la lampe de couleur _Shelly_* : vue programmation - vue en complet
  ],
)
#label("fig:81_Proof_of_concept_LightColor_Complet")



#figure(
  image("/resources/img/81_Proof_of_concept_LightColor_ConstructionTrameInterface.png", width: 120%),
  caption: [
    *Gestion de la lampe de couleur _Shelly_* : vue programmation - création url path pour changement des couleurs
  ],
)
#label("fig:81_Proof_of_concept_LightColor_Complet")

#figure(
  image("/resources/img/81_Proof_of_concept_LightColor_Luminosite_ConstructionTrameInterface.png", width: 120%),
  caption: [
    *Gestion de la lampe de couleur _Shelly_* : vue programmation - création url path pour changement de luminosité
  ],
)
#label("fig:81_Proof_of_concept_LightColor_Luminosite_ConstructionTrameInterface")

#figure(
  image("/resources/img/81_Proof_of_concept_LightColor_garage.png", width: 120%),
  caption: [
    *Gestion de la lampe de couleur _Shelly_* : vue programmation - création url path pour changement des couleurs - *porte garage*
  ],
)
#label("fig:81_Proof_of_concept_LightColor_garage")
== Gestion lumières <sec:lampe2-vs-vue>
#figure(
  image("/resources/img/97_light_general.png", width: 120%),
  caption: [
    *Gestion des lumières garage et cuisine* : vue programmation - vue en complet
  ],
)

#figure(
  image("/resources/img/97_light_cuisineRead.png", width: 120%),
  caption: [
    *Gestion des lumières garage et cuisine* : vue programmation - lecture cuisine
  ],
)
#figure(
  image("/resources/img/97_light_cuisine_write.png", width: 120%),
  caption: [
    *Gestion des lumières garage et cuisine* : vue programmation - écriture cuisine
  ],
)
#figure(
  image("/resources/img/97_light_garage_read.png", width: 120%),
  caption: [
    *Gestion des lumières garage et cuisine* : vue programmation - logique garage
  ],
)
#figure(
  image("/resources/img/97_light_garage_read.png", width: 120%),
  caption: [
    *Gestion des lumières garage et cuisine* : vue programmation - écriture garage
  ],
)