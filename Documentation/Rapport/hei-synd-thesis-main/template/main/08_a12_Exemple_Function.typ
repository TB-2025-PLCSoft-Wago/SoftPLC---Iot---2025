= Exemple création de fonction <sec:creationFonction>

== Fonction OR, AND, XOR

#figure(
  image("/resources/img/94_Fonction_fonc1_xor.png", width: 100%),
  caption: [
    *Fonction OR, AND, XOR* : fonction
  ],
)

#figure(
  image("/resources/img/94_Fonction_main_xor.png", width: 100%),
  caption: [
    *Fonction OR, AND, XOR* : programme _build_ – mode debug
  ],
)

== Imbrication de fonction

L'idée est d'avoir une fonction _TestFunc1_ (@fig:93_Fonction_fonc1_imbriaction-vs-vue) à l'intérieur d'une fonction _TestFunc2_ qui est ensuite _build_ dans le programme principal (@fig:93_Fonction_main_imbriaction-vs-vue).  
Il faut charger _TestFunc1_, puis charger _TestFunc2_, puis charger de nouveau _TestFunc1_.

#figure(
  image("/resources/img/93_Fonction_fonc1_imbriaction.png", width: 100%),
  caption: [
    *Imbrication de fonction* : fonction 1 (TestFunc1)
  ],
)
#label("fig:93_Fonction_fonc1_imbriaction-vs-vue")

#figure(
  image("/resources/img/93_Fonction_fonc2_imbriaction.png", width: 100%),
  caption: [
    *Imbrication de fonction* : fonction 2 (TestFunc2)
  ],
)
#label("fig:93_Fonction_fonc2_imbriaction-vs-vue")

#figure(
  image("/resources/img/93_Fonction_main_imbriaction.png", width: 100%),
  caption: [
    *Imbrication de fonction* : programme _build_
  ],
)
#label("fig:93_Fonction_main_imbriaction-vs-vue")

#figure(
  image("/resources/img/93_Fonction_accordion_imbriaction.png", width: 50%),
  caption: [
    *Imbrication de fonction* : accordion
  ],
)

== Fonction vérification message reçu

L'idée est d'avoir une fonction qui vérifie un message reçu pour savoir s’il comporte une erreur, active _error_ dans ce cas, ou active _isActive_ s’il contient _true_.

#figure(
  image("/resources/img/95_fonction_msgCheck_user.png", width: 100%),
  caption: [
    *Fonction vérification message reçu* : user view
  ],
)

#set page(
  flipped: true,
)
#figure(
  image("/resources/img/95_fonction_msgCheck_fonc.png", width: 90%),
  caption: [
    *Fonction vérification message reçu* : fonction
  ],
)

#set page(
  flipped: false,
)
#figure(
  image("/resources/img/95_fonction_msgCheck_main.png", width: 120%),
  caption: [
    *Fonction vérification message reçu* : programme _build_ – mode debug
  ],
)


