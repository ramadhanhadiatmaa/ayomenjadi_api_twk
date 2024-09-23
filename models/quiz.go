package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Question    string `json:"question"`
	Questionurl string `json:"questionurl"`
	Answurl     string `json:"answurl"`
	Answ        string `json:"answer"`
	Optiona     string `json:"optiona"`
	Optionb     string `json:"optionb"`
	Optionc     string `json:"optionc"`
	Optiond     string `json:"optiond"`
	Optione     string `json:"optione"`
	Optionurla  string `json:"optionurla"`
	Optionurlb  string `json:"optionurlb"`
	Optionurlc  string `json:"optionurlc"`
	Optionurld  string `json:"optionurld"`
	Optionurle  string `json:"optionurle"`
	Result      string `json:"result"`
	Tipe        string `json:"tipe"`
	Val         bool   `json:"value"`
}
