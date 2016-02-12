package main

import (
	"os"
	"encoding/json"
)

type SettingType struct{
	ApiUrl string
	BaseUrl string
	Sock  string
}

func (s *SettingType) ReadSettings() error {
	fileConf, err := os.Open("settings")
	defer fileConf.Close()
	if err != nil {
		return err
	}

	jsonParser := json.NewDecoder(fileConf)
	err = jsonParser.Decode(&s)
	if err != nil {
		return err
	}
	return nil
}

func (s *SettingType) SavSeettings() error{
	jsonConfig, err := json.Marshal(s)
	if err != nil {
		return err
	}
	f, err := os.Create("settings")
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(jsonConfig)
	f.Close()
	return nil
}