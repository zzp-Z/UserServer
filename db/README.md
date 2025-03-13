
```shell
  goctl sql mysql ddl --src=.\model\user.sql -c -d ./crud/ --idea --style=GoZero
```

```shell
  goctl model mysql datasource --url="root:123123@tcp(localhost:3306)/user" -c -d ./crud/ --idea --style GoZero -t "*"
```