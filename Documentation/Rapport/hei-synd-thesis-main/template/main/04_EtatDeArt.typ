#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= Etat de l'art
Durant le cours Projet 4, plusieurs solutions existantes similaires à ce que l’on souhaitait développer ont été identifiées. Cette section présente également le travail réalisé lors du TB 2024, qui précède celui-ci.

== n8n @N8nioN8n2025
C’est un logiciel d’automation présent en ligne sur gitHub. Il permet une programmation en no-code sur page web comme ce qu’on l’on essaie de faire. Il est surtout conçu pour l’automation de tâche simple. Il permet notamment l’automation de chat alimenté par l’IA, c’est-à-dire des réponses automatiques. Ce n’est pas ce qui nous intéresse mais cela peut nous aider à avoir des idées.


#figure(
  image("/resources/img/05_n8n_exemple.png", width: 100%),
  caption: [
    Exemple gitHub n8n
  ],
)

L’activation d’une output en n8n peut ce faire de la manière suivante. Il suffit d’un bloc HTTP Request1 qu’on configure.

#figure(
  image("/resources/img/09_Bloc_HTTP_Request1_n8n.png", width: 60%),
  caption: [
    Bloc HTTP Request1 en n8n  
  ],
)

#figure(
  image("/resources/img/10_configuration_HTTP_Request1_n8n_activation_output_2.png", width: 100%),
  caption: [ configuration HTTP Request1 en n8n pour activation output 2 ],
)

Le code de n8n est disponible sur GitHub @N8nioN8n2025. Une analyse a été faite, mais il utilise une librairie différente de celle que nous utilisons. Nous ne pouvons donc pas nous en inspirer directement. Cependant, il est possible de s’en inspirer pour la création de l’interface graphique et de la logique de programmation.

== total js @JavaScriptLibrariesComponents
C'est un logiciel de programmation en no-code. Il est possible de faire des programmes en JS, mais il y a aussi une interface graphique qui ressemble à ce qu'on voudrait faire. 
#figure(
  image("/resources/img/11_totalJS_exemple.jpg", width: 100%),
  caption: [
    Exemple totalJS
  ],
)


== Analyse critique du code précédent
Le code précédent est un bon point de départ pour la création d'un #gls("HAL") de développement d’automate. Cependant, il y a plusieurs points à améliorer.

Le code est presque totalement fonctionnel malgré quelques petites erreurs. Cependant, la structure actuelle ne permet pas l'intégration de blocs plus complexes que ce qui a été fait. En effet, par exemple, nous pouvons recevoir et transmettre qu'un nombre très limité de paramètres, et la structure de ceux-ci n'est pas très flexible. Le typage des blocs est également très limité. Il utilise le type *float64* pour tous les blocs, ce qui n'est pas adapté à la transmission de données plus complexes. Il serait préférable d'utiliser un type plus flexible, comme un type any, un type générique ou au moins un type *string*. De plus, le code est très difficile à lire et à comprendre car il utilise beaucoup d'imbrications de boucles *for* et *if*. Il manque des commentaires. Il n'a également pas prévu la possibilité d'avoir *plusieurs output* pour un bloc.

Au niveau visuel, la structure doit être améliorée. En effet, on ne peut pas continuer à mettre tous les blocs logiques dans le même fichier. Il faudrait au moins créer des fichiers séparés pour chaque type de bloc logique. La taille des blocs n'est également pas adaptée aux besoins. Le contenu des blocs n’est pas mis à jour après un "restore", il faut implémenter des *useEffect*.

Le code ne permet pas de changer le type d'un bloc de manière dynamique, car c'est le programme #gls("backend") qui l'envoie à l'initialisation du programme. Cela est donc impossible sans de grosses modifications.

Il n'y a pas non plus de méthode permettant de réinitialiser les nodes, ce qui est nécessaire pour pouvoir créer des blocs plus complexes. En effet, il peut être nécessaire de réinitialiser certaints nodes pour pouvoir faire un nouveau _build_ proprement. 


=== Ajouter des blocs simples

Il y a également des points positifs. La structure des "Nodes" est facile à utiliser.  
Pour ajouter un bloc simple, il suffit de créer un fichier dans le dossier *nodes* du programme #gls("backend") (`softplc-main`).  
On peut, par exemple, copier-coller puis renommer un fichier existant selon le type de bloc à ajouter. Il existe trois types de blocs : `LogicalNode`, `OutputNode` et `InputNode`.

#warningbox()[
* nommer un fichier :* il peut parfois y avoir des problèmes si le nom du bloc commence par une lettre en début d'alphabet.
#label("warn-alpha")
]
Vous pouvez copier le fichier `OrNode`, puis le modifier. Pour cela, utilisez la fonctionnalité _rechercher/remplacer_ afin de remplacer `"Or"` par le nom de votre nouveau bloc. Il faut le faire deux fois :  
- une fois avec `"Or"` (majuscule)  
- une fois avec `"or"` (minuscule)  

Les éléments suivants doivent ensuite être adaptés :

1. La fonction `ProcessLogic`, où vous écrirez la logique spécifique de votre bloc.  
2. La variable de type de la structure `nodeDescription`, dont le lien avec la partie visuelle est montré dans la figure @fig:nodedescription-vs-vue.

#figure(
  image("/resources/img/14_LienEntreVueNewProcess.png", width: 100%),
  caption: [
    *nodeDescription* VS vue — modifier le fichier `OrNode` pour créer un nouveau bloc
  ],
)
#label("fig:nodedescription-vs-vue")


#include "/main/03-erreurNonGerer.typ"




