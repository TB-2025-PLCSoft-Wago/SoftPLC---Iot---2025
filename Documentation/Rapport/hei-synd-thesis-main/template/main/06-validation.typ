#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *
#pagebreak()
= #i18n("validation-title", lang:option.lang) <sec:validation>

#option-style(type:option.type)[
  In addition to presenting the *results of your research in relation to your research question*, it is imperative that the validation section of your bachelor's thesis adheres to certain principles to ensure clarity, coherence, and rigor. Here are some additional considerations to enhance the validation process:

  - *Objective Description of Data*: Provide an objective and detailed description of the data used in your analysis.
  - *Utilize Graphs and Tables*: Visual aids such as graphs, charts, and tables can greatly enhance the clarity and impact of your results presentation.
  - *Link Results to Research Questions*: For each result presented, explicitly link it back to the corresponding research question or hypothesis.
  - *Ranking Results by Importance*: Prioritize your results by ranking them in order of importance or relevance to your research objectives.
  - *Confirmation or Rejection of Hypotheses*: Evaluate each result in light of the hypotheses formulated in your thesis.
]

#lorem(50)

#add-chapter(
  after: <sec:validation>,
  before: <sec:conclusion>,
  minitoc-title: i18n("toc-title", lang: option.lang)
)[
  #pagebreak()
 == WDA défauts
Le principal défaut de WDA, en plus d’être lent, est qu’il n’est pas possible d’écrire plusieurs *outputs* en une seule requête. Il est donc nécessaire de faire une requête pour chaque *output* que l’on souhaite écrire, ce qui ajoute environ 500 ms à chaque fois.

Cela peut être problématique si l’on souhaite écrire plusieurs *outputs* en même temps, car cela ralentit le temps de cycle. Par exemple, si l’on souhaite écrire 8 *outputs*, il faudra faire 8 requêtes, ce qui ajoutera environ 4 secondes au temps de cycle. À cela s’ajoutent le temps de la requête pour lire les *inputs*, ainsi que le temps de la requête pour la création de la *monitoring list*.

On ne peut donc pas garantir un temps de cycle inférieur à 5 secondes, ce qui est problématique pour une application qui demande de la rapidité ou un temps de cycle précis. La figure @fig:programmeLentWda-vs-vue présente le programme qui prend le plus de temps, et la figure @fig:programmeLentWdaAnalyse-vs-vue montre le temps de cycle de ce programme selon les étapes effectuées. On remarque de grandes variations du temps de cycle.


  #figure(
    image("/resources/img/58_wdaProblem8Outs.png", width: 100%),
    caption: [
      Programme - 8 outputs
    ],
  )
  #label("fig:programmeLentWda-vs-vue")

  #figure(
    image("/resources/img/58_wdaProblem8OutsAnalyse.png", width: 100%),
    caption: [
      Analyse programme - 8 outputs
    ],
  )
  #label("fig:programmeLentWdaAnalyse-vs-vue")


  

  == Section 2

  #lorem(50)

  == Conclusion

  #lorem(50)
]
