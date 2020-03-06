# remo-to-bq

remoの情報を取得しBigQueryに保存します.
Cloud Functionsで動かす想定で作成しています.

## Local Run
```
$ go run cmd/main.go
or
$ go build
$ ./remo-to-bq
```

## Deploy
env.yaml.sampleを参考にenv.yamlを作成します.

定期実行するためTopicを作成します.
```
$ gcloud pubsub topics create <topic-name>
```

保存先のBigQueryを用意します.
```
$ bq --location=asia-northeast1 mk
  --dataset \
  --default_table_expiration 3600 \
  --description "Nature Remo dataset" \
  <dataset>
$ bq mk --table <dataset>.<table> ./db/sensors.json
$ bq update --expiration 0 <dataset>.<table>
```

Cloud Functionsのデプロイをします.
```
$ gcloud functions deploy <function-name>
  --entry-point Subscription \
  --runtime go111 \
  --trigger-topic <topic-name> \
  --env-vars-file env.yaml \
  --service-account <service-account>
```

定期実行するスケジューラーを登録します.
```
$ gcloud beta scheduler jobs create pubsub <scheduler-name> \
  --schedule '*/5 * * * *' \
  --topic <topic-name> \
  --message-body 'remo-to-bq' \
  --time-zone 'Asia/Tokyo'
```
