package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Seller save information
type Seller struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//Hash function is used for encrypt password form
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword compare password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave function call Hash function for encrypt password
func (s *Seller) BeforeSave() error {
	hashedPassword, err := Hash(s.Password)
	if err != nil {
		return err
	}
	s.Password = string(hashedPassword)
	return nil
}

//Prepare function is used for trim input param space
func (s *Seller) Prepare() {
	s.ID = 0
	s.Nickname = html.EscapeString(strings.TrimSpace(s.Nickname))
	s.Email = html.EscapeString(strings.TrimSpace(s.Email))
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

//Validate function is used for validation on input param value before save to database
func (s *Seller) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if s.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if s.Password == "" {
			return errors.New("Required Password")
		}
		if s.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(s.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if s.Password == "" {
			return errors.New("Required Password")
		}
		if s.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(s.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if s.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if s.Password == "" {
			return errors.New("Required Password")
		}
		if s.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(s.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

//SaveSeller function is used for save data in database
func (s *Seller) SaveSeller(db *gorm.DB) (*Seller, error) {

	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Seller{}, err
	}
	return s, nil
}

//FindAllSellers function is used for fetch all seller table data
func (s *Seller) FindAllSellers(db *gorm.DB) (*[]Seller, error) {
	var err error
	sellers := []Seller{}
	err = db.Debug().Model(&Seller{}).Limit(100).Find(&sellers).Error
	if err != nil {
		return &[]Seller{}, err
	}
	return &sellers, err
}

//FindSellerByID function is used for get specific data
func (s *Seller) FindSellerByID(db *gorm.DB, uid uint32) (*Seller, error) {
	var err error
	err = db.Debug().Model(Seller{}).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return &Seller{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Seller{}, errors.New("User Not Found")
	}
	return s, err
}

//UpdateASeller function is used for update data of seller
func (s *Seller) UpdateASeller(db *gorm.DB, uid uint32) (*Seller, error) {

	// To hash the password
	err := s.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Seller{}).Where("id = ?", uid).Take(&Seller{}).UpdateColumns(
		map[string]interface{}{
			"password":  s.Password,
			"nickname":  s.Nickname,
			"email":     s.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Seller{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Seller{}).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return &Seller{}, err
	}
	return s, nil
}

//DeleteASeller function is used for delete a specific seller
func (s *Seller) DeleteASeller(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Seller{}).Where("id = ?", uid).Take(&Seller{}).Delete(&Seller{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
