
cd ./bin
# go build ../proxy-a
# go build ../proxy-b

./proxy-a conf install
./proxy-b conf install
./proxy-a run &
./proxy-b run



