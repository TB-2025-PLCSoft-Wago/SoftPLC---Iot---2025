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
/*
L’automate a été câblé et configuré. Il est prêt à être utilisé pour la suite du projet.
Les programmes softplc-main et softplcui-main ont pu être testés et fonctionnent comme décrit dans le travail précédent de TB. Toutefois, la partie analogique n’a pas été testée, mais elle ne semble pas fonctionnelle, car elle ne figurait pas parmi les points traités dans le TB précédent.

Par ailleurs, le bloc Appliance Input ne fonctionne pas et fait planter l’interface.
De nombreuses améliorations, décrites dans la partie « @sec:objectif », restent possibles.
*/
=== Fonctionnalités développées :

L’intégration des blocs suivants :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
  - MQTT (@sec:implMQTT)
  - HTTP client (@sec:implHTTPClient)
  - HTTP serveur (@sec:implHTTPServeur)
  - Modbus (@sec:implModbus)
  - Bool to String (@sec:implBlocBooltoString)
  - String to Bool (@sec:implBlocStringtoBool)
  - Comparator GT (@sec:implComparatorGT)
  - Comparator EQ (@sec:implComparatorEQ)
  - Concat (@sec:implConcat)
  - Retain Value (@sec:implRetainValue)
  - Find (@sec:implFind)
  - Counter (@sec:implCounter)
  - SR (@sec:implSR)
  - NOT
  - TOF
  - trigger : RF_trig, Rtrig, Ftrig
]

Ainsi que d’autres fonctionnalités :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
  - L’intégration de #gls("WDA") (@sec:implWDA)
  - La gestion d’erreurs (@sec:implGestionErreur)
  - Les variables (@sec:implVariables)
]

La modification de la vue : *programming view* (@sec:implVueProgrammation) et la création des vues suivantes :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
  - User view (@sec:implVueUser)
  - Debug view (@sec:implVueDebug)
]

Pour réaliser toutes ces fonctionnalités, il a fallu :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
- introduire des types plus complexes,
- permettre la transmission de tableaux de chaînes de caractères entre le #gls("frontend") et le #gls("backend"),
- prendre en charge les blocs ayant plusieurs _outputs_,
- etc.
]

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

La différence est la *slide Bar* car avant si on ouvrait tous on n’avait pas accès aux composants du bas.
/*
#infobox()[
  *Bloc bleu* : il sont le résultat d’un test qui a été fait l’objectif était de voir comment était géré le style css des blocs. Le résultat est qu’il est géré par groupe. Ainsi, tous les Blocs _LogicalNode_ ont le même type. Il faudra donc améliorer la structure pour rendre plus facile l’attribution de style si on veut plus personaliser.
]*/


== Comparaison avec les objectifs initial
/*Les objectifs fixés par pr4 sont atteints. Le programme a pu être testé et permet de créer des programmes très simples. Un nouveau bloc a été ajouté et testé sur l'automate. Le principe de fonctionnement des codes a été vus et il est possible d’ajouter de nouveaux.
*/
Les objectifs du cahier des charges sont remplis. Des blocs de communication complexes ont été créés, tels que MQTT, client/serveur HTTP et MODBUS. De plus, plusieurs nouveaux blocs ont pu être développés, ce qui permet désormais de réaliser bien plus de fonctionnalités logiques, de traiter des chaînes de caractères, et même de travailler avec des tableaux de chaînes de caractères. L'interface REST #gls("WDA") est utilisée. 

D’un point de vue utilisateur, de nombreuses améliorations ont été apportées, notamment :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
    - les contrôles automatisés (copier/coller, annuler/rétablir, couper),
    - l'ajout d'une *slide Bar* dans l'accordion car avant si on ouvrait tous on n’avait pas accès aux composants du bas,
    - l’amélioration du visuel,
    - la gestion de fichiers,
    - la possibilité de colorer les connexions pour mieux se repérer,
    - l’ajout d’un menu déroulant sur certains blocs pour basculer plus rapidement,
    - le redimensionnement dynamique des blocs,
    - la possibilité d’ajouter des commentaires,
    - l’ajout d’une boîte à outils (toolbox).
]

Deux nouvelles vues ont également été ajoutées (*user view* et *debug view*) et leur création nécessite très peu d’effort de la part de l’utilisateur.

À cela s’ajoute la résolution de nombreux bugs et l’ajout de plusieurs mécanismes utiles à une future extension du #gls("HAL").

Finalement, un banc de test de démonstration d’une maison connectée a pu être créé, programmé et testé. Cela prouve le bon fonctionnement des solutions mises en place.

== Difficultés rencontrées

La documentation de WDA n'est pas suffisante pour comprendre le fonctionnement de la _library_. Il y a beaucoup de paramètres différents, mais on ne trouve pas ceux qui nous intéressent, la majorité d'entre eux sont pour modifier des paramètres de la configuration automate. Cependant, il a pu être remarqué que les modèles *741-9402* et *751-9401* ne sont pas les mêmes. La documentation du 741-9402 est plus complète, et l'utilisation des entrées/sorties (I/O) y est clairement expliquée. En revanche, il n'a toujours pas été trouvé de documentation concernant l'utilisation du module CAN.

== Perspectives d'avenir
=== Permettre la création de fonction
Un objectif pour la suite est d'ajouter la possibilité à l'intégrateur de créer ses propres blocs de fonction. Par exemple, on pourrait ajouter une vue similaire à la *programming view*. Dans cette vue, on pourrait créer un graphique avec, en plus, des blocs *function input* et *function output*, à qui l’on attribuerait un nom, comme pour les *variables*. Cela correspondrait aux entrées et sorties du bloc.

Il faut également un moyen de permettre de modifier les paramètres de la fonction. Une possibilité serait d’avoir une coche pour chaque paramètre lorsque cette vue est activée, et ceux cochés apparaîtraient comme paramètres de la fonction.

Une fois la fonction terminée, on lui donne un nom et elle apparaît dans l'accordion avec les autres blocs, prête à être utilisée.

L'avantage est que si une entreprise utilise souvent les mêmes mécanismes, elle s'évite un travail redondant. Cela est également utile si l’on a une appliance complexe qui doit être utilisée plusieurs fois.

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


