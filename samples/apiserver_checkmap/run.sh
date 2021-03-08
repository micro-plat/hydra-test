
name=$(basename $(pwd))

echo $name 

if [ ! -d "./bin" ] ; then    
    mkdir -p bin
fi

echo "编译：$name"
go build -mod=mod


rm -rf ./bin/$name
mv $name ./bin/

cd bin 

echo "启动：$name"
./$name run





