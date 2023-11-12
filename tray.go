package main

import (
	"fmt"
	"github.com/energye/systray"
	"github.com/energye/systray/icon"
	"github.com/skratchdot/open-golang/open"
	"os"
	"path/filepath"
	"time"
)

func startTray() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Tray 托盘")
	systray.SetTooltip("左键打开,右键退出")
	//左键点击 - 打开
	systray.SetOnClick(func(menu systray.IMenu) {
		open.Run(url)
	})
	// 右键单击 - 退出
	systray.SetOnRClick(func(menu systray.IMenu) { systray.Quit() })
}

func setIco(icoPath string) {
	icoData, err := os.ReadFile(filepath.Join(icoPath))
	if err != nil {
		fmt.Println("Failed read ico :", err)
	} else {
		systray.SetIcon(icoData)
	}
}
