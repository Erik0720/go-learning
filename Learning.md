# Go Lernfortschritt

Hintergrund: Erfahrung mit Python, Java, C++, C, Lua (Uni-Niveau, oberflächlich).
Ziel: Go lernen anhand kleiner CLI-Tools, eine Session pro Fortschritt.

Repo: `~/Projects/go-learning/` (Git, Remote: `git@github.com:Erik0720/go-learning.git`). Jedes Tool liegt in einem eigenen Unterordner mit eigener `go.mod`. Kompilierte Binaries sind über `.gitignore` ausgeschlossen, nicht committen.

## Wie eine neue Session ablaufen soll

- **Guided Discovery, keine fertigen Lösungen.** Aufgabe stellen (klein, ein neues Konzept pro Schritt), Nutzer schreibt den Code selbst, dann Review.
- **Review = Datei lesen, nicht nur zuhören.** Vor Feedback immer den aktuellen Code der jeweiligen `main.go` lesen.
- **Bei Bugs: Fragen statt Fixes.** Erst gezielt fragen, was der Nutzer erwartet vs. was passiert, ihn selbst die Ursache finden lassen (funktioniert gut, siehe Permission-Bug unten). Nur bei echtem Feststecken (nicht nur Ungeduld) einen konkreten Hinweis geben.
- **Vorhersage vor Ausführung** bei neu eingeführten Konzepten mit nicht-offensichtlichem Verhalten (z.B. Goroutine-Timing).
- **Neue Syntax/Stdlib-Funktionen direkt erklären** (kein Entdecken-Lassen bei reiner Konvention), aber die *Anwendung* auf die Aufgabe dem Nutzer überlassen.
- **Analogien zum Vorwissen nutzen**: C (`argv`→`os.Args`, `printf`→`Printf`, `struct`, Permission-Bits), Lua (Mehrfachrückgabewerte), Java (Structs als einfache Klassen).
- Kleine Stil-Hinweise (unnötiger `return`, ungenaue Fehlermeldung, toter Code) kurz erwähnen, aber nicht draufrum reiten.
- Nach jedem fertigen Tool: kurze Bilanz der gelernten Konzepte, dann fragen ob weiter oder Pause.

## Behandelte Themen

- Projektsetup (`go mod init`)
- Mehrfache Rückgabewerte + Error-Handling-Konvention (`val, err := f(); if err != nil`)
- `defer` (Aufruf erst bei Funktionsende, egal welcher Pfad) — Reihenfolge relativ zum Error-Check wichtig, wenn die Ressource bei Fehler `nil`/ungültig sein kann
- Slices (`[]string{...}`) und `for _, x := range slice`
- Eigene Funktionen mit Parametern/Rückgabewerten
- Goroutinen (`go f()`) — Achtung: Argumente eines `go`-Aufrufs werden synchron ausgewertet, nur der Aufruf selbst läuft parallel
- `sync.WaitGroup` (`Add`/`Done`/`Wait`) zum Warten auf Goroutinen
- `os.Args` für CLI-Argumente (Slice, kein `argc` nötig, leere Slice ist sicher)
- `os.Open` + `bufio.Scanner` (zeilenweise lesen) vs. `os.ReadFile`/`os.WriteFile` (alles auf einmal)
- `strings.Fields` (Split an Whitespace), `fmt.Printf` mit `%s`/`%d` (Analogie zu C's `printf`)
- Structs (`type Task struct { ... }`) mit typisierten Feldern
- `time.Time` als Standard-Typ für Datum/Zeit (nicht `time.Date`, das ist nur die Konstruktor-Funktion)
- `encoding/json`: `json.Marshal` nutzt standardmäßig den exakten (großgeschriebenen) Go-Feldnamen als Key; Struct-Tags (`` `json:"name"` ``) steuern die Keys
- `json.Unmarshal(data, &tasks)` braucht einen Pointer, da es in die Variable hineinschreibt (anders als `Marshal`)
- Unix-Dateiberechtigungen als drittes Argument von `os.WriteFile` (z.B. `0o644`) — **Stolperfalle:** `os.ModeAppend` ist kein Permission-Literal, sondern ein Sonderbit außerhalb der unteren 9 Bits; als Permission benutzt erzeugt es eine Datei mit `000`-Rechten (unlesbar, auch für den Owner). Debugging-Weg: `ls -la`/`stat` auf die erzeugte Datei, Bitmuster mit einem echten Octal-Literal vergleichen.

## Projekt 1: URL-Health-Checker — fertig

Pfad: `urlcheck/`

Nimmt beliebig viele URLs als CLI-Argumente (`./urlcheck url1 url2 ...`), prüft sie parallel per Goroutinen und gibt den Status-Code/-Text aus. Enthält korrektes Error-Handling, Ressourcen-Cleanup (`resp.Body.Close()`) und Synchronisation über `WaitGroup`.

## Projekt 2: Wortzähler — fertig

Pfad: `wordcount/`

Nimmt einen Dateipfad als CLI-Argument, liest zeilenweise per `bufio.Scanner` und zählt Zeilen/Wörter/Zeichen. Nutzt `fmt.Printf` statt manueller String-Verkettung mit `strconv.Itoa`.

## Projekt 3: Todo-Manager — in Arbeit

Pfad: `todo/`

Ziel: Aufgaben in lokaler JSON-Datei speichern, Befehle wie `add`/`list`/`done`.

Stand:
- `Task`-Struct mit JSON-Tags: `ID int`, `Name string`, `Deadline time.Time`, `Done bool`, alle mit lowercase `json:"..."`-Tags.
- Round-Trip funktioniert: Slice von Tasks → `json.Marshal` → `os.WriteFile("tasks.json", ...)` → `os.ReadFile` → `json.Unmarshal` → Ausgabe. Alle Fehler korrekt geprüft.
- Permission-Bug gefunden und gefixt (siehe Stolperfalle oben), toter Code (leere Hilfsfunktion) entfernt.

Nächste Schritte:
1. Subcommands über `os.Args[1]` (`add`, `list`, `done`) statt jedes Mal dieselben zwei Test-Tasks zu überschreiben.
2. `add`: neue Task an bestehende (aus Datei geladene) Liste anhängen, zurückschreiben.
3. `list`: Tasks aus Datei laden und lesbar ausgeben (nicht nur `fmt.Println(tasks)` roh).
4. `done`: Task per ID/Name als erledigt markieren.
5. Eindeutige IDs vergeben, sobald `add` mehrfach genutzt wird (aktuell hart codiert).
