#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= #i18n("implementation-title", lang:option.lang) <sec:impl>

#option-style(type:option.type)[
  In the implementation phase of your bachelor thesis, you translate the design specifications into tangible, functional artifacts. This section offers insights into the practical execution of your research, detailing the steps taken to realize the proposed system. Here are some ways to enhance and elaborate on this section:

  - *Development Methodology*: Describe the methodology or approach employed in the development process.
  - *Prototyping and Iterative Development*: If applicable, discuss any prototyping or iterative development techniques utilized during the implementation phase.
  - *Coding Practices and Standards*: Provide insights into the coding practices, standards, and conventions adhered to during development.
  - *Testing and Quality Assurance*: Detail the testing strategies and quality assurance measures employed to validate the correctness and robustness of the implemented system.
  - *Performance Optimization*: Address any performance considerations or optimizations made during the implementation phase.
  - *Deployment and Configuration*: Describe the deployment process and configuration management practices involved in deploying the system to production or testing environments.
  - *Documentation and Knowledge Transfer*: Highlight the importance of documentation in facilitating knowledge transfer and ensuring the sustainability of the implemented system.

  Dans la phase de mise en œuvre de votre mémoire de licence, vous traduisez les spécifications de conception en artefacts tangibles et fonctionnels. Cette section offre un aperçu de l’exécution pratique de votre recherche, en détaillant les étapes prises pour réaliser le système proposé. Voici quelques façons d’améliorer et de développer cette section :

  - *Méthodologie de développement* : Décrivez la méthodologie ou l’approche employée dans le processus de développement.
  - *Prototypage et développement itératif* : Le cas échéant, discutez des techniques de prototypage ou de développement itératif utilisées pendant la phase de mise en œuvre.
  - *Pratiques et normes de codage* : Fournir des informations sur les pratiques, normes et conventions de codage respectées lors du développement.
  - *Tests et assurance qualité* : Détailler les stratégies de test et les mesures d’assurance qualité utilisées pour valider l’exactitude et la robustesse du système mis en œuvre.
  - *Optimisation des performances* : Traiter toutes les considérations ou optimisations de performance effectuées pendant la phase de mise en œuvre.
  - *Déploiement et configuration* : Décrivez le processus de déploiement et les pratiques de gestion de la configuration impliquées dans le déploiement du système vers des environnements de production ou de test.
  - *Documentation et transfert de connaissances* : Souligner l’importance de la documentation pour faciliter le transfert de connaissances et assurer la durabilité du système mis en œuvre.
]

Cette section décrit les différentes étapes qui ont été implémentées durant le projet. Elle est divisée en plusieurs sous-sections, chacune abordant un aspect spécifique. Il y est notamment décrit comment les blocs logiques complexes ont été implémentés, comment WDA a été utilisé, l'ajout d'une vue webSocket ainsi que tout ce qui a dû être mis en place pour que tout cela soit réalisable.

#add-chapter(
  after: <sec:impl>,
  before: <sec:validation>,
  minitoc-title: i18n("toc-title", lang: option.lang)
)[
  #pagebreak()
  == WDA
  L'intégration de WDA est une partie importante du projet. Cette section décrit comment WDA a été utilisé pour communiquer avec l'automate et comment il a été intégré dans le programme.

  Pour implémenter WDA, il a fallu modifier le programme backend (`softplc-main`) et plus particulièrement les fichiers `inputUpdate.go` et `outputUpdate.go`.
  
  Un des principaux problèmes de WDA est la vitesse, qui est très lente. Après une commande GET sur une I/O, cela peut prendre jusqu'à environ 500ms avant qu'on ait la réponse. Sachant qu'on a besoin de faire une requête GET pour chaque Input et Output, cela représente au total 22 requêtes. La première version du programme faisait une requête GET pour chaque Input et Output, ce qui prenait plusieurs secondes pour récupérer l'état de tous les Inputs et Outputs. Pour résoudre ce problème, il a été trouvé la solution d'utiliser une *Monitoring Lists* (@sec:monitoring-lists). Cela nous permet de récupérer les I/O en une seule requête GET pour récupérer l'état de tous les Inputs et Outputs en une seule fois. Cela permet de réduire le temps de récupération à environ 500ms.

  Toutes les requêtes via WDA doivent avoir le header @fig:header-vs-vue et l'authentification @fig:Authentification-vs-vue.

#figure(
  image("/resources/img/18_HeaderWDA.png", width: 90%),
  caption: [
    Header WDA 
  ],
)
#label("fig:header-vs-vue")
#figure(
  image("/resources/img/19_autentificationWDA.png", width: 90%),
  caption: [
    Authentification WDA 
  ],
)
#label("fig:Authentification-vs-vue")
=== inputUpdate.go
  Dans `inputUpdate.go`, l'idée étant de créer et mettre à jour la variable tableau *InputsOutputsState* de structure _InputsOutputs_ dont la structure d'un élément est montré @fig:exempleElementInputOutputState-vs-vue, afin de garder le fonctionnement du programme existant, mais cette fois-ci en utilisant WDA. 

#figure(
  image("/resources/img/15_exempleElementInputOutputState.png", width: 50%),
  caption: [
    Exemple d'élément Input/Output State dans le programme backend 
  ],
)
#label("fig:exempleElementInputOutputState-vs-vue")

Au démarrage, `softPLC.go` appelle la fonction *InitInputs* pour créer le client HTTP, appeler la fonction *initClient* et initialiser la variable *InputsOutputsState*, qui est créée par la fonction *mapResultsToInputsOutputs* prenant en paramètre le résultat de la fonction *fetchValues* (@fig:returnFetchVal-vs-vue), responsable de faire la requête GET afin de récupérer les valeurs des I/O.

La fonction *initClient* a pour rôle de *configurer les DIO* (@sec:dio-activation). Par défaut, ils sont configurés en *input* (value = 1 à 8). Il faut donc les configurer en *output* (value = 9 à 16) si on veut avoir que des outputs. De plus, avant de pouvoir lire et écrire des I/O, il faut modifier le *access mode* (@sec:access-mode) en activant le control mode (value = 2). Par défaut, celui-ci est désactivé (value = 0).

#figure(
  image("/resources/img/25_ResultFetchVal.png", width: 50%),
  caption: [
    structure valeur retournée par la fonction fetchValues
  ],
)
#label("fig:returnFetchVal-vs-vue")

Ensuite, la fonction *UpdateInputs* est appelée à chaque cycle par `softPLC.go`. La fonction *UpdateInputs* utilise la fonction *updateInputsOutputsState* pour mettre à jour la variable *InputsOutputsState*. La fonction *updateInputsOutputsState* met à jour la variable *InputsOutputsState* avec les nouvelles valeurs récupérées par la fonction *fetchValues* pour récupérer les valeurs des I/O.

En parallèle, la fonction *CreateMonitoringLists* est appelée périodiquement toutes les 600 secondes par `softPLC.go` pour recréer une monitoring lists (@sec:monitoring-lists).

  === outputUpdate.go
  Dans `outputUpdate.go`, l'idée est de mettre à jour les *outputs* après que le programme ait été exécuté. La logique pour déclencher l'envoi des nouvelles valeurs des *outputs* est restée la même que pour l'ancien programme. La différence est que l'on utilise *WDA* pour envoyer les nouvelles valeurs des *outputs*. C’est-à-dire que l'on crée des requêtes *PATCH* pour les *AO* et *DO*, exemple de requête *PATCH* (@sec:exemplePatchOutput-vs-vue).

  #pagebreak()
  == Interface webSocket
  react-router-dom : pour naviguer entre les routes (avec useNavigate).

  const openView = () => {
        navigate('/websocket');
    };
   #pagebreak()
  == Intégration des blocs logiques simples
  L'intégration des blocs logiques simples est une étape importante du projet. Cette section décrit comment ces blocs logiques simples ont été intégrés dans le programme. 

  === Bloc Bool to String
  Il correspond au bloc *Pulse to frame Sender* du schéma @fig:communication-Bloc-principe. Mais il est utilisable pour d'autres fonctionnalités que la communication.  
Le bloc *Bool to String* permet de transmettre en sortie une chaîne de caractères selon si un booléen est à _true_ en face.  
C'est un bloc de type #gls("strechable"), c'est-à-dire que le nombre d'entrées est dynamique.  
La figure @fig:blocBoolToString-vs-vue montre un exemple de ce à quoi pourrait ressembler le bloc *Bool to String*. Il est possible d'écrire plusieurs lignes de texte, chaque ligne de texte se trouve en face d'une entrée booléenne. Si l'entrée booléenne est à _true_, la ligne de texte correspondante sera transmise en sortie.
 
#figure(
  image("/resources/img/35_exempleDesign_BoolToString.png", width: 50%),
  caption: [
    bloc Bool to String
  ],
)
#label("fig:blocBoolToString-vs-vue")
#infobox()[
Il est possible d'*activer plusieurs lignes de texte* en même temps. Par exemple, si les entrées booléennes *x0*, *x1* et *x2* sont à _true_, les lignes de texte correspondantes seront transmises en sortie sous la forme d'un tableau. Cependant, la sortie est de type *String*, donc les lignes de texte seront séparées par " ,, " (deux virgules entourées par des espaces).

]

=== Bloc String to Bool
Le bloc *String to Bool* permet de recevoir une chaîne de caractères et de la convertir en booléens. Il est représenté par le bloc *Frame to Pulse Receiver* sur le schéma @fig:communication-Bloc-principe. Il est également un bloc de type #gls("strechable") mais au niveau des sorties. La figure @fig:blocStringToBool-vs-vue montre un exemple de ce à quoi pourrait ressembler le bloc *String to Bool*. Il est possible d'écrire plusieurs lignes de texte, chaque ligne de texte se trouve en face d'une sortie booléenne. Si la chaîne de caractères reçue contient la ligne de texte correspondante, la sortie booléenne sera à _true_.
#figure(
  image("/resources/img/36_exempleDesign_StringToBool.png", width: 50%),
  caption: [
    bloc String to Bool
  ],
)
#label("fig:blocStringToBool-vs-vue")

Le fonctionnement du bloc *String to Bool* a été réfléchi de manière à ce qu'il soit le plus simple possible, mais pouvant gérer le plus de situations possible. Le schéma bloc simulant toutes les situations se trouve en @fig:blocFonctionmentStringToBool-vs-vue. Des blocs se trouvent en amont, mais ce qui nous intéresse est ce qu'on retrouve en _input_. Les trois situations de fonctionnement principales sont :
1. Activer un booléen si la chaîne de caractères reçue correspond à une des lignes de texte, (@fig:blocFonctionmentStringToBoolSituation1-vs-vue).
2. Traiter un tableau reçu dans son ordre original. Cela est important pour les blocs de communication, car on peut donner le résultat de leurs messages dans le même ordre que l'ordre de demande de requête, ce qui permet de traiter les messages sans se soucier d'avoir plusieurs messages avec les mêmes valeurs, (@fig:blocFonctionmentStringToBoolSituation2-vs-vue).
3. Traiter des tableaux comme des caractères simples, (@fig:blocFonctionmentStringToBoolSituation3-vs-vue).

Ainsi, il doit être capable de gérer des caractères simples et des pseudos tableaux (lignes de texte séparées par " ,, "). Si la chaîne de caractères reçue contient un caractère simple, la sortie booléenne correspondante sera à _true_, pour autant qu'elle existe.

 
#figure(
  image("/resources/img/37_exempleImplementationFonctionnement_StringToBool.png", width: 100%),
  caption: [
    exemple - fonctionnement String to Bool
  ],
)
#label("fig:blocFonctionmentStringToBool-vs-vue")

#figure(
  image("/resources/img/38_strToBool_situation_1.png", width: 100%),
  caption: [
    exemple - fonctionnement *String to Bool* - situation 1
  ],
)
#label("fig:blocFonctionmentStringToBoolSituation1-vs-vue")
#figure(
  image("/resources/img/38_strToBool_situation_2.png", width: 100%),
  caption: [
    exemple - fonctionnement *String to Bool* - situation 2
  ],
)
#label("fig:blocFonctionmentStringToBoolSituation2-vs-vue")
#figure(
  image("/resources/img/38_strToBool_situation_3.png", width: 80%),
  caption: [
    exemple - fonctionnement *String to Bool* - situation 3
  ],
)
#label("fig:blocFonctionmentStringToBoolSituation3-vs-vue")
=== SR
  Le bloc *SR* est qui permet de maintenir un état booléen. Il suffit d'une impultion sur l'entrée *S* (set) pour mettre la sortie à _true_ et d'une impultion sur l'entrée *R1* (reset) pour mettre la sortie à _false_.
  Le reset à la priorité sur le set, c'est-à-dire que si on a une *S* à _true_ et une *R1* à _true_, la sortie sera à _false_.
#figure(
  image("/resources/img/39_exemple_utilisation_SR.png", width: 100%),
  caption: [
    exemple - utilisation *SR*
  ],
)
#label("fig:blocFonctionnementSR-vs-vue")
#figure(
  image("/resources/img/39_exemple_utilisation_SR_View.png", width: 100%),
  caption: [
    exemple - utilisation *SR* - visualisation des états
  ],
)
#label("fig:blocFonctionnementSR_result-vs-vue")

=== Counter
  Le bloc *Counter* permet de compter et décompter des impulsions. 
  Les entrées sont les suivantes :
  - *Step* : l'entrée pour choisir le pas d'incrémentation ou de décrémentation du compteur, (par défaut, le pas est de 1).
  - *+* : l'entrée pour incrémenter le compteur.
  - *-* : l'entrée pour décrémenter le compteur.
  - *R* : l'entrée pour remettre le compteur à zéro.
  
La sortie est *result* qui peut être de format de type *int* ou *float*, c'est le résultat des opérations.
#label("fig:blocFonctionnementSR-vs-vue")
#figure(
  image("/resources/img/40_exemple_utilisation_Counter_View.png", width: 100%),
  caption: [
    exemple - utilisation *Counter*
  ],
)
#infobox[Noter que l'on utilise des blocs *RF_trig*. Cela permet de générer des impulsions à chaque clic sur les boutons.]

=== Concat
Le bloc *Concat* est utilisable pour assembler des chaînes de caractère de manière dynamique.

#figure(
  image("/resources/img/42_Concact_Exemple_1.png", width: 100%),
  caption: [
    exemple - utilisation *Concat*
  ],
)
#label("fig:blocFonctionnementConcat-vs-vue")
#figure(
  image("/resources/img/42_Concact_Exemple_2.png", width: 80%),
  caption: [
    exemple - utilisation *Concat* - visualisation des états
  ],
)
#label("fig:blocFonctionnementConcat_result-vs-vue")
=== Retain value
Le bloc *Retain Value* sert à bloquer une valeur _strIn_ et à ne la transmettre sur _strOut_ uniquement si _pass_ est à _true_.

#infobox()[ Si _pass_ = true alors  _strOut_ = _strIn_ sinon _strOut_ = \"\"
]
#figure(
  image("/resources/img/43_retainValue.png", width: 25%),
  caption: [
    bloc *Retain Value*
  ],
)
#label("fig:blocRetainValue-vs-vue")
#pagebreak()
=== Variables
Afin de permettre des fonctions de type boucle de contre-réaction, il a été choisi de rajouter un système de variable. Les blocs @fig:blocsVariables-vs-vue s'utilisent de la manière suivante : il faut faire correspondre les noms pour que les valeurs correspondent. Pour un même nom, une seule *output* doit être définie, mais plusieurs *Input* peuvent porter ce nom.

L'exemple @fig:blocsVariablesContreReactionBool-vs-vue montre le fonctionnement d'un programme qui, à chaque cycle, change l'état d'un booléen.

L'exemple @fig:blocsVariablesContreReactionValue-vs-vue montre le fonctionnement d'un programme qui, à chaque cycle, change l'état d'une valeur (*BoolToString_Output*) lorsqu'on a "start". La séquence est la suivante : "empty" → "Hi" → "Bye" → etc.

#figure(
  image("/resources/img/44_blocs_variables_possibles.png", width: 100%),
  caption: [
    blocs variable
  ],
)
#label("fig:blocsVariables-vs-vue")
#figure(
  image("/resources/img/44_blocs_variables_ContreReaction_Bool.png", width: 100%),
  caption: [
    exemple - variable bool - contre-réaction
  ],
)
#label("fig:blocsVariablesContreReactionBool-vs-vue")

#figure(
  image("/resources/img/44_blocs_variables_ContreReaction_Value.png", width: 120%),
  caption: [
    exemple - variable value - contre-réaction
  ],
)
#label("fig:blocsVariablesContreReactionValue-vs-vue")



  == Intégration des blocs logiques complexes
  === MQTT
  === WebSocket
  === HTTP

  == Améliorations interface
  === Rajouter des raccourcies
  Le fichier _useKeyboardShortcuts.tsx_ a été créé pour l'occasion. Il est appelé dans _App.tsx_.  
  Les raccourcis qui ont été rajoutés sont :
  - ctrl + c : copie les nodes/edges sélectionnés  
  - ctrl + v : coller  
  - ctrl + x : couper  
  - ctrl + z : annule la dernière modification (undo)  
  - ctrl + y : rétablit la modification annulée (redo)  

  *undo / redo* : Le principe est d'avoir deux piles _redoStack_ et _undoStack_. On utilise _pushToUndoStack()_ pour créer une pile, et _useDebouncedUndo.tsx_ qui vérifie lorsqu'il y a des modifications et utilise un petit délai pour éviter de pousser plusieurs fois à cause d'une modification mineure survenant au même moment.

  #pagebreak()
    === Nodes 
  Un node standard est constitué de trois éléments principaux, comme le montre @fig:blocCSS_Node_Basique-vs-vue. Pour chacun de ces éléments, une classe CSS a été créée. Il est également possible de spécifier plus particulièrement pour des blocs un peu plus complexes, comme le montre @fig:blocCSS_Communication_Modife-vs-vue, par exemple pour agrandir légèrement pour les blocs de communication. Cette structure permet de changer les couleurs et tailles pour chaque type de node.

  #figure(
  image("/resources/img/45_css_base.png", width: 50%),
  caption: [
    exemple - node standard
  ],
)
#label("fig:blocCSS_Node_Basique-vs-vue")
#figure(
    align(left,
    ```css
      /*Communication*/
        .ntb-Communication{
            height: 60px;
        }
        .react-flow__node .dl-Communication{
            top: 15px;
        }


        .ns-Communication{
            top: 60px;
        }
    ```
    ),
    caption: [*css*, exemple pour agrandir légèrement pour les blocs de communication],
  
  )
#label("fig:blocCSS_Communication_Modife-vs-vue")

Un autre type de node @fig:nodeSelect-vs-vue est celui avec un "menu déroulant". Ils sont conçus pour changer le _Node_ de manière rapide sans refaire les connexions. Cela est possible uniquement si les nodes ont les mêmes entrées et sorties, et qu'elles sont du même type. Cela est donc possible pour les _timer_ et les _trigger_.

  #figure(
    image("/resources/img/50_nodeSelect.png", width: 100%),
    caption: [
      Node avec menu déroulant
    ],
  )
  #label("fig:nodeSelect-vs-vue")
  === Couleurs de connection dynamique
  Pour que l'utilisateur puisse améliorer la lisibilité, il a été rajouté un outil permettant de choisir la couleur des connexions sélectionnées. Cet outil a été inspiré de "Custom Nodes"
 @CustomNodes2025, qui montre un exemple d'utilisation de l'_input_ de type _color_ @fig:colorConnections-vs-vue.
  #figure(
    image("/resources/img/46_colorConnections.png", width: 40%),
    caption: [
      outil modification couleur (_input_ de type _color_)
    ],
  )
  #label("fig:colorConnections-vs-vue")
  === Elément sélectionnées
  Quand une connection est sélectionnée, elle devient plus large @fig:selectConnections-vs-vue. C'est le meilleur critère d'apparence qu'on peu modifier pour garder la même couleur.
  #figure(
    image("/resources/img/47_selectConnection.png", width: 50%),
    caption: [
      connection : select VS not select 
    ],
  )
  #label("fig:selectConnections-vs-vue")

  Quand un Node est sélectionné, il devient coloré sur le tour avec une légère ombre @fig:selectNode-vs-vue.
  #figure(
    image("/resources/img/48_selectNode.png", width: 50%),
    caption: [
      Node : select VS not select 
    ],
  )
  #label("fig:selectNode-vs-vue")
  === Accordion - css
  #figure(
    image("/resources/img/49_accordion.png", width: 100%),
    caption: [
      Accordion - css sélecteurs
    ],
  )
  #label("fig:accordionCss-vs-vue")

  === Boutons
  L'exemple de @CSSButtons a été utilisé.
  Plusieurs classes de boutons ont été définies dans le _css_ (@fig:blocCSS_Bouttons-vs-vue). Le sélecteur de classe de base est ".button". Un exemple d'utilisation serait : 
  #figure(
    align(left,
    ```tsx
    <button className={"button button1"} onClick={openView}>Open view</button>
    ```
    ),
  )
  Tous les styles définis dans les deux sélecteurs de classe CSS (dans l'exemple _.button_ et _.button1_) correspondants seront cumulés, et les styles de _.button1_ peuvent écraser ceux de _.button_ s'ils affectent les mêmes propriétés.

  On peut utiliser les sélecteurs définis dans @fig:blocCSS_Bouttons-vs-vue, de la même manière que _.button1_. Ils sont représentés en @fig:bouttonsVisu-vs-vue.

#figure(
    align(left,
    ```css
      /*Buttons*/
      .button {
          background-color: #04AA6D; /* Green */
          border: none;
          color: white;
          padding: 8px 16px;
          text-align: center;
          text-decoration: none;
          display: inline-block;
          font-size: 16px;
          margin: 4px 4px;
          transition-duration: 0.4s;
          cursor: pointer;
      }
      .button1 {
          background-color: white;
          color: black;
          border: 2px solid #04AA6D;
      }

      .button1:hover {
          background-color: #04AA6D;
          color: white;
      }

      .buttonNode {
          background-color: white;
          color: black;
          border: 1px solid #1a192b;
          font-style: italic;
      }

      .buttonNode:hover {
          background-color: #1a192b;
          color: white;
      }
    ```
    ),
    caption: [*css*, Bouttons],
  
  )
#label("fig:blocCSS_Bouttons-vs-vue")

#figure(
    image("/resources/img/52_boutonsImageListe.png", width: 100%),
    caption: [
      exemple - bouttons - visualisation
    ],
  )
  #label("fig:bouttonsVisu-vs-vue")
#pagebreak()
  == Gestion des erreurs
  Pour envoyer un message d'erreur sur la page de programmation, il faut utiliser dans le programme : *serverResponse.ResponseProcessGraph* = "message à envoyer". Cela doit se faire avant la fin de la vérification, donc pour mettre des messages pour un *Node* précis, il faut utiliser l'appel dans la méthode *GetOutput* de celui-ci. C'est ce qui a été utilisé pour @sec:timer.

  === Outputs
  Il est interdit d'avoir deux fois la même _Output_, même si le type est le même. 
  #figure(
  image("/resources/img/51_erreur_Output.png", width: 100%),
  caption: [
    exemple - erreur *Output*
  ],
  )
  #figure(
  image("/resources/img/51_erreur_viewWebOutput.png", width: 100%),
  caption: [
    exemple - erreur *view Output*
  ],
  )
  #figure(
  image("/resources/img/51_erreur_variableOutput.png", width: 100%),
  caption: [
    exemple - erreur *variable Output*
  ],
  )
  === Timer <sec:timer>
  #figure(
  image("/resources/img/41_erreur_Tof.png", width: 100%),
  caption: [
    exemple - erreur *TOF*
  ],
  )

  #figure(
  image("/resources/img/41_erreur_Ton.png", width: 100%),
  caption: [
    exemple - erreur *TON*
  ],
  )

  

  == Conclusion

]
