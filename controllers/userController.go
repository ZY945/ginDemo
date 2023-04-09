package controllers

import (
	"GinAndSqlx/dao"
	"GinAndSqlx/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUserVoById 查询一行数据(sqlx)
func GetUserVoById(context *gin.Context) {
	idStr := context.Query("id")
	id, _ := strconv.Atoi(idStr)
	vo := dao.SqlxQueryByGet(id)
	context.JSON(http.StatusOK, vo)
}

// List 查询列表(sqlx)
func List(context *gin.Context) {
	_, vos, _ := dao.SqlxList()
	fmt.Println(vos)
	context.JSON(http.StatusOK, vos)
}

// InsertUser 新增(sqlx)
func InsertUser(context *gin.Context) {
	data, _ := context.GetRawData()
	var bo *models.AddUserBo
	//var bo map[string]interface{}
	err := json.Unmarshal(data, &bo)
	if err != nil {
		log.Print("json error:\n", err)
		return
	}
	newId, err := dao.Insert(bo)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"new Id":  newId,
		"bo":      bo,
	})
}

// UpdateUser 修改(sqlx)
func UpdateUser(context *gin.Context) {
	data, _ := context.GetRawData()
	var bo *models.UpdateUserBo
	//var bo map[string]interface{}
	err := json.Unmarshal(data, &bo)
	if err != nil {
		log.Print("json error:\n", err)
		return
	}
	dao.Update(bo)
	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"newBo":   bo,
	})
}

// DelUser 删除(sqlx)
func DelUser(context *gin.Context) {
	idStr := context.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := dao.Del(id)
	if err != nil {
		log.Print("del fail")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "del success",
		"id":      id,
	})
}
func Login(context *gin.Context) {
	data, _ := context.GetRawData()
	var bo *models.LoginUserBo
	//var bo map[string]interface{}
	err := json.Unmarshal(data, &bo)
	if err != nil {
		log.Print("json error:\n", err)
		return
	}
	dao.Login(bo)
	if err != nil {
		log.Print("login fail")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":  "login success",
		"username": bo.UserName,
	})
}
