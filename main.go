package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/wgsuse/ecm-distro-tools-ui/include"
)

//const MinWidth = 1024
//const MinHeight = 768

func main() {
	gtk.Init(&os.Args)

	// Get the GtkBuilder UI definition in the glade file.
	builder := gtk.NewBuilder()
	_, err := builder.AddFromFile("main.glade")
	if err != nil {
		log.Fatalln("Unable to load user interface file", err)
	}
	obj := builder.GetObject("window1")
	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "foo")

	// set title font
	lblTitle := *(*gtk.Label)(unsafe.Pointer(&builder.GetObject("label2").Object))
	lblTitle.ModifyFontEasy("Sans 15")
	lblTitle.SetMarkup(fmt.Sprintf(`<span weight="bold" color="darkgreen" size="xx-large" style="normal" >%v</span>`, lblTitle.GetText()))

	pane := *(*gtk.VPaned)(unsafe.Pointer(&builder.GetObject("vpaned1").Object))
	//check := *(*gtk.CheckMenuItem)(unsafe.Pointer(&builder.GetObject("menu_disable").Object))

	btn1 := *(*gtk.Button)(unsafe.Pointer(&builder.GetObject("button1").Object))
	btn1.SetSensitive(false)
	btn2 := *(*gtk.Button)(unsafe.Pointer(&builder.GetObject("button2").Object))
	btn2.SetSensitive(false)
	btn3 := *(*gtk.Button)(unsafe.Pointer(&builder.GetObject("button3").Object))
	btn3.SetSensitive(false)
	btn4 := *(*gtk.Button)(unsafe.Pointer(&builder.GetObject("button4").Object))
	btn4.SetSensitive(false)

	btnOpen := *(*gtk.Button)(unsafe.Pointer(&builder.GetObject("btn_open").Object))
	entryPath := *(*gtk.Entry)(unsafe.Pointer(&builder.GetObject("entry_path").Object))

	textView := *(*gtk.TextView)(unsafe.Pointer(&builder.GetObject("textview1").Object))

	dialogPrefs := *(*gtk.Dialog)(unsafe.Pointer(&builder.GetObject("dialog_prefs").Object))

	aboutDialog := &gtk.AboutDialog{
		*(*gtk.Dialog)(unsafe.Pointer(&builder.GetObject("aboutdialog1").Object))}

	// connect menus
	builder.ConnectSignalsFull(func(builder *gtk.Builder, obj *glib.GObject, sig, handler string, conn *glib.GObject, flags glib.ConnectFlags, user_data interface{}) {
		switch handler {
		case "on_menu_quit_activate":
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
			}, user_data, flags)
		case "on_button1_clicked":
			// standup
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
			}, user_data, flags)
		case "on_button2_clicked":
			// backport
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
			}, user_data, flags)
		case "on_button3_clicked":
			// gen-release-notes
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
			}, user_data, flags)
		case "on_button4_clicked":
			// k3s-release
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
			}, user_data, flags)
		case "on_btn_open_clicked":
			obj.SignalConnect(sig, func() {
				fileChooserDialog := gtk.NewFileChooserDialog(
					"Choose File...",
					btnOpen.GetTopLevelAsWindow(),
					gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
					gtk.STOCK_OK,
					gtk.RESPONSE_ACCEPT)
				fileChooserDialog.Response(func() {
					folder := fileChooserDialog.GetCurrentFolder()
					if ok, _ := exists(folder); ok {
						entryPath.SetText(folder)
						tools := include.Tools(folder)
						buffer := textView.GetBuffer()
						buffer.SetText("")
						if len(tools) > 0 {
							for _, t := range tools {
								if len(t) == 0 {
									continue
								}
								buffer.InsertAtCursor("Found: " + t + "\n")
							}
							btn1.SetSensitive(true)
							btn2.SetSensitive(true)
							btn3.SetSensitive(true)
							btn4.SetSensitive(true)
						} else {
							buffer.SetText("The ECM distro tools were not found in " + folder)
						}
					}
					fileChooserDialog.Destroy()
				})
				fileChooserDialog.Run()
			}, user_data, flags)
		case "on_menu_disable_activate":
			obj.SignalConnect(sig, func() {
				pane.SetSensitive(!pane.GetSensitive())
			}, user_data, flags)
		case "on_menu_prefs_activate":
			obj.SignalConnect(sig, func() {
				dialogPrefs.Run()
				dialogPrefs.Hide()
			}, user_data, flags)
		case "menu_about_activate":
			obj.SignalConnect(sig, func() {
				aboutDialog.Run()
				aboutDialog.Hide()
			}, user_data, flags)
		}
	}, nil)

	gtk.Main()
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
