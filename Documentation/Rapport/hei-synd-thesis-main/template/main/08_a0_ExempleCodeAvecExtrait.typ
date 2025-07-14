#pagebreak()
= Exemples codes : vue programme + JSON
/*
== Exemple WebSocket <sec:exempleUtililationWebSocket-vs-vue>
#figure(
  image("/resources/img/19_OutputPatchWDA.png", width: 100%),
  caption: [
    *WebSocket* : vue programmation
  ],
)
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
  */

  /*

          ```
    ),
    
  )
  #figure(
   
  align(left,
    ```json

  */
== Exemple mode debug : simple <sec:exempleUtililationDebugSimple-vs-vue>
L'objectif de cet exemple est de démontrer les changements effectués sur le graphique lorsqu'on passe en mode debug. Les modifications sont faites sur les *edges*. La principale différence entre "debug 1" et "debug 2" est l'actionnement du bouton DI1. La valeur affichée est transmise avec la structure suivante @fig:structurDebug-vs-vue. À cela sont rajoutés, le style, l'animation et autres.
#figure(
  align(left,
    ```json
    "data": {
        "label": "valeur affichée"
      },
    ```
  ),
  caption: [*Mode debug* : extrait structure envoi données valeur affichée],
)
#label("fig:structurDebug-vs-vue")

  
#figure(
  image("/resources/img/53_ExempleModeDebugSimpleProgrammation.png", width: 100%),
  caption: [
    *Mode debug* _simple_ : vue programmation
  ],
)
#figure(
  image("/resources/img/53_ExempleModeDebugSimpleDebug1.png", width: 100%),
  caption: [
    *Mode debug* _simple_ : vue debug 1
  ],
)
#figure(
  image("/resources/img/53_ExempleModeDebugSimpleDebug2.png", width: 100%),
  caption: [
    *Mode debug* _simple_ : vue debug 2 (DI1 = true)
  ],
)
#figure(
    align(left,
    ```json
      {
  "edges": [
    {
      "id": "reactflow__edge-0Output-3Input0",
      "selected": false,
      "source": "0",
      "sourceHandle": "Output",
      "style": {
        "stroke": "#000000",
        "strokeWidth": 1
      },
      "target": "3",
      "targetHandle": "Input0",
      "type": "step"
    },
    {
      "id": "reactflow__edge-3Output-1Input",
      "source": "3",
      "sourceHandle": "Output",
      "style": {
        "stroke": "#30362F",
        "strokeWidth": 1
      },
      "target": "1",
      "targetHandle": "Input",
      "type": "step"
    },
    {
      "id": "reactflow__edge-16Output-3Input1",
      "source": "16",
      "sourceHandle": "Output",
      "style": {
        "stroke": "#30362F",
        "strokeWidth": 1
      },
      "target": "3",
      "targetHandle": "Input1",
      "type": "step"
    }
    
  ],
  "nodes": [
    {
        ...
    ```
    ),
    caption: [*Mode debug* _simple_ : Edges graphique format JSON (vue programmation)],
  
  )


  #figure(
    align(left,
    ```json
     "edges": [
    {
      "data": {
        "label": "0"
      },
      "id": "reactflow__edge-0Output-3Input0",
      "label": "???",
      "labelBgBorderRadius": 4,
      "labelBgPadding": [
        8,
        4
      ],
      "labelBgStyle": {
        "color": "#333",
        "fill": "white",
        "fillOpacity": 0.5
      },
      "selected": false,
      "source": "0",
      "sourceHandle": "Output",
      "style": {
        "animation": "dash 1s linear infinite",
        "stroke": "#FF0000",
        "strokeDasharray": "8 4",
        "strokeWidth": 1
      },
      "target": "3",
      "targetHandle": "Input0",
      "type": "customDebugEdge"
    },
    {
      "data": {
        "label": "0"
      },
      "id": "reactflow__edge-3Output-1Input",
      "label": "???",
      "labelBgBorderRadius": 4,
      "labelBgPadding": [
        8,
        4
      ],
      "labelBgStyle": {
        "color": "#333",
        "fill": "white",
        "fillOpacity": 0.5
      },
      "source": "3",
      "sourceHandle": "Output",
      "style": {
        "animation": "dash 1s linear infinite",
        "stroke": "#FF0000",
        "strokeDasharray": "8 4",
        "strokeWidth": 1
      },
      "target": "1",
      "targetHandle": "Input",
      "type": "customDebugEdge"
    },
    ```
    ),
    
  )
  #figure(
   
  align(left,
    ```json
    {
      "data": {
        "label": "0"
      },
      "id": "reactflow__edge-16Output-3Input1",
      "label": "???",
      "labelBgBorderRadius": 4,
      "labelBgPadding": [
        8,
        4
      ],
      "labelBgStyle": {
        "color": "#333",
        "fill": "white",
        "fillOpacity": 0.5
      },
      "source": "16",
      "sourceHandle": "Output",
      "style": {
        "animation": "dash 1s linear infinite",
        "stroke": "#FF0000",
        "strokeDasharray": "8 4",
        "strokeWidth": 1
      },
      "target": "3",
      "targetHandle": "Input1",
      "type": "customDebugEdge"
    }
  ],
  "nodes": [
    {
        ...
    ```
    ),
    caption: [*Mode debug* _simple_ : Edges graphique format JSON (vue debug 1)],
  
  )

#figure(
    align(left,
    ```json
    "edges": [
    {
      "data": {
        "label": "1"
      },
      "id": "reactflow__edge-0Output-3Input0",
      "label": "???",
      "labelBgBorderRadius": 4,
      "labelBgPadding": [
        8,
        4
      ],
      "labelBgStyle": {
        "color": "#333",
        "fill": "white",
        "fillOpacity": 0.5
      },
      "selected": false,
      "source": "0",
      "sourceHandle": "Output",
      "style": {
        "animation": "dash 1s linear infinite",
        "stroke": "#00FF00",
        "strokeDasharray": "8 4",
        "strokeWidth": 1
      },
      "target": "3",
      "targetHandle": "Input0",
      "type": "customDebugEdge"
    },
    {
      "data": {
        "label": "1"
      },
      "id": "reactflow__edge-3Output-1Input",
      "label": "???",
      "labelBgBorderRadius": 4,
      "labelBgPadding": [
        8,
        4
      ],
      "labelBgStyle": {
        "color": "#333",
        "fill": "white",
        "fillOpacity": 0.5
      },
      "source": "3",
      "sourceHandle": "Output",
      "style": {
        "animation": "dash 1s linear infinite",
        "stroke": "#00FF00",
        "strokeDasharray": "8 4",
        "strokeWidth": 1
      },
      "target": "1",
      "targetHandle": "Input",
      "type": "customDebugEdge"
    },
    ```
    ),
    
  )
  #figure(
   
  align(left,
    ```json
    {
      "data": {
        "label": "0"
      },
      "id": "reactflow__edge-16Output-3Input1",
      "label": "???",
      "labelBgBorderRadius": 4,
      "labelBgPadding": [
        8,
        4
      ],
      "labelBgStyle": {
        "color": "#333",
        "fill": "white",
        "fillOpacity": 0.5
      },
      "source": "16",
      "sourceHandle": "Output",
      "style": {
        "animation": "dash 1s linear infinite",
        "stroke": "#FF0000",
        "strokeDasharray": "8 4",
        "strokeWidth": 1
      },
      "target": "3",
      "targetHandle": "Input1",
      "type": "customDebugEdge"
    }
  ],
  "nodes": [
    {
        ...
    ```
    ),
    caption: [*Mode debug* _simple_ : Edges graphique format JSON (vue debug 2)],
  
  )

#pagebreak()
== Exemples MQTT : simple <sec:exempleUtililationMqttSimple-vs-vue>
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
