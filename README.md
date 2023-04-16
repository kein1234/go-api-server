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
curl -X GET "http://localhost:8080/api/v1/users/"
curl -X GET "http://localhost:8080/api/v1/users/1"

# POST
curl -X POST "http://localhost:8080/api/v1/users/" -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john.doe@example.com"}'

# PUT
curl -X PUT "http://localhost:8080/api/v1/users/1" -H "Content-Type: application/json" -d '{"name": "John Doe Updated", "email": "john.doe.updated@example.com"}'

# DELETE
curl -X DELETE "http://localhost:8080/api/v1/users/1"


```