# GoInterpreter

Es wurde ein Interpreter im Rahmen der Veranstaltung [modellbasierte Softwareentwicklung] (https://sulzmann.github.io/ModelBasedSW/imp.html) implementiert. Die Umsetzung enthielt folgende Teile: 
- Evaluator
- Typechecker
- Parser

Die Formen des Evaluators und des Typecheckers waren durch die Vorgaben recht streng gegeben. Bei der Umsetzung des Parsers hatten wir mehr freie Hand. 

Wir haben im WS21/22 eine Zusatzaufgabe mit ähnlichem Thema für die Erstsemester in Programmieren 1 gemacht. In dieser ging es darum auf 5 Blätter verteilt die Grundlagen eines einfachen Plotters zu verstehen und diese mit Hilfestellungen zu imlementieren. Zuerst wurden Expressions erstellt, dann wurden Tokes hinzugefügt und zuletzt ein Parser entworfen. In diesem letzten Schritt sollte eine Implementierung mithilfe des [Shuntingyard-Algorithmus] (https://de.wikipedia.org/wiki/Shunting-yard-Algorithmus) umgesetzt werden. Der Algorithmus führt mathematische Ausdrücke der Infix-Schreibweise in solche in umgekehrte polnische Notation über, welche analog zu einem Syntaxbaum sind. (Und somit von einem Programm gelesen werden können). Als Hilfestellung wurde Pseudocode zur Verfügung gestellt, dieser musste noch in java implementiert werden. 

Nach der Entwicklung dieser Aufgabe hat es uns interessiert ob es auch möglich ist while-Schleifen und if-Abfragen in umgekehrte polnische Notation umzuwandeln. Dies ist uns gelungen und wir haben uns somit für eine ähnliche Vorgehensweise entschieden. 

## Wie funktioniert es nun konkret?

Es folgt eine allgemeine Erläuterung unserer Umsetzung und Implementierung. 

### Evaluator

Der Evaluator enthält am Ende ausführbare Dateien. Diese bestehen aus Statements und Expressions, welche mit Hilfe von stucts implementiert sind. Jedes stuct enthählt eine eigene spezifische eval-Methode. (zB enthählt das plus-stuct eine eval-Methode die zwei Expressions addiert.) So bauen sich verschiedene eval Methoden verschachtelt auf bis alle Regeln abgedeckt sind. 
Mit hilfe der eval-Mehtoden lässt sich das programm ausführen. 

Auf eine Sache möchten wir hier hinweisen: um lokale Variablen zu realisieren, haben wir zwei Änderungen vorgenommen. Zum einen kopieren wir die Map mit Variablennamen und Werten und zum anderen benutzt diese Map nun Pointer für die Werte. 

### Typechecker

Der Typechecker implementiert die selben Logiken wie der Evaluator. Statements geben hier einen Boolean zurück ob der richtige Typ ausgewertet wurde. Expressions geben des ausgewerteten Typ zurück. IllTyped wird zurück gegeben, falls es einen Typfehler gab.

### Parser

Zuerst werden aus dem String einzelne Tokens erstellt. Diese werden nun mit Hilfe des erweiterten Shuntingyard-Algorithmus in umgekehrte polnische Notation umsortiert.

Der Algorithmus arbeitet mit einem Stack und einer Liste. Ziel ist es in der Liste zum Schluss die Tokens in umgekehrter polnischer Notation zu haben. Heirführ werden Zahlen, boolsche Ausdrücke und Variablen direkt in die Liste geschrieben. Operatoren und Statements (zB `+` oder `while`) werden erstmal aus den Stack geschrieben. 

Kommt nun ein neuer Operator mit einer niedrigeren Priorität werden zuerst alle Operatoren mit höhrer Priorität vom Stack in die Liste geschrieben und dann der neue Operator auf den Stack gelegt. 

Eine Ausnahme hiervon bilden Klammern (`()` und `{}`), diese bleiben solange auf dem Stack bis die schließende Klammer kommt. Dann werden alle Operatoren zwischen den Klammern in die Liste geschrieben. 

Am Ende gehen wir die Liste druch und erstellen einen Syntaxbaum aus Expressions und Statements (aus dem Evaluator). 

## Testing

Im Anschluss haben wir uns noch dazu entschieden die Unit-Tests von go zu benuzten um unser Parser-Funktion zu testen. Hierzu wurden einige Tests hinzugefügt, welche in der Datei Helper_test.go zu finden sind. 
