package version

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
const versionNumber = "0.1.1"
const versionName = "Hanover"
const comment = "Testing version endpoint"

var versionDate = "January 2022"
var launchDate = time.Now().Format(time.ANSIC)

var response gin.H = gin.H{"Number": versionNumber,
	"Name": versionName, "Date": versionDate, "LaunchDate": launchDate,
	"Comment": comment}

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, response)
}
