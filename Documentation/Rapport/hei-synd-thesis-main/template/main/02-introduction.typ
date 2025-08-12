#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= #i18n("introduction-title", lang:option.lang) <sec:intro>
Ce document présente les travaux réalisés ainsi que les perspectives à venir pour le projet *PLCSoft* destiné à *WAGO*.

Ce projet consiste en l’ajout de nouvelles fonctionnalités à un *HAL* (Hardware Abstraction Layer) pour le développement d’automates via une interface web. La programmation se fait de manière graphique et modulaire sur une page web, en reliant les différentes fonctions disponibles entre elles. Le but étant de permettre une programmation simple et intuitive. L’objectif principal du projet est l'ajout de nouveaux blocs fonctionnels pouvant être utilisés dans ce cadre de développement et plus spécifiquement l'ajout de blocs de communication réseau. En effet, la communication réseau est essentielle pour le développement de programmes automates modernes et le développement dans des systèmes #gls("iot"), comme nous l'explique l'article @Sharma2019.

"_Scientists claim that the potential benefit derived from this technology will sprout a foreseeable future where the smart objects sense, think and act. Internet of Things is the trending technology and embodies various concepts such as fog computing, edge computing, communication protocols, electronic devices, sensors, geo-location etc._"

Il y a également l'article @s23167194 qui explique l'importance de l'#gls("iot") dans de nombreux domaines comme la santé, l'agriculture, l'industrie 4.0, la domotique, etc. Il est donc important pour *WAGO* de permettre le développement de programmes automates pour ces systèmes.


== Point de départ
Le projet se construit sur la base de deux programmes déjà développés, lors du TB 2024 :
#[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
  -	*Softplcui-main*  : Gérant l’interface web (#gls("frontend")).
  -	*Softplc-main* : Gérant l’activation des entrées / sorties selon le programme build depuis PLC UI (#gls("backend")).

]

Le travail effectué précédemment nous prouve la faisabilité du développement d’un tel #gls("HAL").
Les fonctionnalités implémentées par ce précédent projet sont :
#[
  #set list(marker: ([•], [o]),  spacing: auto, indent: 2em)
  -	Digital input / output
  -	Analogue input / output
  -	Ton (timer retardé à la montée)
]

== Objectif
L’objectif est l’amélioration et l’implémentation de nouvelles fonctionnalités.
Les tâches devant être réalisées sont les suivantes :
#[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
  -	Modifier les programmes actuels pour utiliser la nouvelle interface REST WDA pour piloter les nouveaux automates.
  -	L’implémentation de nouveaux blocs de haut niveau comme CAN, MQTT, WebServer, client/serveur HTTP et autres blocs. Il faudra trouver une solution pour faire ces tâches par programmation en blocs.
  -	Amélioration et extensions du #gls("frontend") web.
  -	Développement d’un banc de test physique et d’une application de démonstration pour une maison connectée.
  -	Documentation, tests et rédaction du rapport, poster et présentation.
]

== Application de démonstration pour maison connectée

L’application de démonstration d'une maison connectée (@fig:applicationMaison-vs-vue) a pour but de démontrer les capacités du #gls("HAL") qui sera développé dans le cadre de ce projet. Il s’agit de réaliser un programme automate permettant de contrôler et de surveiller les différents équipements d’une maison connectée, tels que les lumières, les prises électriques, les capteurs de température, etc.  
Une page web doit permettre de visualiser l’état des équipements et de les contrôler.

La fonctionnalité de l’application de démonstration qui a été choisie et qui sera développée est la suivante :  
Depuis une interface web, il sera possible de contrôler une lampe en réglant son intensité et sa couleur. Depuis cette même interface, il sera également possible de régler une consigne de température. Un capteur de température enverra l'information et une prise électrique s'activera si la température est trop basse. De plus avec un bouton _Shelly_, on pourra contrôler la porte de garage. Ainsi qu'un autre bouton _Shelly_ gérant l'éclairage du garage et de la cuisine selon notre emplacement dans la maison.
   
 #iconbox(linecolor: hei-pink)[Le résultat final est présenté au @sec:validationProofConcept et l’annexe @sec:homeController-vs-vue montre le programme. Cependant, il sera difficile à comprendre sans avoir assimilé les chapitres précédents.]
#figure(
  image("/resources/img/34_ApplicationMaison.png", width: 100%),
  caption: [
    Application de démonstration pour *maison connectée*
  ],
)
#label("fig:applicationMaison-vs-vue")


