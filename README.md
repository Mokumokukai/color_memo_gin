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

## memos
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| GET | api/v1/memos | 自動生成の色と人気の色組み合わせを返答 | 不要|
| POST | api/v1/memos | 色メモの登録 | 必要　|
| POST | api/v1/memos/:memo_id/edit | 色メモの編集 | 必要 |
| POST | api/v1/memos/:memo_id/delete | 色メモの削除 | 必要 |
| POST | api/v1/memos/:memo_id/duplicate | 色メモの複製 | 必要 |


## memo_tags
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| POST | api/v1/memos/:memo_id/tags |　色メモにタグ追加 | 必要 |
| POST | api/v1/memos/:memo_id/tags/:tag_id/delete | メモに紐づけられているタグの削除 | 必須 |


## tags
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| POST | api/v1/tags | タグの作成 | 必須 |
| GET | api/v1/tags | タグ一覧の取得 | 必須 |
| GET | api/v1/tags/:tag_id/memos | タグに紐づけられているメモの取得 | 必須 |
| POST | api/v1/tags/:tag_id/edit | タグ名の変更（実装検討中) | 不要 |
| POST | api/v1/tags/:tag_id/delete | タグ名の削除（実装検討中) | 必須 |


## users
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| GET | api/v1/users | ユーザ一覧 | 不要 |
| GET | api/v1/users/:user_id/memos | ユーザのメモ一覧 | 不要 |
| GET | api/v1/users/:user_id/memos/:memo_id |　色メモの詳細 | 不要 |

## Auth
| メソッド | URI | リソース | 
| - | - | - |
| POST | api/v1/auth/signup | サインアップ | 
| POST | api/v1/auth/login |　ログイン |


--------------------------------------------------------------------


# API 設計詳細
[色メモ関連](./api_memos.md)


## tags


### tag取得API

<details>
<summary>タグ取得一覧API</summary>

#### 説明
登録されているタグの一覧を返す。

#### パス

```/api/v1/tags```


#### メソッド
- GET
    - JSON(res)

#### レスポンス
##### 成功時
 - ステータスコード　200

#### サンプル

```
{
    "tags":{
        [
            {"モノクロ":"01ARZ3NDEKTSV4RRFFQ69G5FAV"},
            {"白黒":"02ARZ3NDBDTSV4RRFFQ69G5FAX"},
            {"ビビッド":"01GA138S72W3GTJPV4DD8G7YX4"}
        ]
    }
}
```


| パラメータ | 型 | 内容 |
| - | - | - |
| tags | array_object | 登録したタグ情報の配列 |

#### 失敗時
##### サンプル
```
{
    "err":"タグが取得できませんでした。"
}
```
#### 注意点
</details>




### tag投稿API

<details>
<summary>タグ取得一覧API</summary>

#### 説明
タグの登録を一つだけ行う。

#### パス

```/api/v1/tags```


#### メソッド
- POST
    - JSON(req,res)
##### サンプル
```
{
    "tag": "モノクロ"
}
```

#### レスポンス
##### 成功時
 - ステータスコード　201 created

#### サンプル

```
{
    "tag":{
        "モノクロ":"01ARZ3NDEKTSV4RRFFQ69G5FAV"
    }
}
```


| パラメータ | 型 | 内容 |
| - | - | - |
| tag | object | 登録したタグ |

#### 失敗時
##### サンプル
```
{
    "err":"すでにそのタグは投稿されています。"
}
```
#### 注意点
</details>

