package main

import (
    "fmt"
    "github.com/AllenDang/giu"
    "image"
	"image/color"
)

// TODO: Add a CLI for created installers.
// TODO: Add support for creating MSI files. (For Windows)
// TODO: Add support for creating .deb packges. (For Debian)

const (
    TitleBarMode_None int = 0
    TitleBarMode_Custom int = 1
    TitleBarMode_Native int = 2
)

type Installer struct {
    Title string
    Pages []Page

    // TitleBarMode:
    // 0: None
    // 1: Custom
    // 2: Native
    TitleBarMode int

    OnCancel func(page Page)
    OnBack func(before Page, after Page)
    OnNext func(before Page, after Page)
    OnFinish func(page Page)
    
    Width int
    Height int

    CancelText string
    BackText string
    NextText string
    FinishText string

    GlobalWidgets []giu.Widget
    Copyright string

    Flags giu.MasterWindowFlags
    
    CurrentPage int
    Window *giu.MasterWindow
}

func (i *Installer) DefaultCancel(page Page) {
    i.Window.Close()
}
func (i *Installer) DefaultBack(before Page, after Page) {
    i.CurrentPage--
}
func (i *Installer) DefaultNext(before Page, after Page) {
    i.CurrentPage++
}
func (i *Installer) DefaultFinish(page Page) {
    i.Window.Close()
}

func (obj *Installer) Init() {
    if obj.OnCancel == nil {
        obj.OnCancel = obj.DefaultCancel
    }
    if obj.OnBack == nil {
        obj.OnBack = obj.DefaultBack
    }
    if obj.OnNext == nil {
        obj.OnNext = obj.DefaultNext
    }
    if obj.OnFinish == nil {
        obj.OnFinish = obj.DefaultFinish
    }
    
    if obj.Width <= 0 {
		obj.Width = 640
	}
	if obj.Height <= 0 {
		obj.Height = 480
	}

    if obj.CancelText == "" {
        obj.CancelText = "Cancel"
    }
    if obj.BackText == "" {
        obj.BackText = "Back"
    }
    if obj.NextText == "" {
        obj.NextText = "Next"
    }
    if obj.FinishText == "" {
        obj.FinishText = "Finish"
    }

    if obj.GlobalWidgets == nil {
        obj.GlobalWidgets = make([]giu.Widget, 0)
    }

    if obj.Copyright == "" {
        obj.Copyright = "Copyright (c) 2022, DaRubyMiner360. All rights reserved."
    }

    if obj.TitleBarMode == TitleBarMode_None || obj.TitleBarMode == TitleBarMode_Custom {
        obj.Flags = obj.Flags | giu.MasterWindowFlagsFrameless
    }

    obj.CurrentPage = 0
    
    for _, p := range obj.Pages {
        p.Installer = obj
        if p.Widgets == nil {
            p.Widgets = make([]giu.Widget, 0)
        }
    }
}

func (i Installer) Run() {
    (&i).Init()

    //giu.Context.FontAtlas.SetDefaultFont("fonts/GidoleFont/Gidole-Regular.ttf", 12)
    fmt.Println(fmt.Sprintf("Creating installer '%s' (%dx%d)", i.Title, i.Width, i.Height))
    wnd := giu.NewMasterWindow(i.Title, i.Width, i.Height, i.Flags)

    (&i).Window = wnd
    
	wnd.Run((&i).loop)
}

func (i *Installer) loop() {
    p := i.Pages[i.CurrentPage]

    // TODO: Make the installer responsive.
    widgets := []giu.Widget{
        giu.Custom(func() {
            if i.TitleBarMode == TitleBarMode_Custom {
                canvas := giu.GetCanvas()
                width, _ := i.Window.GetSize()
                canvas.AddRectFilled(image.Pt(0, 0), image.Pt(width, 20), color.RGBA{200, 75, 75, 255}, 0, 0)
                // TODO: Add buttons on top of the custom title bar.
                
                giu.SetCursorScreenPos(giu.GetCursorScreenPos().Add(image.Pt(0, 15)))
            }
		}),
    }

    // TODO: Add support for adding custom buttons instead of JUST the default ones.
    // TODO: Add suport for adding custom fonts.
    // TODO: Add support for customizing the look and feel of the installer via styling.
    navigation := []giu.Widget{}
    navigation = append(navigation, giu.Widget(giu.Button(i.CancelText).OnClick(func() {
        i.OnCancel(p)
    })))
    if i.CurrentPage > 0 {
        navigation = append(navigation, giu.Widget(giu.Button(i.BackText).OnClick(func() {
            i.OnBack(p, i.Pages[i.CurrentPage-1])
        })))
    }
    if i.CurrentPage < len(i.Pages) - 1 {
        navigation = append(navigation, giu.Widget(giu.Button(i.NextText).OnClick(func() {
            i.OnNext(p, i.Pages[i.CurrentPage+1])
        })))
    } else {
        navigation = append(navigation, giu.Widget(giu.Button(i.FinishText).OnClick(func() {
            i.OnFinish(p)
        })))
    }

    widgets = append(
        widgets,
        p.Widgets...,
    )
    widgets = append(
        widgets,
        i.GlobalWidgets...,
    )
    
    widgets = append(
        widgets,
        giu.Custom(func() {
            _, height := i.Window.GetSize()
            giu.SetCursorScreenPos(image.Pt(giu.GetCursorScreenPos().X, height - 45))
            giu.Align(giu.AlignRight).To(
                giu.Style().SetColor(giu.StyleColorText, color.RGBA{235, 237, 242, 255}).To(
                    giu.Label(i.Copyright),
                ),
            ).Build()
            giu.Separator().Build()
            giu.Row(navigation...).Build()
		}),
    )
    giu.SingleWindow().Layout(
        widgets...
	)
}

type Page struct {
    Title string
    Widgets []giu.Widget

    Installer *Installer
}
