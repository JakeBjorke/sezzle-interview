# getting started

Do these in different shells

to run the sever

```
cd backend
go run main.go
```

If you want to do it in docker....

```
cd backend
docker build -t backend . 
docker run -it -p 8080:8080 backend 
```

to run the UI

```
cd frontend
yarn start
```

**note**
code uses comments for the uri for the server, if you want to run locally you have to comment/uncomment in the below files

```
frontent/src/api/index.js
frontend/src/components/ExpressionInput/ExpressionInput.jsx
```