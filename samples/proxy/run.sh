
cd ./bin

# rm -rf ./proxy-a
# rm -rf ./proxy-b

go build ../proxy-a
go build ../proxy-b



./proxy-a conf install -v
./proxy-b conf install -v
./proxy-a run &
./proxy-b run





