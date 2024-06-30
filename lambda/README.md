build 
```
GOOS=linux GOARCH=amd64 go build -o main main.go
```

deploy local stack
```bash
aws --endpoint-url=http://localhost:4566 lambda create-function \
    --function-name MyLambdaFunction \
    --runtime provided.al2 \
    --role arn:aws:iam::123456789012:role/lambda-role \
    --handler main \
    --zip-file fileb://main

```

request apigateway 
```
curl -X POST https://localhost:4566/restapis/0000/local/_user_request -d '{"id": "1"}'
```