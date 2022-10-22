package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"github.com/2mf8/GoWeb/data"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func PriceAddAndUpdateByItemApi(c *gin.Context) {
	citem := c.Param("item")
	cp, err := data.GetItem(citem)
	if err != nil {
		len := c.Request.ContentLength
		body := make([]byte, len)
		c.Request.Body.Read(body)
		json.Unmarshal(body, &cp)
		cid := cp.Id
		if cp.Item == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "item不能为空",
			})
		}else{
			cu, err := data.GetItem(cp.Item)
			gid := cu.Id
			if err != nil {
				err = cp.ItemCreate()
				if err != nil {
					msg := fmt.Sprintf("创建%s失败", cp.Item)
					c.JSON(http.StatusExpectationFailed, gin.H{
						"msg": msg,
					})
				} else {
					msg := fmt.Sprintf("创建%s成功", cp.Item)
					c.JSON(http.StatusOK, gin.H{
						"msg": msg,
					})
				}
			} else {
				if cid == gid {
					err = cp.ItemUpdate()
				}else{
					err1 := cu.ItemDeleteById()
					fmt.Println(err1)
					err = cp.ItemUpdate()
				}
				if err != nil {
					msg := fmt.Sprintf("更新%s失败", cp.Item)
					c.JSON(http.StatusExpectationFailed, gin.H{
						"msg": msg,
					})
				}else{
					msg := fmt.Sprintf("更新%s成功", cp.Item)
					c.JSON(http.StatusOK, gin.H{
						"msg": msg,
					})
				}
			}
		}
	}else{
		len := c.Request.ContentLength
		body := make([]byte, len)
		c.Request.Body.Read(body)
		json.Unmarshal(body, &cp)
		err = cp.ItemUpdate()
		if err != nil {
			msg := fmt.Sprintf("更新%s失败", cp.Item)
			c.JSON(http.StatusExpectationFailed, gin.H{
				"msg": msg,
			})
		} else {
			msg := fmt.Sprintf("更新%s成功", cp.Item)
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
		}
	}
}

func PriceDeleteByItemApi(c *gin.Context) {
	citem := c.Param("item")
	cp, err := data.GetItem(citem)
	if err != nil {
		fmt.Println(err)
		msg := fmt.Sprintf("获取%s失败", cp.Item)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": msg,
		})
	}else{
		err = cp.ItemDelete()
		if err != nil {
			msg := fmt.Sprintf("删除%s失败", cp.Item)
			c.JSON(http.StatusExpectationFailed, gin.H{
				"msg": msg,
			})
		}else{
			msg := fmt.Sprintf("删除%s成功", cp.Item)
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
		}
	}
}

func PriceGetItemApi(c *gin.Context) {
	citem, _ := url.QueryUnescape(c.Param("item"))
	cp, err := data.GetItem(citem)
	if err != nil {
		fmt.Println(err)
		msg := fmt.Sprintf("获取%s失败", citem)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": msg,
		})
	}else{
	/*op, err := json.MarshalIndent(&cp, "", "\t")
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "转换为JSON失败",
		})
	}
	fmt.Println(string(op))*/
		c.JSON(http.StatusOK, gin.H{
			"data": cp,
		})
	}
}

func PriceGetItemsApi(c *gin.Context) {
	citem := c.Param("key")
	cp, err := data.GetItems(citem)
	//fmt.Println(cp)
	if err != nil {
		fmt.Println(err)
		msg := fmt.Sprintf("获取%s失败", citem)
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": msg,
		})
	}else{
	/*op, err := json.MarshalIndent(&cp, "", "\t")
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "转换为JSON失败",
		})
	}
	fmt.Println(string(op))*/
		c.JSON(http.StatusOK, gin.H{
			"data": cp,
		})
	}
}
