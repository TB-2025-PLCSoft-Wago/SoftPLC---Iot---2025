#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
== HTTP Client : Exemple WDA intégré <sec:httpClientExampleWDA>

L'exemple suivant montre que l'on peut utiliser le bloc *HTTP Client* pour communiquer avec l'automate WAGO via *WDA*. Il est ainsi possible de lire ou d’écrire des entrées/sorties de l’automate, voire celles d’un autre automate.

Le code (@fig:BlocHttpClientWDA_vueProgrammation-vs-vue) permet d’activer et de désactiver la sortie *DIO1*, et également de créer une *monitoring-list*.

L’objectif est de reproduire avec le bloc *HTTP Client* de PLCSoft ce que l’on pourrait faire avec *HTTPie*. Les requêtes HTTPie équivalentes sont illustrées dans les figures suivantes :
- Activation de la sortie DIO1 : @fig:BlocHttpClientWDA_DIO1_HTTPpie-vs-vue
- Désactivation de la sortie DIO1 : @fig:BlocHttpClientWDA_DIO1_HTTPpie_OFF-vs-vue
- Création d’une monitoring-list : @fig:BlocHttpClientWDA_MonitoringList_HTTPpie-vs-vue

#infobox()[
  La sortie *xDone* s’active uniquement lorsqu’une réponse est reçue, ce qui explique pourquoi elle ne s’active pas immédiatement.
]

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_config.png", width: 100%),
  caption: [
    *HTTP Client* : vue programmation – configuration du bloc pour *WDA*
  ],
)

#set page(
  flipped: true,
)

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_DIO1_RequeteSendEtResponse.png", width: 120%),
  caption: [
    *HTTP Client* : HTTPie – activation – requête envoyée et réponse reçue
  ],
)
#label("fig:BlocHttpClientWDA_DIO1_HTTPpie-vs-vue")

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_DIO1_RequeteSendEtResponse_off.png", width: 120%),
  caption: [
    *HTTP Client* : HTTPie – désactivation – requête envoyée et réponse reçue
  ],
)
#label("fig:BlocHttpClientWDA_DIO1_HTTPpie_OFF-vs-vue")

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_MonitoringList_RequeteSendEtResponse.png", width: 120%),
  caption: [
    *HTTP Client* : HTTPie – création d'une monitoring-list – requête envoyée et réponse reçue
  ],
)
#label("fig:BlocHttpClientWDA_MonitoringList_HTTPpie-vs-vue")

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_vueProgramation.png", width: 100%),
  caption: [
    *HTTP Client* : vue programmation – bloc *HTTP Client* configuré pour *WDA*
  ],
)
#label("fig:BlocHttpClientWDA_vueProgrammation-vs-vue")

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_init.png", width: 100%),
  caption: [
    *HTTP Client* : vue debug – état initial du code
  ],
)

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_DO1_ON.png", width: 100%),
  caption: [
    *HTTP Client* : vue debug – activation de la sortie DIO1
  ],
)

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_DO1_OFF.png", width: 100%),
  caption: [
    *HTTP Client* : vue debug – désactivation de la sortie DIO1
  ],
)

#figure(
  image("/resources/img/69_exemple_HTTPClient_WDA_MonitoringList_send.png", width: 100%),
  caption: [
    *HTTP Client* : vue debug – création d'une monitoring-list
  ],
)

#figure(
  image("/resources/img/69_WDA_UserView_MonitoringList.png", width: 100%),
  caption: [
    *HTTP Client* : vue utilisateur – réponse à la création d'une monitoring-list
  ],
)


