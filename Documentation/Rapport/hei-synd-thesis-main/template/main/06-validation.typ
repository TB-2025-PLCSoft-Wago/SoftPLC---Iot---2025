#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= #i18n("validation-title", lang:option.lang) <sec:validation>

#option-style(type:option.type)[
  In addition to presenting the *results of your research in relation to your research question*, it is imperative that the validation section of your bachelor's thesis adheres to certain principles to ensure clarity, coherence, and rigor. Here are some additional considerations to enhance the validation process:

  - *Objective Description of Data*: Provide an objective and detailed description of the data used in your analysis.
  - *Utilize Graphs and Tables*: Visual aids such as graphs, charts, and tables can greatly enhance the clarity and impact of your results presentation.
  - *Link Results to Research Questions*: For each result presented, explicitly link it back to the corresponding research question or hypothesis.
  - *Ranking Results by Importance*: Prioritize your results by ranking them in order of importance or relevance to your research objectives.
  - *Confirmation or Rejection of Hypotheses*: Evaluate each result in light of the hypotheses formulated in your thesis.
]

#lorem(50)

#add-chapter(
  after: <sec:validation>,
  before: <sec:conclusion>,
  minitoc-title: i18n("toc-title", lang: option.lang)
)[
  #pagebreak()
 == WDA défauts
Le principal défaut de WDA, en plus d’être lent, est qu’il n’est pas possible d’écrire plusieurs *outputs* en une seule requête. Il est donc nécessaire de faire une requête pour chaque *output* que l’on souhaite écrire, ce qui ajoute environ 500 ms à chaque fois.

Cela peut être problématique si l’on souhaite écrire plusieurs *outputs* en même temps, car cela ralentit le temps de cycle. Par exemple, si l’on souhaite écrire 8 *outputs*, il faudra faire 8 requêtes, ce qui ajoutera environ 4 secondes au temps de cycle. À cela s’ajoutent le temps de la requête pour lire les *inputs*, ainsi que le temps de la requête pour la création de la *monitoring list*.

On ne peut donc pas garantir un temps de cycle inférieur à 5 secondes, ce qui est problématique pour une application qui demande de la rapidité ou un temps de cycle précis. La figure @fig:programmeLentWda-vs-vue présente le programme qui prend le plus de temps, et la figure @fig:programmeLentWdaAnalyse-vs-vue montre le temps de cycle de ce programme selon les étapes effectuées. On remarque de grandes variations du temps de cycle.


  #figure(
    image("/resources/img/58_wdaProblem8Outs.png", width: 100%),
    caption: [
      Programme - 8 outputs
    ],
  )
  #label("fig:programmeLentWda-vs-vue")

  #figure(
    image("/resources/img/58_wdaProblem8OutsAnalyse.png", width: 100%),
    caption: [
      Analyse programme - 8 outputs
    ],
  )
  #label("fig:programmeLentWdaAnalyse-vs-vue")


  

  == proof of concept : maison intelligente
  Le programme de la maison intelligente est un exemple d’application qui peut être développé grâce aux nouvelles fonctionnalitée développées durant ce TB. Il démontre les possibilitées de connectivité d'appareils qui communique dans différents protocole de communication et prouve que le point principal du cahier des charges est respecté. 
  
  L'exemple permet de contrôler la porte du garage, le chauffage et les lumières et d'une partie de la maison. Il est possible pour l'utilisateur de paramètrer différentes choses comme la couleur et luminosité de la lampe _Shelly_, l'allumage de la lampe manuellement et l'allumage et les consignes de chauffage.

  Le schéma des différents appareils connectée a déjà été défini sur la figure @fig:applicationMaison-vs-vue.


  Le programme est présenté en annexe au @sec:homeController-vs-vue. Il est divisé en plusieurs parties, chacune correspondant à un élément de la maison.

  La section @sec:configurationAppareils-vs-vue présente la configuration des appareils.

  La section @sec:enclenchementChauffage-vs-vue présente le programme responsable de l'enclenchement du chauffage. Il est possible de le contrôler manuellement grâce à l'interface utilisateur ou automatiquement en fonction de la température.

  La section @sec:porteGarage-vs-vue présente le programme responsable de la porte du garage. Il est possible de l'ouvrir et de la fermer grâce à un bouton qui envoie une requête HTTP.

  La section @sec:lampe-vs-vue présente le programme responsable de la lampe _Shelly Bulb_. Il est possible de contrôler la couleur et la luminosité de la lampe grâce à l'interface utilisateur. Il est également possible de l'allumer ou l'éteindre manuellement de puis se même interface.

  La section @sec:lampe2-vs-vue présente le programme responsable des lumières. Il permet de contrôler les lumières selon oû on se trouve dans la maison. Un short press sur le bouton _Shelly_ allume les lampe, dans la pièce ou on se trouve, un long press les éteintes toutes. 


    #table(
  columns: 5,
  [*Bloc*], [*Appareil*], [*Adresse IP*], [*Rôle*], [*documentations*],
  [MODBUS], [Home IO], [localhost:1502], [Simulateur de la maison connectée], [Aperçu @sec:HomeIO-vs-vue],

  [HTTP Server], [Shellybutton], [192.168.39.225], [Ouverture/fermeture porte du garage], [ Manuel @SHELLYBUTTON1USER],

  [HTTP Server], [Shellybutton], [192.168.39.226], [Lampes monitoring], [ Manuel @SHELLYBUTTON1USER],

  [HTTP Client], [Shelly Bulb Duo RGBW], [192.168.39.223], [Lampe de couleur en physique], [ Modèle @DocumentationShellyBulb

  Manuel @WebhooksHTTPSRequests],
  
  [HTTP Client], [My Strom], [192.168.39.222], [Relais chauffage], [@sec:myStromDoc-vs-vue],
  [MQTT], [Shelly H&T], [192.168.39.224], [Capteur de température], [Manuel @ModeDemploiShelly
  
  Doc MQTT @MQTTShellyTechnical],
  
  )

  == Conclusion

  
]
