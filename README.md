# go-serverless-for-vue-calendar
- [x] Go SDKを使用してローカルでDynamoDBへアクセスしてデータをGetItemする
- [x] Api-GatewayでAPIを作成
- [x] 作成したAPIとLambda関数の連携
- [x] Api-GatewayでAPIを作成
- [x] 作成したAPIとLambda関数の連携
- [x] Serverless Frameworkを使用して構成のコード化
- [x] Go SDKを使用してQueryStringを参照してDynamoDBへQueryする
- [x] Go SDKを使用してRequestBodyを参照してDynamoDBへPutItemする

# Deployment
```shell
make deploy
```

# 仕様
### DynamoDB
- TimeStamp      パーティションキー       String
- content        ソートキー              String

### lambda
- GETに対応したLambdaハンドラー（API GatewayからQuery String取得して→DynamoDBで検索かける）
- POSTに対応したLambdaハンドラー（API GatewayからRequest bodyを取得して→DynamoDBで保存）

### API Gateway
- GET /todo?date=2020-10-1 とかで
- POST /todo request body { todo: string, date: timestamp }

### Client
Vue.js（https://github.com/fujisawaryohei/vue-calendar）
- カレンダーの日付をクリックするとGETするようにする
- 「追加する」をクリックするとPOSTするようにする
