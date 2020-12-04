package sqls

const InsertData = `
insert into  TBL_TEST_TABLE (record_id,record_name)values(10001,@record_name)
`
const Getdata = `
select record_id,record_name from TBL_TEST_TABLE 
`

const Update = `
update TBL_TEST_TABLE t
set t.record_name = @record_name
where t.record_id  = 10001
`

const Delete = `
delete from TBL_TEST_TABLE  where record_id in( 10001 ,10002)
`

const SP = `test_procedure(@record_name)`
