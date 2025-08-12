//-------------------------------------
// Document options
//
#let option = (
  type : "final",
  //type : "draft",
  //lang : "en",
  //lang : "de",
  lang : "fr",
  template    : "thesis",
  //template    : "midterm"
)
//-------------------------------------
// Optional generate titlepage image
//
#import "@preview/fractusist:0.1.1":*
#let project-logo= dragon-curve(
  12,
  step-size: 11,
  stroke-style: stroke(
    paint: gradient.radial(..color.map.rocket),
    thickness: 3pt, join: "round"
  ),
  height: 5cm,
  fit: "contain",
)

//-------------------------------------
// Metadata of the document
//
#let doc= (
  title    : "SoftPLC pour l'IoT",
  subtitle : "",
  author: (
    name        : "Marcelin Puippe",
    email       : "marcelin.puippe@hevs.ch",
    degree      : "Bachelor",
    affiliation : "HEI-Vs",
    place       : "Sion",
    url         : "https://synd.hevs.io",
    signature   : image("/resources/img/signature.svg", width:3cm),
  ),
  keywords : ("WAGO", "automate", "HAL", "HTTP", "MQTT", "MODBUS", "frontend", "backend", "IoT", "Docker", "programming view","user view", "debug view", "golang", "react flow"),
  version  : "v0.1.0",
)

#let summary-page = (
   summary-title: "Objectif du projet", // <-- Ton nouveau titre ici
  //logo: image("/resources/img/65_Modbus_ReadBool_exemple_3.png", width: 125%),
  logo: image("/resources/img/86_introPosterImage.png", width: 124%),
  //one sentence with max. 240 characters, with spaces.
  objective: [
    L’entreprise WAGO, qui commercialise des automates, a mandaté la HES-SO afin de réaliser un nouveau HAL (Hardware Abstraction Layer) pour ses nouvelles interfaces des PLC WAGO CC100 (751-9401 et 751-9402). L’objectif est de permettre aux automaticiens de programmer de manière simple via une page web, tout en leur donnant la possibilité de réaliser des tâches complexes telles que la communication HTTP, MQTT, Modbus, ainsi que d'autres fonctions avancées. Cela permettra l’intégration de systèmes IoT, en facilitant la mise en œuvre de communications et de fonctions connectées directement depuis l’interface web.
    /*
    L’objectif est l’amélioration et l’implémentation de nouvelles fonctionnalité. 
    Les tâches devant être réalisées sont :
    #[
    #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Modifier les programmes actuelle pour utiliser la nouvelle interface REST WDA pour piloter l’automate.
    -	L’implémentation de nouveaux blocs de haut niveau comme CAN, MQTT, WebServer, client/serveur HTTP, Modbus et autres bloc. Il faudra trouver une solution pour faire ces tâches par programmation en bloc.
    -	Amélioration et extensions du frontend web.
    -	Développement d’un banc de test physique et d’une application de démonstration pour une maison connectée. 
    //-	Documentation et tests et rédaction du rapport, poster et présentation.
    ]*/
  ],
  //summary max. 1200 characters, with spaces.
  content: [
  Le projet se construit sur la base de deux programmes :
  #[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Softplcui-main  : Gérant l’interface web (3 vues).
    -	Softplc-main : Gérant l’activation des entrées / sorties selon le programme build depuis l'interface.
  ]
  La première étape a été de créer les blocs permettant de lier des booléens à des messages pour l'envoi et la réception, ainsi que plusieurs blocs de logique de base (SR, NOT, trigger, etc.). Ensuite, le premier bloc de communication a été créé. Pour tester plus facilement celui-ci, une *user view* a été créée pour visualiser en temps réel l'état des _outputs_ et gérer les _inputs_ spécifiques à cette vue. De plus, une * debug view* a été créée pour visualiser les valeurs de chaque connexion du graphique programmé. Finalement, les autres blocs de communication ont été créés. Ensuite, un banc de démonstration d'une maison connectée a été créé afin de prouver l'efficacité des blocs. En parallèle de toutes ces étapes, la  *programming view* a été améliorée par l’ajout de raccourcis clavier, l’amélioration de l’aspect visuel, la possibilité d’ajouter des commentaires, l’intégration de mécanismes de gestion de fichiers, etc.
 

  
/*
  Les fonctionnalités implémentées par ce précédent projet sont :
  #[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Digital Input / output
    -	Analogue Input / output 
    -	Ton (timer retardé à la montée)
  ]  
*/
  #figure(
  image("/resources/img/34_ApplicationMaison.png", width: 100%),
  caption: [
    Application de démonstration pour *maison connectée*
  ]
  
  
)

  ],
  address: [HES-SO Valais Wallis • rue de l'Industrie 23 • 1950 Sion \ +41 58 606 85 11 • #link("mailto"+"info@hevs.ch")[info\@hevs.ch] • #link("www.hevs.ch")[www.hevs.ch]]
)

#let professor= (
  affiliation: "HEI-Vs",
  name: "Prof. Métrailler Christopher",
  email: "christopher.metrailler@hevs.ch",
)
#let expert= (
  affiliation: "WAGO Contact SA",
  name: "Stéphane Rey",
  email: "stephane.rey@wago.com",
)
#let school= (
  name: none,
  orientation: none,
  specialisation: none,
)
#if option.lang == "de" {
  school.name = "Hochschule für Ingenieurwissenschaften Wallis, HES-SO"
  school.orientation = "Systemtechnik"
  school.specialisation = "Infotronics"
} else if option.lang == "fr" {
  school.name = "Haute École d'Ingénierie du Valais, HES-SO"
  school.shortname = "HEI-Vs"
  school.orientation = "Systèmes industriels"
  school.specialisation = "Infotronics"
} else {
  school.name = "University of Applied Sciences Western Switzerland, HES-SO Valais Wallis"
  school.shortname = "HEI-Vs"
  school.orientation = "Systems Engineering"
  school.specialisation = "Infotronics"
}

#let date = (
  submission: datetime(year: 2025, month: 8, day: 14),
  mid-term-submission: datetime(year: 2025, month: 2, day: 17),
  today: datetime.today(),
)

#let logos = (
  main: project-logo,
  topleft: if option.lang == "fr" or option.lang == "de" {
    image("/resources/img/logos/hei-defr.svg", width: 6cm)
  } else {
    image("/resources/img/logos/hei-en.svg", width: 6cm)
  },
  topright: image("/resources/img/logos/hesso-logo.svg", width: 4cm),
  bottomleft: image("/resources/img/logos/hevs-pictogram.svg", width: 4cm),
  bottomright: image("/resources/img/logos/swiss_universities-valais-excellence-logo.svg", width: 5cm),
  )
)

//-------------------------------------
// Settings
//
#let tableof = (
  toc: true,
  tof: false,
  tot: false,
  tol: false,
  toe: false,
  maxdepth: 3,
)

#let gloss    = true
#let appendix = false
#let bib = (
  display : true,
  path  : "/tail/bibliography.bib",
  style : "ieee", //"apa", "chicago-author-date", "chicago-notes", "mla"
)
