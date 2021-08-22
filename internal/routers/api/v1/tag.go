/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        tag.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 2:07 下午
 */

package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

func (t Tag) List(c *gin.Context) {}

func (t Tag) Create(c *gin.Context) {}

func (t Tag) Update(c *gin.Context) {}

func (t Tag) Delete(c *gin.Context) {}
