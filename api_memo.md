# 色メモAPI


## colors
| メソッド | URI                                                                   | リソース                               | token必須か |
| -------- | --------------------------------------------------------------------- | -------------------------------------- | ----------- |
| GET      | [api/v1/memos](#apiv1memosget)                                      | 自動生成の色と人気の色組み合わせを返答 | 不要        |
| POST     | [api/v1/memos](#apiv1memos(POST))                                   | 色メモの登録                           | 必要        |
| POST     | [api/v1/memos/:memo_id/edit](#apiv1memosmemoideditpost)           | 色メモの編集                           | 必要        |
| POST     | [api/v1/memos/:memo_id/delete](#apiv1memosmemo_iddeletepost)                                      | 色メモの削除                           | 必要        |
| POST     | [api/v1/memos/:memo_id/duplicate](#apiv1memosmemoidduplicatepost) | 色メモの複製                           | 必要        |

## api/v1/memos(GET)
<details>
<summary>色メモ取得API</summary>

### 説明
ユーザ全ての色メモを人気順で取得を行う。
### パス
```/api/v1/memos```
### メソッド
- GET
  - JSON(res)
###　レスポンス
```
{
    "memos":{
        [
            {
                "colors":{
                    "color1":"#FFFFFF",
                    "color2":"#000000"
                },
                "tags":{
                    [{
                        "モノクロ":"01GA3NHKCNQHE9WJTNR70WF2MX"
                    },{
                        "白黒":"01GA3NJKZXRH4CR7TT28HKS8XN"
                    }]
                },
                "creater_id":"01GA3NKQ5VZE70FY5MRVPYRABY",
                "owner_id":"01GA3NM3ETBHBP0SCXHB4TGYXB",
                "created_at":"2022-07-14T02:40:00Z",
                "updated_at":"2022-07-14T02:40:00Z",
                "id":"01GA3NQ0B4NM2Z6HPW2K9R1DF6"
            },

        ]
    }
}
```

</details>

## api/v1/memos(POST)

<details>
<summary> 色メモ登録API</summary>

### パス

```/api/v1/memos```


### メソッド
- POST
    - JSON(req,res)

### リクエスト

#### サンプル

```
{
    "memo":{
        "colors":{
            "color1":"#FFFFFF",
            "color2":"#000000"
        },
        "tags":{
            ["モノクロ","白黒"]
        }
    }
}
```
tagがない場合
```
{
    "colors":{
        "color1":"#FFFFFF",
        "color2":"#000000"
    }
}
```
color1, color2はサーバ側で名前順でソートする。
### レスポンス
#### 成功時
 - ステータスコード　200

### サンプル

```
{
    "memo":{
        "color":{
            "color1":"#000000",
            "color2":"#FFFFFF"
        },
        "owner_id":"01GA0Y1R6F5HJ6AJVH2RWT863G",
        "craeter_id":"01GA0Y1R6F5HJ6AJVH2RWT863G",
        "id":"01GA0Y1WN23BYR1N4EDZNYBDH5",
        "created_at":"2022-07-14T02:40:00Z",
        "updated_at":"2022-07-14T02:40:00Z",
        "tags":{
            [
                {"モノクロ":"01GA0Y2BBYCJFYY92T237Y17SX"},
                {"白黒":"01GA0Y21VP47YFR89VX5MKNYVJ"}
            ]
        }
    }
}
```


| パラメータ | 型           | 内容                                               |
| ---------- | ------------ | -------------------------------------------------- |
| color      | object       | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1     | string       | 順番が若い方の色                                   |
| color2     | string       | 順番が若くない色                                   |
| owner_id  | string       | 所持者id                                           |
| creater_id | string       | 作成者id                                           |
| created_at | string       | 作成日時                                           |
| updated_at | string       | 更新日時                                           |
| tags       | array_object | 登録したタグ情報の配列                             |

### 失敗時
#### サンプル
```
{
    "err":"すでにその色の組み合わせは登録されています。"
}
```
### 注意点


</details>

## api/v1/memos/:memo_id/edit(POST)
<details>
<summary>色メモ編集API</summary>

### 説明
ログインしているユーザ自身の

### パス

```/api/v1/memos```


### メソッド
- POST
    - JSON(req,res)

### リクエスト

#### サンプル

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
#### 成功時
 - ステータスコード　200

### サンプル

```
{
    "memo":{
        "colors":{
            "color1":"#000000",
            "color2":"#FFFFFF"
        },
        "owner_id":"01JS2Z3NDEKCDS4RRFFQ69G5FAV",
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


| パラメータ | 型           | 内容                                               |
| ---------- | ------------ | -------------------------------------------------- |
| color      | object       | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1     | string       | 順番が若い方の色                                   |
| color2     | string       | 順番が若くない色                                   |
| owner_id  | string       | 所持者id                                           |
| creater_id | string       | 作成者id                                           |
| created_at | string       | 作成日時                                           |
| updated_at | string       | 更新日時                                           |
| tags       | array_object | 登録したタグ情報の配列                             |

### 失敗時
#### サンプル
```
{
    "err":"すでにその色の組み合わせは登録されています。"
}
```
### 注意点
色は名前順でソートされて返されます。



</details>


## api/v1/memos/:memo_id/delete(POST)
<details>
<summary>色メモ削除API</summary>

### 説明
ログインしているユーザ自身の

### パス

```api/v1/memos/:memo_id/delete```


### メソッド
- POST
    - JSON(res)

### リクエスト
なし

### レスポンス
#### 成功時
 - ステータスコード　204 

### サンプル

なし



### 失敗時
#### サンプル
```
{
    "err":"その色メモは存在しません。"
}
```
### 注意点



</details>



## api/v1/memos/:memo_id/duplicate(POST)
<details>

<summary> 色メモ複製API</summary>
### パス

```api/v1/memos/:memo_id/duplicate```


### メソッド
- POST
    - JSON(req,res)

### リクエスト
なし

### レスポンス
#### 成功時
 - ステータスコード　201

### サンプル

```
{
    "memo":{
        "colors":{
            "color1":"#000000",
            "color2":"#FFFFFF"
        },
        "owner_id":"01GA0Y1R6F5HJ6AJVH2RWT863G",
        "craeter_id":"01GA0Y1R6F5HJ6AJVH2RWT863G",
        "id":"01GA0Y1WN23BYR1N4EDZNYBDH5",
        "created_at":"2022-07-14T02:40:00Z",
        "updated_at":"2022-07-14T02:40:00Z",
        "tags":{
            [
                {"モノクロ":"01GA0Y2BBYCJFYY92T237Y17SX"},
                {"白黒":"01GA0Y21VP47YFR89VX5MKNYVJ"}
            ]
        }
    }
}
```


| パラメータ | 型           | 内容                                               |
| ---------- | ------------ | -------------------------------------------------- |
| color      | object       | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1     | string       | 順番が若い方の色                                   |
| color2     | string       | 順番が若くない色                                   |
| owner_id  | string       | 所持者id                                           |
| creater_id | string       | 作成者id                                           |
| created_at | string       | 作成日時                                           |
| updated_at | string       | 更新日時                                           |
| tags       | array_object | 登録したタグ情報の配列                             |

### 失敗時
- status code 409 (Conflict)
#### サンプル
```
{
    "err":"すでにそのメモは複製されています。"
}
```
### 注意点
複製だけできる。
新規タグの追加などはできない（考えていない）
created_atは複製時点の時刻

</details>
