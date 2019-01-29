package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

type Coord struct {
	XMLName xml.Name `xml:"coord" json:"-"`
	Lon     string   `xml:"lon,attr" json:"lon"`
	Lat     string   `xml:"lat,attr" json:"lat"`
}

type Sun struct {
	XMLName xml.Name `xml:"sun" json:"-"`
	Rise    string   `xml:"rise,attr" json:"rise"`
	Set     string   `xml:"set,attr" json:"set"`
}

type City struct {
	XMLName xml.Name `xml:"city" json:"-"`
	Id      string   `xml:"id,attr" json:"id"`
	Name    string   `xml:"name,attr" json:"name"`
	Coord   Coord    `xml:"coord" json:"coord"`
	Country string   `xml:"country" json:"country"`
	Sun     Sun      `xml:"sun" json:"sun"`
}

type Temperature struct {
	XMLName xml.Name `xml:"temperature" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Min     string   `xml:"min,attr" json:"min"`
	Max     string   `xml:"max,attr" json:"max"`
	Unit    string   `xml:"unit,attr" json:"unit"`
}

type Humidity struct {
	XMLName xml.Name `xml:"humidity" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Unit    string   `xml:"unit,attr" json:"unit"`
}

type Pressure struct {
	XMLName xml.Name `xml:"pressure" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Unit    string   `xml:"unit,attr" json:"unit"`
}

type Speed struct {
	XMLName xml.Name `xml:"speed" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Name    string   `xml:"name,attr" json:"name"`
}

type Directtion struct {
	XMLName xml.Name `xml:"direction" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Code    string   `xml:"code,attr" json:"code"`
	Name    string   `xml:"name,attr" json:"name"`
}

type Wind struct {
	XMLName   xml.Name   `xml:"wind" json:"-"`
	Speed     Speed      `xml:"speed" json:"speed"`
	Direction Directtion `xml:"direction" json:"direction"`
}

type Clouds struct {
	XMLName xml.Name `xml:"clouds" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
	Name    string   `xml:"name,attr" json:"name"`
}

type Visibility struct {
	XMLName xml.Name `xml:"visibility" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
}

type Precipitation struct {
	XMLName xml.Name `xml:"precipitation" json:"-"`
	Mode    string   `xml:"mode,attr" json:"mode"`
}

type Weather struct {
	XMLName xml.Name `xml:"weather" json:"-"`
	Number  string   `xml:"number,attr" json:"number"`
	Value   string   `xml:"value,attr" json:"value"`
	Icon    string   `xml:"icon,attr" json:"icon"`
}

type Lastupdate struct {
	XMLName xml.Name `xml:"lastupdate" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
}

type Current struct {
	XMLName       xml.Name      `xml:"current" json:"-"`
	City          City          `xml:"city" json:"city"`
	Temperature   Temperature   `xml:"temperature" json:"temperature"`
	Humidity      Humidity      `xml:"humidity" json:"humidity"`
	Pressure      Pressure      `xml:"pressure" json:"pressure"`
	Wind          Wind          `xml:"wind" json:"wind"`
	Clouds        Clouds        `xml:"clouds" json:"clouds"`
	Visibility    Visibility    `xml:"visibility" json:"visibility"`
	Precipitation Precipitation `xml:"precipitation" json:"precipitation"`
	Weather       Weather       `xml:"weather" json:"weather"`
	Lastupdate    Lastupdate    `xml:"lastupdate" json:"lastupdate"`
}

func (c Current) ToJSON() []byte {
	jsonData, _ := json.MarshalIndent(c, "", "  ")
	return jsonData
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arg := os.Args[1:]

	xmlFile, err := ioutil.ReadFile(arg[0])
	check(err)

	c := Current{}
	xml.Unmarshal(xmlFile, &c)
	jsonDat := c.ToJSON()
	jsonPath := strings.TrimSuffix(arg[0], ".xml") + ".json"
	err = ioutil.WriteFile(jsonPath, jsonDat, 0644)
	check(err)
}
