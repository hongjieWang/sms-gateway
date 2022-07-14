package models

import (
	"go-admin/common/models"
)

type TArticle struct {
    models.Model
    
    ArticleNo string `json:"articleNo" gorm:"type:varchar(32);comment:文章编号"` 
    ArticleType string `json:"articleType" gorm:"type:varchar(32);comment:文章分类"` 
    ArticleName string `json:"articleName" gorm:"type:varchar(100);comment:文章名称"` 
    ArticleTitle string `json:"articleTitle" gorm:"type:varchar(100);comment:文章标题"` 
    ArticleAuthor string `json:"articleAuthor" gorm:"type:varchar(100);comment:文章作者"` 
    ArticleSubtitle string `json:"articleSubtitle" gorm:"type:varchar(200);comment:副标题"` 
    ArticleText string `json:"articleText" gorm:"type:text;comment:文章内容HTML"` 
    Status string `json:"status" gorm:"type:bit(10);comment:状态"` 
    models.ModelTime
    models.ControlBy
}

func (TArticle) TableName() string {
    return "t_article"
}

func (e *TArticle) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TArticle) GetId() interface{} {
	return e.Id
}