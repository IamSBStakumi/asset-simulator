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

- Go 1.26以上

## Setup

```bash
git clone git@github.com:IamSBStakumi/asset-simulator.git
cd asset-simulator
go mod tidy
```

## Usage

将来資産シミュレーションは、以下のように CLI からオプション指定で実行できます。

```bash
go run ./cmd/asset-simulator \
  --mode simulate \
  --principal 1000000 \
  --current-profit 100000 \
  --invested-years 3 \
  --monthly-contribution 50000 \
  --annual-yield-rate 5
```

現在の元本・利益・積立月数・毎月積立額から利回りを逆算する場合は、`return-rate` モードを指定します。

```bash
go run ./cmd/asset-simulator \
  --mode return-rate \
  --principal 1200000 \
  --current-profit 100000 \
  --accumulated-months 24 \
  --monthly-contribution 50000
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

[ディレクトリ構成図](./docs/directory_architecture.md)

## Architecture Policy

本プロジェクトは、将来的な機能追加や API 化を見据えて、モジュラーモノリス構成で設計します。

各機能は `internal/modules/<module_name>` 配下に分離します。

```txt
internal/modules/<module_name>/
├── interface/   # モジュール外に公開する入口
└── internal/    # モジュール内部の実装詳細
    ├── domain/  # ドメインモデル・計算ルール
    └── usecase/ # アプリケーション固有の処理手順
```

MVP段階では `simulation` モジュールのみを作成します。

```txt
cmd/asset-simulator/main.go
  ↓
internal/app/cli.go
  ↓
internal/modules/simulation/interface/service.go
  ↓
internal/modules/simulation/internal/usecase/simulate_assets.go
  ↓
internal/modules/simulation/internal/domain/calculator.go
```

CLI固有の処理は `internal/app` に置きます。

一方で、資産シミュレーションの中核ロジックは `internal/modules/simulation/internal` に閉じ込めます。  
`app` 層や将来追加する API 層からは、`simulation/interface` のみを呼び出す方針です。

これにより、外部からモジュール内部の `domain` や `usecase` に直接依存することを避け、モジュール単位で変更しやすい構成にします。

## Module Design

### interface

`interface` は、モジュール外に公開する入口です。

```txt
internal/modules/simulation/interface/
├── input.go
├── output.go
├── presenter.go
└── service.go
```

主な責務は以下です。

- CLI や API から受け取る入力値の定義
- CLI や API へ返す出力値の定義
- モジュール内部の usecase 呼び出し
- 出力形式の整形

`app` 層からは、基本的にこの `interface` 配下のみを参照します。

例：

```go
import simulation "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/interface"
```

### internal

`internal` は、モジュール内部の実装詳細です。

```txt
internal/modules/simulation/internal/
├── domain/
│   ├── calculator.go
│   ├── input.go
│   └── result.go
└── usecase/
    └── simulate_assets.go
```

`domain` には、資産シミュレーションの計算ルールやドメインモデルを配置します。

`usecase` には、入力値を受け取り、ドメインロジックを使ってシミュレーション結果を組み立てる処理を配置します。

モジュール外からは、以下のような直接参照を避けます。

```go
import "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/internal/domain"
import "github.com/IamSBStakumi/asset-simulator/internal/modules/simulation/internal/usecase"
```

## Future API Design

将来的に API を追加する場合も、`simulation` モジュールの `interface` を利用します。

```txt
cmd/api/main.go
  ↓
internal/app/api.go
  ↓
internal/modules/simulation/interface/service.go
  ↓
internal/modules/simulation/internal/usecase/simulate_assets.go
  ↓
internal/modules/simulation/internal/domain/calculator.go
```

API 用の request / response と CLI 用の入出力形式が異なる場合でも、  
`simulation/interface` を境界として変換することで、`domain` や `usecase` への影響を抑えます。

## Future Modules

今後、機能が増えた場合は、以下のようにモジュールを追加します。

```txt
internal/modules/
├── simulation/
├── scenario/
├── tax/
└── inflation/
```

想定する責務は以下です。

| Module       | Responsibility                       |
| ------------ | ------------------------------------ |
| `simulation` | 将来資産額の計算                     |
| `scenario`   | 複数条件でのシナリオ比較             |
| `tax`        | 税金・控除を考慮した計算             |
| `inflation`  | インフレ率を考慮した実質資産額の算出 |

MVP段階では `simulation` のみを実装します。  
`tax` や `inflation` は、必要になった時点で追加します。

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
- 複数シナリオ比較

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
