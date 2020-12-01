
pwd

BASE_DIR=$(cd -P $(dirname $0);pwd)
echo $BASE_DIR

cd $BASE_DIR

go build

 
sleep 1 

unzip -o dist.zip -d ./src

sleep 1
 
./webserver_static run