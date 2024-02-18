package main

import (
	_ "embed"
	"fmt"
	"github.com/energye/systray"
	"github.com/energye/systray/icon"
	"github.com/skratchdot/open-golang/open"
	"os"
	"path/filepath"
	"time"
)

//go:embed default.ico
var DefaultIco []byte

type WebAppTray struct {
	IconPath *string
}

func (w *WebAppTray) start() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}
	systray.Run(w.onReady, onExit)
}

func (w *WebAppTray) onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Tray 托盘")
	systray.SetTooltip("左键打开,右键退出")
	//左键点击 - 打开
	systray.SetOnClick(func(menu systray.IMenu) {
		open.Run(url)
	})
	// 右键单击 - 退出
	systray.SetOnRClick(func(menu systray.IMenu) { systray.Quit() })
	w.setIco(w.IconPath)
}

func (w *WebAppTray) setIco(icoPath *string) {
	if icoPath == nil || *icoPath == "" {
		// 使用默认的图标
		systray.SetIcon(DefaultIco)
	} else {
		// 使用自定义的ico图标
		icoData, err := os.ReadFile(filepath.Join(*icoPath))
		if err != nil {
			fmt.Println("Failed read ico :", err)
		} else {
			systray.SetIcon(icoData)
		}
	}
}
