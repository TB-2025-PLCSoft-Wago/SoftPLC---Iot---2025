#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
== HTTP Server : Exemples <sec:httpServerExample>

#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_config.png", width: 100%),
  caption: [
    HTTP Server : Vue programmation – configuration du bloc
  ],
)

L’exemple suivant montre comment utiliser le bloc *HTTP Server*.

La figure @fig:BlocHTTPServer_exemple_1_init-vs-vue présente la vue debug initiale du bloc *HTTP Server*.

En @fig:BlocHTTPServer_exemple_1_POST_HTTPIE-vs-vue, on observe une requête *POST* envoyée avec *HTTPie*, tandis que la figure @fig:BlocHTTPServer_exemple_1_POST-vs-vue montre la vue debug au moment de la réception de cette requête.

De même, la figure @fig:BlocHTTPServer_exemple_1_PATCH_HTTPIE-vs-vue illustre une requête *PATCH* envoyée avec *HTTPie*, et @fig:BlocHTTPServer_exemple_1_PATCH-vs-vue montre la réception de cette requête côté debug. On y remarque que la lampe `"DO3"` s’allume, car le paramètre `param3 = true`.

La figure @fig:BlocHTTPServer_exemple_1_PUT_HTTPIE-vs-vue montre une requête *PUT* envoyée avec *HTTPie*. Sa réception est illustrée dans la figure @fig:BlocHTTPServer_exemple_1_PUT-vs-vue. Enfin, en @fig:BlocHTTPServer_exemple_1_PUT_sans-vs-vue, on constate que les paramètres envoyés avec la requête *PUT* ne sont pas actifs. Cela démontre la possibilité d’avoir des comportements dynamiques selon la nature ou la présence des paramètres.

La figure @fig:BlocHTTPServer_exemple_1_GET_SHORT-vs-vue illustre l’envoi d’une requête *GET* avec un chemin arbitraire (ici `short1`). Ce mécanisme est particulièrement utile pour gérer des *appliances* qui n’envoient ni _body_ ni _headers_ dans leurs requêtes.


#set page(
  flipped: true,
)


#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_init.png", width: 100%),
  caption: [
    HTTP Server : Vue debug - initiale 
  ],
)
#label("fig:BlocHTTPServer_exemple_1_init-vs-vue")

#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_POST._httpPIEpng.png", width: 100%),
  caption: [
    HTTP Server : HTTPie - Création d'une ressource avec HTTPie (POST)
  ],
)
#label("fig:BlocHTTPServer_exemple_1_POST_HTTPIE-vs-vue")

#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_POST.png", width: 100%),
  caption: [
    HTTP Server : Vue debug - Création d'une ressource avec HTTPie (POST)
  ],
)
#label("fig:BlocHTTPServer_exemple_1_POST-vs-vue")



#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_PATCH._httpPIE.png", width: 100%),
  caption: [
    HTTP Server : HTTPie - Modification d'une ressource (PATCH)
  ],
)
#label("fig:BlocHTTPServer_exemple_1_PATCH_HTTPIE-vs-vue")

#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_PATCH.png", width: 100%),
  caption: [
    HTTP Server : Vue debug - Modification d'une ressource (PATCH)
  ],
)
#label("fig:BlocHTTPServer_exemple_1_PATCH-vs-vue")


#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_PUT._httpPIE.png", width: 100%),
  caption: [
    HTTP Server : HTTPie - PUT
  ],
)
#label("fig:BlocHTTPServer_exemple_1_PUT_HTTPIE-vs-vue")

#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_PUT.png", width: 100%),
  caption: [
    HTTP Server : Vue debug - PUT
  ],
)
#label("fig:BlocHTTPServer_exemple_1_PUT-vs-vue")
#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_PUT_ParametersPasActif.png", width: 100%),
  caption: [
    HTTP Server : Vue debug - PUT avec les paramètres non actifs
  ],
)
#label("fig:BlocHTTPServer_exemple_1_PUT_sans-vs-vue")


#figure(
  image("/resources/img/72_ServeurHTTP_Exemple_GET_short1.png", width: 100%),
  caption: [
    HTTP Server : Vue debug - GET avec un path quelconque ici _192.168.39.56:8080/short1_
  ],
)
#label("fig:BlocHTTPServer_exemple_1_GET_SHORT-vs-vue")


