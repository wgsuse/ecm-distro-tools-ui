package main

import (
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/wgsuse/ecm-distro-tools-ui/images"
	"github.com/wgsuse/ecm-distro-tools-ui/include"
)

const MinWidth = 1024
const MinHeight = 768

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("ECM Distro Tools UI")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		//fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.NewVBox(false, 1)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.NewVPaned()
	vpaned.SetName("VPaned")
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk.NewFrame("Frame1")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("Frame2")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------

	label := gtk.NewLabel("Go Binding for GTK")
	label.ModifyFontEasy("DejaVu Serif 15")
	framebox1.PackStart(label, false, true, 0)

	//--------------------------------------------------------
	// GtkEntry
	//--------------------------------------------------------
	entry := gtk.NewEntry()
	entry.SetText("Hello world")
	framebox1.Add(entry)

	// Generated with:
	// go run ~/code/go/bin/make_inline_pixbuf logoPNG logo.png > logo.gen.go
	pb := gdkpixbuf.NewPixbufFromData(images.LogoPNG)
	image := gtk.NewImageFromPixbuf(pb)
	framebox1.Add(image)

	//--------------------------------------------------------
	// GtkScale: stepper
	//--------------------------------------------------------
	scale := gtk.NewHScaleWithRange(0, 100, 1)
	scale.Connect("value-changed", func() {
		//fmt.Println("scale:", int(scale.GetValue()))
	})
	framebox2.Add(scale)

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	buttons := gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	button := gtk.NewButtonWithLabel("Button with label")
	button.Clicked(func() {
		//fmt.Println("button clicked:", button.GetLabel())
		messagedialog := gtk.NewMessageDialog(
			button.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			entry.GetText())
		messagedialog.Response(func() {
			//fmt.Println("Dialog OK!")

			//--------------------------------------------------------
			// GtkFileChooserDialog
			//--------------------------------------------------------
			filechooserdialog := gtk.NewFileChooserDialog(
				"Choose File...",
				button.GetTopLevelAsWindow(),
				gtk.FILE_CHOOSER_ACTION_OPEN,
				gtk.STOCK_OK,
				gtk.RESPONSE_ACCEPT)
			filter := gtk.NewFileFilter()
			filter.AddPattern("*.go")
			filechooserdialog.AddFilter(filter)
			filechooserdialog.Response(func() {
				//fmt.Println(filechooserdialog.GetFilename())
				filechooserdialog.Destroy()
			})
			filechooserdialog.Run()
			messagedialog.Destroy()
		})
		messagedialog.Run()
	})
	buttons.Add(button)

	//--------------------------------------------------------
	// GtkFontButton
	//--------------------------------------------------------
	fontbutton := gtk.NewFontButton()
	fontbutton.Connect("font-set", func() {
		//fmt.Println("title:", fontbutton.GetTitle())
		//fmt.Println("fontname:", fontbutton.GetFontName())
		//fmt.Println("use_size:", fontbutton.GetUseSize())
		//fmt.Println("show_size:", fontbutton.GetShowSize())
	})
	buttons.Add(fontbutton)
	framebox2.PackStart(buttons, false, false, 0)

	buttons = gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkToggleButton
	//--------------------------------------------------------
	togglebutton := gtk.NewToggleButtonWithLabel("ToggleButton with label")
	togglebutton.Connect("toggled", func() {
		if togglebutton.GetActive() {
			togglebutton.SetLabel("ToggleButton ON!")
		} else {
			togglebutton.SetLabel("ToggleButton OFF!")
		}
	})
	buttons.Add(togglebutton)

	//--------------------------------------------------------
	// GtkCheckButton
	//--------------------------------------------------------
	checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
	checkbutton.Connect("toggled", func() {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!")
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!")
		}
	})
	buttons.Add(checkbutton)

	//--------------------------------------------------------
	// GtkRadioButton
	//--------------------------------------------------------
	buttonbox := gtk.NewVBox(false, 1)
	radiofirst := gtk.NewRadioButtonWithLabel(nil, "Radio1")
	buttonbox.Add(radiofirst)
	buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
	buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
	buttons.Add(buttonbox)
	//radiobutton.SetMode(false);
	radiofirst.SetActive(true)

	framebox2.PackStart(buttons, false, false, 0)

	//--------------------------------------------------------
	// GtkVSeparator
	//--------------------------------------------------------
	vsep := gtk.NewVSeparator()
	framebox2.PackStart(vsep, false, false, 0)

	//--------------------------------------------------------
	// GtkComboBoxEntry
	//--------------------------------------------------------
	combos := gtk.NewHBox(false, 1)
	comboboxentry := gtk.NewComboBoxText()
	comboboxentry.AppendText("Monkey")
	comboboxentry.AppendText("Tiger")
	comboboxentry.AppendText("Elephant")
	comboboxentry.Connect("changed", func() {
		//fmt.Println("value:", comboboxentry.GetActiveText())
	})
	combos.Add(comboboxentry)

	//--------------------------------------------------------
	// GtkComboBox
	//--------------------------------------------------------
	combobox := gtk.NewComboBoxText()
	combobox.AppendText("Peach")
	combobox.AppendText("Banana")
	combobox.AppendText("Apple")
	combobox.SetActive(1)
	combobox.Connect("changed", func() {
		//fmt.Println("value:", combobox.GetActiveText())
	})
	combos.Add(combobox)

	framebox2.PackStart(combos, false, false, 0)

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	textview := gtk.NewTextView()
	var start, end gtk.TextIter
	buffer := textview.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "Hello ")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "World!")
	tag := buffer.CreateTag("bold", map[string]interface {
	}{"background": "#FF0000", "weight": 700})
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	buffer.ApplyTag(tag, &start, &end)
	swin.Add(textview)
	framebox2.Add(swin)

	buffer.Connect("changed", func() {
		//fmt.Println("changed")
	})

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	var menuitem *gtk.MenuItem
	menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk.MainQuit()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	checkmenuitem := gtk.NewCheckMenuItemWithMnemonic("_Disable")
	checkmenuitem.Connect("activate", func() {
		vpaned.SetSensitive(!checkmenuitem.GetActive())
	})
	submenu.Append(checkmenuitem)

	menuitem = gtk.NewMenuItemWithMnemonic("_Font")
	menuitem.Connect("activate", func() {
		fsd := gtk.NewFontSelectionDialog("Font")
		fsd.SetFontName(fontbutton.GetFontName())
		fsd.Response(func() {
			//fmt.Println(fsd.GetFontName())
			fontbutton.SetFontName(fsd.GetFontName())
			fsd.Destroy()
		})
		fsd.SetTransientFor(window)
		fsd.Run()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.NewMenuItemWithMnemonic("_About")
	menuitem.Connect("activate", makeAbout)
	submenu.Append(menuitem)

	//--------------------------------------------------------
	// GtkStatusbar
	//--------------------------------------------------------
	statusbar := gtk.NewStatusbar()
	context_id := statusbar.GetContextId("go-gtk")
	statusbar.Push(context_id, "GTK binding for Go!")
	statusbar.SetHasResizeGrip(true)

	framebox2.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------

	event := make(chan interface{})

	window.Add(vbox)
	window.SetSizeRequest(MinWidth, MinHeight)
	window.Connect("configure-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event <- *(**gdk.EventConfigure)(unsafe.Pointer(&arg))
	})

	go func() {
		for {
			<-event
			w, h := window.GetSize()
			if w < MinWidth || h < MinHeight {
				window.Resize(MinWidth, MinHeight)
			}
		}
	}()

	window.ShowAll()
	gtk.Main()
}

func makeAbout() {
	dialog := gtk.NewAboutDialog()
	dialog.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	dialog.SetName("ECM Distro Tools UI")
	dialog.SetProgramName("ecm-distro-tools-ui")
	dialog.SetAuthors(include.Authors())
	pb := gdkpixbuf.NewPixbufFromData(images.RancherLogoPNG)
	dialog.SetLogo(pb)
	dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
	dialog.SetWrapLicense(true)
	dialog.Run()
	dialog.Destroy()
}
