## 各ディレクトリの役割

```
project/
|-- cmd/
|   |-- server/
|       |-- main.go
|-- internal/
|   |-- entity/
|       |-- user.go
|   |-- usecase/
|       |-- user_usecase.go
|   |-- repository/
|       |-- user_repository.go
|-- infrastructure/
|   |-- grpc/
|       |-- handler.go
|-- proto/
|   |-- user.proto // gitsubmodule
|-- go.mod
|-- go.sum
```

git submodule add https://github.com/your-org/proto-repo proto


go mod init project
go get google.golang.org/grpc
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go

go install google.golang.org/grpc
go install github.com/golang/protobuf/proto
go install github.com/golang/protobuf/protoc-gen-go
apt-get install -y protobuf-compiler

protoc --go_out=plugins=grpc:. user.proto

