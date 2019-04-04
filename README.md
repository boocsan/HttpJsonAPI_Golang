# HttpJsonAPI_Golang

Golang で Http サーバーを実行する Docker コンテナ

```
git clone https://github.com/iedred7584/HttpJsonAPI_Golang

cd HttpJsonAPI_Golang

docker build -t JsonAPI .

// Random Port
docker run -dP --name JsonAPI JsonAPI

// Static Port
docker run -dp 8888:8888 --name JsonAPI JsonAPI
```
