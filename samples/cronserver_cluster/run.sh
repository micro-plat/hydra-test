
cd ./bin

# rm -rf ./proxy-a
# rm -rf ./proxy-b

go build ../
cp ./cronserver_cluster ./cronserver_cluster1
cp ./cronserver_cluster ./cronserver_cluster2

./cronserver_cluster conf install -v
./cronserver_cluster run &
./cronserver_cluster1 run &
./cronserver_cluster2 run 






