#! /bin/bash

BASE_DIR=$(cd -P $(dirname $0);pwd)
 

delete_go_app() {
    for file in `ls $1`
    do
        if [ "$file" = "vendor" ]
        then 
            return
        fi
        if [ -d $1"/"$file ]  #注意此处之间一定要加上空格，否则会报错
        then
            delete_go_app $1"/"$file
        else
            if [ "$file" = "main.go" ] 
            then
                cd $1
                project_name="${1##*/}"
                echo "remove file $1/$project_name"
                rm -f $project_name
                cd $BASE_DIR
                sleep 0.1
            fi            
        fi
    done
}

#测试目录 test
delete_go_app .

