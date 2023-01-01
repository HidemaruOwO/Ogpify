<div align="center">

![Last commit](https://img.shields.io/github/last-commit/HidemaruOwO/ogp-generate-api?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/HidemaruOwO/ogp-generate-api?style=flat-square)
![Issues](https://img.shields.io/github/issues/HidemaruOwO/ogp-generate-api?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/HidemaruOwO/ogp-generate-api?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/HidemaruOwO/ogp-generate-api/bug?style=flat-square)

# OGP Create 🔖

v1.0-beta2

## なんだこれは

この API に任意のテキストを含めて POST をすると、そのテキストを埋め込んだサイト用の OGP を生成します。

</div>

## 🚀 使い方

### リポジトリのクローン

```bash
git clone https://github.com/HidemaruOwO/OGP-Create
cd ogp-create
```

### 🔨 ビルド

```bash
go build src/ogc.go
```

### 💨 実行

CORS 対応のため、`--page-domain`及び`--api-domain`フラグが必要です。  
もし、このアプリが実行されているサーバーに紐づけられたドメインが `api.ogc.v-sli.me` で、POST するためのページのドメインが `ogc.v-sli.me`の場合は以下のコマンドのようになります。  
順次オプションの値を変更してください。  
Local で完結させたい場合は値はお好きな値を入れてください。

```bash
./ogc --api-domain api.ogc.v-sli.me --page-domain ogc.v-sli.me
```

これでサーバーが起動します。  
試しに何かしら POST してみてください。

```bash
curl -X POST -H "Content-Type: application/json" -d '{"text" : "これはテストです"}' http://127.0.0.1:3090/generate
```

### ❔ ヘルプ

```
Usage:
   [flags]

Flags:
  -a, --api-domain string    API Domain option (Example: api.ogc.v-sli.me)
  -d, --debug                Enable this flag causes logging in debug mode
  -h, --help                 help for this command
  -p, --page-domain string   Domain of the site used for the Post (Example: ogc.v-sli.me)
```

### ✈️ POST データ

```json
{
  "text": "こちらに45文字以内の文章を入力"
}
```

## ⛏️ 開発

このリポジトリの開発ブランチは`develop`ブランチです。<br/>
PR を送る場合は`develop`ブランチに送っていただくと助かります。

```bash
git checkout develop
```

## 📜 ライセンス

MIT License
