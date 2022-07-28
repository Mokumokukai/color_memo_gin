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
