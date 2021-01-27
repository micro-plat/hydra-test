package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/samples/apiserver_db/apiserver_mysql/sqls"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var hydraApp = hydra.NewApp(
	hydra.WithServerTypes(http.API),
	hydra.WithPlatName("hydratest", "Hydra样例"),
	hydra.WithSystemName("apiserver_db_mysql", "Mysql数据库"),
	hydra.WithClusterName("test"),
	hydra.WithRegistry("lm://."),
)

func init() {
	hydra.Conf.Vars().DB().MySQLByConnStr("0.36", "test:123456@tcp(192.168.0.36:3306)/test?charset=utf8")

	hydra.Conf.API(":50022", api.WithTimeout(10, 10)).Header(header.WithHeader("content-type", "application/json"))

	hydraApp.API("/api/mysql/insert", insert)
	hydraApp.API("/api/mysql/update", update)
	hydraApp.API("/api/mysql/getdata", getdata)
	hydraApp.API("/api/mysql/sp", sp)
	hydraApp.API("/api/mysql/delete", delete)

	hydraApp.API("/api/mysql/config", config)
}

// apiserver_db 数据库组件是否正确工作，修改配置是否自动生效（mysql）
// 1. 编译程序： go build
// 2. 启动程序：./apiserver_mysql run

// 3. 请求 http://localhost:50022/api/mysql/insert 添加一条 10001数据
// 4. 请求 http://localhost:50022/api/mysql/update 更新10001数据
// 5. 请求 http://localhost:50022/api/mysql/getdata 获取数据表中所有数据
// 6. 请求 http://localhost:50022/api/mysql/sp 调用存储过错添加一条 10002数据
// 7. 请求 http://localhost:50022/api/mysql/delete 删除所有数据
// 8. 请求 http://localhost:50022/api/mysql/config 修改配置数据库 test==>test2

func main() {

	hydraApp.OnStarting(func(cnf app.IAPPConf) (err error) {
		oracleDB := hydra.C.DB().GetRegularDB("0.36")

		_, err = oracleDB.Execute(sqls.DBTableDrop, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableDrop:%v", err))
		}

		_,  err = oracleDB.Execute(sqls.DBTableInit1, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableInit1:%v", err))
		}

		_,  err = oracleDB.Execute(sqls.DBSPInit, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBSPInit:%v", err))
		}
		return nil

	}, http.API)

	hydraApp.Start()
}

var insert = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")

	effCount,  err := oracleDB.Execute(sqls.InsertData, map[string]interface{}{
		"record_name": "mysql" + time.Now().Format("20060102150405"),
	})
	rows,  err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
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
	rows,  err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Query:%v", err)

	}
	return rows
}

var update = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")

	effCount,  err := oracleDB.Execute(sqls.Update, map[string]interface{}{
		"record_name": "mysql" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		return fmt.Errorf("mysql.Execute:%v", err)
	}
	rows,  err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
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
	effCount,  err := oracleDB.Execute(sqls.Delete, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Execute:%v", err)
	}
	return map[string]interface{}{
		"effect_count": effCount,
	}
}

var sp = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.36")
	effCount,  err := oracleDB.ExecuteSP(sqls.SP, map[string]interface{}{
		"record_name": "mysql" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		return fmt.Errorf("mysql.ExecuteSP:%v", err)
	}
	rows,  err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("mysql.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}
var config = func(ctx hydra.IContext) (r interface{}) {
	regst, err := registry.GetRegistry(global.Def.RegistryAddr, global.Def.Log())
	if err != nil {
		return fmt.Errorf("NewRegistry:%v", err)
	}
	dbpath := "/hydratest/var/db/0.36"
	err = regst.Update(dbpath, `{"provider":"mysql","connString":"test2:123456@tcp(192.168.0.36:3306)/test2?charset=utf8","maxOpen":10,"maxIdle":3,"lifeTime":600}`)
	if err != nil {
		return fmt.Errorf("UpdateDB:%v", err)
	}
	path := "/hydratest/apiserver_db_mysql/api/test/conf"
	err = regst.Update(path, `{"status":"start","address":":50021"}`)
	if err != nil {
		return fmt.Errorf("UpdateConf:%v", err)
	}
	return "success"
}
