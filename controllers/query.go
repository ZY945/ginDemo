package controllers

import (
	"GinAndSqlx/dao/mysql"
	"GinAndSqlx/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserVoById(context *gin.Context) {
	idStr := context.Query("id")
	id, _ := strconv.Atoi(idStr)
	vo := mysql.SqlxqueryByGet(id)
	context.JSON(http.StatusOK, vo)
}

func List(context *gin.Context) {
	_, vos, _ := mysql.SqlxList()
	fmt.Println(vos)
	context.JSON(http.StatusOK, vos)
}

func InsertUser(context *gin.Context) {
	data, _ := context.GetRawData()
	var bo *models.AddUserBo
	//var bo map[string]interface{}
	err := json.Unmarshal(data, &bo)
	if err != nil {
		log.Print("json error:\n", err)
		return
	}
	newId, err := mysql.Insert(bo)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"new Id":  newId,
		"bo":      bo,
	})
}

func UpdateUser(context *gin.Context) {
	data, _ := context.GetRawData()
	var bo *models.UpdateUserBo
	//var bo map[string]interface{}
	err := json.Unmarshal(data, &bo)
	if err != nil {
		log.Print("json error:\n", err)
		return
	}
	mysql.Update(bo)
	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"newBo":   bo,
	})
}

func DelUser(context *gin.Context) {
	idStr := context.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := mysql.Del(id)
	if err != nil {
		log.Print("del fail")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "del success",
		"id":      id,
	})
}
