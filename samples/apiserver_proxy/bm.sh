

while true
do
ab -c 100 -n 1000 http://192.168.4.121:8091/request
sleep 1

done