package database

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/models"
	"main/utils"
	"os"
)

var DB *gorm.DB

func InitDatabase(dbName string) {
	var (
		databaseUser     string = utils.GetValue("DB_USER")
		databasePassword string = utils.GetValue("DB_PASSWORD")
		databaseHost     string = utils.GetValue("DB_HOST")
		databasePort     string = utils.GetValue("DB_PORT")
		databaseName     string = dbName
	)
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_PORT"))
	fmt.Println(os.Getenv("DB_USER"))
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort)
	var err error
	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	if err := DB.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName).Error; err != nil {
		panic("failed to create database: " + err.Error())
	}
	dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the selected database: " + err.Error())
	}
	fmt.Println("Connected to database")
	DB.AutoMigrate(&models.User{}, &models.Item{})
}

func SeedItem() (models.Item, error) {
	item, err := utils.CreateFaker[models.Item]()
	if err != nil {
		return models.Item{}, err
	}
	DB.Create(&item)
	fmt.Println("Item seeded to the database")
	return item, nil
}

func SeedUser() (models.User, error) {
	user, err := utils.CreateFaker[models.User]()
	if err != nil {
		return models.User{}, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	var inputUser = models.User{
		ID:       user.ID,
		Password: string(password),
		Email:    user.Email,
	}
	DB.Create(&inputUser)
	fmt.Println("User seeded to the database")
	return user, nil
}

func CleanSeeders() {
	itemsResult := DB.Exec("TRUNCATE items")
	usersResult := DB.Exec("TRUNCATE users")
	isFailed := itemsResult.Error != nil || usersResult.Error != nil
	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}
	fmt.Println("Seeders are cleaned up successfully")
}
