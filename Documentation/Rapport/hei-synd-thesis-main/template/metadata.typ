//-------------------------------------
// Document options
//
#let option = (
  type : "final",
  //type : "draft",
  //lang : "en",
  //lang : "de",
  lang : "fr",
  //template    : "thesis",
  template    : "midterm"
)
//-------------------------------------
// Optional generate titlepage image
//
#import "@preview/fractusist:0.1.1":*
#let project-logo= dragon-curve(
  12,
  step-size: 10,
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
  keywords : ("HEI-Vs", "Systems Engineering", "Infotronics", "Thesis", "Template"),
  version  : "v0.1.0",
)

#let summary-page = (
   summary-title: "Objectif du projet", // <-- Ton nouveau titre ici
  logo: project-logo,
  //one sentence with max. 240 characters, with spaces.
  objective: [
    L’objectif est l’amélioration et l’implémentation de nouvelles fonctionnalité. 
    Les tâches devant être réalisées sont :
    #[
    #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Modifier les programmes actuelle pour utiliser la nouvelle interface REST WDA pour piloté l’automate.
    -	L’implémentation de nouveaux blocs de haut niveau comme CAN, MQTT, WebServer, client/serveur HTTP et autres bloc. Il faudra trouver une solution pour faire ces tâches par programmation en bloc.
    -	Amélioration et extensions du frontend web.
    -	Développement d’un banc de test physique et d’une application de démonstration pour une maison connectée. 
    -	Documentation et tests et rédaction du rapport, poster et présentation.
    ]
  ],
  //summary max. 1200 characters, with spaces.
  content: [
  Le projet se construit sur la base de deux programmes déjà développés, lors du TB 2024 :
  #[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Softplcui-main  : Gérant l’interface web côté PC.
    -	Softplc-main : Gérant l’activation des entrées / sorties selon le programme build depuis PLC UI.
  ]
  Le travail effectué précédemment nous prouve la faisabilité du développement d’un tel HAL.

  Les fonctionnalités implémentées par ce précédent projet sont :
  #[
  #set list(marker: ([•], [--]),  spacing: auto, indent: 2em)
    -	Digital Input / output
    -	Analogue Input / output 
    -	Ton (timer retardé à la montée)
  ]  

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
  submission: datetime(year: 2025, month: 2, day: 17),
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
