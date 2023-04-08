package mysql

import (
	"GinAndSqlx/models"
	_ "GinAndSqlx/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func getVosByUser(users []*models.User) (userVos []*models.UserVo) {
	for i := range users {
		userVo := &models.UserVo{
			Id:         users[i].Id,
			Name:       users[i].Name.String,
			Age:        users[i].Age.Int64,
			Sex:        users[i].Sex.String,
			Address:    users[i].Address.String,
			Phone:      users[i].Phone.String,
			CreateTime: users[i].CreateTime.String,
		}
		userVos = append(userVos, userVo)
	}
	return
}

func getVoByUser(user *models.User) (userVo *models.UserVo) {
	return &models.UserVo{
		Id:         user.Id,
		Name:       user.Name.String,
		Age:        user.Age.Int64,
		Sex:        user.Sex.String,
		Address:    user.Address.String,
		Phone:      user.Phone.String,
		CreateTime: user.CreateTime.String,
	}
}

// 查询一行数据
func SqlxqueryByGet(id int) (vo *models.UserVo) {
	//Get方法

	sqlStr := "SELECT * FROM user WHERE id = ?"
	//var u User
	//var u *User = &User{}
	u := new(models.User)
	vo = new(models.UserVo)
	err := db.Get(u, sqlStr, id)
	if err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}
	vo = getVoByUser(u)
	return
}

func SqlxList() (users []*models.User, userVos []*models.UserVo, err error) {

	query := "select * from user"
	err = db.Select(&users, query)
	if err != nil {
		fmt.Printf("error:", err)
	}
	userVos = getVosByUser(users)

	return
}

// in查询
// 查询id在给定id集合中的数据。
func GetInIds(ids []int) (users []models.User, err error) {

	query, args, err := sqlx.In("select id, name, age from user where age in (?)", ids)
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)

	return
}

// 删除--Exec
func Del(id int) (err error) {

	sqlStr := "delete from user where id = ?"
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("Del failed, err:%v\n", err)
		return
	}
	//受影响的行数
	row_affect, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("del success, affected rows:%d\n", row_affect)
	return
}

// 新增--Exec
func Insert(bo *models.AddUserBo) (insertId int64, err error) {

	sqlStr := "insert user(name, age, sex, address, phone, create_time) VALUE (?,?,?,?,?,?)"

	//QueryRowx方法可以查询,需要注意scan后的参数应与查询返回的数量一致,且一一对应
	result, err := db.Exec(sqlStr, bo.Name, bo.Age, bo.Sex, bo.Address, bo.Phone, bo.CreateTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	row_affect, err := result.RowsAffected()
	fmt.Printf("insert success, affected rows:%d\n", row_affect)
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", insertId)
	return
}

// 修改
func Update(bo *models.UpdateUserBo) {

	sqlStr := "update user set age=?,address=?,phone=? where id = ?"
	exec, err := db.Exec(sqlStr, bo.Age, bo.Address, bo.Phone, bo.Id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := exec.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}
