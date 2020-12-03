package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/samples/apiserver_db/apiserver_mysql/sqls"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/vars/db/mysql"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest"),
	hydra.WithSystemName("apiserver_db_oracle"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.Vars().DB().MySQL("0.36", mysql.New("root:rTo0CesHi2018Qx@tcp(192.168.0.36:3306)/test?charset=utf8"))

	hydra.Conf.API(":50021", api.WithTimeout(10, 10)).Header(header.WithHeader("content-type", "application/json"))

	hydraApp.API("/api/mysql/insert", insert) 
	hydraApp.API("/api/mysql/update", update)
	hydraApp.API("/api/mysql/getdata", getdata)
	hydraApp.API("/api/mysql/sp", sp)
	hydraApp.API("/api/mysql/delete", delete)
}

// apiserver_db 数据库组件是否正确工作，修改配置是否自动生效（mysql）
// 1. 编译程序： go build
// 2. 启动程序：./apiserver_mysql run

// 3. 请求 http://localhost:50020/api/mysql/insert 添加一条 10001数据
// 4. 请求 http://localhost:50020/api/mysql/update 更新10001数据
// 5. 请求 http://localhost:50020/api/mysql/getdata 获取数据表中所有数据
// 6. 请求 http://localhost:50020/api/mysql/sp 调用存储过错添加一条 10002数据
// 7. 请求 http://localhost:50020/api/mysql/delete 删除所有数据
func main() {
	hydraApp.OnStarting(func(cnf app.IAPPConf) (err error) {
		oracleDB := hydra.C.DB().GetRegularDB("0.36")

		_, _, _, err = oracleDB.Execute(sqls.DBTableDrop, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableDrop:%v", err))
		}
		
		_, _, _, err = oracleDB.Execute(sqls.DBTableInit1, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableInit1:%v", err))
		}
		  
		_, _, _, err = oracleDB.Execute(sqls.DBSPInit, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBSPInit:%v", err))
 		}
		return nil

	}, http.API)

	hydraApp.Start()
}

var insert = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")

	effCount, _, _, err := oracleDB.Execute(sqls.InsertData, map[string]interface{}{
		"record_name": "mysql" + time.Now().Format("20060102150405"),
	}) 
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}

var getdata = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Query:%v", err)

	}
	return rows
}

var update = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")

	effCount, _, _, err := oracleDB.Execute(sqls.Update, map[string]interface{}{
		"record_name": "mysql" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		return fmt.Errorf("mysql.Execute:%v", err)
	}
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}

var delete = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")
	effCount, _, _, err := oracleDB.Execute(sqls.Delete, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Execute:%v", err)
	}
	return map[string]interface{}{
		"effect_count": effCount,
	}
}

var sp = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")
	effCount, _, err := oracleDB.ExecuteSP(sqls.SP, map[string]interface{}{
		"record_name": "mysql" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		return fmt.Errorf("mysql.ExecuteSP:%v", err)
	}
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}
