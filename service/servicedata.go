package service

import "ProjectSavePassword/model"

func SearchData(IdUser uint) ([]model.SaveData, error) {
	var data []model.SaveData
	if err := db.Where("user_id = ?", IdUser).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func AddData(data *model.SaveData) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
