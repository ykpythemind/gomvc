# gomvc


- db
  - migration
    - https://github.com/thuss/standalone-migrations を使う (railsのmigrationの部分のみ抜き出したもの)


## development

```
bundle install
rake db:migrate

DATABASE_URL="$(pwd)/../db/development.sqlite3" go run cmd/server/main.go
```
