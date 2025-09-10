package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=deneme port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Tabloları oluştur
	db.AutoMigrate(&TBlog{}, &TBlog1{}, &TCoach{}, &TAreasOfExpertise{}, &TAreasOfCoach{})
}

type TBlog struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;type:int"`
	Title       string    `gorm:"type:varchar(75)"`
	Description string    `gorm:"type:varchar(165)"`
	Content     string    `gorm:"type:text"`
	Tags        string    `gorm:"type:varchar(255)"`
	Data        []byte    `gorm:"type:bytea"`
	FileType    string    `gorm:"type:varchar(50)"`
	Views       uint      `gorm:"not null;default:0"`
	Slug        string    `gorm:"type:varchar(165);uniqueIndex:idx_user_slug"`
	Active      bool      `gorm:"type:bool;default:false"`
	Approval    bool      `gorm:"type:bool;default:false"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type TBlog1 struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;type:int"`
	Title       string    `gorm:"type:varchar(75)"`
	Description string    `gorm:"type:varchar(165)"`
	Content     string    `gorm:"type:text"`
	Tags        string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(165);uniqueIndex:idx_user_slug"`
	Active      bool      `gorm:"type:bool;default:false"`
	Approval    bool      `gorm:"type:bool;default:false"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type TCoach struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;type:int"`
	Name      string    `gorm:"type:varchar(150)"`
	Content   string    `gorm:"type:text"`
	Active    bool      `gorm:"type:bool;default:false"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type TAreasOfExpertise struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;type:int"`
	Name      string    `gorm:"type:varchar(150)"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type TAreasOfCoach struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement;type:int"`
	TCoachId            uint      `gorm:"type:int"`
	TAreasOfExpertiseId uint      `gorm:"type:int"`
	CreatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

// Optional: Add relationships if needed
type TCoachWithRelations struct {
	TCoach
	AreasOfExpertise []TAreasOfExpertise `gorm:"many2many:t_areas_of_coaches;"`
}
