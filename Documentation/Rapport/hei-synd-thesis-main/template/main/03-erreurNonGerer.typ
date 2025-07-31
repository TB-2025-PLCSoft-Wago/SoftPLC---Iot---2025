#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
=== Erreurs non gérées <sec:erreurNonGerer>
#infobox()[Certaines erreurs n’ont pas été traitées lors de l’ancien TB. Cette section présente certaines de ces erreurs qui devront être réglées.]
==== Manque lien
Le problème est que le programme plcSoft plante au lieu d’afficher simplement une erreur et de ne pas _build_ le programme dans l’automate.

Cependant, le save est possible et le restore peut être fait après avoir relancé le programme.
/*
*Remarque* Bloc bleu : il y a des bloc bleu qui est le résultat d’un test explication dans la synthèse.
*/
#figure(
  image("/resources/img/02_Erreur_manque_lien .png", width: 100%),
  caption: [
    Défaut manque lien entre OR et TON
  ],
)

#figure(
  image("/resources/img/03_message_plcsoft_save.png", width: 100%),
  caption: [
    message plcsoft save défaut
  ],
)

#figure(
  image("/resources/img/04_Erreur_Build_manque_lien_entre_OR_TON.png", width: 80%),
  caption: [
    Erreur Build manque lien entre OR et TON
  ],
)

==== Plusieurs Output pour une Input
#importantbox()[Cela devrait être faisable, pourtant c’est interdit.]
#figure(
  image("/resources/img/2output_1input.png", width: 100%),
  caption: [
    Erreur 2 output pour un input, avec une output directement sur l'input
  ],
)


