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
