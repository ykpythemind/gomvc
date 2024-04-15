# gomvc


- db
  - migration
    - https://github.com/thuss/standalone-migrations を使う (railsのmigrationの部分のみ抜き出したもの)


## development

```
bundle install
rake db:migrate

DATABASE_URL="$(pwd)/../db/development.sqlite3" air -c air.toml
```

### test

```
rake db:migrate RAILS_ENV=test
TEST_DATABASE_URL="$(pwd)/../db/test.sqlite3" go test ./...
```
