package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbHost     = "172.17.0.1"
	dbPort     = 3307
	dbUsername = "root"
	dbPassword = "root123"
	dbSchema   = "products"
)

// Default konversi yang digunakan untuk binding field ke column database adalah snack case FirstName -> first_name
type User struct {
	ID        uint64       `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement;not null"` // Eksplisit mapping name field column:id
	Email     string       `gorm:"type:varchar(100);not null"`
	FirstName string       `gorm:"type:varchar(100);not null"` // nama field first_name
	LastName  string       `gorm:"type:varchar(100);not null"`
	Password  string       `gorm:"type:varchar(128);not null"`
	Address   string       `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time    `gorm:"type:datetime;not null"` // default field name: create_at
	UpdatedAt time.Time    `gorm:"type:datetime;not null"`
	DeletedAt sql.NullTime `gorm:"type:datetime"`
}

// func (User) TableName() string {
// 	return "mobile_users"
// }

// https://gorm.io/docs/index.html

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbSchema,
	)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// Logger:                 logger.New(
		// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
		// 	logger.Config{
		// 		SlowThreshold:             0,
		// 		Colorful:                  false,
		// 		IgnoreRecordNotFoundError: true,
		// 		ParameterizedQueries:      false,
		// 		LogLevel:                  0,
		// 	},
		// ),
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalln("Open connection:", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalln("ping database:", err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	var user = User{
		Email:     "example@email.com",
		FirstName: "john",
		LastName:  "wick",
		Password:  "ABC",
	}
	db.Create(&user)
	fmt.Println("ID User:", user.ID)

	// Read
	var result User
	/*
		Benar : db.First(&result, "field1 = ? AND field2 = ?", field1, field2)
		Salah : db.First(&result, fmt.Sprintf("field1 = '%s' AND field2 = %d", field1, field2))
	*/
	db.First(&result, user.ID) // Find user with integer primary key
	fmt.Printf("First #1: %+v \n", result)

	result = User{}
	db.First(&result, "first_name = ?", "john") // Ketika tidak diketemukan maka akan muncul error gorm.ErrRecordNotFound
	fmt.Printf("First #2: %+v \n", result)

	// Tidak akan muncul error, result data field bernilai zero value
	result = User{}
	// result.CreatedAt.IsZero() == true -> 0001-01-01 00:00:00 +0000 UTC
	if err := db.Find(&result, "first_name = ?", "john1").Error; err != nil {
		log.Fatalln("find user:", err.Error())
	}
	fmt.Printf("Find #1: %+v \n", result)

	// Update - update user address
	res := db.Model(&User{}).Where("id = ?", user.ID).Update("address", "Bondowoso")
	fmt.Println("Affected Rows #1:", res.RowsAffected)

	res2 := db.Model(&User{}).Where("id = ?", user.ID).Updates(User{Password: "ABC", CreatedAt: time.Now(), UpdatedAt: time.Now()}) // Non-zero fields
	fmt.Println("Affected Rows #2:", res2.RowsAffected)

	res3 := db.Model(&User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{"password": ""})
	fmt.Println("Affected Rows #3:", res3.RowsAffected)

	// Delete data
	db.Delete(&User{}, user.ID)
}
