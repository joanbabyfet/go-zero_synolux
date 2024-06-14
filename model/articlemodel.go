package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		withSession(session sqlx.Session) ArticleModel
		//添加新方法
		FindAll(ctx context.Context, title string, limit int) ([]*Article, error)
		PageList(ctx context.Context, page int, page_size int, title string) ([]*Article, int, error)
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn),
	}
}

func (m *customArticleModel) withSession(session sqlx.Session) ArticleModel {
	return NewArticleModel(sqlx.NewSqlConnFromSession(session))
}

// 获取列表
func (m *customArticleModel) FindAll(ctx context.Context, title string, limit int) ([]*Article, error) {
	var resp []*Article

	//组装条件
	where := "delete_time = 0 AND status = 1 " //未删除且启用
	if len(title) > 0 {
		where += "AND title LIKE %" + title + "%"
	}

	//获取数据
	dataQuery := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d", articleRows, m.table, where, limit)
	err := m.conn.QueryRowsCtx(ctx, &resp, dataQuery)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// 获取分页列表
func (m *customArticleModel) PageList(ctx context.Context, page int, page_size int, title string) ([]*Article, int, error) {
	var resp []*Article
	total := 0 //总条数
	offset := (page - 1) * page_size

	//组装条件
	where := "delete_time = 0 AND status = 1 " //未删除且启用
	if len(title) > 0 {
		where += "AND title LIKE %" + title + "%"
	}

	//获取总条数
	countQuery := fmt.Sprintf("SELECT COUNT(*) AS `count` FROM %s WHERE %s", m.table, where)
	_ = m.conn.QueryRowCtx(ctx, &total, countQuery)

	//获取数据
	dataQuery := fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY create_time DESC LIMIT %d, %d", articleRows, m.table, where, offset, page_size)
	err := m.conn.QueryRowsCtx(ctx, &resp, dataQuery)

	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}
