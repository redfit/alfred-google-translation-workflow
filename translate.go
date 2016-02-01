package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/json"

	"github.com/spf13/viper"
	"github.com/codegangsta/cli"
	"github.com/pascalw/go-alfred"
)

func cmdTranslate(c *cli.Context) {

	if len(c.Args()) < 2 { return }

	query := c.Args()
	alfred.InitTerms(query)

	sourceLang := query[0]
	targetLang := query[1]
	text := strings.Join(query[2:], " ")

	result := translate(sourceLang, targetLang, text)
	response := alfred.NewResponse()

	for i, item := range result.Data.Translations {
		response.AddItem(&alfred.AlfredResponseItem{
			Valid: true,
			Uid: strconv.Itoa(i),
			Title: item.TranslatedText,
			Arg: strings.Join(query, " "),
			Subtitle: "翻訳結果",
		})
	}
	response.Print()
}

type Translations struct {
	TranslatedText string `json:"translatedText"`
}

type Data struct {
	Translations []Translations `json:"translations"`
}

type JsonData struct {
	Data Data `json:"data"`
}

const baseUrl = "https://www.googleapis.com/language/translate/v2?q=%s&source=%s&target=%s&key=%s"

func buildUrl(source string, target string, text string) string {
	return fmt.Sprintf(baseUrl,
		url.QueryEscape(text), source, target, viper.GetString("accessToken"))
}

func translate(source string, target string, text string) JsonData {
	resp, err := http.Get(buildUrl(source, target, text))

	if err != nil { panic(err) }

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var jsonData JsonData
	json.Unmarshal(body, &jsonData)

	return jsonData
}
