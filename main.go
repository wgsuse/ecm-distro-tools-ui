package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/wgsuse/ecm-distro-tools-ui/images"
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

	// connect menus
	builder.ConnectSignalsFull(func(builder *gtk.Builder, obj *glib.GObject, sig, handler string, conn *glib.GObject, flags glib.ConnectFlags, user_data interface{}) {
		switch handler {
		case "on_menu_quit_activate":
			obj.SignalConnect(sig, func() {
				gtk.MainQuit()
			}, user_data, flags)
		case "menu_about_activate":
			obj.SignalConnect(sig, makeAbout, user_data, flags)
		}
	}, nil)

	gtk.Main()

	//
	//vbox := gtk.NewVBox(false, 1)
	//
	//menubar := gtk.NewMenuBar()
	//vbox.PackStart(menubar, false, false, 0)
	//
	//vpaned := gtk.NewVPaned()
	//vpaned.SetName("VPaned")
	//vbox.Add(vpaned)
	//
	//frame1 := gtk.NewFrame("")
	//frame1.SetBorderWidth(0)
	//framebox1 := gtk.NewVBox(false, 1)
	//framebox1.SetBorderWidth(0)
	//frame1.Add(framebox1)
	//
	//frame2 := gtk.NewFrame("")
	//framebox2 := gtk.NewVBox(false, 1)
	//frame2.Add(framebox2)
	//
	//vpaned.Pack1(frame1, false, false)
	//vpaned.Pack2(frame2, true, false)
	//
	//// Add logo
	//label := gtk.NewLabel("ECM Distro Tools UI")
	//label.ModifyFontEasy("Sans Serif 15")
	//framebox1.PackStart(label, false, true, 0)
	//// Generated with:
	//// go run ~/code/go/bin/make_inline_pixbuf logoPNG logo.png > logo.gen.go
	//pb := gdkpixbuf.NewPixbufFromData(images.LogoPNG)
	//image := gtk.NewImageFromPixbuf(pb)
	//framebox1.Add(image)
	//
	//// Add tools path
	//cmdBox := gtk.NewHBox(true, 1)
	//lblOpen := gtk.NewLabel("ECM Distro Tools path")
	//btnOpen := gtk.NewFileChooserButton("Find", gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
	//cmdBox.PackStart(lblOpen, false, false, 0)
	//cmdBox.PackEnd(btnOpen, true, true, 0)
	//framebox1.Add(cmdBox)
	//
	////--------------------------------------------------------
	//// GtkScale: stepper
	////--------------------------------------------------------
	//scale := gtk.NewHScaleWithRange(0, 100, 1)
	//scale.Connect("value-changed", func() {
	//	//fmt.Println("scale:", int(scale.GetValue()))
	//})
	//framebox2.Add(scale)
	//
	////--------------------------------------------------------
	//// GtkHBox
	////--------------------------------------------------------
	//buttons := gtk.NewHBox(false, 1)
	//
	////--------------------------------------------------------
	//// GtkButton
	////--------------------------------------------------------
	//button := gtk.NewButtonWithLabel("Button with label")
	//button.Clicked(func() {
	//	//fmt.Println("button clicked:", button.GetLabel())
	//	messagedialog := gtk.NewMessageDialog(
	//		button.GetTopLevelAsWindow(),
	//		gtk.DIALOG_MODAL,
	//		gtk.MESSAGE_INFO,
	//		gtk.BUTTONS_OK,
	//		btnOpen.GetFilename())
	//	messagedialog.Response(func() {
	//		//fmt.Println("Dialog OK!")
	//
	//		//--------------------------------------------------------
	//		// GtkFileChooserDialog
	//		//--------------------------------------------------------
	//		filechooserdialog := gtk.NewFileChooserDialog(
	//			"Choose File...",
	//			button.GetTopLevelAsWindow(),
	//			gtk.FILE_CHOOSER_ACTION_OPEN,
	//			gtk.STOCK_OK,
	//			gtk.RESPONSE_ACCEPT)
	//		filter := gtk.NewFileFilter()
	//		filter.AddPattern("*.go")
	//		filechooserdialog.AddFilter(filter)
	//		filechooserdialog.Response(func() {
	//			//fmt.Println(filechooserdialog.GetFilename())
	//			filechooserdialog.Destroy()
	//		})
	//		filechooserdialog.Run()
	//		messagedialog.Destroy()
	//	})
	//	messagedialog.Run()
	//})
	//buttons.Add(button)
	//
	////--------------------------------------------------------
	//// GtkFontButton
	////--------------------------------------------------------
	//fontbutton := gtk.NewFontButton()
	//fontbutton.Connect("font-set", func() {
	//	//fmt.Println("title:", fontbutton.GetTitle())
	//	//fmt.Println("fontname:", fontbutton.GetFontName())
	//	//fmt.Println("use_size:", fontbutton.GetUseSize())
	//	//fmt.Println("show_size:", fontbutton.GetShowSize())
	//})
	//buttons.Add(fontbutton)
	//framebox2.PackStart(buttons, false, false, 0)
	//
	//buttons = gtk.NewHBox(false, 1)
	//
	////--------------------------------------------------------
	//// GtkToggleButton
	////--------------------------------------------------------
	//togglebutton := gtk.NewToggleButtonWithLabel("ToggleButton with label")
	//togglebutton.Connect("toggled", func() {
	//	if togglebutton.GetActive() {
	//		togglebutton.SetLabel("ToggleButton ON!")
	//	} else {
	//		togglebutton.SetLabel("ToggleButton OFF!")
	//	}
	//})
	//buttons.Add(togglebutton)
	//
	////--------------------------------------------------------
	//// GtkCheckButton
	////--------------------------------------------------------
	//checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
	//checkbutton.Connect("toggled", func() {
	//	if checkbutton.GetActive() {
	//		checkbutton.SetLabel("CheckButton CHECKED!")
	//	} else {
	//		checkbutton.SetLabel("CheckButton UNCHECKED!")
	//	}
	//})
	//buttons.Add(checkbutton)
	//
	////--------------------------------------------------------
	//// GtkRadioButton
	////--------------------------------------------------------
	//buttonbox := gtk.NewVBox(false, 1)
	//radiofirst := gtk.NewRadioButtonWithLabel(nil, "Radio1")
	//buttonbox.Add(radiofirst)
	//buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
	//buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
	//buttons.Add(buttonbox)
	////radiobutton.SetMode(false);
	//radiofirst.SetActive(true)
	//
	//framebox2.PackStart(buttons, false, false, 0)
	//
	////--------------------------------------------------------
	//// GtkVSeparator
	////--------------------------------------------------------
	//vsep := gtk.NewVSeparator()
	//framebox2.PackStart(vsep, false, false, 0)
	//
	////--------------------------------------------------------
	//// GtkComboBoxEntry
	////--------------------------------------------------------
	//combos := gtk.NewHBox(false, 1)
	//comboboxentry := gtk.NewComboBoxText()
	//comboboxentry.AppendText("Monkey")
	//comboboxentry.AppendText("Tiger")
	//comboboxentry.AppendText("Elephant")
	//comboboxentry.Connect("changed", func() {
	//	//fmt.Println("value:", comboboxentry.GetActiveText())
	//})
	//combos.Add(comboboxentry)
	//
	////--------------------------------------------------------
	//// GtkComboBox
	////--------------------------------------------------------
	//combobox := gtk.NewComboBoxText()
	//combobox.AppendText("Peach")
	//combobox.AppendText("Banana")
	//combobox.AppendText("Apple")
	//combobox.SetActive(1)
	//combobox.Connect("changed", func() {
	//	//fmt.Println("value:", combobox.GetActiveText())
	//})
	//combos.Add(combobox)
	//
	//framebox2.PackStart(combos, false, false, 0)
	//
	////--------------------------------------------------------
	//// GtkTextView
	////--------------------------------------------------------
	//swin := gtk.NewScrolledWindow(nil, nil)
	//swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	//swin.SetShadowType(gtk.SHADOW_IN)
	//textview := gtk.NewTextView()
	//var start, end gtk.TextIter
	//buffer := textview.GetBuffer()
	//buffer.GetStartIter(&start)
	//buffer.Insert(&start, "Hello ")
	//buffer.GetEndIter(&end)
	//buffer.Insert(&end, "World!")
	//tag := buffer.CreateTag("bold", map[string]interface {
	//}{"background": "#FF0000", "weight": 700})
	//buffer.GetStartIter(&start)
	//buffer.GetEndIter(&end)
	//buffer.ApplyTag(tag, &start, &end)
	//swin.Add(textview)
	//framebox2.Add(swin)
	//
	//buffer.Connect("changed", func() {
	//	//fmt.Println("changed")
	//})
	//
	////--------------------------------------------------------
	//// GtkMenuItem
	////--------------------------------------------------------
	//cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	//menubar.Append(cascademenu)
	//submenu := gtk.NewMenu()
	//cascademenu.SetSubmenu(submenu)
	//
	//var menuitem *gtk.MenuItem
	//menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	//menuitem.Connect("activate", func() {
	//	gtk.MainQuit()
	//})
	//submenu.Append(menuitem)
	//
	//cascademenu = gtk.NewMenuItemWithMnemonic("_View")
	//menubar.Append(cascademenu)
	//submenu = gtk.NewMenu()
	//cascademenu.SetSubmenu(submenu)
	//
	//checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	//checkmenuitem.Connect("activate", func() {
	//	vpaned.SetSensitive(!checkmenuitem.GetActive())
	//})
	//submenu.Append(checkmenuitem)
	//
	//menuitem = gtk.NewMenuItemWithMnemonic("_Font")
	//menuitem.Connect("activate", func() {
	//	fsd := gtk.NewFontSelectionDialog("Font")
	//	fsd.SetFontName(fontbutton.GetFontName())
	//	fsd.Response(func() {
	//		//fmt.Println(fsd.GetFontName())
	//		fontbutton.SetFontName(fsd.GetFontName())
	//		fsd.Destroy()
	//	})
	//	fsd.SetTransientFor(window)
	//	fsd.Run()
	//})
	//submenu.Append(menuitem)
	//
	//cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
	//menubar.Append(cascademenu)
	//submenu = gtk.NewMenu()
	//cascademenu.SetSubmenu(submenu)
	//
	//menuitem = gtk.NewMenuItemWithMnemonic("_About")
	//menuitem.Connect("activate", makeAbout)
	//submenu.Append(menuitem)
	//
	////--------------------------------------------------------
	//// GtkStatusbar
	////--------------------------------------------------------
	//statusbar := gtk.NewStatusbar()
	//contextId := statusbar.GetContextId("go-gtk")
	//statusbar.Push(contextId, "Ready!")
	//statusbar.SetHasResizeGrip(true)
	//
	//framebox2.PackStart(statusbar, false, false, 0)
	//
	////--------------------------------------------------------
	//// Event
	////--------------------------------------------------------
	//
	//event := make(chan interface{})
	//
	//window.Add(vbox)
	//window.SetSizeRequest(MinWidth, MinHeight)
	//
	//// listen for window resizing events
	//window.Connect("configure-event", func(ctx *glib.CallbackContext) {
	//	arg := ctx.Args(0)
	//	event <- *(**gdk.EventConfigure)(unsafe.Pointer(&arg))
	//})
	//
	//// global event handler
	//go func() {
	//	for {
	//		<-event
	//		w, h := window.GetSize()
	//		// prevent resizing the window to less than the minimum dimensions
	//		if w < MinWidth || h < MinHeight {
	//			window.Resize(MinWidth, MinHeight)
	//		}
	//	}
	//}()
	//

}

func makeAbout() {
	dialog := gtk.NewAboutDialog()
	dialog.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	dialog.SetName("ECM Distro Tools UI")
	dialog.SetProgramName("ECM Distro Tools UI")
	dialog.SetAuthors(include.Authors())
	pb := gdkpixbuf.NewPixbufFromData(images.RancherLogoPNG)
	dialog.SetLogo(pb)
	dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
	dialog.SetWrapLicense(true)
	dialog.Run()
	dialog.Destroy()
}
