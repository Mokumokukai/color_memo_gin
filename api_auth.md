# Auth
| メソッド | URI | リソース | 
| - | - | - |
| POST | api/v1/auth/signup | サインアップ | 
| POST | api/v1/auth/login |　ログイン |


## api/v1/auth/signup
<details>
<summary>サインアップ</summary>


### メソッド
- POST
    - JSON(req,res)

```
{
    "user":{
        "name":"name",
        "password":"password",
        "email":"test@test.com"
    }
}
```

### レスポンス
#### 成功時
 - ステータスコード　201 Created

### サンプル
```
{
    "user":{
        "name":"name",
        "email":"test@test.com",
        "id":"01GA3TBQ6F1XP8Y75T274XT4FQ",
        "create_at":"2022-07-14T02:40:00Z",
        "updated_at":"2022-07-14T02:40:00Z"
    }
}
```

### 失敗時
 - ステータスコード　401 Unauthorized
#### サンプル
```
{
    "err":"そのユーザ名はすでに存在します。"
}
```
```
{
    "err":"そのメールアドレスはすでに登録されています。"
}
```
```
{
    "err":"そのパスワードは短すぎです。"
}
```

### 注意点
パスワードや名前にバリデーションをつける予定です。

</details>


## api/v1/auth/login

<details>
<summary>ログイン</summary>


### メソッド
- POST
    - JSON(req,res)

### リクエスト
```
{
    "user":{
        "name":"name",
        "password":"password"
    }
}
```

### レスポンス
```
{
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```
[参照](https://jwt.io/#debugger)
#### 成功時
 - ステータスコード　201 created

### サンプル


### 失敗時
 - ステータスコード　404 NotFound
 - ステータスコード　401 Unauthorized
#### サンプル
```
{
    "err":"指定されたユーザは存在しません。"
}
```

```
{
    "err":"ユーザ名もしくはパスワードが間違っています。"
}
```

### 注意点
</details>