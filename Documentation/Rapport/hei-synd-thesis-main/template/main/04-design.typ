#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= #i18n("design-title", lang:option.lang) <sec:design>

#option-style(type:option.type)[
  In the design section of your bachelor thesis, you have the opportunity to provide a detailed blueprint of the system you intend to develop or analyze. This section serves as the foundation upon which your implementation will be built. Here's how you can enrich and expand upon this section:

  - *System Overview*: Begin by providing a comprehensive overview of the system under consideration.
  - *Requirements Specification*: Outline the specific requirements that your system must fulfill.
  - *Architecture and Design Principles*: Delve into the architectural design of your system, elucidating the underlying principles and design decisions that govern its structure.
  - *Technology Stack*: Detail the technologies and tools that will be employed in the development of your system.
  - *Data Management and Storage*: If your system involves the management or manipulation of data, provide insights into how data will be structured, stored, and accessed.
  - * User Interface (UI) Design*: If applicable, describe the user interface of your system, focusing on usability, accessibility, and user experience (UX) design principles.
  - *Integration and Interoperability*: Address how your system will integrate with existing systems or external services, if relevant.

  Dans la section de conception de votre bachelor thesis, vous avez l’opportunité de fournir un plan détaillé du système que vous souhaitez développer ou analyser. Cette section sert de base sur laquelle votre mise en œuvre sera construite. Voici comment vous pouvez enrichir et développer cette section :

  - *Aperçu du système* : Commencez par donner un aperçu complet du système à l’étude.
  - *Spécification des exigences* : Décrivez les exigences spécifiques que votre système doit respecter.
  - *Principes d’architecture et de conception* : Plongez dans la conception architecturale de votre système, en élucidant les principes sous-jacents et les décisions de conception qui régissent sa structure.
  - *Technology Stack* : Détaillez les technologies et les outils qui seront utilisés dans le développement de votre système.
  - *Gestion et stockage des données* : Si votre système implique la gestion ou la manipulation de données, fournissez des informations sur la façon dont les données seront structurées, stockées et accessibles.
  - *Conception de l’interface utilisateur (UI) *: Le cas échéant, décrivez l’interface utilisateur de votre système en mettant l’accent sur la convivialité, l’accessibilité et les principes de conception de l’expérience utilisateur.
  - *Intégration et interopérabilité* : Expliquez comment votre système s’intégrera aux systèmes existants ou aux services externes, le cas échéant.
]

#lorem(50)

#add-chapter(
  after: <sec:design>,
  before: <sec:impl>,
  minitoc-title: i18n("toc-title", lang: option.lang)
)[
  
  //section environment de développement
  #include "/main/01-Environnement.typ"
  #pagebreak()
  == Aperçu du système
  /*Le système que nous allons développer est un #gls("HAL") (Hardware Abstraction Layer) pour la programmation d'automates WAGO. Il permettra de créer des programmes PLC en utilisant une interface graphique, sans avoir à écrire de code. Le système sera divisé en plusieurs parties, chacune ayant un rôle spécifique.*/
  L'aperçu du système est présenté dans le schéma de principe @fig:schemaPrincipe-vs-vue. Il montre les différentes parties du système et comment elles interagissent entre elles. Le schéma n'est pas exhaustif, mais il donne une idée générale des grands principes du système. 
  #figure(
  image("/resources/img/26_schemaPrincipe.png", width: 130%),
  caption: [
    aperçu du système - schéma de principe
  ],
)
#label("fig:schemaPrincipe-vs-vue")

  #pagebreak()
== Bloc de haut niveau

Il existe plusieurs manière d’aborder le problème. Une des approches est de repérer les points commun entre ces blocs de haut niveau pour essayer d’en tirer une forme générique. On remarque que tous ces blocs ont pour objectif de transmettre et recevoir des données. Il faudra donc commencer par le développement de bloc commun pour une communication. Il faut également des blocs permettant de travailler avec des STRING. Le schéma figure 1 montre le concept d’une telle structure avec tous les blocs qui devront être développé autour pour pouvoir créer une communication. 
#figure(
  image("/resources/img/01_communication_principe.png", width: 110%),
  caption: [
    Communication principe
  ],
)
#label("fig:communication-Bloc-principe")
#ideabox()[
L’idée étant d’avoir un bloc communication qui s’occupe de la configuration étant différente pour MQTT, HTTP, CAN, etc. Sur le quel, on pourra double cliqués pour accéder à la page de configuration. Sur ce bloc de communication, on pourait ensuite venir lier nos 2 blocs permettant la transition de boolean vers nos trame. Par le future, en mode debug, l’utilisateur pourra voir l’état de la communication grace au "bloc de communication" et voir ce que la logique combinatoire transmet comme trame grâce au bloc en vert.
]

  #pagebreak()


  == User Interface (UI) Design
  L'interface utilisateur utilise _React_, la documentation @UsingTypeScriptReact peut aider. On utilise également _React flow_, dont la documentation @Quickstart2025, qui s'occupe plus particulièrement des mécanismes liés aux Nodes, Edges, Handle, etc.

  === Vue programmation

  La vue programmation permet de créer des programmes PLC en utilisant une interface graphique. Elle est représentée par la "Programming Page" sur le schéma @fig:schemaPrincipe-vs-vue. Cette vue permet de créer par *drag and drop* des blocs logiques, des *Inputs* et des *Outputs*, et de les connecter entre eux pour créer un programme PLC.

  === Vue WebSocket <sec:websocketVUE>

  La vue WebSocket permet de visualiser en temps réel l'état des Inputs et Outputs spécifiques à cette vue. Sur le schéma @fig:schemaPrincipe-vs-vue, elle est représentée par la "View Page".

  Cette vue s'accompagne de blocs logiques qui permettent sa création. Ces blocs sont représentés dans la figure @fig:blocWebSocket-vs-vue. Ils permettent de créer des Inputs et Outputs spécifiques à cette vue, permettant ainsi de visualiser l'état de n'importe quelle connexion dans le programme PLC. Un exemple de ce à quoi pourrait ressembler la vue WebSocket est présenté dans la figure @fig:vuePrincipeWebSocket-vs-vue, et le programme qui crée cette vue est présenté dans la figure @fig:ProgrammePrincipeWebSocket-vs-vue. Les parties _recevoir_ et _envoyer_ les messages sont prévues pour du débogage. Les autres parties sont prévues pour contrôler et tester le programme PLC. Remarquez l’impact du champ _Appliance Name_, qui permet de créer des groupes dans la vue WebSocket. Cela facilite le regroupement des Inputs et Outputs par groupe, ce qui est très utile pour la visualisation. L'idée a été inspirée de "Remote Controller" vue en cours de Data Engineering, où une page WebSocket est générée à partir d'un fichier JSON, dont voici le lien https://cyberlearn.hes-so.ch/pluginfile.php/3312976/mod_resource/content/2/index.html. 
  
  #figure(
  image("/resources/img/27_blocsWebSocket_InOut.png", width: 80%),
  caption: [
    blocs vue WebSocket - Inputs et Outputs
  ],
)
#label("fig:blocWebSocket-vs-vue")


#figure(
  image("/resources/img/29_programmePrincipeWebSocket_1.png", width: 120%),
  /*caption: [
    exemple programme WebSocket
  ],*/
)

#figure(
  image("/resources/img/29_programmePrincipeWebSocket_2.png", width: 110%),
  caption: [
    exemple programme WebSocket
  ],
)
#label("fig:ProgrammePrincipeWebSocket-vs-vue")


#figure(
  image("/resources/img/28_VisuWebSocket_Principe.png", width: 100%),
  caption: [
    exemple vue WebSocket (commentaires en brun)
  ],
)
#label("fig:vuePrincipeWebSocket-vs-vue")



#pagebreak()
== Gestion et stockage des données
La gestion et le stockage des données sont des aspects importantes du système. Dans notre cas, les données sont stockées et transmises en fromat JSON. 
=== Node MQTT 
Aucun des paramètres n'est obligatoire, par défaut le port est *1883* et le serveur tourne sur l'automate. Les "Settings" rentrés dans la vue sont données par *parameterValueData*.
#figure(
  image("/resources/img/33_mqtt_settingExemple.png", width: 90%),
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
    caption: [mqtt, extrait de la structure JSON d'un exemple],
  
  )
=== Node HTTP client
Le package Go @HttpPackageNetb a été trouvé.
Pour le Node HTTP client, il est possible de configurer les paramètres suivants :
- *URL* : l'URL de la requête HTTP.
- *user* : l'utilisateur pour l'authentification HTTP.
- *password* : le mot de passe pour l'authentification HTTP.
- *Headers* : les en-têtes HTTP à envoyer avec la requête.
#infobox()[les _Headers_ sont des paires clé-valeur, par exemple : `{"Content-Type": "application/json"}`. Il faut donc deux paramètres pour chaque Header. De plus, il faut que ce soit possible de mettre plusieurs Headers. ]
  
Le bloc peut prendre dynamiquement les paramètres suivants :
- *xSend* : un booléen pour envoyer lorsque qu'il passe à _true_.
- *url path* : la suite du chemin de l'URL de la requête HTTP.Il est ajouté à la suite du paramètre _URL_ pour donner l'URL final.
- *Method* : la méthode HTTP à utiliser (GET, POST, PATCH, PUT, DELETE, HEAD, OPTIONS), par défaut GET.
- *Body* : le corps de la requête HTTP, qui peut être au format JSON ou autre.
Le bloc nous retournera les paramètres suivants :
- *xDone* : un booléen pour indiquer si la requête a été effectuée avec succès.
- *Response* : la réponse de la requête HTTP.
/*
#figure(
  image("/resources/img/34_http_settingExemple.png", width: 90%),
  caption: [
    exemple de paramétrage de "HTTP"
  ],
)*/
=== Node HTTP serveur
Le package Go @HttpPackageNetb a été trouvé.
L'exemple @soysouvanhClientsServeursHTTP permet d'en comprendre d'avantage sur la création d'un serveur HTTP en Go. Pour le déploiement d'un serveur HTTP sur Docker, il a été trouvé la documentation @nicholsonCraignicholsonSimplehttp2023.

Le Node HTTP serveur permet de créer un serveur HTTP qui écoute les requêtes entrantes. Le but étant que l'on puisse recevoir une requête venant de n'importe où, par exemple une _appliance_ HTTP qui veut activer une sortie automate. Il doit être permettre de créer une resource (POST), de modifier une resource (PUT, PATCH), de lire une resource (GET) et de supprimer une resource (DELETE).

=== Node webSocket output
Pour cette fois, nous n'utilisons pas "parameterValueData" car les outputs ont un fonctionnement différent dans le programme qui rend cette implémentation plus complexe. De plus, il n'est pas nécessaire car, généralement, les outputs n'ont pas de paramètres. Il est donc possible de se passer de cette fonctionnalité en utilisant "selectedServiceData" et "selectedSubServiceData".

#figure(
  image("/resources/img/32_utilisationViewOutputWebSocket.png", width: 50%),
  caption: [
    exemple de paramétrage de "view Output Bool"
  ],
)
 #figure(
    align(left,
    ```rust
      "label": "view Output Bool",
        "outputHandle": [],
        "parameterNameData": null,
        "primaryType": "outputNode",
        "selectedFriendlyNameData": "default",
        "selectedServiceData": "test",
        "selectedSubServiceData": "led",
    ```
    ),
    caption: [view Output Bool, extrait de la structure JSON d'un exemple],
  
  )

=== WebSocket data
Le *WebSocket* est un protocole de communication bidirectionnelle qui permet d'envoyer et de recevoir des données. Il est utilisé pour la vue *WebSocket* décrite dans le @sec:websocketVUE.

Dans les figures @fig:vuePrincipeWebSocketInput-vs-vue et @fig:vuePrincipeWebSocketOutput-vs-vue, on peut voir le principe de transmission des données. Les schémas permettent de visualiser toutes les structures de données nécessaires, ainsi que les fonctions (en brun) et les variables (en violet) les plus utiles.

#figure(
  image("/resources/img/30_WebsocketInput_dataTransmission.png", width: 100%),
  caption: [
    principe de transmission des données via WebSocket pour les inputs
  ],
)
#label("fig:vuePrincipeWebSocketInput-vs-vue")
#figure(
  image("/resources/img/31_WebsocketOutput_dataTransmission.png", width: 100%),
  caption: [
    principe de transmission des données via WebSocket pour les outputs
  ],
)
#label("fig:vuePrincipeWebSocketOutput-vs-vue")
  == Conclusion

/*
== Section 1
  Liste :
  #[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    - Top-level
      - Nested
      - Items
    - Items
  ]
  Top-level
    - Nested
    - Items
*/
  
]

  