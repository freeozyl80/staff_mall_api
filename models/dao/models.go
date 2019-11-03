package dao

import (
	"fmt"
	"log"
	"staff-mall-center/pkg/setting"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int   `gorm:"primary_key;auto_increment:10" json:"id"`
	CreatedOn  int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
	DeletedOn  int64 `json:"deleted_on"`
}

func Setup() {
	fmt.Println("开始连接数据库")
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		fmt.Printf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)  //用于设置最大打开的连接数，默认值为0表示不限制。
	db.DB().SetMaxOpenConns(100) //于设置闲置的连接数

	if !db.HasTable(&User{}) {
		log.Println("用户表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&User{})
	}
	if !db.HasTable(&Auth{}) {
		log.Println("权限表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Auth{})
	}
	if !db.HasTable(&Firm{}) {
		log.Println("公司表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Firm{})
	}
	if !db.HasTable(&Product{}) {
		log.Println("商品表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Product{})
	}
	if !db.HasTable(&Category{}) {
		log.Println("品类表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Category{})
	}
	if !db.HasTable(&Staff{}) {
		log.Println("员工表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Staff{})
	}
	if !db.HasTable(&Order{}) {
		log.Println("订单表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Order{})
	}
	if !db.HasTable(&Supplier{}) {
		log.Println("供应商表创建ing")
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&Supplier{})
	}
}
func IsEstablish() bool {
	if !db.HasTable(&User{}) {
		return false
	}
	if !db.HasTable(&Auth{}) {
		return false
	}
	if !db.HasTable(&Firm{}) {
		return false
	}
	if !db.HasTable(&Product{}) {
		return false
	}
	if !db.HasTable(&Category{}) {
		return false
	}
	if !db.HasTable(&Staff{}) {
		return false
	}
	if !db.HasTable(&Order{}) {
		return false
	}
	if !db.HasTable(&Supplier{}) {
		return false
	}
	return true
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("created_on"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("modified_on"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("deleted_on")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(), //返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func BulkInsertOnDuplicateUpdate(db *gorm.DB, objArr []interface{}, updates string) ([]int, error) {
	ids := []int{}
	// If there is no data, nothing to do.
	if len(objArr) == 0 {
		return ids, nil
	}

	mainObj := objArr[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			placeholders = append(placeholders, scope.AddToVars(fields[i].Field.Interface()))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}

	mainScope.Raw(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s",
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
		updates,
	))
	if res, err := mainScope.SQLDB().Exec(mainScope.SQL, mainScope.SQLVars...); err != nil {
		return ids, err
	} else {
		lastid, _ := res.LastInsertId()
		len, _ := res.RowsAffected()
		lastid2 := int(lastid)
		len2 := int(len)

		for i := 0; i < len2; i++ {
			ids = append(ids, lastid2)
			lastid2 = lastid2 + 1
		}
		return ids, err
	}
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
