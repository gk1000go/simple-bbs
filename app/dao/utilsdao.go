package dao

//func fl1() string {
//	return aputils.File_line(1)
//}
//func fl2() string {
//	return aputils.File_line(2)
//}
//func flError(s string) error {
//	msg := fmt.Sprintf("%s\n%s\n",fl1(),s)
//	log.Errorln(msg)
//	return errors.New(s)
//}
//func hasErr(err error) bool {
//	if err != nil {
//		log.Errorf("%s\n%s\n",fl2(),err.Error())
//	}
//	return (err == nil)
//}
// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
//reflect.TypeOf(s).FieldByName()
//reflect.TypeOf(s).Elem().FieldByName() // pointer
//func FetchData(row *sqlx.Rows,result interface{}) error {
//	if result==nil {
//		return flError("Not support result type<nil>")
//	}
//	typ := reflect.TypeOf(result)
//	if typ.Kind() == reflect.Ptr {
//		newData := reflect.New(typ.Elem())
//		if hasErr(row.StructScan(newData.Interface())) {
//			return row.Err()
//		}
//		result = newData
//	}else if(typ.Kind() == reflect.Slice) {
//		if ret,e := row.SliceScan();hasErr(e) {
//			return e
//		}else{
//			result = ret
//		}
//	}else if(typ.Kind() == reflect.Map) {
//		result = make(map[string]interface{})
//		if hasErr(row.MapScan(result.(map[string]interface{}))) {
//			return row.Err()
//		}
//	}else{
//		return flError("Not support result type")
//	}
//	return nil
//}

//func doQuery(sql string,param []interface{}) (*sqlx.Rows, error) {
//	db := framework.GetConn()
//	return db.Queryx(sql,param)
//}
//func GetOne(sql string,param []interface{},result interface{}) error {
//	db := framework.GetConn()
//	rows,err := db.Queryx(sql,param)
//	if hasErr(err){return err}
//	defer rows.Close()
//	if rows.Next() {
//		return FetchData(rows,result)
//	}else{
//		return flError("No Data.")
//	}
//}

//func GetList(sql string,param []interface{},result interface{}) error {
//	db := framework.GetConn()
//	rows,err := db.Queryx(sql,param)
//	if hasErr(err){
//		return err
//	}
//	defer rows.Close()
//
//	typ := reflect.TypeOf(result)
//	if(typ.Kind() == reflect.Slice) {
//		if ret,e := row.SliceScan();hasErr(e) {
//			return e
//		}else{
//			result = ret
//		}
//	}else if(typ.Kind() == reflect.Map) {
//		if hasErr(row.MapScan(result.(map[string]interface{}))) {
//			return row.Err()
//		}
//	}else{
//		return flError("Not support result type")
//	}
//	return nil
//}
//
//func (d *utilDAO)GetGroupList(sql string) (result map[uint]*model.Group) {
//	rows,err := d.Db.Queryx(sql)
//	if d.GetError(err){
//		return nil
//	}
//	defer rows.Close()
//	result = make(map[uint]*model.Group)
//	mg := &model.Group{}
//	for rows.Next(){
//		if err := rows.StructScan(mg);err!=nil{
//			log.Warnln("GetGroupList/StructScan:",err.Error())
//		}else{
//			result[mg.Gid] = mg
//		}
//	}
//	return
//}
//func (d *utilDAO)GetForumList(sql string) (result map[uint]*model.Forum) {
//	rows,err := d.Db.Queryx(sql)
//	if d.GetError(err){
//		return nil
//	}
//	defer rows.Close()
//	result = make(map[uint]*model.Forum)
//	mf := &model.Forum{}
//	for rows.Next(){
//		if err := rows.StructScan(mf);err!=nil{
//			log.Warnln("GetForumList/StructScan:",err.Error())
//		}else{
//			result[mf.Fid] = mf
//		}
//	}
//	return
//}
//func (d *utilDAO)GetForumAccessList(sql string) (result map[uint]*model.Forum_access) {
//	//rows,err := d.Db.Queryx(sql)
//	//if d.GetError(err){
//	//	return nil
//	//}
//	//defer rows.Close()
//	//result = make(map[uint]*model.Forum)
//	//mf := &model.Forum{}
//	//for rows.Next(){
//	//	if err := rows.StructScan(mf);err!=nil{
//	//		log.Warnln("GetForumList/StructScan:",err.Error())
//	//	}else{
//	//		result[mf.Fid] = mf
//	//	}
//	//}
//	return
//}
/*
Save は SQL を実行する際にすべてのフィールドを含みます。
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)

//// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
 */


/*
// ひとつのフィールドを更新します
db.Model(&user).Update("name", "hello")
//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 条件付きでひとつのフィールドを更新します
db.Model(&user).Where("active = ?", true).Update("name", "hello")
//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

// `map` で複数のフィールドを更新します(対象のフィールドのみ)
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
//// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// `struct` で複数のフィールドを更新します(空ではないフィールドのみ)
db.Model(&user).Updates(User{Name: "hello", Age: 18})
//// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;



もしupdate時に特定のフィールドのみを更新する、もしくは無視する時に、selectとomitが使えます
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
//// UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
 */

/*
Hook なしでの更新
上記の更新処理は、BeforeUpdate, AfterUpdateメソッドを実行します。その結果更新時にUpdatedAtのタイムスタンプや 持っている
Associations が更新されます。もしそれらのメソッドを呼びたくない場合はUpdateColumnとUpdateColumnsが使えます。
// Update single attribute, similar with `Update`
db.Model(&user).UpdateColumn("name", "hello")
//// UPDATE users SET name='hello' WHERE id = 111;

// Update multiple attributes, similar with `Updates`
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
//// UPDATE users SET name='hello', age=18 WHERE id = 111;


バッチでの更新
フックはバッチアップデート時は実行されません。
db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

// Update with struct only works with none zero values, or use map[string]interface{}
db.Model(User{}).Updates(User{Name: "hello", Age: 18})
//// UPDATE users SET name='hello', age=18;

// Get updated records count with `RowsAffected`
db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected


Hook での値変更
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
  if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
    scope.SetColumn("EncryptedPassword", pw)
  }
}
 */