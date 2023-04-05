package controllers

import (
	"github.com/dsxg666/notebook/internal/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type BackController struct{}

func (ctl BackController) Login(c *gin.Context) {
	c.HTML(200, "back/login.html", nil)
}

func (ctl BackController) LoginP(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "xxx" && password == "xxxx" {
		session.Set("isLogin", "true")
		_ = session.Save()
		c.Redirect(301, "/admin")
	} else {
		c.Redirect(301, "/login")
	}
}

func (ctl BackController) Admin(c *gin.Context) {
	session := sessions.Default(c)
	isI := session.Get("isLogin")
	if isI == nil {
		c.HTML(200, "back/login.html", nil)
	} else {
		is := isI.(string)
		if is == "true" {
			c.HTML(200, "back/admin.html", nil)
		} else {
			c.HTML(200, "back/login.html", nil)
		}
	}
}

func (ctl BackController) Editor(c *gin.Context) {
	session := sessions.Default(c)
	isI := session.Get("isLogin")
	if isI == nil {
		c.HTML(200, "back/login.html", nil)
	} else {
		is := isI.(string)
		if is == "true" {
			id := c.PostForm("id")
			a := db.Article{Id: id}
			a2 := a.Select()
			c.HTML(200, "back/editor.html", gin.H{"content": a2.Content, "id": id, "title": a2.Title})
		} else {
			c.HTML(200, "back/login.html", nil)
		}
	}
}

func (ctl BackController) AddArticle(c *gin.Context) {
	session := sessions.Default(c)
	isI := session.Get("isLogin")
	if isI == nil {
		c.HTML(200, "back/login.html", nil)
	} else {
		is := isI.(string)
		if is == "true" {
			category := c.PostForm("category")
			title := c.PostForm("title")
			content := c.PostForm("content")
			a := db.Article{Title: title, Content: content, Category: category}
			a.Add()
			c.JSON(http.StatusOK, nil)
		} else {
			c.HTML(200, "back/login.html", nil)
		}
	}
}

func (ctl BackController) ModifyArticle(c *gin.Context) {
	session := sessions.Default(c)
	isI := session.Get("isLogin")
	if isI == nil {
		c.HTML(200, "back/login.html", nil)
	} else {
		is := isI.(string)
		if is == "true" {
			id := c.PostForm("id")
			title := c.PostForm("title")
			content := c.PostForm("content")
			a := db.Article{Title: title, Content: content, Id: id}
			a.Modify()
			c.JSON(http.StatusOK, nil)
		} else {
			c.HTML(200, "back/login.html", nil)
		}
	}
}

func (ctl BackController) SendPicture(c *gin.Context) {
	session := sessions.Default(c)
	isI := session.Get("isLogin")
	if isI == nil {
		c.HTML(200, "back/login.html", nil)
	} else {
		is := isI.(string)
		if is == "true" {
			file, _ := c.FormFile("file")
			dst := path.Join("./static/img/upload", file.Filename)
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, nil)
		} else {
			c.HTML(200, "back/login.html", nil)
		}
	}
}
