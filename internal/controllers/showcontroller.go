package controllers

import (
	"fmt"
	"github.com/dsxg666/notebook/global"
	"github.com/dsxg666/notebook/internal/db"
	"github.com/dsxg666/notebook/pkg/email"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ShowController struct{}

func (ctl ShowController) Home(c *gin.Context) {
	c.HTML(200, "show/home.html", nil)
}

func (ctl ShowController) Contact(c *gin.Context) {
	c.HTML(200, "show/contact.html", nil)
}

func (ctl ShowController) Intro(c *gin.Context) {
	c.HTML(200, "show/author.html", nil)
}

func (ctl ShowController) SendMsg(c *gin.Context) {
	name := c.PostForm("name")
	subject := c.PostForm("subject")
	content := c.PostForm("content")
	hisMail := c.PostForm("email")
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	err := mailer.SendMail("2637046983@qq.com", subject, fmt.Sprintf("发送人名称: %s ;邮箱: %s ;内容: %s",
		name, hisMail, content))
	if err != nil {
		global.SugarLogger.Error("mailer.SendMail err:", err)
	}
	c.JSON(http.StatusOK, nil)
}

func handle(category string) []db.Article {
	a := db.Article{Category: category}
	arr := a.SelectAll()
	for i := 0; i < len(arr); i++ {
		arr[i].Time = strings.ReplaceAll(arr[i].Time[:10], "-", "/")
	}
	return arr
}

func (ctl ShowController) Category(c *gin.Context) {
	id := c.Param("id")
	if id == "1" {
		arr := handle("Golang")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Golang",
			"categoryB": "技术类",
			"intro":     "Go是富有表现力的、简洁的、干净的、高效的。它的并发机制使编写多核和网络机器的程序变得容易，而它的新型系统使程序结构灵活和模块化。Go可以快速编译为机器码，同时还具有垃圾收集的便利性和运行时反射的能力。它是一种快速的、静态类型的编译语言，但感觉上像是一种动态类型的解释型语言。",
		})
	} else if id == "2" {
		arr := handle("Blockchain")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Blockchain",
			"categoryB": "技术类",
			"intro":     "去中心化？点对点？账本？数据库？一起走进区块链！",
		})
	} else if id == "3" {
		arr := handle("Linux")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Linux",
			"categoryB": "技术类",
			"intro":     "Linux，全称GNU/Linux，是一种免费使用和自由传播的类UNIX操作系统，其内核由林纳斯·本纳第克特·托瓦兹于1991年10月5日首次发布，它主要受到Minix和Unix思想的启发，是一个基于POSIX的多用户、多任务、支持多线程和多CPU的操作系统。",
		})
	} else if id == "4" {
		arr := handle("Tour")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "旅游",
			"categoryB": "生活类",
			"intro":     "旅行可以让自己变得更美好。在旅行的同时，逆增长的是见识，充实的是灵魂，会更加完善自己，眼界开阔，心思透亮。所见所闻会让你懵懂的灵魂变得知性、豁达。",
		})
	} else if id == "5" {
		arr := handle("MyArticle")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "我的文章",
			"categoryB": "生活类",
			"intro":     "学疏才浅，拙作数篇，希望喜欢，嘻嘻嘻😁",
		})
	} else if id == "6" {
		arr := handle("Painting")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "我是大画家",
			"categoryB": "生活类",
			"intro":     "画画使心灵安宁，画画是一件陶冶情操的事情，是自己和观者的对话，也是自己和内心的交互。在难得的闲暇时光,画画,静心,养气质。乃人生中的一大美事。",
		})
	} else if id == "7" {
		arr := handle("Literature")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "文学",
			"categoryB": "文学类",
			"intro":     "经典是具有典范性、权威性的经久不衰的万世之作，经过历史选择出来的最有价值的、最能表现行业的精髓的、最具代表性的、最完美的作品。",
		})
	} else if id == "8" {
		arr := handle("Frontend")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Frontend",
			"categoryB": "技术类",
			"intro":     "哈哈，虽然我是前端菜狗，但谁说后端不能有一颗称为前端大师的心呢。",
		})
	} else if id == "9" {
		arr := handle("ComputerBasis")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "ComputerBasis",
			"categoryB": "技术类",
			"intro":     "基础不牢地动山摇！基础不牢地动山摇！基础不牢地动山摇！",
		})
	} else if id == "10" {
		arr := handle("OtherTechnology")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "OtherTechnology",
			"categoryB": "技术类",
			"intro":     "技术的发展是很快的，计算机软件尤其如此，要追逐时代的步伐，就必须持续学习。",
		})
	} else if id == "11" {
		arr := handle("Java")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Java",
			"categoryB": "技术类",
			"intro":     "Java是一种高级的、基于类的、面向对象的编程语言。Java是一门非常成熟、可靠的语言，超过3亿设备使用Java运行。",
		})
	} else if id == "12" {
		arr := handle("Python")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Python",
			"categoryB": "技术类",
			"intro":     "Python是一种高级的、通用的编程语言。它的设计理念强调代码的可读性，并使用有意义的缩进。它的语言结构和面向对象的方法旨在帮助程序员为小型和大型项目编写清晰的、合乎逻辑的代码。",
		})
	} else if id == "13" {
		arr := handle("C")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "C",
			"categoryB": "技术类",
			"intro":     "C语言是一种通用的编程语言，于1972年开发，到现在仍然很受欢迎。C非常强大，它被常用于开发操作系统、数据库等等。",
		})
	} else if id == "14" {
		arr := handle("Android")
		c.HTML(200, "show/category.html", gin.H{
			"arr":       arr,
			"categoryS": "Android",
			"categoryB": "技术类",
			"intro":     "Android是Google开发的基于Linux平台的开源手机操作系统。它包括操作系统、用户界面和应用程序等移动电话工作所需的全部软件，而且不存在任何以往阻碍移动产业创新的专有权障碍。",
		})
	}
}

func (ctl ShowController) Article(c *gin.Context) {
	id := c.Param("id")
	a := db.Article{Id: id}
	a1 := a.Select()
	a.UpdateClick()
	c.HTML(200, "show/article.html", gin.H{"article": a1})
}

func (ctl ShowController) Search(c *gin.Context) {
	id := c.PostForm("id")
	a := db.Article{Id: id}
	a1 := a.Select()
	if a1.Content == "" {
		c.HTML(200, "show/home.html", gin.H{"err": "err"})
	} else {
		a.UpdateClick()
		a1.Time = strings.ReplaceAll(a1.Time[:10], "-", "/")
		c.HTML(200, "show/article.html", gin.H{"article": a1})
	}
}
