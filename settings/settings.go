package settings

import (
	"os"
	"encoding/json"
	"bytes"
)

type SettingsType struct{
	ApiUrl string
	BaseUrl string
	Sock  string
}


func (s *SettingsType)ReadSettings() error {
	fileSettings, err := os.Open("dwm.settings")
	defer fileSettings.Close()
	if err != nil {
		return err
	}

	jsonSettings := json.NewDecoder(fileSettings)
	err = jsonSettings.Decode(&s)
	if err != nil {
		return err
	}
	return nil
}

func (s *SettingsType)SavSettings() error{
	jsonSettings, err := json.Marshal(s)
	if err != nil {
		return err
	}
	var out bytes.Buffer
	json.Indent(&out, jsonSettings, "", "\t")
	f, err := os.Create("dwm.settings")
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(out.Bytes())
	f.Close()
	return nil
}