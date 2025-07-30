#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
#heading(numbering:none)[#i18n("abstract-title", lang:option.lang)] <sec:abstract>

#option-style(type:option.type)[
  The abstract serves as a concise summary of your entire thesis, encapsulating key elements on a single page such as:
  - General background information
  - Objective(s)
  - Approach and method
  - Conclusions
]

L’objectif est de permettre aux automaticiens de programmer, les automates WAGO CC 100, de manière simple via une page web, tout en leur donnant la possibilité de réaliser des tâches complexes telles que la communication HTTP, MQTT, Modbus, ainsi que d'autres fonctions avancées. Cela permettra l’intégration de systèmes IoT, en facilitant la mise en œuvre de communications et de fonctions connectées directement depuis l’interface web.

Pour la récupération des I/O, L'interface REST #gls("WDA") a été utilisé.

Le projet se construit sur la base de deux programmes :
  #[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Softplcui-main  : Gérant l’interface web (3 vues).
    -	Softplc-main : Gérant l’activation des I/O selon le programme build depuis l'interface.
  ]
  La première étape a été de créer les blocs permettant de lier des booléens à des messages pour l'envoi (bool to string) et la réception (string to bool), ainsi que plusieurs blocs de logique de base (SR, NOT, trigger, etc.). Ensuite, le premier bloc de communication (MQTT) a été créé. Pour tester plus facilement celui-ci, une *user view* a été créée pour visualiser en temps réel l'état des _outputs_ et gérer les _inputs_ spécifiques à cette vue. De plus, une *debug view* a été créée pour visualiser les valeurs de chaque connexion du graphique programmé. Finalement, les blocs de communication (HTTP Client, HTTP Server, MODBUS) ont été créés. Ensuite, un banc de démonstration d'une maison connectée a été créé afin de prouver l'efficacité des blocs et de la solution. 
  
  En parallèle de toutes ces étapes, la *programming view* a été améliorée par l’ajout de raccourcis clavier, l’amélioration de l’aspect visuel, la possibilité d’ajouter des commentaires, l’intégration de mécanismes de gestion de fichiers, la possibilité de choisir la couleur des connections pour mieux se repérer.

  Le code a été conçu afin d'éviter tous plantage.

  Le code peut maintenant être charger dans l'automate avec DOCKER pour être utilisé pour la création de programme IoT.
#v(2em)
#if doc.at("keywords", default:none) != none {[

  _*#i18n("keywords", lang: option.lang)*_:

  #enumerating-items(
    items: doc.keywords,
    italic: true
  )
]}

/*
#v(2em)

#if doc.at("keywords", default:none) != none {[

  _*#i18n("keywords", lang: option.lang)*_:

  #enumerating-items(
    items: doc.keywords,
    italic: true
  )
]}
*/