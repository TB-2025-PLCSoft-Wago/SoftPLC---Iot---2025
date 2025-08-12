/*
#pagebreak()
= Exemples codes : vue programme + JSON

== Exemple WebSocket <sec:exempleUtililationWebSocket-vs-vue>
#figure(
  image("/resources/img/19_OutputPatchWDA.png", width: 100%),
  caption: [
    *WebSocket* : vue programmation
  ],
)
#figure(
    align(left,
    ```json
      {
        "data": {
          "id": "0-0-io-channels-9-dovalue",
          "type": "parameters",
          "attributes": {
            "value": true
          }
        }
      }
    ```
    ),
    caption: [*WebSocket* : graphique format JSON],
  
  )
  #label("fig:toto-vs-vue")
  */

  /*

          ```
    ),
    
  )
  #figure(
   
  align(left,
    ```json

  */
#set page(
  flipped: true,
)
== _Home-IO_ <sec:HomeIO-vs-vue>

#figure(
  image("/resources/img/66_homIOGarageRegister.png", width: 120%),
  caption: [
    _Home-IO_ : registres de l’interface de garage
  ],
)
#label("fig:homIOGarageRegister-vs-vue")

#figure(
  image("/resources/img/66_homIOGarageRegitre_emplacementBtn_1.png", width: 120%),
  caption: [
    _Home-IO_ : emplacement des boutons dans le garage – 1
  ],
)
#label("fig:homIOGarageRegitre_emplacementBtn_1-vs-vue")

== Exemples Modbus Read Bool <sec:exempleUtililationModbusReadBool-vs-vue>

#figure(
  image("/resources/img/65_configModbus_ReadBool.png", width: 80%),
  caption: [
    Modbus Read Bool : vue programmation – configuration du bloc
  ],
)
#label("fig:ModbusReadBoolConfiguration-vs-vue")

=== Exemple 1 – Modbus Read Bool – sans Quantity

Dans cet exemple, l’entrée _Quantity_ n’étant pas renseignée, on lit un seul registre aux adresses 0, 1 et 3 de l’unité « garage ». Cela correspond à la lecture des états des boutons présentés en @fig:homIOGarageRegitre_emplacementBtn_1-vs-vue.

#figure(
  image("/resources/img/65_Modbus_ReadBool_exemple_sansQuantity.png", width: 100%),
  caption: [
    Modbus Read Bool : vue debug – exemple 1 sans *Quantity*
  ],
)
#label("fig:65_Modbus_ReadBool_exemple_sansQuantity-vs-vue")

#pagebreak()

=== Exemple 2 – Modbus Read Bool – avec Quantity

Cet exemple réalise exactement la même opération que l’exemple précédent, mais cette fois avec l’entrée _Quantity_ renseignée. On lit donc toujours les registres aux adresses 0, 1 et 3 de l’unité « garage ». À noter que le bloc "constant value Input" peut être remplacé par un bloc de type "bool to string" et inversement.

#figure(
  image("/resources/img/65_Modbus_ReadBool_exemple_2_avecQuantity.png", width: 80%),
  caption: [
    Modbus Read Bool : vue debug – exemple 2 avec *Quantity*
  ],
)
#label("fig:65_Modbus_ReadBool_exemple_2_avecQuantity-vs-vue")

#pagebreak()

=== Exemple 3 – Modbus Read Bool – démonstration complète de la différence entre Quantity et Addresses

Cet exemple lit les registres aux adresses 0, 1, 3, 4 et 5 de l’unité « garage ». Le _Quantity_ fixé à 7 n’est pas pris en compte car toutes les adresses sont spécifiées via _Addresses_.

#figure(
  image("/resources/img/65_Modbus_ReadBool_exemple_3.png", width: 80%),
  caption: [
    Modbus Read Bool : vue debug – exemple 3 avec *Quantity*
  ],
)
#label("fig:65_Modbus_ReadBool_exemple_3-vs-vue")

=== Exemple 4 – Modbus Read Bool – sans Addresses

Dans cet exemple, les registres aux adresses 0 et 1 de l’unité « garage » sont lus. Le _Quantity_ fixé à 2 est pris en compte car _Addresses_ n’est pas défini. Par défaut, l’adresse utilisée est donc *0*.

#figure(
  image("/resources/img/65_Modbus_ReadBool_exemple_4.png", width: 80%),
  caption: [
    Modbus Read Bool : vue debug – exemple 4 avec *Quantity*
  ],
)
#label("fig:65_Modbus_ReadBool_exemple_4-vs-vue")

== Exemple Modbus Read Value <sec:exempleUtililationModbusReadValue-vs-vue>

Avec _Home-IO_, une seule valeur est retournée, donc la différence entre _Quantity_ et _Addresses_ n’a pas d’effet visible. Toutefois, la logique reste la même que pour _Modbus Read Bool_.

#figure(
  image("/resources/img/67_Modbus_ReadValue_exemple_1.png", width: 80%),
  caption: [
    Modbus Read Value : vue debug – exemple 1 sans *Quantity*
  ],
)
#label("fig:67_Modbus_ReadValue_exemple_1-vs-vue")

== Exemple Modbus Write Bool <sec:exempleUtililationModbusWriteBool-vs-vue>

Cet exemple permet d’allumer les lampes du garage et du passage de l’unité « garage » via la vue *User WebSocket*. En @fig:67_Modbus_WriteBool_Resultat_exemple_OFF-vs-vue, la sortie _ValueReceived_ affiche _1 ,, 1_, ce qui signifie que deux requêtes Modbus ont été faites et que, dans chaque cas, un seul _coil_ a été écrit.

#figure(
  image("/resources/img/67_Modbus_WriteBool_exemple_1.png", width: 90%),
  caption: [
    Modbus Write Bool : vue debug – lampes allumées ("on")
  ],
)
#label("fig:67_Modbus_WriteBool_exemple_ON-vs-vue")

#figure(
  image("/resources/img/67_Modbus_WriteBool_exemple_1_resultatUser_ON.png", width: 100%),
  caption: [
    Modbus Write Bool : résultat du point de vue utilisateur – lampes "on"
  ],
)
#label("fig:67_Modbus_WriteBool_Resultat_exemple_ON-vs-vue")

#figure(
  image("/resources/img/67_Modbus_WriteBool_exemple_1_OFF.png", width: 100%),
  caption: [
    Modbus Write Bool : vue debug – lampes éteintes ("off")
  ],
)
#label("fig:67_Modbus_WriteBool_exemple_OFF-vs-vue")

#figure(
  image("/resources/img/67_Modbus_WriteBool_exemple_1_resultatUser_OFF.png", width: 100%),
  caption: [
    Modbus Write Bool : résultat du point de vue utilisateur – lampes "off"
  ],
)
#label("fig:67_Modbus_WriteBool_Resultat_exemple_OFF-vs-vue")

#pagebreak()

== Exemple Modbus Write Value <sec:exempleUtililationModbusWriteValue-vs-vue>

Cet exemple permet de régler l’intensité des lampes du garage via la vue *User WebSocket*.

#figure(
  image("/resources/img/68_Modbus_WriteValue_exemple_1_prog.png", width: 100%),
  caption: [
    Modbus Write Value : vue debug – configuration du bloc
  ],
)
#label("fig:67_Modbus_WriteValue_exemple_1-vs-vue")

#figure(
  image("/resources/img/68_Modbus_WriteValue_exemple_ResultatUser.png", width: 100%),
  caption: [
    Modbus Write Value : résultat du point de vue utilisateur – lampes à intensités variables
  ],
)
#label("fig:67_Modbus_WriteValue_exemple_1_userResult-vs-vue")

#pagebreak()

#set page(
  flipped: false,
)


