package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/hugozhu/doctorapp/lib"
	"gopkg.in/yaml.v2"
)

type Command struct {
	Command string `yaml:"command"`
}

//type Config struct {
//	Config []Command
//}

const (
	YamlFilePath = "config.yaml"
)

var command = Command{} //将用来承载yaml文件解析出的内容

func init() {
	//解析yaml文件，并根据msg中记录的dir 与yaml解析出来的command组装命令
	yamlFile, err := ioutil.ReadFile(YamlFilePath)
	if err != nil {
		log.Printf("Get error when open yaml file : %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &command)
	if err != nil {
		log.Printf("Get error when parse yaml file : %v", err)
	}
}

func main() {
	a := app.New() //初始化fyne app 并设置窗体信息
	w := a.NewWindow("HHO 商品包检查工具")
	w.SetFixedSize(true)
	w.Resize(
		fyne.Size{
			Width:  800,
			Height: 500,
		},
	)
	a.Settings().SetTheme(&lib.MyTheme{})
	//used to print res
	grid := widget.NewLabel("请选择一个商品包检查")
	grid.Alignment = fyne.TextAlignLeading

	//打开文件夹控件
	msg := widget.NewLabel("")
	folderOpen := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) { //该函数用来打开一个供用户选择文件夹的交互窗口
		if uri == nil {
			return
		}
		msg.SetText(uri.Path())
		fmt.Println(msg.Text)
		canvas.Refresh(grid)
		//build and run command
		res := executeCommand(command.Command, msg.Text)
		fmt.Println(res)
		grid.SetText(res)
		grid.Wrapping = fyne.TextWrapBreak
	}, w)
	uri, _ := storage.ListerForURI(storage.NewFileURI("."))
	folderOpen.SetLocation(uri)

	//选择文件夹按钮
	chooseFolderBtn := widget.NewButton("选择", func() {
		folderOpen.Show()
	})

	btns := container.NewHBox(chooseFolderBtn)
	top := container.NewVBox(
		msg,
		container.NewCenter(btns),
		canvas.NewLine(color.Gray{}),
	)
	bottom := container.NewScroll(grid)
	bottom.SetMinSize(fyne.Size{Height: 500})
	w.SetContent(container.NewBorder(top, bottom, nil, nil))
	w.ShowAndRun()
}

func executeCommand(command string, folder string) string {
	//build and run command
	cmd := exec.Command(command, folder)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Get error when execute command : %v", err)
	}
	return string(output)
}
