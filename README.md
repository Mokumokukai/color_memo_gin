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

# API 設計

## color
| メソッド | URI | リソース |
| - | - | - |
| GET | api/v1/colors | 自動生成の色と人気の色組み合わせを返答 |
| POST | api/v1/color | 色メモの登録 |
| POST | api/v1/color/:id/edit | 色メモの編集 |
| POST | api/v1/color/:id/delete | 色メモの削除 |
| POST | api/v1/color/:id/duplicate | 色メモの複製 |
## tags
| メソッド | URI | リソース |
| - | - | - |
| GET | api/v1/colors/:color_id/tags | タグ一覧 |
| POST | api/v1/colors/:color_id/tags |　タグ登録 |
| POST | api/v1/color/:color_id/tags/:tag_id/edit | タグの編集 |
| POST | api/v1/color/:color_id/tags/:tag_id/delete | タグの削除 |


## Auth
| メソッド | URI | リソース |
| - | - | - |
| POST | api/v1/auth/signup | サインアップ |
| POST | api/v1/auth/login |　ログイン |