#! /bin/bash

BASE_DIR=$(cd -P $(dirname $0);pwd)
 

build_go_app() {
    for file in `ls $1`
    do
        if [ "$file" = "vendor" ]
        then 
            continue
        fi
        if [ -d $1"/"$file ]  #注意此处之间一定要加上空格，否则会报错
        then
            build_go_app $1"/"$file
        else
            if [ "$file" = "main.go" ] 
            then
                cd $1
                echo "go build $1"
                go build -mod=mod
                cd $BASE_DIR
                sleep 0.1
            fi            
        fi
    done
}

#测试目录 test
build_go_app .


