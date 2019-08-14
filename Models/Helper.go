package Models

type (
	Helper struct {
		ID             uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
		Username       string `gorm:"not null" json:"username"`
		Firstname      string `gorm:"not null" json:"firstname"`
		Lastname       string `gorm:"not null" json:"lastname"`
		Email          string `gorm:"not null" json:"email"`
		Dob            string `gorm:"not null" json:"dob"`
		Password       string `gorm:"not null" json:"password"`
		IsSocMed       bool   `gorm:"not null" json:"is_soc_med"`
		ProfileImage   string `gorm:"not null" json:"profile_image"`
		AutographImage string `gorm:"not null" json:"autograph_image"`
	}
)
