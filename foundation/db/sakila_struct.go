package db

import (
	"time"
)

// Actor -
type Actor struct {
	ActorId    int       `xorm:"not null pk autoincr SMALLINT(5)"`
	FirstName  string    `xorm:"not null VARCHAR(45)"`
	LastName   string    `xorm:"not null index VARCHAR(45)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Address -
type Address struct {
	AddressId  int       `xorm:"not null pk autoincr SMALLINT(5)"`
	Address    string    `xorm:"not null VARCHAR(50)"`
	Address2   string    `xorm:"VARCHAR(50)"`
	District   string    `xorm:"not null VARCHAR(20)"`
	CityId     int       `xorm:"not null index SMALLINT(5)"`
	PostalCode string    `xorm:"VARCHAR(10)"`
	Phone      string    `xorm:"not null VARCHAR(20)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Category -
type Category struct {
	CategoryId int       `xorm:"not null pk autoincr TINYINT(3)"`
	Name       string    `xorm:"not null VARCHAR(25)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// City -
type City struct {
	CityId     int       `xorm:"not null pk autoincr SMALLINT(5)"`
	City       string    `xorm:"not null VARCHAR(50)"`
	CountryId  int       `xorm:"not null index SMALLINT(5)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Country -
type Country struct {
	CountryId  int       `xorm:"not null pk autoincr SMALLINT(5)"`
	Country    string    `xorm:"not null VARCHAR(50)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Customer -
type Customer struct {
	CustomerId int       `xorm:"not null pk autoincr SMALLINT(5)"`
	StoreId    int       `xorm:"not null index TINYINT(3)"`
	FirstName  string    `xorm:"not null VARCHAR(45)"`
	LastName   string    `xorm:"not null index VARCHAR(45)"`
	Email      string    `xorm:"VARCHAR(50)"`
	AddressId  int       `xorm:"not null index SMALLINT(5)"`
	Active     int       `xorm:"not null default 1 TINYINT(1)"`
	CreateDate time.Time `xorm:"not null DATETIME"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Film -
type Film struct {
	FilmId             int       `xorm:"not null pk autoincr SMALLINT(5)"`
	Title              string    `xorm:"not null index VARCHAR(255)"`
	Description        string    `xorm:"TEXT"`
	LanguageId         int       `xorm:"not null index TINYINT(3)"`
	OriginalLanguageId int       `xorm:"index TINYINT(3)"`
	RentalDuration     int       `xorm:"not null default 3 TINYINT(3)"`
	RentalRate         string    `xorm:"not null default 4.99 DECIMAL(4,2)"`
	Length             int       `xorm:"SMALLINT(5)"`
	ReplacementCost    string    `xorm:"not null default 19.99 DECIMAL(5,2)"`
	Rating             string    `xorm:"default 'G' ENUM('G','PG','PG-13','R','NC-17')"`
	SpecialFeatures    string    `xorm:"SET('Trailers','Commentaries','Deleted Scenes','Behind the Scenes')"`
	LastUpdate         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// FilmActor -
type FilmActor struct {
	ActorId    int       `xorm:"not null pk SMALLINT(5)"`
	FilmId     int       `xorm:"not null pk index SMALLINT(5)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// FilmCategory -
type FilmCategory struct {
	FilmId     int       `xorm:"not null pk SMALLINT(5)"`
	CategoryId int       `xorm:"not null pk index TINYINT(3)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// FilmText -
type FilmText struct {
	FilmId      int    `xorm:"not null pk SMALLINT(6)"`
	Title       string `xorm:"not null index(idx_title_description) VARCHAR(255)"`
	Description string `xorm:"index(idx_title_description) TEXT"`
}

// Inventory -
type Inventory struct {
	InventoryId int       `xorm:"not null pk autoincr MEDIUMINT(8)"`
	FilmId      int       `xorm:"not null index(idx_store_id_film_id) index SMALLINT(5)"`
	StoreId     int       `xorm:"not null index(idx_store_id_film_id) TINYINT(3)"`
	LastUpdate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Language -
type Language struct {
	LanguageId int       `xorm:"not null pk autoincr TINYINT(3)"`
	Name       string    `xorm:"not null CHAR(20)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Payment -
type Payment struct {
	PaymentId   int       `xorm:"not null pk autoincr SMALLINT(5)"`
	CustomerId  int       `xorm:"not null index SMALLINT(5)"`
	StaffId     int       `xorm:"not null index TINYINT(3)"`
	RentalId    int       `xorm:"index INT(11)"`
	Amount      string    `xorm:"not null index(idx_payment_date) DECIMAL(5,2)"`
	PaymentDate time.Time `xorm:"not null index(idx_payment_date) DATETIME"`
	LastUpdate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' index(idx_payment_date) TIMESTAMP"`
}

// Rental -
type Rental struct {
	RentalId    int       `xorm:"not null pk autoincr INT(11)"`
	RentalDate  time.Time `xorm:"not null index(idx_rental_date) DATETIME"`
	InventoryId int       `xorm:"not null index index(idx_rental_date) MEDIUMINT(8)"`
	CustomerId  int       `xorm:"not null index(idx_rental_date) index SMALLINT(5)"`
	ReturnDate  time.Time `xorm:"DATETIME"`
	StaffId     int       `xorm:"not null index TINYINT(3)"`
	LastUpdate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Staff -
type Staff struct {
	StaffId    int       `xorm:"not null pk autoincr TINYINT(3)"`
	FirstName  string    `xorm:"not null VARCHAR(45)"`
	LastName   string    `xorm:"not null VARCHAR(45)"`
	AddressId  int       `xorm:"not null index SMALLINT(5)"`
	Picture    []byte    `xorm:"BLOB"`
	Email      string    `xorm:"VARCHAR(50)"`
	StoreId    int       `xorm:"not null index TINYINT(3)"`
	Active     int       `xorm:"not null default 1 TINYINT(1)"`
	Username   string    `xorm:"not null VARCHAR(16)"`
	Password   string    `xorm:"VARCHAR(40)"`
	LastUpdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

// Store -
type Store struct {
	StoreId        int       `xorm:"not null pk autoincr TINYINT(3)"`
	ManagerStaffId int       `xorm:"not null unique TINYINT(3)"`
	AddressId      int       `xorm:"not null index SMALLINT(5)"`
	LastUpdate     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
