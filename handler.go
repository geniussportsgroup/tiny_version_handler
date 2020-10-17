package version

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var launchDate = time.Now().Format(time.ANSIC)

var response gin.H

type Info struct {
	Name    string
	Date    string
	Number  string
	Comment string
}

func init() {

	file, err := ioutil.ReadFile("version.json")
	if err != nil {
		panic("cannot open version.json")
	}

	info := Info{}
	err = json.Unmarshal([]byte(file), &info)

	response = gin.H{
		"Number":     info.Number,
		"Name":       info.Number,
		"Date":       info.Date,
		"LaunchDate": launchDate,
		"Comment":    info.Comment,
	}
}

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, response)
}
