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
  - D+S1 (@sec:implDeleteShow)
  - SR Value (@sec:implSRValue)
  - Counter (@sec:implCounter)
  - SR (@sec:implSR)
  - NOT
  - TOF
  - trigger : RF_trig, Rtrig, Ftrig (@sec:implTrigger) 
]

Ainsi que d’autres fonctionnalités :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
  - L’intégration de #gls("WDA") (@sec:implWDA)
  - La gestion d’erreurs (@sec:implGestionErreur)
  - Les variables (@sec:implVariables)
  - Permettre la création de fonctions personnalisée (@sec:CreatFonct)
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

De plus, les erreurs du TB précédent (2024) ont été corrigées (@sec:erreurNonGerer) et les éléments à régler mentionnés dans (@sec:codePrecedent) ont été pris en compte.


=== Changement de l’interface

 #figure(
  image("/resources/img/12_VisuAvant.png", width: 100%),
  caption: [
    Interface programmatation - avant
  ],
)

#figure(
  image("/resources/img/13_VisuApres.png", width: 100%),
  caption: [
    Interface programmatation -  après
  ],
)

#figure(
  image("/resources/img/85_interfaceDebugConclusion.png", width: 100%),
  caption: [
    nouvelle interface debug -  après
  ],
)

#figure(
  image("/resources/img/85_interfaceUserConclusion.png", width: 100%),
  caption: [
    nouvelle interface user -  après
  ],
)
/*
Il y a de nombreuse différences :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
- introduire des types plus complexes,
- permettre la transmission de tableaux de chaînes de caractères entre le #gls("frontend") et le #gls("backend"),
- prendre en charge les blocs ayant plusieurs _outputs_,
- etc.
]
La différence est la *slide Bar* car avant si on ouvrait tous on n’avait pas accès aux composants du bas.
*/
/*
#infobox()[
  *Bloc bleu* : il sont le résultat d’un test qui a été fait l’objectif était de voir comment était géré le style css des blocs. Le résultat est qu’il est géré par groupe. Ainsi, tous les Blocs _LogicalNode_ ont le même type. Il faudra donc améliorer la structure pour rendre plus facile l’attribution de style si on veut plus personaliser.
]*/


== Comparaison avec les objectifs initiaux
/*Les objectifs fixés par pr4 sont atteints. Le programme a pu être testé et permet de créer des programmes très simples. Un nouveau bloc a été ajouté et testé sur l'automate. Le principe de fonctionnement des codes a été vu et il est possible d'en ajouter de nouveaux.
*/
Les objectifs du cahier des charges sont remplis, sauf pour CAN. Des blocs de communication complexes ont été créés, tels que MQTT, client/serveur HTTP et MODBUS. De plus, plusieurs nouveaux blocs ont pu être développés, ce qui permet désormais de réaliser bien plus de fonctionnalités logiques, de traiter des chaînes de caractères, et même de travailler avec des tableaux de chaînes de caractères. L'interface REST #gls("WDA") est utilisée. 

D’un point de vue utilisateur, de nombreuses améliorations ont été apportées, notamment :
#[
  #set list(marker: ([•], [--]), spacing: auto, indent: 2em)
    - les contrôles automatisés (copier/coller, annuler/rétablir, couper),
    - l'ajout d'une *slide Bar* dans l'accordion car avant si on ouvrait tout on n’avait pas accès aux composants du bas,
    - interdire la rétroaction,
    - l’amélioration du visuel,
    - la gestion de fichiers,
    - la possibilité de colorer les connexions pour mieux se repérer,
    - l’ajout d’un menu déroulant sur certains blocs pour basculer plus rapidement,
    - le redimensionnement dynamique des blocs,
    - la possibilité d’ajouter des commentaires,
    - la récupération des valeurs des blocs après un _restore_ grâce à des _useEffect_,
    - l’ajout d’une boîte à outils (toolbox).
]

Deux nouvelles vues ont également été ajoutées (*user view* et *debug view*) et leur création nécessite très peu d’efforts de la part de l’utilisateur.

À cela s’ajoute la résolution de nombreux bugs et l’ajout de plusieurs mécanismes utiles à une future extension du #gls("HAL").

Finalement, un banc de test de démonstration d’une maison connectée a pu être créé, programmé et testé. Cela prouve le bon fonctionnement des solutions mises en place. 

== Difficultés rencontrées

La documentation de WDA n'est pas suffisante pour comprendre le fonctionnement de la _library_. Il y a beaucoup de paramètres différents, mais on ne trouve pas ceux qui nous intéressent, la majorité d'entre eux sont pour modifier des paramètres de la configuration automate. Cependant, il a pu être remarqué que les modèles *741-9402* et *751-9401* ne sont pas les mêmes. La documentation du *741-9402* est plus complète, et l'utilisation des entrées/sorties (I/O) y est clairement expliquée. En revanche, il n'a toujours pas été trouvé de documentation concernant l'utilisation du module CAN.

Pour l’implémentation de la fonction undo/redo (ctrl + z/y), la difficulté a été de ne pas prendre trop d’évènements. En effet, *React Flow* a tendance à envoyer beaucoup d’évènements pour la moindre action.

De plus, il a fallu s’assurer que le programme ne puisse pas planter et que les fonctionnalités qui plantent soient redémarrées.

Pour la création de fonctions, il a fallu trouver un moyen de transmettre les données 
vers _nodeFunctions_ autorisé par _Golang_.

== Perspectives d'avenir
=== Améliorer la création de fonctions
Un objectif pour la suite est d’optimiser la possibilité pour l’intégrateur de créer ses propres blocs de fonction. 
Cette fonctionnalité, déjà développée (@sec:CreatFonct), nécessite toutefois des améliorations pour un meilleur fonctionnement. Une phase de test approfondie sera aussi indispensable.
/*
Par exemple, il faut ajouter la prise en charge du multi-instance pour les blocs internes qui dépendent de leur état précédent, 
comme les SR ou les Trigger. Concrètement, chaque bloc de la fonction implémentée devrait disposer de sa propre instance, 
et non partager un seul état commun. 
*/
Par exemple, il faudrait offrir un moyen de modifier facilement les paramètres d’une fonction.  
Une possibilité serait d’afficher une case à cocher pour chaque paramètre lorsque la vue est activée.
Les paramètres cochés apparaîtraient alors comme variables configurables de la fonction.  
Une autre idée consisterait à créer un tableau de variables propre au graphique affiché, 
comportant deux colonnes : *nom* et *valeur*.

Les noms définis dans ce tableau pourraient ensuite être réutilisés dans les _settings_, constantes ou autres éléments, le programme remplaçant automatiquement les noms par leurs valeurs.  
Dans le cas d’une fonction, les noms sans valeur attribuée apparaîtraient dans les _settings_ du bloc, où l’utilisateur pourrait alors définir les valeurs.  

Cette approche est particulièrement utile lorsqu’une logique complexe, telle qu’une _appliance_, 
doit être réutilisée à plusieurs reprises dans un système.  
Il suffirait alors de créer un bloc commun, avec seulement dans ses _settings_ les paramètres 
comme l’adresse IP, le numéro de registre à lire/écrire ou encore un identifiant permettant l’affichage dans la _user view_ (@sec:implVueUser).

Enfin, il serait pratique, qu'en mode _programmatation_ ou _debug_ (@sec:implVueDebug), il soit permis l’ouverture des fonctions, par exemple via un nouvel outil (_tool_) dédié.

=== Ajout de nouveau bloc : calendrier
Il pourrait être intéressant d’ajouter des blocs permettant de connaître le jour, l’heure 
ou de vérifier si l’on se trouve dans une plage de dates données.  
Ce type de fonctionnalité serait particulièrement utile dans les systèmes dépendant de l’heure ou de la période (été/hiver), comme par exemple un système de chauffage alimenté par panneaux solaires.
 

=== Idées d’amélioration et extensions du #gls("frontend") web <sec:objectif>
Il y a de nombreuse possibilités d’amélioration pour l’interface utilisateur.
#[
  #set list(marker: ([--], [•]),  spacing: auto, indent: 2em)
 //-	Interdire les liens qui passent sur un bloc, ajouter une intelligence de connexion.
//-	Interdire les boucles de rétroaction (comme dans Codesys) ou les gérer proprement.
//-	Ajouter des blocs logiques contenant un champ (pour les Inputs, c’est déjà en partie fait, mais non fonctionnel, et il n’y a pas de système de seuil).
//-	Améliorer la nomenclature : éviter d’utiliser "Output" pour l’analogique et le digital, et "Input" pour les constantes. Une idée serait d’ajouter un menu déroulant sur le bloc pour choisir le type.
//-	Permettre de coder le #gls("frontend") indépendamment du #gls("backend"), c’est-à-dire générer les accordéons à partir d’un fichier, qui peut être mis à jour lorsqu’on est connecté.
-	Ajout de raccourcis clavier :
//	- Rendre la touche _Delete_ fonctionnel.
//	-	Ctrl + C / V / A / Z / Y.
//	-	Clic + glisser = multi-sélection.
	//-	Touche O pour placer un Output, I pour un Input.
	//-	Touche Espace pour placer un composant identique au précédent.
	//-	Shift + clic gauche + glisser pour dupliquer.
	-	Une idée intéressante : une touche (Ctrl + Alt + C) pour ajouter automatiquement tous les blocs nécessaires autour d’un bloc ou groupe sélectionné, avec des valeurs par défaut. Par exemple, on sélectionne un bloc TON, on appuie sur la touche, et le système ajoute automatiquement une constante de 1 seconde, une entrée DIO1 (ou la suivante si déjà utilisée), et une sortie DO1. Les valeurs par défaut ne sont pas obligatoires, on peut faire sans. Mettre une touche dédiée pour activer ou désactiver les valeurs par défaut de cette fonctionnalitée.
//-	L’ordre des Inputs, Outputs, blocs logiques, etc. dans l’accordion n’est jamais le même, ce qui rend l’utilisation plus pénible car on ne peut pas s'habituer.
//-	Amélioration de l’aspect visuel : couleurs et autres éléments graphiques.
-	Affichage des raccourcis clavier.
-	Aide avec des exemples pour les blocs.
- Mettre en évidence les blocs qui posent problèmes.
//	-	Choix de l’emplacement d’enregistrement par l’utilisateur
//	-	Ajout d’un mode Debug
]


