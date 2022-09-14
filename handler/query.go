package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TurboHsu/Posta/sql"
	"github.com/gin-gonic/gin"
)

func QueryTest(c *gin.Context) {
	var opt []byte
	raw := sql.RunQuery(c.Query("q"))
	opt, _ = json.Marshal(raw)
	c.String(http.StatusOK, string(opt))
}
