# HelloGo

## Description

Hello Rest API by Go Gin/MySQL.

This API for **managing notes(memos)**.

## How To Use

```zsh
docker-compose up -d
```

```zsh
# Read all memos
curl -X GET localhost:8080/memos

# Read id memos
curl -X GET localhost:8080/memos/:id

# Create memos
curl -X POST localhost:8080/memos -H "Content-Type: application/json" -d '{"Title": "hoge", "Text": "hogehoge"}'

# Update memos
curl -X PUT localhost:8080/memos/:id -H "Content-Type: application/json" -d '{"Title": "fuga", "Text": "fugafuga"}'

# Delete memos
curl -X DELETE localhost:8080/memos/:id
```
