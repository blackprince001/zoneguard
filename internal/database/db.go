package database

import (
	"gorm.io/gorm"
)

type IplocatorInterface interface {
	CreateIplocatorItem(geo_ip *Iplocator) error
	GetIplocatorItem(id uint) error
}

type IplocatorRepo struct {
	db *gorm.DB
}

func NewIplocatorRepo(db *gorm.DB) IplocatorInterface {
	return &IplocatorRepo{db: db}
}

func (r *IplocatorRepo) CreateIplocatorItem(geo_ip *Iplocator) error {
	return r.db.Create(geo_ip).Error
}

// GetIplocatorItem implements IplocatorInterface.
func (r *IplocatorRepo) GetIplocatorItem(id uint) error {
	panic("unimplemented")
}
