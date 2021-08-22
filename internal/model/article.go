/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        article.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 1:29 下午
 */

package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
