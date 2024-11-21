# MY_TRAINING_GOLANG

## 起動方法

```bash
go run main.go
```

起動しない場合は`go mod tidy`を実行してみる
`go mod tidy`が実行できないときは、go のバージョンが低い可能性がある。
`go.mod`に`go 1.20`と書いてある。この場合、go バージョン 1.20 以上でないと`go mod tidy`が実行できない

## 使用方法

```bash
curl http://localhost:7000/version
```

## linter, formatter

それぞれ実行しても良いが、Makefile を作成し、コマンドから一括で実行できるようにしている。

導入

```bash
make setup
```

linter & formatter 実行

```bash
make lintAll

# ファイル名指定
make lint FILENAME=/target/directory/file.go
```

### revive

導入

```bash
go install github.com/mgechev/revive@latest
```

```bash
revive -config reviveConfig.toml -formatter friendly ./...
```

### staticcheck

導入

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```

```bash
staticcheck ./...
```

pre-commit hook について
clone 直後、pre-commit hook が働かない場合、以下のコマンドを実行すること

## pre-commit hook

### githooks の参照先を変更

```bash
git config --local core.hooksPath .githooks
```

### 実行権限を与える

```bash
chmod +x .githooks/pre-commit
```

## コード自動生成

generated ディレクトリ内のファイルは自動生成したファイルである。
自動生成実行時に上書きされるため、編集することは望ましくない。

自動生成の方法を以下に記載。

### Install oapi-codegen

- コード自動生成に使用するツール
- go のインストールが前提

[oapi-codegen GitHub](https://github.com/deepmap/oapi-codegen/tree/master)

```bash
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

```bash
oapi-codegen -version
```

### Generate Code

oapi-codegen を使用する時、1 ファイルにまとめていないと失敗するため、openapi.yaml と components 内部の統合ファイルを生成する。
`@redocly/cli`を使用すると、自動で生成してくれる。

```bash
npx @redocly/cli bundle openapi.yaml -o ./generated/allinone.gen.yaml # 出力ファイル名は任意の名称で良い
```

openapi.yaml の参照先を自動で読み込んで、allinone.gen.yaml として出力してくれる。

- config.yaml: 設定ファイル
  記法例

```config.yaml
package: api                       # package名を指定。import時に使用する
output: ./generated/openapi.gen.go # 生成コードの出力先。generatedディレクトリが存在しないと失敗する
generate:
  echo-server: true                # echoサーバーを使用する。他のサーバーを使用する場合、置き換える。
  models: true                     # モデル（型定義）を作成するか否か
output-options:
  skip-prune: true
```

oapi-codegen を使用して、go コードの自動生成をする。

```bash
oapi-codegen -config config.yaml generated/allinone.gen.yaml
```

## 各ディレクトリ,ファイル説明

### generated

- openapi.yaml から自動生成されたファイルの保存先
- 自動生成の度に上書きされるので、このファイルは編集しない

### materials

- PDF 作成の素材やテストデータをまとめておく

### config.yaml

- コード自動生成の設定ファイル

### go.mod

- 依存ライブラリの情報
- node でいう package.json と類似の役割

### go.sum

- node でいう yarn.lock と類似の役割

### main.go

- main ファイル

### openapi.yaml

- API の仕様書代わり
- 各エンドポイントで使用できるメソッドやリクエスト/レスポンスを定義
- リクエスト/レスポンスの変更、エンドポイントの追加などあれば、まずはこのファイルを編集する
