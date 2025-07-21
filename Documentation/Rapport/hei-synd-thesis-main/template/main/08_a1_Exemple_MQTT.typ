
== MQTT configuration détaillé <sec:mqttConfiguration>
Aucun des paramètres n'est obligatoire, par défaut le port est *1883* et le serveur tourne sur l'automate. Les "Settings" rentrés dans la vue sont données par *parameterValueData*.
#figure(
  image("/resources/img/61_BlocMqtt_Configuration.png", width: 90%),
  caption: [
    exemple de paramétrage de "Mqtt"
  ],
)
#figure(
    align(left,
    ```rust
      "inputHandle": [
          {
            "dataType": "bool",
            "name": "xEnable"
          },
          {
            "dataType": "value",
            "name": "topicToSend"
          },
          {
            "dataType": "value",
            "name": "msgToSend"
          },
          {
            "dataType": "value",
            "name": "topicToReceive"
          }
        ],
        "label": "MQTT",
        "outputHandle": [
          {
            "dataType": "bool",
            "name": "xDone"
          },
          {
            "dataType": "value",
            "name": "msg"
          }
        ],
        "parameterNameData": [
          "broker",
          "port",
          "user",
          "password"
        ],
        "parameterValueData": [
          "broker.hivemq.com",
          "1883",
          "",
          ""
        ]
    ```
    ),
    caption: [*MQTT* : extrait de la structure JSON d'un exemple],
  
  )

  == Exemples MQTT <sec:exempleUtililationMqttSimple-vs-vue>
#figure(
  image("/resources/img/61_BlocMqtt_Configuration.png", width: 100%),
  caption: [
    Settings MQTT utilisés pour les exemples 
  ],
)
#label("fig:BlocMqtt_Configuration-vs-vue")
=== Exemple 1
Dans cet exemple, on reçoit les messages envoyés sur les 3 topics topic/test/1, topic/test/2, topic/test/3 (@fig:BlocMqtt_exemple_1_msg-vs-vue). Cet exemple montre l'envoi d'un message sur le topic "topic/test/100" (@fig:BlocMqtt_exemple_1_envoyer_1-vs-vue) et sur les topics "topic/test/101", "topic/test/102" en simultané (@fig:BlocMqtt_exemple_1_envoyer_2-vs-vue). Il montre également ce qui se passe lorsque l'on reçoit plusieurs topics en même temps (@fig:BlocMqtt_exemple_1_msg_simultané-vs-vue). On remarque également que les messages reçus sont affichés dans le même ordre que les topics à recevoir, ce qui facilite son utilisation.

#figure(
  image("/resources/img/62_BlocMqtt_exemple_TestClient_2.png", width: 100%),
  caption: [
    MQTT.cool : visualisation des tests clients réalisés pour l'exemple 1
  ],
)  
#set page(
  flipped: true,
)

#pagebreak()

#figure(
  image("/resources/img/61_BlocMqtt_exemple_init.png", width: 100%),
  caption: [
    MQTT : Vue debug initiale 
  ],
)
#label("fig:BlocMqtt_exemple_1_init-vs-vue")

#figure(
  image("/resources/img/61_BlocMqtt_exemple_msg_1.png", width: 100%),
  caption: [
    MQTT : Vue debug message reçu ("1er message")
  ],
)
#label("fig:BlocMqtt_exemple_1_msg-vs-vue")

#figure(
  image("/resources/img/61_BlocMqtt_exemple_Rtrig.png", width: 100%),
  caption: [
    MQTT : Vue debug impulsions _Rising edge_ -> message envoyé "Hi" sur le topic "topic/test/100" 
  ],
)
#label("fig:BlocMqtt_exemple_1_envoyer_1-vs-vue")

#figure(
  image("/resources/img/61_BlocMqtt_exemple_Ftrig.png", width: 100%),
  caption: [
    MQTT : Vue debug impulsions _Falling edge_ -> messages envoyés : "Bye" sur le topic "topic/test/101" et "Have a nice day !" sur le topic "topic/test/102"
  ],
)
#label("fig:BlocMqtt_exemple_1_envoyer_2-vs-vue")

#figure(
  image("/resources/img/61_BlocMqtt_exemple_msg_simultané.png", width: 100%),
  caption: [
    MQTT : Vue debug message reçu des 3 topics en simultané
  ],
)
#label("fig:BlocMqtt_exemple_1_msg_simultané-vs-vue")

=== Exemple 2
Cet exemple est similaire au précédent, mais cette fois-ci, on traite la sortie *"msgLastReceived"* pour activer des sorties quand on reçoit un message *1 topic _x_* sur le topic x et désactiver les sorties quand on reçoit un message *0 topic _x_* sur le topic x. Cet exemple démontre très bien la puissance que peut avoir notre programme.

#figure(
  image("/resources/img/62_BlocMqtt_exemple_traitementSortie.png", width: 100%),
  caption: [
    MQTT : Exemple traitement sortie *"msgLastReceived"*
  ],
)
#label("fig:BlocMqtt_exemple_2_traitementSortie-vs-vue")