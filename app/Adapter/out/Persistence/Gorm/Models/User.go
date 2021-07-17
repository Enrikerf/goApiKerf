package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID uuid.UUID `gorm:"primary_key; unique; 
                      type:uuid; column:id; 
                      default:uuid_generate_v4()"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	//if !user.Validate("") {
	//	err = errors.New("can't save invalid data")
	//}
	return
}

func (user *User) Prepare() {
	user.Nickname = html.EscapeString(strings.TrimSpace(user.Nickname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) FindAllUsers(db *gorm.DB) ([]User, error) {
	var err error
	users := []User{}
	err = db.Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, err
}

func (user *User) FindUserByID(db *gorm.DB, uid string) (User, error) {
	var err error
	modelUser := User{}
	err = db.Model(User{}).Where("id = ?", uid).Take(&modelUser).Error
	if err != nil {
		return User{}, err
	}
	return modelUser, err
}

func (user *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := user.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  user.Password,
			"nickname":  user.Nickname,
			"email":     user.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
