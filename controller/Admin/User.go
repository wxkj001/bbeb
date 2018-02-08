package Admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/wxkj001/bbeb/core"
)

type User struct {
	*core.FrameWork
}


func (f *User) user(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"a": "1"})
}
