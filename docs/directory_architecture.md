# ディレクトリ構成図

```txt
asset-simulator/
├── cmd/
│   └── asset-simulator/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── cli.go
│   ├── modules/
│   │   └── simulation/
│   │       ├── interface/
│   │       │   ├── input.go
│   │       │   ├── output.go
│   │       │   ├── presenter.go
│   │       │   └── service.go
│   │       └── internal/
│   │           ├── domain/
│   │           │   ├── calculator.go
│   │           │   ├── input.go
│   │           │   └── result.go
│   │           └── usecase/
│   │               └── simulate_assets.go
│   └── shared/
│       └── money/
│           └── formatter.go
├── docs/
│   ├── mvp_requirements.md
│   └── flowchart.md
├── go.mod
├── README.md
└── .gitignore
```
