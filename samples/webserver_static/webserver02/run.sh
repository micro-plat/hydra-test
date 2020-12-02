
BASE_DIR=$(cd -P $(dirname $0);pwd)
 
cd $BASE_DIR

go build
 
sleep 1 

#unzip -o dist.zip -d ./src

sleep 1
 
project_name="${BASE_DIR##*/}"

./$project_name run