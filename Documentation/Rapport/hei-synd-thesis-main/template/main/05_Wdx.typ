#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= WDx
Documentation API : https://192.168.37.134/openapi/wda.openapi.html

Documentation : https://downloadcenter.wago.com/wago/null/details/m2ddbfyuihrn44rp2h

Json Parameter : https://192.168.37.134/wda/parameter-definitions?page[limit]=20000

== Accéder IO 
=== Sans wda 
Lien pour accédéer sur page web : 192.168.37.134:8888/api/v1/hal/io

Résultat affiché sur la page web :

{"di":[false,false,false,false,false,false,false,false],"do":[false,false,false,false],"ai":[0.336,0.343],"ao":[0,0],"temp":[16778.26508951407,16778.26508951407]}

*Activer output : *

Dans le fichier « OutputUpdate.go » de softplc-main.
Pour active DO1 : PUT http://192.168.37.134:8888/api/v1/hal/do/0
#figure(
  image("/resources/img/06_OutputUpdate_do0.png", width: 100%),
  caption: [
    programme : OutputUpdate.go
  ],
)
 
C’est à partir de la ligne 54 qu’on a la lampe allumé.
Plus de détaille dans dossier « autre » puis dossier « Request ».

#figure(
  image("/resources/img/07_wireShark_Http_stream.png", width: 100%),
  caption: [
    wireShark Http stream
  ],
)

=== Avec wda
751-9402 : https://192.168.37.134/wda/parameters/0-0-io-channelcompositions-1-channels

#figure(
  image("/resources/img/08_751-9402_acceder_IO.png", width: 100%),
  caption: [
    automate 751-9402 accéder aux IO
  ],
)
 
C’est possible avec un automate 751-9402  @WAGODownloadCentera,  cependant on a le modèle 751-9401 @WAGODownloadCenter. Cela ne signifie pas que c’est impossible, mais que ce n'est pas documenté.


