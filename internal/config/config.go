package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/iskhakmuhamad/todo-api/internal/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	Port           string
	JWTSecret      string
	JWTExpireHours int
	RunSeeder      bool
}

func Load() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	jwtExpire, _ := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "24"))
	runSeeder, _ := strconv.ParseBool(getEnv("RUN_SEEDER", "false"))

	return &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "todoapp"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		Port:           getEnv("PORT", "3000"),
		JWTSecret:      getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpireHours: jwtExpire,
		RunSeeder:      runSeeder,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ConnectDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Category{},
		&domain.Todo{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed")
}
