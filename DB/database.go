package DB

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	Tel          string `gorm:"type:char(11)"`
	Salt         string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(30)"`
	LoginToken   string `gorm:"type:varchar(255)"`
	RefreshToken string `gorm:"type:varchar(255)"`

	StorageID uint    // 外键字段
	Storage   Storage `gorm:"foreignkey:StorageID"`

	PermissionID uint       // 外键字段
	Permission   Permission `gorm:"foreignkey:PermissionID"`
}

type Storage struct {
	gorm.Model

	LimitID uint  // 外键字段
	Limit   Limit `gorm:"foreignkey:LimitID"`

	RootFolderID uint   // 外键字段
	RootFolder   Folder `gorm:"foreignkey:RootFolderID"` // 表示根目录
}

type Limit struct {
	gorm.Model
	Capacity       int `gorm:"not null"`
	Use            int `gorm:"not null"`
	ShareNumber    int `gorm:"not null"`
	UseShareNumber int `gorm:"not null"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"not null"`
	C    bool   `gorm:"not null"`
	D    bool   `gorm:"not null"`
	U    bool   `gorm:"not null"`
	R    bool   `gorm:"not null"`
}

type Folder struct {
	gorm.Model
	Name string `gorm:"not null"`
	Size int    `gorm:"not null"`

	ParentID *uint   // 外键字段，指向父目录
	Parent   *Folder `gorm:"foreignkey:ParentID"`

	Children []Folder `gorm:"foreignkey:ParentID"` // 子目录
	Files    []File   `gorm:"many2many:folder_files;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type File struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Size     int    `gorm:"not null"`
	FolderID uint   // 外键字段，指向所属目录
	Folder   Folder `gorm:"foreignkey:FolderID"`
}

func AutoMigrate() []any {
	return []any{&User{}, &Storage{}, &Limit{}, &Permission{}, &Folder{}, &File{}}
}
