# go-serverless-for-vue-calendar
- [x] Go SDKを使用してローカルでDynamoDBへアクセスしてデータをGETする
- [ ] Api-GatewayでAPIを作成
- [ ] 作成したAPIとLambda関数の連携
- [ ] Go SDKを使用してローカルでDynamoDBへPUTする
- [ ] Api-GatewayでAPIを作成
- [ ] 作成したAPIとLambda関数の連携

# 仕様
### DynamoDB
- ID            パーティションキー       String
- TimeStamp     レンジキー            String
- content                            String

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
