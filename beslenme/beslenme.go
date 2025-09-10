package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=nutrition port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Tabloları oluştur
	db.AutoMigrate(
		&TDietType{},
		&TDisease{},
		&TNutritionPlan{},
		&TMeal{},
		&TRecipe{},
	)
}

// Diyet tipleri
type TDietType struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Hastalıklar
type TDisease struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Beslenme planı
type TNutritionPlan struct {
	Goal          string      `gorm:"primaryKey;type:varchar(50)"` // Zayıflama / Kas / Kilo Alma / Sıkılaşma
	DiseaseID     uint        `gorm:"primaryKey;type:int"`         // Hastalık
	Title         string      `gorm:"type:varchar(150)"`
	Description   string      `gorm:"type:text"`
	DailyCalories uint        `gorm:"type:int"`
	DietTypes     []TDietType `gorm:"many2many:t_plan_diettypes;"`
	Meals         []TMeal     `gorm:"foreignKey:NutritionPlanGoal,NutritionPlanDiseaseID;references:Goal,DiseaseID"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`

	// Foreign key ilişkisi
	Disease TDisease `gorm:"foreignKey:DiseaseID;references:ID"`
}

// Öğünler
type TMeal struct {
	ID                     uint      `gorm:"primaryKey;autoIncrement"`
	NutritionPlanGoal      string    `gorm:"type:varchar(50)"`
	NutritionPlanDiseaseID uint      `gorm:"type:int"`
	Name                   string    `gorm:"type:varchar(50)"`
	TimeOfDay              string    `gorm:"type:varchar(10)"`
	Recipes                []TRecipe `gorm:"foreignKey:MealID"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime"`
}

// Tarifler
type TRecipe struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	MealID       uint      `gorm:"type:int"`
	Title        string    `gorm:"type:varchar(150)"`
	Ingredients  string    `gorm:"type:text"`
	Instructions string    `gorm:"type:text"`
	Calories     uint      `gorm:"type:int"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
