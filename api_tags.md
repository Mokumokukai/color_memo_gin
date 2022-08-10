# tags API

## tags
| メソッド | URI                                                   | リソース                         | token必須か |
| -------- | ----------------------------------------------------- | -------------------------------- | ----------- |
| GET      | [api/v1/tags](#apiv1tagsget)                          | タグ一覧の取得                   | 必須        |
| POST     | [api/v1/tags](#apiv1tagspost)                         | タグの作成                       | 必須        |
| GET      | [api/v1/tags/:tag_id/memos](#apiv1tagstag_idmemosget) | タグに紐づけられているメモの取得 | 必須        |


## api/v1/tags(GET)
<details>
<summary>タグ取得一覧API</summary>

### 説明
登録されているタグの一覧を返す。
### メソッド
- GET
    - JSON(res)


### レスポンス
#### 成功時
 - ステータスコード　200

### サンプル

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
#### 注意点
</details>


## api/v1/tags(POST)
<details>
<summary>タグの新規作成</summary>

### 説明
タグを新規作成する。
### メソッド
- POST
    - JSON(req,res)


### レスポンス
#### 成功時
 - ステータスコード　201 Created

### サンプル

```
{
    "tag":{
        "モノクロ":"01ARZ3NDEKTSV4RRFFQ69G5FAV"
    }
}
```
#### 注意点
</details>

## api/v1/tags/:tag_id/memos(GET)


<details>
<summary>タグに紐づけられているメモの取得</summary>

### 説明
タグを登録しているメモを人気順で返す。
### メソッド
- GET
    - JSON(res)


### レスポンス
#### 成功時
 - ステータスコード　200

### サンプル

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
#### 注意点
</details>