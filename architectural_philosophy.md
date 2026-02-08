- Directory walking → only finds files
- Worker pool → handles concurrency
- Analyzer → inspects file contents
- Result aggregator → collects scan results
- Main → orchestration only
---

- No analyzer should know about directory walking.
- No walker should know about PDF heuristics.
- No worker should know about JSON formatting or CLI specifics.

<i> That separation is what keeps systems from turning into spaghetti. </i>

### Code architecture (Mermaid)

```mermaid
%%{init: { 'themeVariables': { 'fontSize': '16px' } }}%%
classDiagram
    direction TB

    class Main {
        +main()
    }
    class Walker {
        +WalkDirectory(...)
    }
    class WorkerPool {
        +StartWorkerPool(...)
    }
    class Analyzer {
        <<interface>>
        +Supports(...)
        +Analyze(...)
    }
    class WordAnalyzer {
    }
    class PDFAnalyzer {
    }
    class ScanResult {
    }

    Main --> Walker : uses
    Main --> WorkerPool : uses
    WorkerPool --> Analyzer : depends on
    Analyzer <|.. WordAnalyzer
    Analyzer <|.. PDFAnalyzer
    WorkerPool --> ScanResult : produces
    Main --> ScanResult : aggregates
```
