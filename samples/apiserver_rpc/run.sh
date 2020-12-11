BASE_DIR=$(cd -P $(dirname $0);pwd)
 
cd $BASE_DIR

go build -mod=mod

sleep 1

localip=`/sbin/ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"`
echo 浏览器请求：$localip:50008/api/requestbyctx
echo 浏览器请求：$localip:50008/api/swap
echo 浏览器请求：$localip:50008/api/request
echo 浏览器请求：$localip:50008/api/remoterpc
echo 浏览器请求：$localip:50008/api/remoterpcip

project_name="${BASE_DIR##*/}"

./$project_name run

