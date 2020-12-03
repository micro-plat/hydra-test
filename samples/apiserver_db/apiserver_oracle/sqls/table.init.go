package sqls

const DBTableDrop = `
 drop table TBL_TEST_TABLE
`


const DBTableInit1 = `
create table TBL_TEST_TABLE
(
  RECORD_ID   NUMBER(20) not null,
  RECORD_NAME VARCHAR2(32)
)
tablespace USERS
  pctfree 10
  initrans 1
  maxtrans 255
  storage
  (
    initial 64
    minextents 1
    maxextents unlimited
  )
`
const DBTableInit2 = `
comment on table TBL_TEST_TABLE
  is 'test table'`

const DBTableInit3 = `
alter table TBL_TEST_TABLE
  add constraint PK_TEST_ID primary key (RECORD_ID)
  using index 
  tablespace USERS
  pctfree 10
  initrans 2
  maxtrans 255
`

const DBSPInit = `
create or replace procedure test_procedure(v_record_name varchar2) is
begin

  delete tbl_test_table t where t.record_id = 10002;
  insert into tbl_test_table
    (record_id, record_name)
  values
    (10002, v_record_name);

end;
`
