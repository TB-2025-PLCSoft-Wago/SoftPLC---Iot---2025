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
  L'intégration de #gls("WDA") est une partie importante du projet. Cette section décrit comment WDA a été utilisé pour communiquer avec l'automate et comment il a été intégré dans le programme.

  C'est pour utiliser WDA qu'il a été choisi de remplacer l'automate 751-9401 par un 751-9402. L'analyse de la raison de cette décision se trouve @sec:WDA_AnalysePR4. Vous y trouverez également une analyse de comment était récuperé les I/O dans le programme du TB 2024 qui utilisait une méthode plus rapide.

  Pour implémenter WDA, il a fallu modifier le programme backend (`softplc-main`) et plus particulièrement les fichiers `inputUpdate.go` et `outputUpdate.go`.
  
  Un des principaux problèmes de WDA est la vitesse, qui est très lente. Après une commande GET sur une I/O, cela peut prendre jusqu'à environ 500ms avant qu'on ait la réponse. Sachant qu'on a besoin de faire une requête GET pour chaque Input et Output, cela représente au total 22 requêtes. La première version du programme faisait une requête GET pour chaque Input et Output, ce qui prenait plusieurs secondes pour récupérer l'état de tous les Inputs et Outputs. Pour résoudre ce problème, il a été trouvé la solution d'utiliser une *Monitoring Lists* (@sec:monitoring-lists). Cela nous permet de récupérer l'état de des I/O en une seule requête GET. Cela permet de réduire le temps de récupération à environ 500ms. Cependant, la *Monitoring Lists* a un temps de vie limité, donc il faut la recréer périodiquement. Cela est fait toutes les 600 secondes dans le programme.

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
  image("/resources/img/42_Concact_Exemple_2.png", width: 50%),
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
  image("/resources/img/43_retainValue.png", width: 20%),
  caption: [
    bloc *Retain Value*
  ],
)
#label("fig:blocRetainValue-vs-vue")

=== Comparator EQ
Le bloc *EQ* permet de comparer si deux valeur sont égale et d'activer la sortie si c'est le cas.
#figure(
  image("/resources/img/56_ComparatorEQNode.png", width: 100%),
  caption: [
    bloc *x = y*
  ],
)
=== Comparator GT
Le bloc *GT* permet de comparer deux valeurs. Si la première valeur est plus grande ou égale, alors la sortie est activée.  
Le bloc compare les valeurs si elles peuvent être converties en _float_, sinon il compare leur taille (longueur).

La figure @fig:ComparatorGTNode-vs-vue montre la différence entre ces deux cas.

#figure(
  image("/resources/img/56_ComparatorGTNode.png", width: 100%),
  caption: [
    bloc *x > y*
  ],
)
#label("fig:ComparatorGTNode-vs-vue")
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



== Intégration des blocs logiques de communication

Cette section décrit comment les blocs logiques de communication ont été intégrés dans le programme. Ces blocs peuvent être utilisés pour créer des communications MQTT, HTTP et MODBUS.

Le principe des blocs logiques de communication a été défini dans le schéma @fig:communication-Bloc-principe. Cependant, il a été adapté aux besoins du projet. En effet, il est parfois nécessaire d’avoir plus d’entrées ou de sorties que ce qui a été défini dans le schéma. De plus, afin d’éviter que l’utilisation ne devienne trop complexe, il a parfois été décidé de séparer le client et le serveur, comme c’est le cas pour HTTP. Toutefois, le principe de base avec l’utilisation de blocs permettant la conversion de booléens en chaînes de caractères et inversement est resté le même. Il a même été élargi pour permettre l’utilisation de constantes, de résultats de concaténation, de variables, etc. En bref, tous les blocs dont la sortie est de type *value* sont compatibles.

La figure @fig:BlocMqttHttpClientServeur_fermer-vs-vue montre la *vue programmation* sans les _settings_.  
La figure @fig:BlocMqttHttpClientServeur_ouvert-vs-vue montre les _settings_ dans la *vue programmation*. Il est possible de les ouvrir en cliquant sur le bouton "_settings_" du bloc. Il est également possible d’ouvrir les _settings_ de plusieurs blocs en même temps, même s’ils sont du même type. Cela est pratique si l’on souhaite modifier plusieurs blocs logiques de communication simultanément, copier les _settings_ d’un bloc à l’autre, ou comparer les _settings_ de plusieurs blocs.

Dans la *vue programmation*, les blocs logiques de communication sont définis par le composant *CommunicationHandles* dans le fichier _CommunicationHandles.tsx_. C’est notamment là que l’on retrouve la logique liée à l’affichage de la vue "_settings_". Un élément React basé sur le composant *CommunicationHandles* est créé dans le fichier _LogicalNode.tsx_, où sont générés dynamiquement tous les _Logical Nodes_. On y calcule également le nombre d’éléments de chaque type à afficher (inputs/outputs).

// TO DO : mettre dans gestion data  
Le tableau *parameterNameData* contient les noms des _settings_ affichés en _placeholder_, et le tableau *parameterValueData* contient les valeurs des _settings_ des blocs logiques de communication.

Dans la partie *backend*, les blocs logiques de communication se trouvent à la fin du package _nodes_. Ils ont tous été conçus pour ne pas faire planter le reste du programme en cas de perte de connexion ou de problème de communication. Pour cela, ils utilisent la méthode *go func()* afin de lancer une *goroutine* qui s’occupe de la communication. Cela permet de ne pas bloquer le reste du programme en cas d’erreur de communication.
 
#figure(
  image("/resources/img/60_BlocMqttHttpClientServeur_fermer.png", width: 100%),
  caption: [
    blocs logiques complexes - Bloc MQTT, HTTP client et HTTP Serveur
  ],
)
#label("fig:BlocMqttHttpClientServeur_fermer-vs-vue")
#figure(
  image("/resources/img/60_BlocMqttHttpClientServeur_Ouvert.png", width: 100%),
  caption: [
    blocs logiques complexes - Bloc MQTT, HTTP client et HTTP Serveur - Settings
  ],
)
#label("fig:BlocMqttHttpClientServeur_ouvert-vs-vue")

#figure(
  image("/resources/img/60_BlocModbusAll_Fermer.png", width: 100%),
  caption: [
    blocs logiques complexes - Bloc MQTT, HTTP client et HTTP Serveur
  ],
)
#pagebreak()
  === MQTT
  Le bloc *MQTT* permet de communiquer avec un broker MQTT. Il est possible de publier des messages sur des topics ou de s'abonner à des topics pour recevoir des messages. 
  
  Le bloc *MQTT* a les *inputs* suivantes :
  - *xEnable* : l'entrée pour activer le bloc. Si cette entrée est à _false_, le bloc ne fera rien.
  - *topicToSend* : les topics sur lesquels publier (exemple : topic/test1 ,, topic/test2).
  - *msgToSend* : le message à publier sur les topics de *topicToSend*.
  - *topicToReceive* : les topics sur lesquels s'abonner (exemple : topic/test1 ,, topic/test2).

  Les *outputs* sont :
  - *xReceive* : impulsions lorsque le bloc a reçu un message sur un des topics auxquels on s'est abonné.
  - *msgLastReceived* : le dernier message reçu sur les topics auxquels on s'est abonné.
  
   #iconbox(linecolor: hei-pink)[Des *exemples* d'utilisations du bloc *MQTT* sont présentés en annexe au @sec:exempleUtililationMqttSimple-vs-vue. ]
  

  Lorsqu'un message est reçu, la fonction _messageHandler()_ est appelée. Cette fonction s'occupe de ne pas rater de messages, cependant c'est ensuite la fonction _ProcessLogic()_ qui s'occupe de traiter les messages reçus, gérer dynamiquement les _subscribes_ et _unsubscribes_ et écrire les *outputs*.
  
   L'ordre des topics données sur *topicToReceive* est le même que l'ordre des messages reçus sur *msgLastReceived*. Cela permet de traiter les messages dans le même ordre que les topics auxquels on s'est abonné.

   La fonction _makeConnectLostHandler(n \*MqttNode)_ permet de gérer la perte de connexion avec le broker MQTT. Elle s'assure de relancer la connexion et de réabonner aux topics si la connexion a un problème.
  
  === Node HTTP client
Le package Go @HttpPackageNetb a été utilisé.
Pour le Node HTTP client, il est possible de configurer les paramètres suivants :
- *url* : l'URL de la requête HTTP.
- *user* : l'utilisateur pour l'authentification HTTP.
- *password* : le mot de passe pour l'authentification HTTP.
- *Headers* : les en-têtes HTTP à envoyer avec la requête.
  - *key* x : le nom de l'en-tête x.
  - *value* x : la valeur de l'en-tête x.
#infobox()[les _Headers_ sont des paires clé-valeur, par exemple : `{"Content-Type": "application/json"}`. Il faut donc deux paramètres pour chaque Header. De plus, il est possible de mettre plusieurs Headers. ]
  
Le bloc peut prendre dynamiquement les _inputs_ suivantes :
- *xSend* : un booléen pour envoyer lorsque qu'il passe à _true_.
- *url path* : la suite du chemin de l'URL de la requête HTTP. Il est ajouté à la suite du paramètre _URL_ pour donner l'URL final.
- *Method* : la méthode HTTP à utiliser (GET, POST, PATCH, PUT, DELETE, HEAD, OPTIONS), par défaut GET.
- *Body* : le corps de la requête HTTP, qui peut être au format JSON ou autre.
Le bloc nous retournera les paramètres suivants :
- *xDone* : un booléen pour indiquer si la requête a été effectuée avec succès.
- *Response* : la réponse de la requête HTTP.
  
   #iconbox(linecolor: hei-pink)[Des *exemples* d'utilisations du bloc *HTTP client* sont présentés en annexe /* et TO DO : PLUS SIMPLE */ au @sec:httpClientExampleWDA qui montre comment l'utiliser avec WDA.]

/*
#figure(
  image("/resources/img/34_http_settingExemple.png", width: 90%),
  caption: [
    exemple de paramétrage de "HTTP"
  ],
)*/
#pagebreak()
=== Node HTTP serveur

Le package Go @HttpPackageNetb a été utilisé.  
L’exemple @soysouvanhClientsServeursHTTP permet d’en comprendre davantage sur la création d’un serveur HTTP en Go. Pour le déploiement d’un serveur HTTP sur Docker, la documentation @nicholsonCraignicholsonSimplehttp2023 a été trouvée.

Le Node HTTP serveur permet de créer un serveur HTTP qui écoute les requêtes entrantes. Le but est de pouvoir recevoir une requête venant de n’importe où, par exemple une *appliance* HTTP qui veut activer une sortie automate. Il doit être possible de créer une ressource (POST), de modifier une ressource (PATCH), de lire une ressource (GET) et de supprimer une ressource (DELETE). 

Cette ressource peut être créée avec une requête POST sur `http://192.168.39.56:8080/flatten`, par exemple.  
Cette ressource s'appelle *flatten* car elle permet d’aplatir les données pour les rendre plus digestes pour une utilisation dans le programme de l’automaticien. Il est possible de créer plusieurs ressources, mais elles doivent être uniques, c’est pourquoi on retourne un *id* après un POST, suivi de *result*: { ... } qui contient les paramètres. La figure @fig:RequetePostServerHttp-vs-vue montre un exemple de requête POST sur le serveur HTTP.

#figure(
  image("/resources/img/71_ServeurHTTP_ExemplePost.png", width: 100%),
  caption: [
    exemple de requête POST sur le serveur HTTP
  ],
)
#label("fig:RequetePostServerHttp-vs-vue") 

On doit également pouvoir recevoir des requêtes HTTP qui ne possèdent pas de _body_ ou de _headers_. Ce qui est le cas pour les requêtes envoyées par certaines *appliances*, comme les boutons _shelly_, ce qui est possible via l'output *Received URL path*.

Le node utilise *bus.go* pour ne pas perdre d’événements. Les messages stockés dans le bus sont ensuite traités tranquillement dans la fonction _ProcessLogic()_ du node.

Les ressources sont stockées dans une variable *storage*, commune à tous les nodes HTTP serveur. L’accès concurrent à *storage* est protégé par un verrou (`sync.Mutex`).

Le node supporte l’authentification HTTP Basic.

Pour le Node HTTP Server, il est possible de configurer les paramètres suivants :
- *url* : l'URL du serveur HTTP (par défaut : localhost:8080).
- *user* : l'utilisateur pour l'authentification HTTP des requêtes autorisées.
- *password* : le mot de passe pour l'authentification HTTP des requêtes autorisées.

Le bloc peut prendre dynamiquement les _inputs_ suivants :
- *Parameters to receive* : les paramètres à recevoir dans la requête HTTP. Ils sont séparés par des virgules (exemple : `param1 ,, param2 ,, param3`). Ils sont liés à la sortie _Values received_.

Les *outputs* sont :
- *xDone* : un booléen pour indiquer si la requête a été effectuée avec succès.
- *Values received* : les valeurs dans le _body_ reçues des paramètres donnés dans _Parameters to receive_. Elles sont dans le même ordre que les _Parameters to receive_. C’est-à-dire que si on a _param1 ,, param3 ,, data-attributes-value_ dans *Parameters to receive* et qu’on exécute la requête POST @fig:RequetePostServerHttp-vs-vue, alors *Values received* sera _value1 ,, value3 ,, true_. Elle fonctionne également pour les requêtes PATCH et PUT.
- *Resource ID* : l’identifiant de la ressource créée lors d’une requête POST, ou modifiée lors d’un PATCH, ou supprimée lors d’un DELETE.
- *Received URL path* : le chemin de l’URL de la requête HTTP reçue, sans le _host_ et le _port_. Par exemple, si on reçoit la requête GET : http://192.168.39.56:8080/short1, alors _Received URL path_ sera _short1_.

#iconbox(linecolor: hei-pink)[Des *exemples* d’utilisations du bloc *HTTP serveur* sont présentés en annexe au @sec:httpServerExample qui montre comment l’utiliser.]


Exemples de requêtes HTTP (ici _localhost:8080_ est équivalent à _192.168.39.56:8080_) :
- GET : http://192.168.39.56:8080/short1, alors _Received URL path_ sera égal à _short1_ (@fig:RequeteGetShortServerHttp-vs-vue).
- POST : http://192.168.39.56:8080/flatten, alors une ressource sera créée avec un _id_ (@fig:RequetePostServerHttp-vs-vue).
- PATCH : http://localhost:8080/parameters/flatten/7 (deux possibilités équivalentes : @fig:RequetePatchServerHttp-vs-vue et @fig:RequetePatchServerHttp_v2-vs-vue), alors la ressource avec l’_id_ 7 sera modifiée.
- GET : http://localhost:8080/parameters/flatten/7, alors la ressource avec l’_id_ 7 sera lue (@fig:RequeteGETServerHttpResource-vs-vue).
- DELETE : http://localhost:8080/parameters/flatten/7, alors la ressource avec l’_id_ 7 sera supprimée (@fig:RequeteDeleteServerHttp-vs-vue).
- PUT : http://localhost:8080/message, alors si les paramètres du body sont parmi les _Parameters to receive_, alors ils seront renvoyés dans _Values received_ (@fig:RequetePutServerHttp-vs-vue).


#figure(
  image("/resources/img/71_ServeurHTTP_ExemplePatch.png", width: 100%),
  caption: [
    exemple de requête PATCH sur le serveur HTTP
  ],
)
#label("fig:RequetePatchServerHttp-vs-vue")
#figure(
  image("/resources/img/71_ServeurHTTP_ExemplePatch_v2.png", width: 100%),
  caption: [
    exemple de requête PATCH sur le serveur HTTP version 2 
  ],
)
#label("fig:RequetePatchServerHttp_v2-vs-vue") 

#figure(
  image("/resources/img/71_ServeurHTTP_ExempleGet.png", width: 100%),
  caption: [
    exemple de requête GET sur le serveur HTTP pour lire une ressource
  ],
)
#label("fig:RequeteGETServerHttpResource-vs-vue")


#figure(
  image("/resources/img/71_ServeurHTTP_ExempleDelete.png", width: 100%),
  caption: [
    exemple de requête DELETE sur le serveur HTTP pour supprimer la ressource 7
  ],
)
#label("fig:RequeteDeleteServerHttp-vs-vue")

#figure(
  image("/resources/img/71_ServeurHTTP_ExemplePut.png", width: 100%),
  caption: [
    exemple de requête PUT sur le serveur HTTP
  ],
)
#label("fig:RequetePutServerHttp-vs-vue")

#figure(
  image("/resources/img/71_ServeurHTTP_ExempleGet_short.png", width: 100%),
  caption: [
    exemple de requête GET sur le serveur HTTP avec un path quelconque
  ],
)
#label("fig:RequeteGetShortServerHttp-vs-vue")
#pagebreak()

  === MODBUS
Les documentations utilisées pour la création des blocs MODBUS sont :
- @CodesFonctionsModbus : pour la compréhension du protocole MODBUS.
- @GoburrowModbus2025 : pour comprendre comment l'implementer en Go.
#figure(
  image("/resources/img/63_accordionModbus.png", width: 30%),
  caption: [
    Accordion – Modbus
  ],
)
#label("fig:accordionModbus-vs-vue")

Les blocs Modbus ont été réalisés dans le but de permettre la lecture et l’écriture de toutes les valeurs de _Home-IO_, ce qui peut expliquer certains choix de fonctions utilisés. Les blocs logiques de communication MODBUS développés sont les suivants :

- *Modbus Read Bool* : permet de lire des valeurs booléennes dans un dispositif esclave MODBUS, correspondant au code fonction 0x02. Dans le programme Go, la fonction _ReadDiscreteInputs_ est utilisée pour la lecture. La fonction _ReadCoils_ n’a pas été retenue car elle ne convient pas à _Home-IO_. Des exemples d’utilisation sont présentés en annexe à la section @sec:exempleUtililationModbusReadBool-vs-vue.

- *Modbus Read Value* : permet de lire des valeurs entières dans un dispositif esclave MODBUS, correspondant au code fonction 0x04. Dans le programme Go, la fonction _ReadInputRegisters_ est utilisée. Voir la section @sec:exempleUtililationModbusReadValue-vs-vue pour des exemples.

- *Modbus Write Bool* : permet d’écrire des valeurs booléennes dans un dispositif esclave MODBUS, correspondant au code fonction 0x15. Dans le programme Go, la fonction _WriteMultipleCoils_ est utilisée. Voir la section @sec:exempleUtililationModbusWriteBool-vs-vue.

- *Modbus Write Value* : permet d’écrire des valeurs entières dans un dispositif esclave MODBUS, correspondant au code fonction 0x06. Dans le programme Go, la fonction _WriteSingleRegister_ est utilisée. Voir la section @sec:exempleUtililationModbusWriteValue-vs-vue.

Ces blocs se configurent à l’aide du *host* et du *port* du serveur MODBUS.

Il existe deux types de blocs Modbus : les blocs de lecture (_Read_) et les blocs d’écriture (_Write_). Les deux types possèdent les *inputs* suivantes :

- *xEnable* : permet d’activer le bloc.
- *UnitID* : identifiant de l’esclave MODBUS auquel accéder.
- *Addresses* : adresses des registres à lire ou écrire. On peut spécifier plusieurs adresses séparées par des virgules (ex. : 0 ,, 2 ,, 4). Il faudra définir _Quantity_ pour la lecture ou fournir _NewValues_ pour l’écriture. Le comportement de ces blocs dépend de la relation entre *Addresses* et *Quantity* ou *NewValues* (@fig:ModbusGestionQuantity-vs-vue et @fig:ModbusGestionNewValues-vs-vue).

Les blocs de lecture ont également l’*input* :

- *Quantity* (@fig:ModbusGestionQuantity-vs-vue) : nombre de registres à lire. Par défaut, cette valeur est fixée à 1. Plusieurs valeurs peuvent être fournies, séparées par des virgules (ex. : 0 ,, 2 ,, 4).

Les blocs d’écriture ont également l’*input* :

- *NewValues* (@fig:ModbusGestionNewValues-vs-vue) : valeurs à écrire dans les registres. Plusieurs valeurs peuvent être fournies, séparées par des virgules (ex. : 0 ,, 2 ,, 4).

Et les *outputs* suivantes :

- *xDone* : activée si la communication avec l’esclave est établie et qu’aucune erreur ne s’est produite.
- *ValuesReceived* : valeur(s) reçue(s) pour les blocs de lecture. Pour les blocs d’écriture, cette sortie contient la réponse du serveur, et peut aussi signaler les erreurs de communication.



#figure(
  image("/resources/img/64_ModbusRead_gestionQuantity.png", width: 80%),
  caption: [
    Modbus Read – Comportement en fonction de la relation entre la quantité (*Quantity*) et le nombre d’adresses (*Addresses*)
  ],
)
#label("fig:ModbusGestionQuantity-vs-vue")

#figure(
  image("/resources/img/64_ModbusRead_gestionNewValues.png", width: 80%),
  caption: [
    Modbus Write – Comportement en fonction de la relation entre *NewValues* et le nombre d’adresses (*Addresses*)
  ],
)
#label("fig:ModbusGestionNewValues-vs-vue")

  == Vue programmation
  La vue programmation permet de créer des programmes PLC en utilisant une interface graphique. Elle est représentée par la "Programming Page" sur le schéma @fig:schemaPrincipe-vs-vue. Cette vue permet de créer par *drag and drop* des blocs logiques, des *Inputs* et des *Outputs*, et de les connecter entre eux pour créer un programme PLC.


 Du côté *frontend*, la majorité de la logique est centralisée dans le fichier *_App.tsx_*. Les fichiers _InputNode.tsx_, _LogicalNode.tsx_ et _OutputNode.tsx_ jouent également un rôle important, car ils gèrent l'affichage des _nodes_ selon leur _primaryType_. 

On retrouve ensuite différents fichiers dans le dossier "handles", dédiés à des types particuliers, comme par exemple les _nodes_ de communication.

La section @sec:ameliorationInterface-vs-vue décrit plus en détail les aspects liés à l'interface visuelle.

#iconbox(linecolor: hei-pink)[Tous les nodes réalisés sont présentés en annexe au @sec:EnsembleBlocs.]

 == Vue User

#figure(
  image("/resources/img/74_exempleBidonUserView.png", width: 80%),
  caption: [
    vue User : bref aperçu
  ],
)

Le rôle de cette vue a déjà été expliqué @sec:websocketVUE.  
Elle est implémentée côté *backend* dans *serverWebSocket.go*, suivant les schémas du chapitre précédent en @fig:vuePrincipeWebSocketInput-vs-vue et @fig:vuePrincipeWebSocketOutput-vs-vue.

Du côté *frontend*, la vue est créée dans le fichier *user.tsx* du dossier *webSocketInterface*. C’est là qu’on gère l’affichage des éléments triés dans les appliances.

Pour passer d’une vue à l’autre, on utilise react-router-dom @ReactRouterUseNavigate : pour naviguer entre les routes (avec `useNavigate`).


#figure(
    align(left,
    ```tsx
      const navigate = useNavigate();
      const openView = () => { navigate('/websocket'); };
      ...
      const navigate = useNavigate();
      const goBackView = () => { navigate(-1); };
    ```
    ),
    caption: [*Vue User* : navigation entre les pages],
  )
  
  == Vue mode debug

Le rôle de cette vue a déjà été expliqué @sec:modeDebugDesign.  
Elle est implémentée côté *backend* dans *serverWebSocket.go*, suivant les schémas du chapitre précédent @sec:debugModeData.

Après avoir dû déboguer avec cet outil, il a finalement été choisi d’afficher la valeur des connexions par défaut, et que l’outil *display connection* permette de cacher celles qui prennent trop de place, par exemple. Cependant, côté *backend*, dans la fonction _DebugMode_, il y a la première version en commentaire.

Les outils communiquent avec le *backend* par la fonction _handleIncomingMessage_, responsable de recevoir les messages _webSocket_. C’est à ce moment-là qu’on ajoute ou enlève les *edges* de la variable _toDebugList_.

Les connexions reçoivent l’animation "dash 1s linear infinite" pour donner l’impression d’un flux de données.

Du côté *frontend*, la vue est créée grâce au fichier *debug.tsx* du dossier *webSocketInterface*.

La vue debug est également réduite comparée à la vue programmatation, en utilisant une sous-classe _css_ "hide-when-debug".

#figure(
  image("/resources/img/76_debugSansDebug.png", width: 100%),
  caption: [
    vue programmatation : plus d'éléments visible que vue debug
  ],
)
#figure(
  image("/resources/img/76_debugAvecDebug.png", width: 100%),
  caption: [
    vue debug : moins d'éléments visible que vue debug
  ],
)


  == Améliorations interface <sec:ameliorationInterface-vs-vue>
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
    image("/resources/img/52_boutonsImageListe.png", width: 80%),
    caption: [
      exemple - bouttons - visualisation
    ],
  )
  #label("fig:bouttonsVisu-vs-vue")
=== Tools <sec:toolsMenu>
Pour permettre de modifier la manière d’interagir avec le graphique, un menu déroulant appelé *Tool* a été ajouté. Il permet de choisir entre différents outils.

Plusieurs outils ont été définis :
- *DisplayConnectionDebug* : permet de sélectionner les connexions à afficher en mode debug.
- *comment* : permet d’ajouter un commentaire à l’endroit où l’on clique.

Les outils peuvent fonctionner de deux manières :
- En récupérant les informations depuis le backend, comme illustré en @fig:MessageWebSocketClicEdge-vs-vue.
- En les utilisant directement depuis le frontend, comme illustré en @fig:MessageWebSocketClicComent-vs-vue.

#figure(
  image("/resources/img/57_toolsAll.png", width: 40%),
  caption: [
    menu déroulant *Tool*
  ],
  )
#figure(
    align(left,
    ```go
    map[source:40 sourceHandle:Output tool:DisplayConnectionDebug type:edge_clicked] 
    ```
    ),
    caption: [*Message WebSocket au backend* : clic effectué sur un edge avec l'outil "DisplayConnectionDebug"],
  
  )
  #label("fig:MessageWebSocketClicEdge-vs-vue")

  #figure(
    align(left,
    ```tsx
    /* add comment */
    const onPaneClick = useCallback(
        (event: React.MouseEvent) => {
            if (tool === 'comment') {
              ...
    ```
    ),
    caption: [*Extrait code App.tsx* : Exemple utilisation de l'outil "comment"],
  
  )
  #label("fig:MessageWebSocketClicComent-vs-vue")
 #figure(
  image("/resources/img/75_exempleCommentaire.png", width: 100%),
  caption: [
   *Tool :* exemple commentaires
  ],
  )
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
