mkdir bin
sleep 0.1

cd ./bin

echo "go build -mod=mod ../http_client"
go build -mod=mod ../http_client

echo "go build -mod=mod ../http_client"
go build -mod=mod ../http_server

echo "server run"
./http_client run & 
./http_server run 