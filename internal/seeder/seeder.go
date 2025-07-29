package seeder

import (
	"log"
	"time"

	"github.com/iskhakmuhamad/todo-api/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) error {
	log.Println("Starting database seeding...")

	// Seed Users
	if err := seedUsers(db); err != nil {
		return err
	}

	// Seed Categories
	if err := seedCategories(db); err != nil {
		return err
	}

	// Seed Todos
	if err := seedTodos(db); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully!")
	return nil
}

func seedUsers(db *gorm.DB) error {
	log.Println("Seeding users...")

	// Check if users already exist
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		log.Println("Users already exist, skipping user seeding")
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []domain.User{
		{
			Email:    "john@example.com",
			Username: "john_doe",
			Password: string(hashedPassword),
		},
		{
			Email:    "jane@example.com",
			Username: "jane_smith",
			Password: string(hashedPassword),
		},
		{
			Email:    "admin@example.com",
			Username: "admin",
			Password: string(hashedPassword),
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Error creating user %s: %v", user.Email, err)
			return err
		}
		log.Printf("Created user: %s", user.Email)
	}

	return nil
}

func seedCategories(db *gorm.DB) error {
	log.Println("Seeding categories...")

	// Check if categories already exist
	var count int64
	db.Model(&domain.Category{}).Count(&count)
	if count > 0 {
		log.Println("Categories already exist, skipping category seeding")
		return nil
	}

	// Get first user for categories
	var firstUser domain.User
	if err := db.First(&firstUser).Error; err != nil {
		log.Println("No users found, skipping category seeding")
		return nil
	}

	categories := []domain.Category{
		{
			UserID:      firstUser.ID,
			Name:        "Work",
			Description: "Work related tasks and projects",
			Color:       "#FF5722",
		},
		{
			UserID:      firstUser.ID,
			Name:        "Personal",
			Description: "Personal tasks and activities",
			Color:       "#4CAF50",
		},
		{
			UserID:      firstUser.ID,
			Name:        "Shopping",
			Description: "Shopping lists and errands",
			Color:       "#2196F3",
		},
		{
			UserID:      firstUser.ID,
			Name:        "Health",
			Description: "Health and fitness related tasks",
			Color:       "#9C27B0",
		},
		{
			UserID:      firstUser.ID,
			Name:        "Learning",
			Description: "Educational and learning activities",
			Color:       "#FF9800",
		},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			log.Printf("Error creating category %s: %v", category.Name, err)
			return err
		}
		log.Printf("Created category: %s", category.Name)
	}

	return nil
}

func seedTodos(db *gorm.DB) error {
	log.Println("Seeding todos...")

	// Check if todos already exist
	var count int64
	db.Model(&domain.Todo{}).Count(&count)
	if count > 0 {
		log.Println("Todos already exist, skipping todo seeding")
		return nil
	}

	// Get first user and categories
	var firstUser domain.User
	if err := db.First(&firstUser).Error; err != nil {
		log.Println("No users found, skipping todo seeding")
		return nil
	}

	var categories []domain.Category
	db.Where("user_id = ?", firstUser.ID).Find(&categories)
	if len(categories) == 0 {
		log.Println("No categories found, creating todos without categories")
	}

	// Helper function to get category ID
	getCategoryID := func(name string) *uint {
		for _, cat := range categories {
			if cat.Name == name {
				return &cat.ID
			}
		}
		return nil
	}

	// Create sample deadlines
	tomorrow := time.Now().AddDate(0, 0, 1)
	nextWeek := time.Now().AddDate(0, 0, 7)
	nextMonth := time.Now().AddDate(0, 1, 0)

	todos := []domain.Todo{
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Work"),
			Title:       "Complete project proposal",
			Description: "Finalize and submit the Q4 project proposal to management",
			Priority:    domain.PriorityHigh,
			Status:      domain.StatusTodo,
			Deadline:    &nextWeek,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Work"),
			Title:       "Review team performance",
			Description: "Conduct quarterly performance reviews for team members",
			Priority:    domain.PriorityMedium,
			Status:      domain.StatusTodo,
			Deadline:    &nextMonth,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Personal"),
			Title:       "Schedule doctor appointment",
			Description: "Book annual health checkup with family doctor",
			Priority:    domain.PriorityMedium,
			Status:      domain.StatusTodo,
			Deadline:    &tomorrow,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Personal"),
			Title:       "Organize home office",
			Description: "Clean and reorganize workspace for better productivity",
			Priority:    domain.PriorityLow,
			Status:      domain.StatusDone,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Shopping"),
			Title:       "Buy groceries",
			Description: "Weekly grocery shopping - milk, bread, vegetables, fruits",
			Priority:    domain.PriorityMedium,
			Status:      domain.StatusTodo,
			Deadline:    &tomorrow,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Shopping"),
			Title:       "Purchase birthday gift",
			Description: "Find and buy birthday gift for mom's birthday next week",
			Priority:    domain.PriorityHigh,
			Status:      domain.StatusTodo,
			Deadline:    &nextWeek,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Health"),
			Title:       "Morning jog",
			Description: "30-minute jog in the park to maintain fitness routine",
			Priority:    domain.PriorityMedium,
			Status:      domain.StatusDone,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Health"),
			Title:       "Prepare healthy meal plan",
			Description: "Plan healthy meals for the upcoming week",
			Priority:    domain.PriorityMedium,
			Status:      domain.StatusTodo,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Learning"),
			Title:       "Complete online course",
			Description: "Finish the Advanced Go Programming course on Udemy",
			Priority:    domain.PriorityHigh,
			Status:      domain.StatusTodo,
			Deadline:    &nextMonth,
		},
		{
			UserID:      firstUser.ID,
			CategoryID:  getCategoryID("Learning"),
			Title:       "Read technical book",
			Description: "Read 'Clean Architecture' by Robert C. Martin",
			Priority:    domain.PriorityLow,
			Status:      domain.StatusTodo,
		},
		{
			UserID:      firstUser.ID,
			Title:       "Update resume",
			Description: "Update professional resume with recent projects and skills",
			Priority:    domain.PriorityLow,
			Status:      domain.StatusTodo,
		},
		{
			UserID:      firstUser.ID,
			Title:       "Backup important files",
			Description: "Create backup of important documents and photos",
			Priority:    domain.PriorityMedium,
			Status:      domain.StatusDone,
		},
	}

	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			log.Printf("Error creating todo %s: %v", todo.Title, err)
			return err
		}
		log.Printf("Created todo: %s", todo.Title)
	}

	return nil
}
