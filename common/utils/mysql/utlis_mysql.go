package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//初始化连接池
var db *sql.DB

//创建一个结构
type Users struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	//数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	//生出句柄
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	//检测连接是否出现错误
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s invalid ,err'=", err)
	}
	fmt.Println("connect database succsed")
	return nil
}

//查询单个对象
func querySingle(id int) {
	var u1 Users
	sqlStr := "select Id,Name,Age from Users where Id=?"
	err := db.QueryRow(sqlStr, id).Scan(&u1.Id, &u1.Name, &u1.Age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
	}
	fmt.Println("u1.name=", u1.Name)
	fmt.Printf("u1:%v\n", u1)
}

//查询多个对象
func queryList() {
	sqlStr := "select Id,Name,Age from Users "
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
	}
	//关闭rows释放持有的数据库链接
	defer rows.Close()
	for rows.Next() {
		var u Users
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("u1:%v\n", u)
	}
}

//插入对象
func insert() {
	sqlStr := "insert into Users(Id,Name,Age) values(?,?,?)"
	ret, err := db.Exec(sqlStr, 4, "Bosh", 35)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
	}
	//新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed,err:%v\n", err)
	}
	fmt.Printf("insert success ,the insetId=%d ,\n", theID)
}

//更新对象
func update() {
	sqlStr := "update Users set Name=? where id=?"
	ret, err := db.Exec(sqlStr, "bosh_update", 4)
	if err != nil {
		fmt.Printf("update failed err=%v \n", err)
		return
	}
	//受影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected failed err=%v \n", err)
		return
	}
	fmt.Printf("update success ,affected=%d\n", n)
}

//删除对象
func delete() {
	sqlStr := "delete from Users where id=?"
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete failed err=%v \n", err)
		return
	}
	//受影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected failed err=%v \n", err)
		return
	}
	fmt.Printf("delete success ,affected=%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("初始化连接失败 err=%v\n", err)
	}
	querySingle(2)
	insert()
	update()
	delete()
	queryList()
}
