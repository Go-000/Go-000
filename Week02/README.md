## 课后作业

### 问题:我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


### 回答

> 从api层到业务Service，再到dao层。取决业务如何去定义CRUD操作的意义。下面个人从包容性强到弱进行理解


- 包容性很强的情况
    - Create：创建出错，在包容性很强情况下也应该抛错给上层
    - Read：读取不到数据，可能就是没数据，在大前提下，不需要抛错，应当理解成一个正常的情况，记录一下log即可
    - Update：更新失败，可能就是没有符合条件的数据，在大前提下，也可以理解成为一个正常的情况，记录下log即可
    - Delete：删除失败，同样是没符合条件的数据，在大前提下，也可以理解为对资源删除操作成功，只不过是一个空操作


- 包容性弱的情况
    - Create：创建出错，在包容性很强情况下也应该抛错给上层
    - Read：读取不到数据，那就需要告诉上层dao层操作失败
    - Update：更新失败，那就需要告诉上层dao层操作失败
    - Delete：删除失败，那就需要告诉上层dao层操作失败


### 个人理解代码

```

type ErrNoData struct{}

func (e *ErrNoData) Error() string{

}

// service

type UserService struct{

}

func GetUserById(id int) (model.User,error){
   user,err := dao.FindUserById(id)
   // 如果是严格模式
   if errors.Is(err,sql.ErrNoRows) {
       return user,errors.Wrap(err,fmt.Sprintf("user_id: %v has no data", id))
   }
   // 如果是宽松模式下
   if errors.Is(err,sql.ErrNoRows){
	fmt.Sprintf("user_id:%v has no data",id)
	return user,nil
   }
   // 这里处理非sql.ErrNoRows
   if !errors.Is(err,sqlErrNoRows){
       return nil,err
   }
}

func UpdateUserNameById(id int,name string)(model,error){
   return dao.UpdateUserNameById(id,name)
}


// dao

type User struct{
}

func FindUserById(id int)(User,error){
    user := new(User)
    err := DB.QueryRow("select id,username,password from users where id=?", id).Scan(user)
    return user,err
}
