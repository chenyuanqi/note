package controllers

import (
	"blog/app/models/article"
	"blog/app/models/category"
	"blog/app/requests"
	"blog/pkg/flash"
	"blog/pkg/route"
	"blog/pkg/view"

	"fmt"
	"net/http"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
	BaseController
}

// Create 文章分类创建页面
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create", "categories._form_field")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// 2. 表单验证
	errors := requests.ValidateCategoryForm(_category)

	// 3. 检测错误
	if len(errors) == 0 {
		// 创建文章分类
		_category.Create()
		if _category.ID > 0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
			// indexURL := route.Name2URL("categories.show", "id", _category.GetStringID())
			// http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

// Show 显示分类下的文章列表
func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的数据
	_category, err := category.Get(id)

	// 3. 获取结果集
	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)

	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {

		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

func (cc *CategoriesController) Edit(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_category, err := category.Get(id)

	// 3. 如果出现错误
	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		// 4. 读取成功，显示编辑文章表单
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   view.D{},
		}, "categories.edit", "categories._form_field")
	}
}

// Update 更新文章
func (cc *CategoriesController) Update(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_category, err := category.Get(id)

	// 3. 如果出现错误
	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		// 4. 未出现错误
		// 4.1 表单验证
		_category.Name = r.PostFormValue("name")

		errors := requests.ValidateCategoryForm(_category)
		if len(errors) == 0 {
			// 4.2 表单验证通过，更新数据
			rowsAffected, err := _category.Update()

			if err != nil {
				// 数据库错误
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
				return
			}

			// √ 更新成功，跳转到文章详情页
			if rowsAffected > 0 {
				flash.Warning("分类更新成功！")
				http.Redirect(w, r, route.Name2URL("home"), http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {
			// 4.3 表单验证不通过，显示理由
			view.Render(w, view.D{
				"Category": _category,
				"Errors":   errors,
			}, "categories.edit", "categories._form_field")
		}
	}
}
