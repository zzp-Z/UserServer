
```shell
goctl model mysql ddl --src=.\model\user.sql -c -d ./crud/ --idea --style=GoZero
```

```shell
goctl model mysql datasource --url="root:123456@tcp(AiCodeStudio.top:3306)/user_server" -c -d ./crud/ --idea --style GoZero -t "*"
```