package sqls

const DBTableDrop = `
drop table if exists TBL_TEST_TABLE
`


const DBTableInit1 = `
CREATE TABLE TBL_TEST_TABLE (
  RECORD_ID int(11) NOT NULL,
  RECORD_NAME varchar(32) DEFAULT NULL,
  PRIMARY KEY (RECORD_ID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表格'

`
 

const DBSPInit = `
CREATE DEFINER="root"@"192.%" PROCEDURE test_procedure(v_record_name varchar(32))
BEGIN
	delete from TBL_TEST_TABLE where record_id = 10002;
	insert into TBL_TEST_TABLE(record_id,record_name)values(10002,v_record_name);  
END
`
