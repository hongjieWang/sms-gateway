package dto

import ("time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TArticleGetPageReq struct {
	dto.Pagination     `search:"-"`
    TArticleOrder
}

type TArticleOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:t_article"`
    ArticleNo string `form:"articleNoOrder"  search:"type:order;column:article_no;table:t_article"`
    ArticleType string `form:"articleTypeOrder"  search:"type:order;column:article_type;table:t_article"`
    ArticleName string `form:"articleNameOrder"  search:"type:order;column:article_name;table:t_article"`
    ArticleTitle string `form:"articleTitleOrder"  search:"type:order;column:article_title;table:t_article"`
    ArticleAuthor string `form:"articleAuthorOrder"  search:"type:order;column:article_author;table:t_article"`
    ArticleSubtitle string `form:"articleSubtitleOrder"  search:"type:order;column:article_subtitle;table:t_article"`
    ArticleText string `form:"articleTextOrder"  search:"type:order;column:article_text;table:t_article"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:t_article"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:t_article"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:t_article"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:t_article"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:t_article"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:t_article"`
    
}

func (m *TArticleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TArticleInsertReq struct {
    Id int `json:"-" comment:"主键"` // 主键
    ArticleNo string `json:"articleNo" comment:"文章编号"`
    ArticleType string `json:"articleType" comment:"文章分类"`
    ArticleName string `json:"articleName" comment:"文章名称"`
    ArticleTitle string `json:"articleTitle" comment:"文章标题"`
    ArticleAuthor string `json:"articleAuthor" comment:"文章作者"`
    ArticleSubtitle string `json:"articleSubtitle" comment:"副标题"`
    ArticleText string `json:"articleText" comment:"文章内容HTML"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *TArticleInsertReq) Generate(model *models.TArticle)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ArticleNo = s.ArticleNo
    model.ArticleType = s.ArticleType
    model.ArticleName = s.ArticleName
    model.ArticleTitle = s.ArticleTitle
    model.ArticleAuthor = s.ArticleAuthor
    model.ArticleSubtitle = s.ArticleSubtitle
    model.ArticleText = s.ArticleText
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *TArticleInsertReq) GetId() interface{} {
	return s.Id
}

type TArticleUpdateReq struct {
    Id int `uri:"id" comment:"主键"` // 主键
    ArticleNo string `json:"articleNo" comment:"文章编号"`
    ArticleType string `json:"articleType" comment:"文章分类"`
    ArticleName string `json:"articleName" comment:"文章名称"`
    ArticleTitle string `json:"articleTitle" comment:"文章标题"`
    ArticleAuthor string `json:"articleAuthor" comment:"文章作者"`
    ArticleSubtitle string `json:"articleSubtitle" comment:"副标题"`
    ArticleText string `json:"articleText" comment:"文章内容HTML"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *TArticleUpdateReq) Generate(model *models.TArticle)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.ArticleNo = s.ArticleNo
    model.ArticleType = s.ArticleType
    model.ArticleName = s.ArticleName
    model.ArticleTitle = s.ArticleTitle
    model.ArticleAuthor = s.ArticleAuthor
    model.ArticleSubtitle = s.ArticleSubtitle
    model.ArticleText = s.ArticleText
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *TArticleUpdateReq) GetId() interface{} {
	return s.Id
}

// TArticleGetReq 功能获取请求参数
type TArticleGetReq struct {
     Id int `uri:"id"`
}
func (s *TArticleGetReq) GetId() interface{} {
	return s.Id
}

// TArticleDeleteReq 功能删除请求参数
type TArticleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TArticleDeleteReq) GetId() interface{} {
	return s.Ids
}