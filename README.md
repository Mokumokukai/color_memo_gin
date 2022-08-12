# 初めにすること
.envファイルを作成し、以下の設定を行う。
```
MYSQL_ROOT_PASSWORD=
MYSQL_DATABASE=
MYSQL_USER=
MYSQL_PASSWORD=
```

# build
docker-compose build

docker-compose up -d
# コンテナに入る
docker-compose exec go ash

[設計書について](./設計書/README.MD)
# migration
## up
migrate -path src/database/migrations -database ${DB_PATH} up 
## down 
migrate -path src/database/migrations -database ${DB_PATH} down
## migration ファイル作成
migrate create -ext sql -dir src/database/migrations -seq ```create_model```
