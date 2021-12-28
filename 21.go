package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Структура, описывающая формат данных для примера
type Cat struct {
	Name string
	Age  int
}

// Структура storage может хранить какие-то данные
type storage struct {
	data []byte
}

// storeXMLData записывает xml в storage
func (s *storage) storeXMLData(xg XMLGetter) error {
	data, err := xg.GetXML()
	if err != nil {
		return err
	}
	s.data = data
	return nil
}

// Интерфейс получения данных XML
type XMLGetter interface {
	GetXML() ([]byte, error)
}

type xmlData struct {
	xmlBlob []byte
}

func (x *xmlData) GetXML() ([]byte, error) {
	return x.xmlBlob, nil
}

// Структура с данными Json
type jsonData struct {
	jsonBlob []byte
}

func (j *jsonData) GetJson() []byte {
	return j.jsonBlob
}

// Адаптер, реализующий XMLGetter
type jsonAdapter struct {
	*jsonData
}

// Вызывает GetJson у jsonData и возвращает XML
func (ja *jsonAdapter) GetXML() ([]byte, error) {
	var s Cat
	jsonBlob := ja.GetJson()
	if err := json.Unmarshal(jsonBlob, &s); err != nil {
		return nil, err
	}
	res, err := xml.Marshal(&s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {

	storage := &storage{}

	xd := &xmlData{
		[]byte(`<Cat><Name>pushok</Name><Age>3</Age></Cat>`),
	}

	storage.storeXMLData(xd)
	fmt.Println(string(storage.data))

	jd := &jsonData{
		[]byte(`{"Name":"barsik", "Age":5}`),
	}
	ja := &jsonAdapter{jd}

	storage.storeXMLData(ja)
	fmt.Println(string(storage.data))

}
