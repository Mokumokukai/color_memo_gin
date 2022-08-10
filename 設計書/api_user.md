
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
    "users":{[
        {
            "name":"test1",
            "id":"01GA3TZNS0B3Y4BFB134P6RNH0",
            "created_at":"2022-07-14T02:40:00Z",
            "updated_at":"2022-07-14T02:40:00Z",
            "memo_num":5,
        },
        {
            "name":"test2",
            "id":"01GA3V1FV9FFJHB1KDFJNTPQJ5",
            "created_at":"2022-07-14T02:40:00Z",
            "updated_at":"2022-07-14T02:40:00Z",
            "memo_num":2,
        }
    ]}
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

