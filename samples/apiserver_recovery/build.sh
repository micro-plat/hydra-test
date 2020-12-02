go build

sleep 1

localip=`/sbin/ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"`
echo 浏览器请求：$localip:50010/api/localrpc

./apiserver_recovery run

