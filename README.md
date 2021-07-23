### API
```
註冊會員
host：/api/v1/member-register
params：{
    account：
    password：
}
return 
```

### project 架構
- app
    - controllers
        邏輯層(定義路由操作)
    - models
        資料層(定義資料表)
- config
    路由設定和資料庫設定
- db
    資料庫相關操作
- utils
    log，或者共通的工具