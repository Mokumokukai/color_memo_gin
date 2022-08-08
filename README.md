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


--------------------------------------------------------------------


# API 設計詳細

## colors

### 色メモ登録API
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


## tags


