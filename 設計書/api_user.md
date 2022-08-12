
# users
| メソッド | URI                                                    | リソース         | token必須か |
| -------- | ------------------------------------------------------ | ---------------- | ----------- |
| GET      | [api/v1/users](#apiv1users)                            | ユーザ一覧       | 不要        |
| GET      | [api/v1/users/:user_id/memos](#apiv1usersuser_idmemos) | ユーザのメモ一覧 | 不要        |


## api/v1/users
<details>
<summary>ユーザ一覧</summary>


### メソッド
- GET
    - JSON(res)


#### 成功時
 - ステータスコード　200 created


### レスポンス
#### サンプル

```
{
    "users":[
        {
            "name":"test1",
            "id":"uOUDFFD",
            "created_at":"2022-07-14T02:40:00Z",
            "updated_at":"2022-07-14T02:40:00Z",
            "memo_num":5,
        },
        {
            "name":"test2",
            "id":"TpsyDfW",
            "created_at":"2022-07-14T02:40:00Z",
            "updated_at":"2022-07-14T02:40:00Z",
            "memo_num":2,
        }
    ]
}
```




### 失敗時
 - ステータスコード　404 NotFound
#### サンプル
```
{
    "err":"指定されたユーザは存在しません。"
}
```


### 注意点
</details>


## api/v1/users/:user_id/memos


<details>
<summary>ユーザメモ一覧</summary>


### メソッド
- POST
    - JSON(req,res)

### レスポンス
ステータスコード　200 OK
#### サンプル
```
{
    "memos":{
        [
            {
                "color1":"#FFFFFF",
                "color2":"#000000",
                "tags": [
                    {
                        "id": "Qn5AuIK",
                        "name": "モノクロ"
                    }
                ],
                "creater_id":"01GA3NKQ5VZE70FY5MRVPYRABY",
                "owner_id":"01GA3NM3ETBHBP0SCXHB4TGYXB",
                "created_at":"2022-07-14T02:40:00Z",
                "updated_at":"2022-07-14T02:40:00Z",
                "id":"r4B6cWM"
            },

        ]
    }
}
```


### 失敗時
 - ステータスコード　404 NotFound
#### サンプル
```
{
    "err":"指定されたユーザは存在しません。"
}
```


### 注意点
</details>

