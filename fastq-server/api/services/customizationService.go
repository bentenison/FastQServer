package services

import (
	"encoding/json"
	"log"
	"os"
)

type Customization struct {
	TTSLanguage    string `json:"tts_language,omitempty"`
	Template       string `json:"template,omitempty"`
	TicketLanguage string `json:"ticket_language,omitempty"`
}

func ChangeTTSLanguage(code string) error {
	var cts Customization
	dataBytes, err := os.ReadFile("./customization.json")
	if err != nil {
		log.Println("error while reading customization config", err)
		return err
	}
	err = json.Unmarshal(dataBytes, &cts)
	if err != nil {
		log.Println("error while unmarshal customization config", err)
		return err
	}
	cts.TTSLanguage = code
	data, err := json.Marshal(cts)
	if err != nil {
		log.Println("error while marshaling customization config", err)
		return err
	}
	err = os.WriteFile("./customization.json", data, 0777)
	if err != nil {
		log.Println("error while writing customization config", err)
		return err
	}
	return nil
}
func ChangeTemplate(temp string) error {
	var cts Customization
	dataBytes, err := os.ReadFile("./customization.json")
	if err != nil {
		log.Println("error while reading customization config", err)
		return err
	}
	err = json.Unmarshal(dataBytes, &cts)
	if err != nil {
		log.Println("error while unmarshal customization config", err)
		return err
	}
	cts.Template = temp
	data, err := json.Marshal(cts)
	if err != nil {
		log.Println("error while marshaling customization config", err)
		return err
	}
	err = os.WriteFile("./customization.json", data, 0777)
	if err != nil {
		log.Println("error while writing customization config", err)
		return err
	}
	return nil
}
func GetClientCustomozations() (Customization, error) {
	var cts Customization
	dataBytes, err := os.ReadFile("./customization.json")
	if err != nil {
		log.Println("error while reading customization config", err)
		return cts, err
	}
	err = json.Unmarshal(dataBytes, &cts)
	if err != nil {
		log.Println("error while unmarshal customization config", err)
		return cts, err
	}

	return cts, nil
}
