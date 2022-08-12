# 色メモAPI


## colors
| メソッド | URI                                                               | リソース                               | token必須か |
| -------- | ----------------------------------------------------------------- | -------------------------------------- | ----------- |
| GET      | [api/v1/memos](#apiv1memosget)                                    | 自動生成の色と人気の色組み合わせを返答 | 不要        |
| POST     | [api/v1/memos](#apiv1memos(POST))                                 | 色メモの登録                           | 必要        |
| POST     | [api/v1/memos/:memo_id/edit](#apiv1memosmemoideditpost)           | 色メモの編集                           | 必要        |
| POST     | [api/v1/memos/:memo_id/delete](#apiv1memosmemo_iddeletepost)      | 色メモの削除                           | 必要        |
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
    "memos":[
        {
            "color1":"#000000",
            "color2":"#FFFFFF",
            "owner_id":"-HgND-k",
            "craeter_id":"yJpTc4d",
            "id":"byNWNoK",
            "created_at": "2022-08-11T20:03:04.827Z",
            "updated_at": "2022-08-11T20:03:04.827Z",
            "tags":{
                [
                    {
                        "id": "Qn5AuIK",
                        "name": "モノクロ"
                    }
                    {
                        "id": "1hQpW2c",
                        "name": "モノクロ"
                    }
                ]
            }
        }
    ]
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
        "color1":"#FFFFFF",
        "color2":"#000000",
        "tags":[
            "ifFxSWn","hFYcyr_"
        ]
    }
}
```
tagがない場合
```
{
    "memo":{
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
        "color1":"#000000",
        "color2":"#FFFFFF",
        "owner_id":"-HgND-k",
        "craeter_id":"-HgND-k",
        "id":"byNWNoK",
        "created_at": "2022-08-11T20:03:04.827Z",
        "updated_at": "2022-08-11T20:03:04.827Z",
        "tags":[
            {
                "id": "Qn5AuIK",
                "name": "モノクロ"
            }
            {
                "id": "1hQpW2c",
                "name": "モノクロ"
            }
        ]
    }
}
```


| パラメータ | 型           | 内容                                               |
| ---------- | ------------ | -------------------------------------------------- |
| color      | object       | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1     | string       | 順番が若い方の色                                   |
| color2     | string       | 順番が若くない色                                   |
| owner_id   | string       | 所持者id                                           |
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
    "memo":{
        "color1":"#FFFFFF",
        "color2":"#000000",
        "tags":[
            "ifFxSWn","hFYcyr_"
        ]
    }

}
```
tagがない場合
```
{
    "memo":{
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
        "color1":"#000000",
        "color2":"#FFFFFF",
        "owner_id":"CHte4Lj",
        "craeter_id":"CHte4Lj",
        "id":"bdaqr2D",
        "created_at": "2022-08-11T20:03:04.827Z",
        "updated_at": "2022-08-11T20:03:04.827Z",
        "tags":[
            {
                "id": "Qn5AuIK",
                "name": "モノクロ"
            }
            {
                "id": "1hQpW2c",
                "name": "モノクロ"
            }
        ]
    }
}
```


| パラメータ | 型           | 内容                                               |
| ---------- | ------------ | -------------------------------------------------- |
| color      | object       | ソートされた色の組み合わせと作成者などの情報を格納 |
| color1     | string       | 順番が若い方の色                                   |
| color2     | string       | 順番が若くない色                                   |
| owner_id   | string       | 所持者id                                           |
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
        "color1":"#000000",
        "color2":"#FFFFFF",
        "owner_id":"-HgND-k",
        "craeter_id":"yJpTc4d",
        "id":"byNWNoK",
        "created_at": "2022-08-11T20:03:04.827Z",
        "updated_at": "2022-08-11T20:03:04.827Z",
        "tags":[
            {
                "id": "Qn5AuIK",
                "name": "モノクロ"
            }
            {
                "id": "1hQpW2c",
                "name": "モノクロ"
            }
        ]
    }
}
```


| パラメータ | 型           | 内容                   |
| ---------- | ------------ | ---------------------- |
| color1     | string       | 順番が若い方の色       |
| color2     | string       | 順番が若くない色       |
| owner_id   | string       | 所持者id               |
| creater_id | string       | 作成者id               |
| created_at | string       | 作成日時               |
| updated_at | string       | 更新日時               |
| tags       | array_object | 登録したタグ情報の配列 |

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
