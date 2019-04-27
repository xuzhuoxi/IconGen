//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package main

import (
	"github.com/xuzhuoxi/IconGen/src/lib"
	_ "github.com/xuzhuoxi/IconGen/src/lib/png"
	_ "github.com/xuzhuoxi/IconGen/src/lib/jpeg"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"os"
	"strconv"
	"fmt"
	"image/jpeg"
)

func main() {
	logger := logx.NewLogger()
	logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	cfg, err := lib.ParseFlag()
	if err != nil {
		logger.Error(err)
		return
	}
	if !osxu.IsExist(cfg.InPath) {
		logger.Error("InPath does net Exist! ")
		return
	}
	handle := func(filePath string) {
		_, fileName := osxu.SplitFilePath(filePath)
		baseName, extName := osxu.SplitFileName(fileName)
		img, err := lib.LoadImage(filePath)
		if nil != err {
			return
		}
		fm := cfg.OutFormat
		if "" == fm {
			fm = extName
		}
		for _, size := range cfg.OutSizes {
			newImg, _ := lib.ResizeImage(img, uint(size), uint(size))
			sizeStr := strconv.Itoa(size)
			fileName := fmt.Sprintf("%s_%sx%s.%s", baseName, sizeStr, sizeStr, fm)
			lib.SaveImage(newImg, cfg.OutPath+fileName, lib.ImageFormat(fm), &jpeg.Options{Quality: cfg.OutRatio})
		}
	}
	if !osxu.IsFolder(cfg.InPath) {
		handle(cfg.InPath)
	} else {
		list, err := osxu.GetFolderFileList(cfg.InPath, false, func(fileInfo os.FileInfo) bool {
			extName := osxu.GetExtensionName(fileInfo.Name())
			if !lib.CheckFormat(extName) {
				return false
			}
			return true
		})
		if nil != err {
			logger.Error(err)
			return
		}
		for _, file := range list {
			handle(file.FullPath())
		}
	}
}