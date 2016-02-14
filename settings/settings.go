package settings

import (
	"os"
	"encoding/json"
)

type SettingsType struct{
	ApiUrl string
	BaseUrl string
	Sock  string
}


func (s *SettingsType)ReadSettings() error {
	fileConf, err := os.Open("dwm.settings")
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

func (s *SettingsType)SavSettings() error{
	jsonConfig, err := json.Marshal(s)
	if err != nil {
		return err
	}
	f, err := os.Create("dwm.settings")
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(jsonConfig)
	f.Close()
	return nil
}