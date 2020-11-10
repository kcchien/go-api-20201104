package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

func main() {

	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

// Get 取得全部資料
func Get(c *gin.Context) {
	// Just return all data
	c.JSON(http.StatusOK, Data)
}

// GetOne 取得單一筆資料
func GetOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Return filtered data
	c.JSON(http.StatusOK, Filter(Data, id))
}

// Post 新增資料
func Post(c *gin.Context) {
	var p Role

	// Sort Data slice first, make sure the orders was corrected
	sort.SliceStable(Data, func(i, j int) bool {
		return Data[i].ID < Data[j].ID
	})

	// Get last element's ID in slice and plus 1 for new data ID
	p.ID = Data[len(Data)-1].ID + 1

	c.ShouldBind(&p)

	Data = append(Data, p)

	c.JSON(http.StatusOK, p)
}

// Put 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	i := IndexOf(Data, id)

	if i > 0 {
		var p Role
		c.ShouldBind(&p)
		Data[i].Name = p.Name
		Data[i].Summary = p.Summary
		c.JSON(http.StatusOK, Data[i])
	} else {
		c.JSON(http.StatusOK, "")
	}
}

// Delete 刪除資料
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	i := IndexOf(Data, id)

	if i > 0 {
		Data = append(Data[:i], Data[i+1:]...)
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusNoContent, "")
	}
}

// Filter Filter correspond role
func Filter(d []Role, id int) []Role {
	result := make([]Role, 0)

	for _, v := range d {
		if v.ID == uint(id) {
			result = append(result, v)
		}
	}

	return result
}

// IndexOf Return index of Role slice
func IndexOf(d []Role, id int) int {

	result := -1

	for i, role := range d {
		if role.ID == uint(id) {
			return int(i)
		}
	}

	return result
}
