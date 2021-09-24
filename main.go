package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"net/http"
)

const (
	css     = ".pp_last_activity_text"
	vk      = "https://m.vk.com/"
	address = ":8080"
)

var (
	onlyForAuth = errors.New("no access to page, probably only available for authenticated users")
)

type activity struct {
	Activity string `json:"activity"`
}

func getActivity(id string) (activity string, err error) {
	res, err := http.Get(vk + id)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	activity = doc.Find(css).Text()
	if activity == "" {
		return "", onlyForAuth
	}

	return activity, nil
}

func returnActivity(c echo.Context) error {
	a, err := getActivity(c.QueryParam("id"))
	if err == onlyForAuth {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, activity{a})
}

func main() {
	e := echo.New()
	e.GET("/", returnActivity)
	e.Logger.Fatal(e.Start(address))
}
