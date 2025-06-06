#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= #i18n("introduction-title", lang:option.lang) <sec:intro>
Ce document présente les travaux réalisés ainsi que les perspectives à venir pour le projet PLCSoft destiné à WAGO. Ce projet consiste en l’ajout de nouvelles fonctionnalités à un HAL (Hardware Abstraction Layer) pour le développement d’automates via une interface web. La programmation se fait de manière graphique et modulaire sur une page web, en reliant les différentes fonctions disponibles entre elles. L’objectif principal du projet est d’ajouter de nouveaux blocs fonctionnels pouvant être utilisés dans ce cadre de développement.

= Point de départ
Le projet se construit sur la base de deux programmes déjà développés, lors du TB 2024 :
#[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
  -	*Softplcui-main*  : Gérant l’interface web côté PC.
  -	*Softplc-main* : Gérant l’activation des entrées / sorties selon le programme build depuis PLC UI.

]

Le travail effectué précédemment nous prouve la faisabilité du développement d’un tel HAL.
Les fonctionnalités implémentées dans par ce précédent projet sont :
#[
  #set list(marker: ([•], [o]),  spacing: auto, indent: 2em)
  -	Digital Input / output
  -	Analogue Input / output
  -	Ton (timer retardé à la montée)
]

= Objectif
L’objectif est l’amélioration et l’implémentation de nouvelles fonctionnalité.
Les tâches devant être réalisées sont :
#[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
  -	Modifier les programmes actuelle pour utiliser la nouvelle interface REST WDA pour piloté les nouveaux automates.
  -	L’implémentation de nouveaux blocs de haut niveau comme CAN, MQTT, WebServer, client/serveur HTTP et autres bloc. Il faudra trouver une solution pour faire ces tâches par programmation en bloc.
  -	Amélioration et extensions du frontend web.
  -	Développement d’un banc de test physique et d’une application de démonstration pour une maison connectée.
  -	Documentation et tests et rédaction du rapport, poster et présentation.
]

#pagebreak()
= L’implémentation bloc de haut niveau

Il existe plusieurs manière d’aborder le problème. Une des approches est de repérer les points commun entre ces blocs de haut niveau pour essayer d’en tirer une forme générique. On remarque que tous ces blocs ont pour objectif de transmettre et recevoir des données. Il faudra donc commencer par le développement de bloc commun pour une communication. Il faut également des blocs permettant de travailler avec des STRING. Le schéma figure 1 montre le concept d’une telle structure avec tous les blocs qui devront être développé autour pour pouvoir créer une communication. 
#figure(
  image("/resources/img/01_communication_principe.png", width: 110%),
  caption: [
    Communication principe
  ],
)
#ideabox()[
L’idée étant d’avoir un bloc communication qui s’occupe de la configuration étant différente pour can, Mqtt etc. Sur le quel, on poura double cliqués pour accéder à la page de configuration. Sur ce bloc de communication, on pourait ensuite venir lier nos 2 blocs permettant la transition de boolean vers nos trame. Par le future, en mode debug, l’utilisateur pourra voir l’état de la communication grace au "bloc de communication" et voir ce que la logique combinatoire transmet comme trame grâce au bloc en vert.
]
