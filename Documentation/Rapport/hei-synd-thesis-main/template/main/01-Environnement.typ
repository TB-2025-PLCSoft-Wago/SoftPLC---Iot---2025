#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
== Environnement de développement

#table(
  columns: 3,
  [*Appareil*], [*Adresse IP*], [*Utiliser dans soft*],
  [Automate WAGO CC100], [192.168.37.134], [softplc-main],
  [PC], [localhost], [softplcui-main],
)
//TO DO : Rajouter les autres appareils

La salle *23.N320* utilisé pour connecté l’automate WAGO CC100 sur le réseau (23.N32x).

#table(
  columns: 3,
  [*Logiciel*], [*Appareil*], [*Commentaires*],
  [Goland], [PC], [Développement logiciel],
  [Umlet], [PC], [Réalisation de schéma basé développement],
  [Firefox], [PC], [Permettant meilleure visualisation de WDx],
  [HTTPie], [PC], [Permettant de tester les requêtes HTTP],
  [miro], [site], [Réalisation de schéma, prise de note, réflexion, analyse],
  [ChatGPT  @ChatGPT], [IA], [Correction orthographique, aide débogage, code de certaines petites fonctions.],
  [WDx], [Automate], [
    Dénomination générale pour parler de WDM + WDD + WDA :

    - WDM = WAGO Device Model (standard utilisé pour les WDD)
    - WDD = WAGO Device Description (manifeste décrivant ce que le produit met à disposition)
    - WDA = WAGO Device Access (accès aux paramètres et IO)
  ],
  [WDA], [Automate], [
    Nouvelle librairie, interface REST, WAGO accessible par web en JSON.
    Permet la récupération des informations de l’automate et le contrôle des entrées/sorties par requête HTTP en format JSON.

    https://192.168.37.134/wda/parameters?page[limit]=20000
    https://192.168.1.126/wda/parameter-definitions?page[limit]=20000
  ],
)

#set table(
  stroke: none,
  gutter: 0.2em,
  fill: (x, y) =>
    if x == 0 or y == 0 { gray },
  inset: (right: 1.5em),
)

#show table.cell: it => {
  if it.x == 0 or it.y == 0 {
    set text(white)
    strong(it)
  } else if it.body == [] {
    // Replace empty cells with 'N/A'
    pad(..it.inset)[_N/A_]
  } else {
    it
  }
}

#let a = table.cell(
  fill: green.lighten(60%),
)[A]
#let b = table.cell(
  fill: aqua.lighten(60%),
)[B]




