package storage

import "gorm.io/gorm"

// Resume model for PostgreSQL
type ResumeDTO struct {
	gorm.Model
	Name             string `json:"name"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	Designation      string `json:"designation"`
	Experience       int    `json:"experience"`
	HighestEducation string `json:"highest_education"`
	Location         string `json:"location"`
	Skills           string `json:"skills"`
}

func SaveResume(resume *ResumeDTO) error {
	return DB.Create(resume).Error
}

// Get all resumes from the database
func GetAllResumes() ([]ResumeDTO, error) {
	var resumes []ResumeDTO
	err := DB.Find(&resumes).Error
	return resumes, err
}

// Search resumes by skill
func SearchResumesBySkill(skill string) ([]ResumeDTO, error) {
	var resumes []ResumeDTO
	err := DB.Where("skills ILIKE ?", "%"+skill+"%").Find(&resumes).Error
	return resumes, err
}
