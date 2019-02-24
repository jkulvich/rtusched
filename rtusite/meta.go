package rtusite

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
)

//ReadMeta - получает информацию о институтах и ссылки на документы с расписанием
func ReadMeta(logx *logrus.Logger, site string) ([]Meta, error) {
	const startParseNodePrefix = "Место проведения занятий:"
	var meta = make([]Meta, 0)

	resp, err := http.Get(site)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected http status 200, but got: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	regexpTimeFinder := regexp.MustCompile(`(\d{1,2}:\d{2})`)

	// Поиск компонентов страницы
	doc.Find("#tab-content .uk-active").
		Find("a[uk-toggle]").
		Each(func(i int, institute *goquery.Selection) {

			// Название института
			name := institute.Text()

			// Местоположение кампусов
			places := make([]string, 0)
			placesHtml, err := institute.Next().Find("div").First().Find("div").Eq(1).Html()
			if err != nil {
				logx.Warnf(`can't read places for "%s": %s`, name, err)
			} else {
				places = strings.Split(placesHtml, "<br/>")
				for i, place := range places {
					places[i] = strings.TrimSpace(place)
				}
			}

			// Расписание звонков
			bells := make([]LessonBell, 0)
			bellsHtml, err := institute.Next().Find("div").Eq(3).Find("div").Eq(1).Html()
			if err != nil {
				logx.Warnf(`can't read bells for "%s": %s`, name, err)
			} else {
				bellsLines := strings.Split(bellsHtml, "<br/>")
				for i, bellLine := range bellsLines {
					times := regexpTimeFinder.FindAllStringSubmatch(bellLine, -1)
					if len(times) != 2 {
						if len(times) != 0 {
							logx.Warnf(`can't parse bells time for: "%s", expected 2 times for bell %d, got: %d`, name, i, len(times))
						}
					} else {
						bells = append(bells, LessonBell{
							Start: times[0][0],
							End:   times[1][0],
						})
					}
				}
			}

			// Расписание занятий
			forms := make([]LessonForm, 0)
			institute.Next().Find("div").Eq(10).Find("tbody tr").Each(func(_ int, lessonForm *goquery.Selection) {
				columns := lessonForm.Find("td")
				form := columns.First().Text()
				links := make([]string, columns.Length()-1)
				for j := 1; j < columns.Length(); j++ {
					links[j-1], _ = columns.Eq(j).Find("a").Attr("href")
				}
				forms = append(forms, LessonForm{
					Form:  form,
					Links: links,
				})
			})

			meta = append(meta, Meta{
				Name:   name,
				Places: places,
				Bells:  bells,
				Forms:  forms,
			})
		})

	return meta, nil
}
