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

  

  === Vue User WebSocket <sec:websocketVUE>

  La vue user WebSocket permet de visualiser en temps réel l'état des _Outputs_ et de gérer les des _Inputs_ spécifiques à cette vue. Sur le schéma @fig:schemaPrincipe-vs-vue, elle est représentée par la "View Page".

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
    Exemple programme WebSocket
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
=== Vue mode debug <sec:modeDebugDesign>
Il est également intéressant d'avoir une vue "debug" qui nous permettrait de visualiser les états des I/O de chaque node et même de changer la couleur des edges pour les booléens (rouge = off et vert = on). Pour ce faire, nous utiliserons la technologie des WebSocket également utilisée pour la vue WebSocket (@sec:websocketVUE). L'idée étant de transmettre un graphique modifié à chaque cycle, la modification du graphique se fera surtout au niveau des _Edges_. Un exemple d'un petit programme qui montre comment sont transmises et affichées les données est détaillé en annexe (@sec:exempleUtililationDebugSimple-vs-vue).

En figure @fig:vuePrincipeModedebug-vs-vue se trouve un exemple démontrant entièrement le fonctionnement de ce mode debug. Les edges qui affichent "???" sont ceux qu'on n'a pas sélectionnés avec l'outil.
#figure(
  image("/resources/img/54_ExempleModeDebugCompletProgrammation.png", width: 100%),
  caption: [
    exemple vue Mode debug 
  ],
)
#label("fig:vuePrincipeModedebug-vs-vue")


L'outil en question est représenté par une loupe. Quand il est sélectionné, il ajoute ou enlève les edges de la liste des edges dont on veut afficher les valeurs dans la vue de debug. La boîte à outil est expliquée plue en détail dans la section @sec:toolsMenu.
#figure(
  image("/resources/img/54_ExempleModeDebugTool.png", width: 50%),
  caption: [
    Mode debug tool : display connection
  ],
)

#label("fig:ToolModedebug-vs-vue")

La *transmission des données* ainsi que les mécanismes principaux sont décrit dans le schéma se trouvant @sec:debugModeData.
#pagebreak()
== Gestion et stockage des données

La gestion et le stockage des données sont des aspects importants du système. Dans notre cas, les données sont stockées et transmises au format JSON.




=== Node View User WebSocket output
//TO DO : Mettre les autres
Pour cette fois, nous n’utilisons pas *parameterValueData* car les *outputs* ont un fonctionnement différent dans le programme, ce qui rend cette implémentation plus complexe. De plus, cela n’est pas nécessaire car, généralement, les *outputs* n’ont pas de paramètres. Il est donc possible de se passer de cette fonctionnalité en utilisant *selectedServiceData* et *selectedSubServiceData*.


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

=== View User WebSocket data
Le *WebSocket* est un protocole de communication bidirectionnelle qui permet d'envoyer et de recevoir des données en temps réel. Il est utilisé pour la vue *User* décrite dans le @sec:websocketVUE.

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

=== Mode debug data <sec:debugModeData>
Le schéma @fig:vuePrincipeModeDebugData-vs-vue permet de comprendre les mécanismes principaux du mode debug qui utilise également la technologie WebSocket.
#figure(
  image("/resources/img/55_ModeDebugPrincipe.png", width: 100%),
  caption: [
    Principe de transmission des données via WebSocket et server echo pour le mode debug
  ],
)
#label("fig:vuePrincipeModeDebugData-vs-vue")

=== Node : blocs complexes – transmission des _Settings_

Pour les blocs complexes, nous devons pouvoir transmettre autant de données que nécessaire. Pour cela, deux tableaux de chaînes (_String_) sont utilisés :

- Un tableau *parameterValueData* contenant les #underline("valeurs") des paramètres (définis dans le frontend).
- Un tableau *parameterNameData* contenant les #underline("noms") des paramètres (définis dans le backend).

Ces tableaux permettent la transmission des données entre le backend et le frontend pour les _nodes_ qui le nécessitent, comme les blocs de communication (MQTT, HTTP, MODBUS), ainsi que les blocs *string to bool* et *bool to string*.


 #iconbox(linecolor: hei-pink)[Un exemple de la structure JSON d’un bloc _MQTT_ est présenté en annexe, à la section @sec:mqttConfiguration.]

//TO DO : Mettre en annexe comment rajouter parameterValueData et parameterValueData
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

  