echo "go build -mod=mod"
go build -mod=mod

echo "start"
./task_conf run &

sleep 5

for k in $( seq 1 10 )
do
    echo "add cron:$k"
    curl http://localhost:8070/cron/add
    sleep 15

    echo "remove cron:$k"
    curl http://localhost:8070/cron/delete
    sleep 15

    echo "add2 cron:$k"
    curl http://localhost:8070/cron/add
    sleep 15

    echo "update cron:$k"
    curl http://localhost:8070/cron/update
    sleep 15
done