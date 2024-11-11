package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 存一个全局的上下文
var thisCtx *context.Context

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	thisCtx = &a.ctx
}

// OpenUrl  用浏览器打开地址
func (a *App) OpenUrl(url string) bool {
	////截图一张全屏的
	//img, _ := screenshot.CaptureScreen()
	//// 创建一个缓冲区来保存 PNG 编码的图像
	//var buf bytes.Buffer
	//_ = png.Encode(&buf, img)
	//
	//// 将 PNG 编码的字节切片编码为 Base64 字符串
	//imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	//
	//runtime.WindowFullscreen(*thisCtx)           //全屏
	//runtime.WindowSetAlwaysOnTop(*thisCtx, true) //窗口置顶
	//return imgBase64
	runtime.BrowserOpenURL(*thisCtx, url)
	return true
}
