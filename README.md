# Deployment
```shell
make deploy
```
# 仕様

## リポジトリ

https://github.com/fujisawaryohei/go-serverless-for-vue-calendar

## APIGateway（Lambda プロキシ統合を使用）

| HTTP メソッド | パス  | 仕様                                                                                            |
| ------------- | ----- | ----------------------------------------------------------------------------------------------- |
| GET           | /todo | DynamoDB の検索で使用するキーの値をクエリ文字列を使用してクライアント側で指定して検索結果を返す |
| POST          | /todo | DynamoDB に保存する API                                                                         |

## DynamoDB

| 項目      | データ型 | キー               |
| --------- | -------- | ------------------ |
| TimeStamp | String   | パーティションキー |
| content   | String   | ソートキー         |

## Client
リポジトリ: https://github.com/fujisawaryohei/vue-calendar
