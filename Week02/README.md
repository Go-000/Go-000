##学习笔记

#### 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

sql.ErrNoRows 是标准库的函数，dao层需要对其做wrap包装。但不应该包含具体的数据操作的信息 

大致代码如下：


```
//service UserService
// GetUser Get User
func (u *UserService) GetUser(id uint64) (model.User, error) {
	b := biz.UserBiz{}
	user, err := b.GetUser(id)
	if err != nil {
		return user, fmt.Errorf("service: error occured. \r\n", err)
	}
	fmt.Printf("Found user: %#v\r\n", user)
	return user, err
}

//biz UserBiz
// GetUser Get User
func (u *UserBiz) GetUser(id uint64) (model.User, error) {
	d := dao.UserDao{}
	return d.GetUser(id)
}

//dao UserDao
// GetUser Get User
func (d *UserDao) GetUser(id uint64) (model.User, error) {
	var u model.User
	errScan := db.QueryRow("select accountid,firstName,lastName,"+
		"userno from user where user_no = ? limit 1", id).Scan(&u.AccountID,
		&u.FirstName, &u.LastName, &u.UserNo)
	if errScan != nil {
		return u, errors.Wrap(errScan, "Scan error")
	}

	defer db.Close()
	return u, nil
}

```
