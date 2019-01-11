// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leagues

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
)

var (
	greenOut  = color.New(color.FgGreen).SprintFunc()
	yellowOut = color.New(color.FgYellow).SprintFunc()
	redOut    = color.New(color.FgRed).SprintFunc()
)

func fetch(uri string, data url.Values) ([]byte, error) {
	u, _ := url.ParseRequestURI(uri)
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	log.Debug().Str("league", "euska").Msgf("HTTP Request URI: %s %s", urlStr, data)

	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Http request to %s failed: %s", r.URL, err.Error())
	}
	defer resp.Body.Close()
	log.Debug().Str("league", "euska").Msgf("HTTP Response Status: %s", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, fmt.Errorf("errorination happened reading the body: %s", err.Error())
	}
	return body, nil
}

// Display will fetch results and print them
func Display(uri string, challengeID string, disciplineID string, levelID string) error {
	data := url.Values{}
	log.Debug().Str("league", "euska").Msgf("Search: challenge:%s discipline:%s level:%s", challengeID, disciplineID, levelID)
	data.Add("InSel", "")
	data.Add("InCompet", challengeID)
	data.Add("InSpec", disciplineID)
	data.Add("InVille", "0")
	data.Add("InClub", "0")
	data.Add("InDate", "")
	data.Add("InDatef", "")
	data.Add("InCat", levelID)
	data.Add("InPhase", "0")
	data.Add("InPoule", "0")
	data.Add("InGroupe", "0")
	data.Add("InVoir", "Voir les résultats")

	body, err := fetch(uri, data)
	if err != nil {
		return err
	}
	z := html.NewTokenizer(strings.NewReader(string(body)))

	content := []string{"", "", "", "", ""}
	i := -1
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Club 1", "Club 2", "Score", "Commentaire"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	for {
		// token type
		tokenType := z.Next()
		if tokenType == html.ErrorToken {
			break
		}
		// token := z.Token()
		switch tokenType {
		case html.StartTagToken: // <tag>
			t := z.Token()
			if t.Data == "tr" {
				i = -1

			} else if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					if len(t.Attr) > 0 {
						if t.Attr[0].Val == "L0" || t.Attr[0].Val == "forfait" { // Text to extract
							text := (string)(z.Text())
							value := strings.TrimSpace(text)
							if len(value) > 0 {
								i = i + 1
								// fmt.Printf("%d Attr::::::::::: %s :: %s\n", i, value, t.Attr)
								if i == 0 {
									content[i] = yellowOut(value)
								} else if i == 3 {
									content[i] = redOut(value)
								} else {
									content[i] = value
								}
							}
						} else if t.Attr[0].Val == "mTitreSmall" {
							text := (string)(z.Text())
							value := strings.TrimSpace(text)
							if len(value) > 0 {
								i = i + 1
								// fmt.Printf("%d Attr::::::::::: %s :: %s\n", i, value, t.Attr)
								content[i] = greenOut(value)
							}
						}
					}
				}

			} else if t.Data == "li" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					value := strings.TrimSpace(text)
					// fmt.Printf("%s\n%s", content[i], value)
					content[i] = fmt.Sprintf("%s\n%s", content[i], value)
				}

			}
		case html.TextToken: // text between start and end tag
		case html.EndTagToken: // </tag>
			t := z.Token()
			if t.Data == "tr" {
				if len(content[0]) > 0 {
					// fmt.Printf("==> %d\n", len(content))
					// for rank, elem := range content {
					// 	fmt.Printf("%d = %s\n", rank, elem)
					// }
					table.Append(content)
					content = []string{"", "", "", "", ""}
				}
			}

		case html.SelfClosingTagToken: // <tag/>
		}
	}

	table.Render()
	return nil
}
