# GoInterpreter

Es wurde ein Interpreter im Rahmen der Veranstaltung [modellbasierte Softwareentwicklung] (https://sulzmann.github.io/ModelBasedSW/imp.html) implementiert. Die Umsetzung enthielt folgende Teile: 
- Evaluator
- Typechecker
- Parser
Die Formen des Evaluators und des Typecheckers waren durch die Vorgaben recht streng gegeben. Bei der Umsetzung des Parsers hatten wir mehr freie Hand. 

Wir haben im WS21/22 eine Zusatzaufgabe mit ähnlichem Thema für die Erstsemester in Programmieren 1 gemacht. In dieser ging es darum auf 5 Blätter verteilt die Grundlagen eines einfachen Plotters zu verstehen und diese mit Hilfestellungen zu imlementieren. Zuerst wurden Expressions erstellt, dann wurden Tokes hinzugefügt und zuletzt ein Parser entworfen. In diesem letzten Schritt sollte eine Implementierung mithilfe des [Shuntingyard-Algorithmus] (https://de.wikipedia.org/wiki/Shunting-yard-Algorithmus) umgesetzt werden. Der Algorithmus führt mathematische Ausdrücke der Infix-Schreibweise in solche in umgekehrte polnische Notation über, welche analog zu einem Syntaxbaum sind. (Und somit von einem Programm gelesen werden können). Als Hilfestellung wurde Pseudocode zur Verfügung gestellt, dieser musste noch in java implementiert werden. 

Nach der Entwicklung dieser Aufgabe hat es uns interessiert ob es auch möglich ist while-Schleifen und if-Abfragen in umgekehrte polnische Notation umzuwandeln. Dies ist uns gelungen und wir haben uns somit für eine ähnliche Vorgehensweise entschieden. 

## Wie funktioniert es nun konkret?



## Testing

Im Anschluss haben wir uns noch dazu entschieden die Unit-Tests von go zu benuzten um unser Parser-Funktion zu testen. Hierzu wurden einige Tests hinzugefügt, welche in der Datei Helper_test.go zu finden sind. 
