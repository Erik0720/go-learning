# Go Lernfortschritt

Hintergrund: Erfahrung mit Python, Java, C++, C, Lua (Uni-Niveau, oberflächlich).
Ziel: Go lernen anhand kleiner CLI-Tools.

## Behandelte Themen

- Projektsetup (`go mod init`)
- Mehrfache Rückgabewerte + Error-Handling-Konvention (`val, err := f(); if err != nil`)
- `defer` (Aufruf erst bei Funktionsende, egal welcher Pfad)
- Slices (`[]string{...}`) und `for _, x := range slice`
- Eigene Funktionen mit Parametern/Rückgabewerten
- Goroutinen (`go f()`) — Achtung: Argumente eines `go`-Aufrufs werden synchron ausgewertet, nur der Aufruf selbst läuft parallel
- `sync.WaitGroup` (`Add`/`Done`/`Wait`) zum Warten auf Goroutinen
- `os.Args` für CLI-Argumente (Slice, kein `argc` nötig, leere Slice ist sicher)
- Structs (`type Task struct { ... }`) mit typisierten Feldern
- `time.Time` als Standard-Typ für Datum/Zeit (nicht `time.Date`, das ist nur die Konstruktor-Funktion)
- `encoding/json`: `json.Marshal` nutzt standardmäßig den exakten (großgeschriebenen) Go-Feldnamen als Key
- Struct-Tags (z.B. `` `json:"name"` ``) zur Steuerung von JSON-Keys, ausgelesen per Reflection

## Projekt 1: URL-Health-Checker — fertig

Pfad: `~/Projects/urlcheck/`

Nimmt beliebig viele URLs als CLI-Argumente (`./urlcheck url1 url2 ...`), prüft sie parallel per Goroutinen und gibt den Status-Code/-Text aus. Enthält korrektes Error-Handling, Ressourcen-Cleanup (`resp.Body.Close()`) und Synchronisation über `WaitGroup`.

## Projekt 2: Wortzähler — fertig

Pfad: `~/Projects/wordcount/`

Nimmt einen Dateipfad als CLI-Argument, liest zeilenweise per `bufio.Scanner` und zählt Zeilen/Wörter/Zeichen. Neu gelernt: `os.Open` + `defer file.Close()` (gleiches Muster wie `resp.Body.Close()`), `bufio.Scanner`, `strings.Fields`, `fmt.Printf` mit `%s`/`%d` (Analogie zu C's `printf`).

## Projekt 3: Todo-Manager — in Arbeit

Pfad: `~/Projects/go-learning/todo/`

Speichert Aufgaben in lokaler JSON-Datei, Befehle wie `add`/`list`/`done`. Neue Themen: JSON (`encoding/json`), Structs, Subcommands.

Stand: `Task`-Struct definiert (`ID int`, `Name string`, `Deadline time.Time`, `Done bool`), `json.Marshal` einmal ohne Tags getestet (Ergebnis: großgeschriebene Keys). Nächster Schritt: Struct-Tags für kleingeschriebene JSON-Keys ergänzen, danach Subcommands (`add`/`list`/`done`) und Persistenz in Datei.

## Didaktisches Format der Session

Guided Discovery: keine fertigen Lösungen, sondern Aufgaben + gezielte Fragen bei Bugs/Designentscheidungen, Vorhersage-vor-Ausführung bei neuen Konzepten (z.B. Goroutine-Verhalten).
