package infrastructure

//mysqlに依存したdb設定
//依存性の注入
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	Port       string
	DBName     string
	UserName   string
	Password   string
	Connection *gorm.DB
}

//依存性の注入をここで行う。
func newDB(db *DB) *DB {
	dsn := db.UserName + ":" + db.Password + "@" + "tcp(" + db.Host + ":" + db.Port + ")/" + db.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	db.Connection = gormDB
	return db

}

func NewDB() *DB {
	c := NewConfig()

	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Port:     c.DB.Production.Port,
		DBName:   c.DB.Production.DBName,
		UserName: c.DB.Production.UserName,
		Password: c.DB.Production.Password,
	})

}

func NewTestDB() *DB {
	c := NewConfig()

	return newDB(&DB{
		Host:     c.DB.Test.Host,
		Port:     c.DB.Test.Port,
		DBName:   c.DB.Test.DBName,
		UserName: c.DB.Test.UserName,
		Password: c.DB.Test.Password,
	})
}

func CloseDB(db *DB) *DB {
	return db
}
