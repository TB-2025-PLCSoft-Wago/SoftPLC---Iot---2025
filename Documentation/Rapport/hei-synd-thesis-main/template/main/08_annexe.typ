#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#import "@preview/muchpdf:0.1.0": muchpdf

#pagebreak()
= Annexe <sec:annexe>
#add-chapter(
  after: <sec:annexe>,
  before: <sec:endAnnex>,
  minitoc-title: i18n("toc-title", lang: option.lang)
)[
#set page(
  flipped: true,
  )

  #pagebreak()
== Planning

#pagebreak()

#muchpdf(
  read("/resources/img/Part_01_Planning_V1.pdf", encoding: none),
  
)
#pagebreak()

#figure(
  image("/resources/img/Part_02_p1_Planning_V1.png", width: 100%),

)
#pagebreak()
#figure(
  image("/resources/img/Part_02_p2_Planning_V1.png", width: 100%),

)
#pagebreak()


#set page(
  flipped: false,
  )
== WDA Monitoring Lists <sec:monitoring-lists>
Documentation (accessible uniquement si l’automate est disponible) : https://192.168.37.134/openapi/WDA.openapi.html#tag/Monitoring-Lists

=== Création
#infobox()[Remarquer que nous insérons un *timeout* de 600 secondes, ce qui correspond à 10 minutes. Il est possible de le modifier si besoin. A noter qu'il faudra *recréer* la Monitoring List à la fin de ce délai pour continuer à monitorer les entrées et sorties de l'automate.

De plus, il faudra mémoriser l'ID de la Monitoring List créée, car il sera nécessaire pour les requêtes GET et DELETE. Il est possible de créer plusieurs Monitoring Lists, mais il faudra alors faire attention à ne pas dépasser le nombre maximum autorisé par l'automate.]
#figure(
  image("/resources/img/16_postMonitoring.png", width: 100%),
  caption: [
    commande post Monitoring List 
  ],
)
  #figure(
    align(left,
    ```rust
      {
        "data": {
          "type": "monitoring-lists",
          "attributes": {
            "timeout": 600
          },
          "relationships": {
            "parameters": {
              "data": [
                { "id": "0-0-io-channels-21-divalue", "type": "parameters" },
                { "id": "0-0-io-channels-22-divalue", "type": "parameters" },
                { "id": "0-0-io-channels-23-divalue", "type": "parameters" },
                { "id": "0-0-io-channels-9-dovalue", "type": "parameters" },
                { "id": "0-0-io-channels-10-dovalue", "type": "parameters" },
                { "id": "0-0-io-channels-11-dovalue", "type": "parameters" }
              ]
            }
          }
            
        }
      }
    ```
    ),
    caption: [*body send*, exemple de création d'une Monitoring List pour les 3 premières Inputs et 3 premières Outputs de l'automate], 
  )

  #figure(
    align(left,
    ```rust
      {
        "data": {
          "attributes": {
            "timeout": 600
          },
          "id": "26",
          "links": {
            "self": "/WDA/monitoring-lists/26"
          },
          "relationships": {
            "parameters": {
              "links": {
                "related": "/WDA/monitoring-lists/26/parameters"
              }
            }
          },
          "type": "monitoring-lists"
        },
        "jsonapi": {
          "version": "1.0"
        },
        "meta": {
          "doc": "/openapi/WDA.openapi.html#operation/createMonitoringList",
          "version": "1.4.1"
  }
}
    ```
    ),
    caption: [*body response*, exemple de création d'une Monitoring List pour les 3 premières Inputs et 3 premières Outputs de l'automate], 
  )
=== Utilisation

#infobox()[Pour utiliser la Monitoring List, il faut faire une requête GET sur l'URL de la Monitoring List créée précédemment. Il est possible de récupérer l'état de toutes les entrées et sorties en une seule requête GET.

On remarque qu'il y a beaucoup d'informations qui ne nous intéressent pas dans la réponse. Il y a que "path" et "value" de "attributes" qui nous intéressent. ]
#figure(
  image("/resources/img/17_GetMonitoring.png", width: 100%),
  caption: [
    commande GET Monitoring List 
  ],
)
#figure(
    align(left,
    ```rust
      {
        "data": [
          {
            "attributes": {
              "dataRank": "scalar",
              "dataType": "boolean",
              "path": "IO/Channels/21/DIValue",
              "value": false
            },
            "id": "0-0-io-channels-21-divalue",
            "links": {
              "self": "/WDA/parameters/0-0-io-channels-21-divalue"
            },
            "relationships": {
              "definition": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-21-divalue/definition"
                }
              },
              "device": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-21-divalue/device"
                }
              }
            },
            "type": "parameters"
          },
          {
            "attributes": {
              "dataRank": "scalar",
              "dataType": "boolean",
              "path": "IO/Channels/22/DIValue",
              "value": false
            },
            "id": "0-0-io-channels-22-divalue",
            "links": {
              "self": "/WDA/parameters/0-0-io-channels-22-divalue"
            },
            "relationships": {
              "definition": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-22-divalue/definition"
                }
              },
              "device": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-22-divalue/device"
                }
              }
            },
            "type": "parameters"
          },
          
    ```
    ),
     caption: [*body response*, exemple d'utilisation d'une Monitoring List pour les 3 premières Inputs et 3 premières Outputs de l'automate],
  )
  #figure(
    align(left,
    ```rust
          {
            "attributes": {
              "dataRank": "scalar",
              "dataType": "boolean",
              "path": "IO/Channels/23/DIValue",
              "value": false
            },
            "id": "0-0-io-channels-23-divalue",
            "links": {
              "self": "/WDA/parameters/0-0-io-channels-23-divalue"
            },
            "relationships": {
              "definition": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-23-divalue/definition"
                }
              },
              "device": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-23-divalue/device"
                }
              }
            },
            "type": "parameters"
          },
          {
            "attributes": {
              "dataRank": "scalar",
              "dataType": "boolean",
              "path": "IO/Channels/9/DOValue",
              "value": false
            },
            "id": "0-0-io-channels-9-dovalue",
            "links": {
              "self": "/WDA/parameters/0-0-io-channels-9-dovalue"
            },
            "relationships": {
              "definition": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-9-dovalue/definition"
                }
              },
              "device": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-9-dovalue/device"
                }
              }
            },
            "type": "parameters"
          },

    ```
    ),
  caption: [*body response*, exemple d'utilisation d'une Monitoring List pour les 3 premières Inputs et 3 premières Outputs de l'automate],
  )
  #figure(
    align(left,
    ```rust
          {
            "attributes": {
              "dataRank": "scalar",
              "dataType": "boolean",
              "path": "IO/Channels/10/DOValue",
              "value": false
            },
            "id": "0-0-io-channels-10-dovalue",
            "links": {
              "self": "/WDA/parameters/0-0-io-channels-10-dovalue"
            },
            "relationships": {
              "definition": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-10-dovalue/definition"
                }
              },
              "device": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-10-dovalue/device"
                }
              }
            },
            "type": "parameters"
          },
          {
            "attributes": {
              "dataRank": "scalar",
              "dataType": "boolean",
              "path": "IO/Channels/11/DOValue",
              "value": false
            },
            "id": "0-0-io-channels-11-dovalue",
            "links": {
              "self": "/WDA/parameters/0-0-io-channels-11-dovalue"
            },
            "relationships": {
              "definition": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-11-dovalue/definition"
                }
              },
              "device": {
                "links": {
                  "related": "/WDA/parameters/0-0-io-channels-11-dovalue/device"
                }
              }
            },
            "type": "parameters"
          }
        ],
        
    ```
    ),
  caption: [*body response*, exemple d'utilisation d'une Monitoring List pour les 3 premières Inputs et 3 premières Outputs de l'automate],
  )
  #figure(
    align(left,
    ```rust
      "jsonapi": {
          "version": "1.0"
        },
        "links": {
          "first": "/WDA/monitoring-lists/26/parameters?page[limit]=255&page[offset]=0",
          "last": "/WDA/monitoring-lists/26/parameters?page[limit]=255&page[offset]=0",
          "self": "/WDA/monitoring-lists/26/parameters?page[limit]=255&page[offset]=0"
        },
        "meta": {
          "doc": "/openapi/WDA.openapi.html#operation/getMonitoringListParameters",
          "version": "1.4.1"
        }
      }
    ```
    ),
  caption: [*body response*, exemple d'utilisation d'une Monitoring List pour les 3 premières Inputs et 3 premières Outputs de l'automate],
  )

#pagebreak()

== WDA access mode <sec:access-mode>
Afin de pouvoir *écrire* ou *lire* les entrées et sorties de l'automate, il faut activer le mode d'accès WDA correspond. 
Le paramètre "Value" peut prendre une des 3 valeurs entières suivantes :
-	0 : no access (no read/write access)
- 1 : monitor mode (read-only mode)
- 2 : control mode (read/write mode)

#figure(
  image("/resources/img/22_accessModeWDA.png", width: 100%),
  caption: [
    commande PATCH pour changer le mode d'accès WDA 
  ],
)

#figure(
    align(left,
    ```rust
      {
        "data": {
          "id": "0-0-io-iocheckaccessmode",
          "type": "parameters",
          "attributes": {
            "value": 2
          }
        }
      }
    ```
    ),
    caption: [*body send*, exemple changer le mode d'accès WDA à "control mode" (read/write mode) ],
  
  )

#pagebreak()

== WDA I/O
=== Configuration des DIO (activation des Outputs) via WDA <sec:dio-activation>
Le bornier X5 de l'automate est composé de *8 DIO*. On peut donc choisir lesquels configurer en *Input* et lesquels configurer en *Output*. Cela peut se faire via une commande *PATCH* comme le montre @fig:patchActiveOutput-vs-vue.  

Le paramètre *"value"* doit prendre un tableau de valeurs entières dont chaque index correspond à une *DIO*.  
Par exemple, si l'on veut configurer *DIO1* pour être une *Output*, on doit mettre 9 à l'index 0 du tableau.  
Si l'on veut que ce soit une *Input*, on doit mettre 1 à l'index 0 du tableau.  

Les figures @fig:valeurActivationWDAOutput-vs-vue et @fig:valeurActivationWDAOutput2-vs-vue montrent les valeurs d'activation des *Inputs* et *Outputs* pour chaque *DIO*.


 #align(center,
 table(
   columns: 2,
   stroke: none,
   align: center + horizon,
   [
     #figure(
       image("/resources/img/23_X5_DI.png", height: 4cm),
       caption: [valeurs d'activation des inputs]
     )
     #label("fig:valeurActivationWDAOutput-vs-vue")
   ],
   [
     #figure(
       image("/resources/img/23_X5_DO.png", height: 4cm),
       caption: [valeurs d'activation des outputs]
     )
     #label("fig:valeurActivationWDAOutput2-vs-vue")
   ],
 )
 )

#figure(
  image("/resources/img/21_OutputsPatchWDA.png", width: 100%),
  caption: [
    commande PATCH pour activer toutes les outputs via WDA 
  ],
)
#label("fig:patchActiveOutput-vs-vue")
#figure(
    align(left,
    ```rust
      {
        "data": {
          "id": "0-0-io-channelcompositions-4-channels",
          "type": "parameters",
          "attributes": {
            "value": [
                9,
                10,
                11,
                12,
                13,
                14,
                15,
                16
              ]
          }
        }
      }

    ```
    ),
    caption: [*body send*, exemple pour passer toutes les DIO en output],
  
  )

=== Activation d'une Output via WDA <sec:exemplePatchOutput-vs-vue>
#figure(
  image("/resources/img/19_OutputPatchWDA.png", width: 100%),
  caption: [
    commande PATCH pour activer une output via WDA 
  ],
)
#figure(
    align(left,
    ```rust
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
    caption: [*body send*, exemple pour passer la valeur d'une output à true (ici DIO1) dans l'automate via WDA],
  
  )
#figure(
  image("/resources/img/20_OutputPatchWDARespond.png", width: 60%),
  caption: [
    *Response : * commande PATCH pour activer une output via WDA 
  ],
)

===
    





<sec:endAnnex>
]


