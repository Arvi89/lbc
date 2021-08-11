package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	MaxWorkers int64 `env:"MAX_WORKERS" envDefault:"2"`
	DbUsername string `env:"DB_USERNAME" envDefault:"lbc"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"lbc"`
	DbHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbPort     string `env:"DB_PORT" envDefault:"3307"`
	DbName     string `env:"DB_NAME" envDefault:"lbc"`
	Db         *gorm.DB
}

func NewApp() *App {
	_ = godotenv.Load()

	app := App{}
	if err := env.Parse(&app); err != nil {
		fmt.Println("Error loading env")
	}

	app.initDb()

	return &app
}

func (app *App) initDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.DbUsername, app.DbPassword, app.DbHost, app.DbPort, app.DbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app.Db = db
}
