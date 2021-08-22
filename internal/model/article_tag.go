/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        article_tag.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 1:32 下午
 */

package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
