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


  == Interface webSocket
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
  == Intégration des blocs logiques complexes
  === MQTT
  === WebSocket
  === HTTP
  == Conclusion

]
