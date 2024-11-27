package database

import (
	"gorm.io/gorm"
)

type IplocatorInterface interface {
	CreateIplocatorItem(geo_ip *Iplocator) error
	GetIplocatorItem(id uint64) (*Iplocator, error)
}

type IplocatorRepo struct {
	db *gorm.DB
}

func NewIplocatorRepo(db *gorm.DB) IplocatorRepo {
	return IplocatorRepo{db: db}
}

func (r *IplocatorRepo) CreateIplocatorItem(geo_ip *Iplocator) error {
	return r.db.Create(geo_ip).Error
}

func (r *IplocatorRepo) GetIplocatorItem(id uint64) (*Iplocator, error) {
	var geo_ip Iplocator
	err := r.db.Where("related_id = ?", id).Find(&geo_ip).Error
	if err != nil {
		return nil, err
	}
	return &geo_ip, nil
}
