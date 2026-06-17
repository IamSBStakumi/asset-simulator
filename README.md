# Asset Simulator

将来の資産額をシミュレーションする Go 製 CLI ツールです。

元本、現在の利益、積立期間、毎月の積立額、想定利回りなどを入力し、  
1年後・3年後・5年後・10年後・20年後・30年後・40年後の予想資産額を出力します。

## 目的

投資信託や積立投資を行った場合に、将来的な資産推移をざっくり把握するためのツールです。

MVPでは CLI として実装しますが、将来的には以下の拡張を想定しています。

- API 化
- Web UI 化
- 複数シナリオ比較
- 税金・手数料を考慮したシミュレーション
- インフレ率を考慮した実質資産額の算出

## Features

MVPで想定している機能は以下です。

- 元本を入力できる
- 現在の利益を任意で入力できる
- これまでの積立期間を任意で入力できる
- 毎月の積立額を入力できる
- 年利回りを入力できる
- 指定年数後の予想資産額を出力できる

出力対象の年数は以下です。

- 1年後
- 3年後
- 5年後
- 10年後
- 20年後
- 30年後
- 40年後

## Requirements

- Go 1.22 以上

## Setup

```bash
git clone git@github.com:IamSBStakumi/asset-simulator.git
cd asset-simulator
go mod tidy
```

## Usage

MVP段階では、以下のようなCLI実行を想定しています。

```bash
go run ./cmd/asset-simulator
```

将来的には、以下のようなオプション指定で実行できる形を目指します。

```bash
asset-simulator \
  --principal 1000000 \
  --current-profit 100000 \
  --monthly-contribution 50000 \
  --annual-return-rate 5
```

想定出力例：

```txt
1年後: 1,730,000円
3年後: 3,060,000円
5年後: 4,530,000円
10年後: 8,980,000円
20年後: 22,110,000円
30年後: 43,550,000円
40年後: 78,600,000円
```

※ 実際の計算結果は、計算ロジックの実装内容により変わります。

## Project Structure

```txt
asset-simulator/
├── cmd/
│   └── asset-simulator/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── cli.go
│   ├── domain/
│   │   └── simulation.go
│   └── usecase/
│       └── simulate_assets.go
├── docs/
│   ├── requirements.md
│   └── flowchart.md
├── go.mod
├── README.md
└── .gitignore
```

## Architecture Policy

CLI固有の処理と、資産シミュレーションの計算ロジックを分離します。

```txt
cmd/main.go
  ↓
internal/app/cli.go
  ↓
internal/usecase/simulate_assets.go
  ↓
internal/domain/simulation.go
```

この構成にすることで、将来的にAPI化する場合でも、  
`usecase` と `domain` のロジックを再利用しやすくします。

API化する場合の想定構成：

```txt
cmd/api/main.go
  ↓
internal/app/api.go
  ↓
internal/usecase/simulate_assets.go
  ↓
internal/domain/simulation.go
```

## MVP Scope

最初のMVPでは、以下を対象とします。

- CLIから入力値を受け取る
- 複利計算で将来資産額を算出する
- 1, 3, 5, 10, 20, 30, 40年後の資産額を表示する
- テスト可能な形で計算ロジックを分離する

以下はMVPでは対象外とします。

- Web UI
- APIサーバー
- ユーザー管理
- データ保存
- 税金計算
- 手数料計算
- 為替考慮
- インフレ率考慮

## Calculation Policy

MVPでは、以下のような単純化した前提で計算します。

- 年利回りは一定とする
- 毎月一定額を積み立てる
- 毎月複利で増加する
- 税金・手数料は考慮しない
- 元本と現在の利益の合計を初期資産額とする

初期資産額：

```txt
initialAssets = principal + currentProfit
```

月利：

```txt
monthlyRate = annualReturnRate / 12
```

将来資産額：

```txt
futureAssets = initialAssets を毎月複利運用しつつ、毎月積立額を加算する
```

## Development

実行：

```bash
go run ./cmd/asset-simulator
```

テスト：

```bash
go test ./...
```

フォーマット：

```bash
gofmt -w .
```

## Roadmap

### Phase 1: MVP CLI

- [ ] プロジェクト初期構成作成
- [ ] README作成
- [ ] 入力値の型定義
- [ ] シミュレーション計算ロジック実装
- [ ] CLIから実行
- [ ] ユニットテスト追加

### Phase 2: CLI改善

- [ ] flagによる入力対応
- [ ] 入力バリデーション
- [ ] 表形式での出力
- [ ] JSON出力オプション

### Phase 3: API化

- [ ] HTTP API追加
- [ ] request / response 型定義
- [ ] usecase層の再利用
- [ ] APIテスト追加

### Phase 4: Web UI

- [ ] フロントエンド追加
- [ ] グラフ表示
- [ ] 複数シナリオ比較
- [ ] 入力値の保存

## License

未定
