package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
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
		case "on_folder_button_file_set":
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
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
