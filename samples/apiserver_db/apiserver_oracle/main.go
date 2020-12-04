package main

import (
	"fmt"
	"time"

	_ "github.com/mattn/go-oci8"
	"gopkg.in/yaml.v2"

	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/registry"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra-test/samples/apiserver_db/apiserver_oracle/sqls"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/vars/db/oracle"
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
	hydra.Conf.Vars().DB().MySQL("0.136", oracle.New("test/123456@orcl136"))

	hydra.Conf.API(":50021", api.WithTimeout(10, 10)).Header(header.WithHeader("content-type", "application/json"))

	hydraApp.API("/api/oracle/insert", insert)
	hydraApp.API("/api/oracle/update", update)
	hydraApp.API("/api/oracle/getdata", getdata)
	hydraApp.API("/api/oracle/sp", sp)
	hydraApp.API("/api/oracle/delete", delete)

	hydraApp.API("/api/oracle/config", config)

}

// apiserver_db 数据库组件是否正确工作，修改配置是否自动生效（mysql）
// 1. 编译程序： go build
// 2. 启动程序：./apiserver_mysql run

// 3. 请求 http://localhost:50021/api/oracle/insert 添加一条 10001数据
// 4. 请求 http://localhost:50021/api/oracle/update 更新10001数据
// 5. 请求 http://localhost:50021/api/oracle/getdata 获取数据表中所有数据
// 6. 请求 http://localhost:50021/api/oracle/sp 调用存储过错添加一条 10002数据
// 7. 请求 http://localhost:50021/api/oracle/delete 删除所有数据
// 8. 请求 http://localhost:50021/api/oracle/config 修改数据库配置，并重启
func main() {

	datarows := []types.XMap{
		types.XMap{
			"a1": "a1",
			"a2": "a2",
		},
		types.XMap{
			"b1": "b1",
			"b2": "b2",
		},
	}

	bytes, err := yaml.Marshal(datarows)
	if err != nil {
		fmt.Println("marshal:", err)
		return
	}
	fmt.Println(string(bytes))

	hydraApp.OnStarting(func(cnf app.IAPPConf) (err error) {
		oracleDB := hydra.C.DB().GetRegularDB("0.136")

		_, _, _, err = oracleDB.Execute(sqls.DBTableDrop, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableDrop:%v", err))
		}

		_, _, _, err = oracleDB.Execute(sqls.DBTableInit1, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableInit1:%v", err))
		}

		_, _, _, err = oracleDB.Execute(sqls.DBTableInit2, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableInit2:%v", err))
		}
		_, _, _, err = oracleDB.Execute(sqls.DBTableInit3, nil)
		if err != nil {
			fmt.Println(fmt.Errorf("DBTableInit3:%v", err))
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
	oracleDB := hydra.C.DB().GetRegularDB("0.136")

	effCount, _, _, err := oracleDB.Execute(sqls.InsertData, map[string]interface{}{
		"record_name": "Oracle" + time.Now().Format("20060102150405"),
	})
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Oracle.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}

var getdata = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.136")
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Oracle.Query:%v", err)

	}
	return rows
}

var update = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.136")

	effCount, _, _, err := oracleDB.Execute(sqls.Update, map[string]interface{}{
		"record_name": "Oracle" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		return fmt.Errorf("Oracle.Execute:%v", err)
	}
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Oracle.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}

var delete = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.136")
	effCount, _, _, err := oracleDB.Execute(sqls.Delete, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Oracle.Execute:%v", err)
	}
	return map[string]interface{}{
		"effect_count": effCount,
	}
}

var sp = func(ctx hydra.IContext) (r interface{}) {
	oracleDB := hydra.C.DB().GetRegularDB("0.136")
	effCount, _, err := oracleDB.ExecuteSP(sqls.SP, map[string]interface{}{
		"record_name": "Oracle" + time.Now().Format("20060102150405"),
	})
	if err != nil {
		return fmt.Errorf("Oracle.ExecuteSP:%v", err)
	}
	rows, _, _, err := oracleDB.Query(sqls.Getdata, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("Oracle.Query:%v", err)

	}
	return map[string]interface{}{
		"effect_count": effCount,
		"data_rows":    rows,
	}
}

var config = func(ctx hydra.IContext) (r interface{}) {
	regst, err := registry.NewRegistry(global.Def.RegistryAddr, global.Def.Log())
	if err != nil {
		return fmt.Errorf("NewRegistry:%v", err)
	}
	dbpath := "/hydratest/var/db/0.136"
	err = regst.Update(dbpath, `{"provider":"oracle","connString":"test/123456@orcl136","maxOpen":10,"maxIdle":3,"lifeTime":600}`)
	if err != nil {
		return fmt.Errorf("UpdateDB:%v", err)
	}
	path := "/hydratest/apiserver_db_oracle/api/test/conf"
	err = regst.Update(path, `{"status":"start","address":":50021"}`)
	if err != nil {
		return fmt.Errorf("UpdateConf:%v", err)
	}
	return "success"
}
