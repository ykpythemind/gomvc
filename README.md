# gomvc

goでつくった単純なAPIサーバのサンプル

- シンプルなmodel/view/controllerの構成で実装する
- database migration ... https://github.com/thuss/standalone-migrations を使う (railsのmigrationの部分のみ抜き出したもの)
- rails likeなrequest testを中心としたテスト設計


## development

```
bundle install
# rake db:create # <- 初回
rake db:migrate

DATABASE_URL="$(pwd)/../db/development.sqlite3" go run cmd/server/main.go
```

(TODO: air.tomlが壊れている）

### test

```
rake db:migrate RAILS_ENV=test
TEST_DATABASE_URL="$(pwd)/../db/test.sqlite3" go test ./...
```
