#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#set heading(numbering: "1.")
#pagebreak()
= #i18n("conclusion-title", lang:option.lang) <sec:conclusion>

#option-style(type:option.type)[
  In the concluding section of your bachelor's thesis, you consolidate the essence of your research journey, encapsulating the most pivotal insights garnered throughout your study. Here's how to enhance and structure your conclusion:

  - *Project Summary*: Offer a succinct recapitulation of the core elements of your project, including its objectives, methodologies employed, and the main findings obtained.
  - *Comparison with Initial Objectives*: Reflect upon how your research outcomes align with the initial objectives set forth at the outset of your thesis.
  - *Encountered Difficulties*: Acknowledge and address any challenges or obstacles encountered during the course of your research.
  - *Future Perspectives*: Offer insights into potential avenues for future research or practical applications stemming from your findings.

  While you keep the conclusion of your bachelor thesis short and to the point, you deal with your results in more details in the discussion. There is no new informations in the conclusion.
]

== Résumé du projet

L’automate a été câblé et configuré. Il est prêt à être utilisé pour la suite du projet.
Les programmes softplc-main et softplcui-main ont pu être testés et fonctionnent comme décrit dans le travail précédent de TB. Toutefois, la partie analogique n’a pas été testée, mais elle ne semble pas fonctionnelle, car elle ne figurait pas parmi les points traités dans le TB précédent.

Par ailleurs, le bloc Appliance Input ne fonctionne pas et fait planter l’interface.
De nombreuses améliorations, décrites dans la partie « @sec:objectif », restent possibles.

=== Fonctionnalités développées :
- Un nouveau bloc timer de type *TOR* a été ajouté. Cela prouve la faisabilité de l’ajout de blocs simples et montre que le programme est à la fois robuste et adaptable.

- Ajout de la fonctionnalité de sauvegarde/restauration via un fichier Très pratique, elle évite de devoir réécrire le code à chaque modification ou chargement du programme.

- Ajout d’une *SlideBar* @AccordionReactBootstrap, une fonctionnalité nécessaire pour l’ergonomie de l’interface, perment de voir les éléments plus bas.

Il a été décidé de concentrer les efforts sur l’ajout de fonctionnalités permettant de gagner du temps lors du développement et des tests.


=== changement de l’interface

 #figure(
  image("/resources/img/12_VisuAvant.png", width: 100%),
  caption: [
    Interface avant
  ],
)

#figure(
  image("/resources/img/13_VisuApres.png", width: 100%),
  caption: [
    Interface après
  ],
)

La différence est la *slide Bar* car avant si on ouvrait tous on n’avait pas accès aux composants du bas. On voit également le nouveau bloc de type *TOR* qui a été ajouté.

#infobox()[
  *Bloc bleu* : il sont le résultat d’un test qui a été fait l’objectif était de voir comment était géré le style css des blocs. Le résultat est qu’il est géré par groupe. Ainsi, tous les Blocs _LogicalNode_ ont le même type. Il faudra donc améliorer la structure pour rendre plus facile l’attribution de style si on veut plus personaliser.
]


== Comparaison avec les objectifs initial
Les objectifs fixés par pr4 sont atteints. Le programme a pu être testé et permet de créer des programmes très simples. Un nouveau bloc a été ajouté et testé sur l'automate. Le principe de fonctionnement des codes a été vus et il est possible d’ajouter de nouveaux. Cependant, le programme n’est pas parfait. Il reste des points à améliorer, notamment des erreurs .


== Difficultés rencontrées

La documentation de WDA n'est pas suffisante pour comprendre le fonctionnement de la _library_. Il y a beaucoup de paramètres différents, mais on ne trouve pas ceux qui nous intéressent, la majorité d'entre eux sont pour modifier des paramètres de la configuration automate. Cependant, il a pu être remarqué que les modèles *741-9402* et *751-9401* ne sont pas les mêmes. La documentation du 741-9402 est plus complète, et l'utilisation des entrées/sorties (I/O) y est clairement expliquée. En revanche, il n'a toujours pas été trouvé de documentation concernant l'utilisation du module CAN.

== Perspectives d'avenir
=== Idées d’amélioration et extensions du #gls("frontend") web <sec:objectif>
Il y a de nombreuse possibilités d’amélioration pour l’interface utilisateur.
#[
  #set list(marker: ([--], [•]),  spacing: auto, indent: 2em)
 -	Interdire les liens qui passent sur un bloc, ajouter une intelligence de connexion.
-	Interdire les boucles de rétroaction (comme dans Codesys) ou les gérer proprement.
-	Ajouter des blocs logiques contenant un champ (pour les Inputs, c’est déjà en partie fait, mais non fonctionnel, et il n’y a pas de système de seuil).
-	Améliorer la nomenclature : éviter d’utiliser "Output" pour l’analogique et le digital, et "Input" pour les constantes. Une idée serait d’ajouter un menu déroulant sur le bloc pour choisir le type.
-	Permettre de coder le #gls("frontend") indépendamment du #gls("backend"), c’est-à-dire générer les accordéons à partir d’un fichier, qui peut être mis à jour lorsqu’on est connecté.
-	Ajout de raccourcis clavier :
	- Rendre la touche _Delete_ fonctionnel.
	-	Ctrl + C / V / A / Z / Y.
	-	Clic + glisser = multi-sélection.
	-	Touche O pour placer un Output, I pour un Input.
	-	Touche Espace pour placer un composant identique au précédent.
	-	Shift + clic gauche + glisser pour dupliquer.
	-	Une autre idée intéressante : une touche (Ctrl + Alt + C) pour ajouter automatiquement tous les blocs nécessaires autour d’un bloc ou groupe sélectionné, avec des valeurs par défaut. Par exemple, on sélectionne un bloc TON, on appuie sur la touche, et le système ajoute automatiquement une constante de 1 seconde, une entrée DIO1 (ou la suivante si déjà utilisée), et une sortie DO1. Les valeurs par défaut ne sont pas obligatoires, on peut faire sans.
	-	Touche dédiée pour activer ou désactiver les valeurs par défaut.
-	L’ordre des Inputs, Outputs, blocs logiques, etc. dans l’accordion n’est jamais le même, ce qui rend l’utilisation plus pénible car on ne peut pas s'habituer.
-	Amélioration de l’aspect visuel : couleurs et autres éléments graphiques.
-	Ajout d’une barre de menu en haut :
	-	Affichage des raccourcis clavier
	-	Aide
	-	Choix de l’emplacement d’enregistrement par l’utilisateur
	-	Ajout d’un mode Debug
]


