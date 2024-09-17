package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Qustion string `json:"qustion"`
	Qustionurl string `json:"qustionurl"`
	Answurl string `json:"answurl"`
	Answ string `json:"answer"`
	Optiona   string `json:"optiona"`
	Optionb   string `json:"optionb"`
	Optionc   string `json:"optionc"`
	Optiond   string `json:"optiond"`
	Optione   string `json:"optione"`
	Optionurla   string `json:"optionurla"`
	Optionurlb   string `json:"optionurlb"`
	Optionurlc   string `json:"optionurlc"`
	Optionurld   string `json:"optionurld"`
	Optionurle   string `json:"optionurle"`
	Tipe   int    `json:"tipe"`
}