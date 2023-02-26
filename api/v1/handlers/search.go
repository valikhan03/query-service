package handlers

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	searchQuery := c.Query("q")
	p := c.Query("p")
	var page int
	var err error
	if p != ""{
		page, err = strconv.Atoi(p)
		if err != nil{
			log.Println(err)
			page = 1
		}
	} else{
		page = 1
	}
	
	ctx := context.TODO();
	res, err := h.service.Search(ctx, searchQuery, page)
	if err != nil{
		log.Println(err)
		c.AbortWithError(500, err)
	}

	c.JSON(200, res)
}