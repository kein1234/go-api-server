# golangでAPIサーバを立ち上げるサンプル

## 手順

1. ".env"ファイル作成
```
DB_HOST=db
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=sample_db
```
2. docker-compose実行

```bash
$ docker-compose up --build
```


3. 確認方法

```bash
# GET
curl http://localhost:8080/users  

# POST
curl -X POST -H "Content-Type: application/json" -d '{"name": "Tom Smith", "email": "tom.smith@example.com"}' http://localhost:8080/users

# PUT
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Tom Smith Updated", "email": "tom.smith.updated@example.com"}' http://localhost:8080/users/1

# DELETE
curl -X DELETE http://localhost:8080/users/1


```