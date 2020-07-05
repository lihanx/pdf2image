package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 页面
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", gin.H{
		"code": 0,
		"message": "ok",
	})
}

// 全部页面批量转换，提供zip格式
func PDF2Zip(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"err": err.Error(),
		})
		return
	}
	// 获得文件名称
	ext := path.Ext(file.Filename)
	basename := strings.TrimSuffix(file.Filename, ext)
	// 保存上传文件
	filepath := path.Join(TEMPDIR, file.Filename)
	_ = c.SaveUploadedFile(file, filepath)
	// 创建图片目录
	dirName := fmt.Sprintf("%s_%d_%d", basename, time.Now().Unix(), rand.Intn(1000))
	imageDir := path.Join(TEMPDIR, dirName)
	if err := os.Mkdir(imageDir, os.ModePerm); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"err": err.Error(),
		})
		return
	}
	// 开始转换
	if err := pdf2images(filepath, imageDir, basename); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 2,
			"err": err.Error(),
		})
		return
	}
	// 打包图片
	zipFileName, err := zipImage(imageDir, basename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 3,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"filename": zipFileName,
	})
	return
}

// 第一页单张图片转换，提供jpg格式
func PDF2JPG(c *gin.Context) {
	file, _ := c.FormFile("file")
	// 获得文件名称
	ext := path.Ext(file.Filename)
	basename := strings.TrimSuffix(file.Filename, ext)
	// 保存上传文件
	filepath := path.Join(TEMPDIR, file.Filename)
	defer os.Remove(filepath)
	_ = c.SaveUploadedFile(file, filepath)
	// PDF -> JPG
	jpgName, err := pdf2jpg(filepath, DATADIR, basename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"filename": jpgName,
	})
}