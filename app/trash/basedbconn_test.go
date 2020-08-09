package trash

import (
	"fmt"
	"testing"
)

func TestDb_cond_to_sqladd_ex(t *testing.T){
	data1 :=[]DbDataWithOpEx{
		DbDataWithOpEx{"IN",map[string]interface{}{
			"co1": []interface{}{1, 2, 3, 4},
			"co2": []interface{}{"a", "b"},
		}},
		DbDataWithOpEx{"AND", map[string]interface{}{
			"col1": OpData{">=","test1"},
			"col2": OpData{"<",2},
		}},
		DbDataWithOpEx{"OR",map[string]interface{}{
			"col1": OpData{">=","testy"},
			"col2": OpData{"<",99},
		}},
		DbDataWithOpEx{"OR",map[string]interface{}{
			"col1": OpData{"%","testx"},
			"col2": OpData{"%-","testy"},
			"col3": OpData{"-%","testz"},
		}},
		DbDataWithOpEx{"AND",map[string]interface{}{
			"col1": OpData{"%","testx"},
			"col2": OpData{"%-","testy"},
			"col3": OpData{"-%","testz"},
		}},
		DbDataWithOpEx{"AND",map[string]interface{}{
			"col1": OpData{"%","testx"},
		}},
	}
	str := db_cond_to_sqladd_ex(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOpEx{
		DbDataWithOpEx{"IN",map[string]interface{}{
			"co2": []interface{}{"a", "b"},
		}},
		DbDataWithOpEx{"AND", map[string]interface{}{
			"col2": OpData{"<",2},
		}},
		DbDataWithOpEx{"OR",map[string]interface{}{
			"col1": OpData{">=","testy"},
		}},
		DbDataWithOpEx{"", map[string]interface{}{
			"col2": OpData{"",11},
			"col3": OpData{"","aa"},
		}},
		DbDataWithOpEx{"", map[string]interface{}{
			"colx": OpData{"",88},
		}},
	}
	str = db_cond_to_sqladd_ex(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOpEx{
		DbDataWithOpEx{"", map[string]interface{}{
			"colx": OpData{"",88},
		}},
	}
	str = db_cond_to_sqladd_ex(data1)
	fmt.Println(str)
}

func TestDb_cond_to_sqladd(t *testing.T) {

	data1 :=[]DbDataWithOp{
		DbDataWithOp{"IN",nil,"col1",[]interface{}{1,2}},
		DbDataWithOp{"IN",nil,"col2",[]interface{}{1.1,2.2}},
		DbDataWithOp{"IN",nil,"col3",[]interface{}{"1.1.1","2.2.2"}},
		DbDataWithOp{"IN",nil,"col4",[]interface{}{9}},
	}
	str := db_cond_to_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOp{
		DbDataWithOp{"IN",nil,"col4",[]interface{}{9}},
	}
	str = db_cond_to_sqladd(data1)
	fmt.Println(str)


	data1 =[]DbDataWithOp{
		DbDataWithOp{"AND",[]string{"=",">="},[]string{"col4","col5"},[]interface{}{9,"abc"}},
	}
	str = db_cond_to_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOp{
		DbDataWithOp{"",[]string{"=",">="},[]string{"col4"},[]interface{}{91}},
	}
	str = db_cond_to_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOp{
		DbDataWithOp{"",[]string{"=",">="},[]string{"col4"},[]interface{}{91}},
		DbDataWithOp{"",[]string{"=",">="},[]string{"col5"},[]interface{}{92}},
		DbDataWithOp{"",[]string{"=",">="},[]string{"col6"},[]interface{}{93}},
	}
	str = db_cond_to_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOp{
		DbDataWithOp{"OR",[]string{"=",">="},[]string{"col4","col5"},[]interface{}{19,"abc1"}},
	}
	str = db_cond_to_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbDataWithOp{
		DbDataWithOp{"OR",[]string{"="},[]string{"col4"},[]interface{}{19,"abc1"}},
		DbDataWithOp{"AND",[]string{"=",">"},[]string{"col4","col5"},[]interface{}{18,"abc1"}},
		DbDataWithOp{"IN",nil,"col3",[]interface{}{"1.1.1","2.2.2"}},
	}
	str = db_cond_to_sqladd(data1)
	fmt.Println(str)
}

func TestDb_orderby_to_sqladd(t *testing.T) {
	data1 :=[]DbData{
		DbData{"IN1",1},
		DbData{"IN2",0},
	}
	str := Db_orderby_to_sqladd(data1)
	fmt.Println(str)
}

func TestDb_array_to_update_sqladd(t *testing.T) {
	data1 :=[]DbData{
		DbData{"IN1",1},
		DbData{"IN2","trt"},
	}
	str := Db_array_to_update_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbData{
		DbData{"IN3",1},
	}
	str = Db_array_to_update_sqladd(data1)
	fmt.Println(str)
}

func TestDb_array_to_insert_sqladd(t *testing.T) {
	data1 :=[]DbData{
		DbData{"IN1","uuuuuu"},
		DbData{"IN2",0},
	}
	str := Db_array_to_insert_sqladd(data1)
	fmt.Println(str)

	data1 =[]DbData{
		DbData{"IN4","ssssss"},
	}
	str = Db_array_to_insert_sqladd(data1)
	fmt.Println(str)
}