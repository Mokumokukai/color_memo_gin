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

# API 設計一覧
## 注意点
未定なもの
 - バリデーション
 - urlクエリ設定
 - オプション
 - エラーについて
   - エラーを複数返すか、一番最初に出たエラーのみを返すか考え中
     - 今のところはエラーは一つのみ返す。
## memos
| メソッド | URI                             | リソース                               | token必須か |
| -------- | ------------------------------- | -------------------------------------- | ----------- |
| GET      | api/v1/memos                    | 自動生成の色と人気の色組み合わせを返答 | 不要        |
| POST     | api/v1/memos                    | 色メモの登録                           | 必要        |
| POST     | api/v1/memos/:memo_id/edit      | 色メモの編集                           | 必要        |
| POST     | api/v1/memos/:memo_id/delete    | 色メモの削除                           | 必要        |
| POST     | api/v1/memos/:memo_id/duplicate | 色メモの複製                           | 必要        |



[色メモ関連](./api_memos.md)

## memo_tags
| メソッド | URI                                       | リソース                         | token必須か |
| -------- | ----------------------------------------- | -------------------------------- | ----------- |
| POST     | api/v1/memos/:memo_id/tags                | 色メモにタグ追加                 | 必要        |
| POST     | api/v1/memos/:memo_id/tags/:tag_id/delete | メモに紐づけられているタグの削除 | 必須        |



[メモタグ関連](./api_memo_tag.md)


## tags
| メソッド | URI                       | リソース                         | token必須か |
| -------- | ------------------------- | -------------------------------- | ----------- |
| POST     | api/v1/tags               | タグの作成                       | 必須        |
| GET      | api/v1/tags               | タグ一覧の取得                   | 必須        |
| GET      | api/v1/tags/:tag_id/memos | タグに紐づけられているメモの取得 | 必須        |



[タグ関連](./api_tags.md)

## users
| メソッド | URI                         | リソース         | token必須か |
| -------- | --------------------------- | ---------------- | ----------- |
| GET      | api/v1/users                | ユーザ一覧       | 不要        |
| GET      | api/v1/users/:user_id/memos | ユーザのメモ一覧 | 不要        |


[ユーザ関連](./api_user.md)

## Auth
| メソッド | URI                | リソース     |
| -------- | ------------------ | ------------ |
| POST     | api/v1/auth/signup | サインアップ |
| POST     | api/v1/auth/login  | ログイン     |


[ユーザ認証関連](./api_auth.md)

--------------------------------------------------------------------

