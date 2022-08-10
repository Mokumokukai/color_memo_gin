# memo_tags
| メソッド | URI                                       | リソース                         | token必須か |
| -------- | ----------------------------------------- | -------------------------------- | ----------- |
| POST     | api/v1/memos/:memo_id/tags                | 色メモにタグ追加                 | 必要        |
| POST     | api/v1/memos/:memo_id/tags/:tag_id/delete | メモに紐づけられているタグの削除 | 必須        |


##  api/v1/memos/:memo_id/tags
<details>
<summary>色メモにタグ追加</summary>

### 説明
色メモにタグを追加
### メソッド
- POST
  - JSON(req,res)

```
{
    "tags":{
        ["モノクロ","白黒"]
    }
}
```


```
{
    "tags":{
        ["モノクロ"]
    }
}
```
### レスポンス
#### 成功時
 - ステータスコード　200

### サンプル

```
{
    "tags":{
        [
            {"モノクロ":"01GA0Y2BBYCJFYY92T237Y17SX"},
            {"白黒":"01GA0Y21VP47YFR89VX5MKNYVJ"}
        ]
    }
}
```

### 失敗時
 - ステータスコード　404 NotFound
 - ステータスコード　401 Unauthorized
#### サンプル
```
{
    "err":"指定されたメモの投稿がありません。"
}
```
### 注意点
登録するタグが一つでも配列に入れてpostする。

</details>

## api/v1/memos/:memo_id/tags/:tag_id/delete


<details>
<summary>色メモからタグ削除</summary>


### メソッド
- POST
    - JSON(res)



### レスポンス
#### 成功時
 - ステータスコード　204

### サンプル


### 失敗時
 - ステータスコード　404 NotFound
 - ステータスコード　401 Unauthorized
#### サンプル
```
{
    "err":"指定されたメモの投稿がありません。"
}
```

```
{
    "err":"指定されたタグは存在しません。"
}
```

### 注意点
</details>