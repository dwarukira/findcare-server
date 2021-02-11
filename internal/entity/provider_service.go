package entity

type ProviderService struct {
	ProviderID uint    `gorm:"primaryKey"`
	ServiceID  uint    `gorm:"primaryKey"`
	Price      float32 `json:"Price"`
}
