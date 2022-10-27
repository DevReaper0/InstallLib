package main

import "github.com/AllenDang/giu"

func SimpleExample() {
	Installer{
        Title: "Installer",
        Pages: []Page{
            {
                Title:"Page 1",
                Widgets:[]giu.Widget{
                    giu.Label("Hello world from InstallLib (1)"),
                },
            },
            {
                Title:"Page 2",
                Widgets:[]giu.Widget{
                    giu.Label("Hello world from InstallLib (2)"),
                },
            },
            {
                Title:"Page 3",
                Widgets:[]giu.Widget{
                    giu.Label("Hello world from InstallLib (3)"),
                },
            },
        },
        OnCancel:Cancel,
        OnFinish:Finish,
        Width:640,
        Height:480,
        GlobalWidgets:[]giu.Widget{
            giu.PrepareMsgbox(),
        },
    }.Run()
}

func Cancel(page Page) {
    giu.Msgbox("Cancel", "Are you sure you want to cancel?").Buttons(giu.MsgboxButtonsYesNo).ResultCallback(func(result giu.DialogResult) {
        if result == giu.DialogResultYes {
            page.Installer.DefaultCancel(page)
        }
    })
}
func Finish(page Page) {
    giu.Msgbox("Finish", "Are you sure you're ready to finish?").Buttons(giu.MsgboxButtonsYesNo).ResultCallback(func(result giu.DialogResult) {
        if result == giu.DialogResultYes {
            page.Installer.DefaultFinish(page)
        }
    })
}

func main() {
	SimpleExample()
}
