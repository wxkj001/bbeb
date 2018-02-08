package Admin

import (
	"github.com/wxkj001/bbeb/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	*core.FrameWork
}

func (this *Article)ArticleGet(g *gin.Context)  {
	g.HTML(http.StatusOK,"ArticleGet.html",gin.H{})
}
