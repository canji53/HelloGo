package memo

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "db"
    "entity"
)

type Service struct{}

type Memo entity.Memo


// Get All
func (s Service) GetAll() ([]Memo, error) {
    db := db.GetDB()
    var m []Memo

    if err := db.Find(&m).Error; err != nil {
        return nil, err
    }

    return m, nil
}

// Create
func (s Service) CreateModel(c *gin.Context) (Memo, error) {
    db := db.GetDB()
    var m Memo

    if err := c.BindJSON(&m); err != nil {
        return m, err
    }

	// Create UUID4
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	uuid := u.String()

    // Insert field
    m.ID = uuid
    m.Updated = time.Now()

    if err := db.Create(&m).Error; err != nil {
        return m, err
    }

    return m, nil
}

// GetByID
func (s Service) GetByID(id string) (Memo, error) {
    db := db.GetDB()
    var m Memo

    if err := db.Where("id = ?", id).First(&m).Error; err != nil {
        return m, err
    }

    return m, nil
}

// UpdateByID
func (s Service) UpdateByID(id string, c *gin.Context) (Memo, error) {
    db := db.GetDB()
    var m Memo

    if err := db.Where("id = ?", id).First(&m).Error; err != nil {
        return m, err
    }

    if err := c.BindJSON(&m); err != nil {
        return m, err
    }

    // Update updated time
    m.Updated = time.Now()

    db.Save(&m)

    return m, nil
}

// DeleteByID
func (s Service) DeleteByID(id string) error {
    db := db.GetDB()
    var m Memo

    if err := db.Where("id = ?", id).Delete(&m).Error; err != nil {
        return err
    }

    return nil
}
