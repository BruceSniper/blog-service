/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        tag.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 1:26 下午
 */

package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
