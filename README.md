# HttpJsonAPI_Golang

Golang で Http サーバーを実行する Docker コンテナ

```
git clone https://github.com/iedred7584/HttpJsonAPI_Golang

cd HttpJsonAPI_Golang

docker build -t jsonapi .

// Random Port
docker run -dP --name jsonapi jsonapi

// Static Port
docker run -dp 8888:8888 --name jsonapi jsonapi
```
