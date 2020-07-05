package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"gopkg.in/gographics/imagick.v2/imagick"
	)


func InitWorkdir() {
	_ = os.Mkdir(TEMPDIR, os.ModePerm)
	_ = os.Mkdir(DATADIR, os.ModePerm)
}

func pdf2jpg(filepath, imageDir, basename string) (string, error) {
	var err error
	// 初始化
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	// 设置分辨率
	if err = mw.SetResolution(150, 150); err != nil {
		return "", err
	}
	// 读取PDF
	err = mw.ReadImage(filepath)
	mw.SetFirstIterator()
	imageName := fmt.Sprintf("%s.jpg", basename)
	imagePath := path.Join(imageDir, imageName)
	_ = mw.WriteImage(imagePath)
	return imageName, nil
}

func pdf2images(filepath, imageDir, basename string) error {
	var err error
	// 初始化
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	// 设置分辨率
	if err = mw.SetResolution(100, 100); err != nil {
		return err
	}
	// 读取PDF
	err = mw.ReadImage(filepath)
	// 开始转换
	total := int(mw.GetNumberImages())
	var wg sync.WaitGroup
	for i := 0; i < total; i++ {
		wg.Add(1)
		_ = mw.SetIteratorIndex(i)
		image := mw.GetImage()
		go func(image *imagick.MagickWand, i int) {
			//_ = image.WriteImage(path.Join(imageDir, fmt.Sprintf("%s-%d.jpg", basename, i)))
			_ = image.WriteImage(path.Join(imageDir, fmt.Sprintf("%d.jpg", i+1)))
			defer wg.Done()
		}(image, i)
	}
	wg.Wait()
	// 转换结束，删除pdf文件
	defer os.Remove(filepath)
	return nil
}

func zipImage(imageDir, basename string) (string, error) {
	var err error
	zipFileName := fmt.Sprintf("%s%s", basename, ".zip")
	zipFilepath := path.Join(DATADIR, zipFileName)
	zipFile, _ := os.OpenFile(zipFilepath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer zipFile.Close()
	writer := zip.NewWriter(zipFile)
	defer writer.Close()
	files, err := ioutil.ReadDir(imageDir)
	defer os.RemoveAll(imageDir)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		w, _ := writer.Create(file.Name())
		imagePath := path.Join(imageDir, file.Name())
		body, _ := ioutil.ReadFile(imagePath)
		_, err = w.Write(body)
	}
	return zipFileName, err
}