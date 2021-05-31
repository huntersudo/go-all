

http://xorm.topgoer.com/

### 1. xorm
   xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。xorm的目标并不是让你完全不去学习SQL，我们认为SQL并不会为ORM所替代，但是ORM将可以解决绝大部分的简单SQL需求。xorm支持两种风格的混用。

#### 1.1. 特性
- 支持Struct和数据库表之间的灵活映射，并支持自动同步
- 事务支持
- 同时支持原始SQL语句和ORM操作的混合执行
- 使用连写来简化调用
- 支持使用Id, In, Where, Limit, Join, Having, Table, SQL, Cols等函数和结构体等方式作为条件
- 支持级联加载Struct
- Schema支持（仅Postgres）
- 支持缓存
- 支持根据数据库自动生成xorm的结构体
- 支持记录版本（即乐观锁）
- 内置SQL Builder支持
- 通过EngineGroup支持读写分离和负载均衡

#### 1.2. 驱动支持
xorm当前支持的驱动和数据库如下：
- Mysql: github.com/go-sql-driver/mysql
- MyMysql: github.com/ziutek/mymysql/godrv
- Postgres: github.com/lib/pq
- Tidb: github.com/pingcap/tidb
- SQLite: github.com/mattn/go-sqlite3
- MsSql: github.com/denisenkom/go-mssqldb
- MsSql: github.com/lunny/godbc
- Oracle: github.com/mattn/go-oci8 (试验性支持)
- ql: github.com/cznic/ql (试验性支持)

####  1.3. 安装
go get xorm.io/xorm
#### 1.4. 文档
GoWalker代码文档

Godoc代码文档