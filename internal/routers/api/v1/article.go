/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        article.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 2:07 下午
 */

package v1

import "github.com/gin-gonic/gin"

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {}

func (a Article) List(c *gin.Context) {}

func (a Article) Create(c *gin.Context) {}

func (a Article) Update(c *gin.Context) {}

func (a Article) Delete(c *gin.Context) {}
