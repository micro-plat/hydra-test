echo "go build -mod=mod"
go build -mod=mod

echo "start"
./registry_repeat run &


for k in $( seq 1 10 )
do
    echo "add cron:$k"
    curl http://localhost:8070/cron/add
    sleep 5

    echo "remove cron:$k"
    curl http://localhost:8070/cron/remove
    sleep 5
done