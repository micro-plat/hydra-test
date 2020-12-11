

while true
do
ab -c 100 -n 1000 http://localhost:8091/request
sleep 1

done