#import "@preview/hei-synd-thesis:0.1.1": *
#import "/metadata.typ": *

#let entry-list = (
  (
    key: "hei",
    short: "HEI",
    long: "Haute École d'Ingénierie",
    group: "University"
  ),
  (
    key: "synd",
    short: "SYND",
    long: "Systems Engineering",
    group: "University"
  ),
  (
    key: "it",
    short: "IT",
    long: "Infotronics",
    group: "University"
  ),
  (
    key: "rust",
    short: "Rust",
    plural: "Rust programs",
    long: "Rust Programming Language",
    description: "Rust is a modern systems programming language focused on safety, speed, and concurrency. It prevents common programming errors such as null pointer dereferencing and data races at compile time, making it a preferred choice for performance-critical applications.",
    group: "Programming Language"
  ),
  (
    key: "hal",
    short: "HAL",
    long: "Hardware Abstraction Layer",
    description: "A HAL is a layer of software that abstracts the hardware details of a computer system, allowing higher-level software to interact with the hardware without needing to know the specifics of the hardware implementation.",
    group: "Software"
  ),
  (
    key: "iot",
    short: "IoT",
    long: "Internet of Things",
    description: "The Internet of Things (IoT) refers to the network of physical objects embedded with sensors, software, and other technologies to connect and exchange data with other devices and systems over the internet.",
    group: "Technology"
  ),
  (
    key: "wago",
    short: "WAGO",
    long: "WAGO GmbH & Co. KG",
    description: "WAGO is a global leader in electrical interconnection and automation technology, known for its innovative solutions in industrial automation, building automation, and process control.",
    group: "Company"
  ),
  (
    key: "plc",
    short: "PLC",
    long: "Programmable Logic Controller",
    description: "A PLC is an industrial digital computer designed for the control of manufacturing processes, such as assembly lines, or robotic devices.",
    group: "Technology"
  ),
  (
    key: "HAL",
    short: "HAL",
    long: "Hardware Abstraction Layer",
    description: "A HAL is a layer of software that abstracts the hardware details of a computer system, allowing higher-level software to interact with the hardware without needing to know the specifics of the hardware implementation.",
    group: "Software"
  ),
  (
    key: "I/O",
    short: "I/O",
    long: "Inputs et Output",
  
  ),
  (
    key : "strechable",
    short: "strechable",
    description: "Strechable fait référence à la capacité d'un bloc d'étendre son nombre d'entrées ou sorties en fonction des besoins. Cela permet une flexibilité dans la conception. C'est un paramètre pouvant être choisi lors de la création d'un bloc.",
    group: "termes généraux"
  ),
  (
    key: "WDA",
    short: "WDA",
    long: "WAGO Device Access (accès aux paramètres et IO)",
    description: "WDA est une interface REST qui permet d'accéder aux paramètres et aux entrées/sorties des automates WAGO. Elle facilite la communication entre le logiciel de contrôle et l'automate en utilisant des requêtes HTTP pour récupérer et modifier les données.",
    group: "termes généraux"
  ),
  (
    key: "vue",
    short: "vue",
    long: "Vue",
    description: "Une vue est une représentation graphique d'un programme ou d'une partie d'un programme. Elle permet de visualiser les entrées, sorties et autres éléments du programme.",
    group: "termes généraux"
  ),
)



#let make_glossary(
  gloss:true,
  title: i18n("gloss-title", lang: option.lang),
) = {[
  #if gloss == true {[
    //#pagebreak()
    #set heading(numbering: none)
    = #title <sec:glossary>
    #print-glossary(
      entry-list,
      // show all term even if they are not referenced, default to true
      show-all: false,
      // disable the back ref at the end of the descriptions
      disable-back-references: false,
    )
  ]} else{[
    #set text(size: 0pt)
    #title <sec:glossary>
    #print-glossary(
      entry-list,
      // show all term even if they are not referenced, default to true
      show-all: false,
      // disable the back ref at the end of the descriptions
      disable-back-references: false,
    )
  ]}
]}
