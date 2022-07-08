package gateway

import (
	"github.com/semirm-dev/seeba/etl"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetMusic(search etl.Search) gin.HandlerFunc {
	return func(c *gin.Context) {
		musicData, err := search.All()
		if err != nil {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(
			http.StatusOK,
			musicData,
		)
	}
}
