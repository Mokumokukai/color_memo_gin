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

## colors
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| GET | api/v1/colors | 自動生成の色と人気の色組み合わせを返答 | 不要|
| POST | api/v1/color | 色メモの登録 | 必要　|
| POST | api/v1/color/:color_id/edit | 色メモの編集 | 必要 |
| POST | api/v1/color/:color_id/delete | 色メモの削除 | 必要 |
| POST | api/v1/color/:color_id/duplicate | 色メモの複製 | 必要 |
## tags
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| GET | api/v1/colors/:color_id/tags | 色メモのタグ一覧 | 不要 |
| POST | api/v1/colors/:color_id/tags |　色メモのタグ登録 | 必要 |
| POST | api/v1/color/:color_id/tags/:tag_id/edit | タグの編集 | 必須 |
| POST | api/v1/color/:color_id/tags/:tag_id/delete | メモに紐づけられているタグの削除 | 必須 |
| POST | api/v1/tags | タグの作成 | 必須 |
| GET | api/v1/tags | タグ一覧の取得 | 必須 |
| POST | api/v1/tags/:tag_id/edit | タグ名の変更（実装検討中) | 不要 |
| POST | api/v1/tags/:tag_id/delete | タグ名の削除（実装検討中) | 必須 |



## users
| メソッド | URI | リソース | token必須か |
| - | - | - | - |
| GET | api/v1/users | ユーザ一覧 | 不要 |
| GET | api/v1/users/:user_id/colors | ユーザのメモ一覧 | 不要 |
| GET | api/v1/users/:user_id/colors/:color_id |　色メモの詳細 | 不要 |

## Auth
| メソッド | URI | リソース | 
| - | - | - |
| POST | api/v1/auth/signup | サインアップ | 
| POST | api/v1/auth/login |　ログイン |


--------------------------------------------------------------------


# API 設計詳細

## colors
### 色メモ登録API

<details>
<summary> 色メモ登録API</summary>

#### パス

```/api/v1/colors```


#### メソッド
- POST
    - JSON(req,res)

#### リクエスト

##### サンプル

```
{
    "color":{
        "color1":"#FFFFFF",
        "color2":"#000000"
    },
    "tags":{
        ["モノクロ","白黒"]
    }
}
```
tagがない場合
```
{
    "color":{
        "color1":"#FFFFFF",
        "color2":"#000000"
    }
}
```
color1, color2はサーバ側で名前順でソートする。
#### レスポンス
##### 成功時
 - ステータスコード　200

#### サンプル

```
{
    "color":{
        "color1":"#000000",
        "color2":"#FFFFFF",
        "author_id":"01GA0Y1R6F5HJ6AJVH2RWT863G",
        "craeter_id":"01GA0Y1R6F5HJ6AJVH2RWT863G",
        "id":"01GA0Y1WN23BYR1N4EDZNYBDH5",
        "created_at":"2022-07-14T02:40:00Z",
        "updated_at":"2022-07-14T02:40:00Z":
    },
    "tags":{
        [
            {"モノクロ":"01GA0Y2BBYCJFYY92T237Y17SX"},
            {"白黒":"01GA0Y21VP47YFR89VX5MKNYVJ"}
        ]
    }
}
```


| パラメータ | 型 | 内容 |
| - | - | - |
| color | object | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1 | string | 順番が若い方の色 |
| color2 | string | 順番が若くない色 |
| author_id | string |　所持者id |
| creater_id | string | 作成者id |
| created_at | string | 作成日時 |
| updated_at | string | 更新日時 |
| tags | array_object | 登録したタグ情報の配列 |

#### 失敗時
##### サンプル
```
{
    "err":"すでにその色の組み合わせは登録されています。"
}
```
#### 注意点


</details>

### 色メモ編集API
<details>
<summary>色メモ取得一覧API</summary>

#### 説明
ログインしているユーザ自身の

#### パス

```/api/v1/colors```


#### メソッド
- POST
    - JSON(req,res)

#### リクエスト

##### サンプル

```
{
    "color":{
        "color1":"#FFFFFF",
        "color2":"#000000"
    },
    "tags":{
        ["モノクロ","白黒"]
    }
}
```
tagがない場合
```
{
    "color":{
        "color1":"#FFFFFF",
        "color2":"#000000"
    }
}
```
color1, color2はサーバ側で名前順でソートする。
### レスポンス
##### 成功時
 - ステータスコード　200

#### サンプル

```
{
    "color":{
        "color1":"#000000",
        "color2":"#FFFFFF",
        "author_id":"01JS2Z3NDEKCDS4RRFFQ69G5FAV",
        "craeter_id":"01JS2Z3NDEKCDS4RRFFQ69G5FAV",
        "id":"23SS2Z3NDEKCDS4RRFFQ69G5FAV",
        "created_at":"2022-07-14T02:40:00Z",
        "updated_at":"2022-07-14T02:40:00Z":
    },
    "tags":{
        [
            {"モノクロ":"01ARZ3NDEKTSV4RRFFQ69G5FAV"},
            {"白黒":"02ARZ3NDBDTSV4RRFFQ69G5FAX"}
        ]
    }
}
```


| パラメータ | 型 | 内容 |
| - | - | - |
| color | object | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1 | string | 順番が若い方の色 |
| color2 | string | 順番が若くない色 |
| author_id | string |　所持者id |
| creater_id | string | 作成者id |
| created_at | string | 作成日時 |
| updated_at | string | 更新日時 |
| tags | array_object | 登録したタグ情報の配列 |

#### 失敗時
##### サンプル
```
{
    "err":"すでにその色の組み合わせは登録されています。"
}
```
#### 注意点
色は名前順でソートされて返されます。
</details>

----------------------

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

