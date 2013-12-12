package main

import "../lib/wx"

type myframe struct {
	frame wx.Frame
	statusbar wx.StatusBar
	menubar wx.MenuBar
	notebook wx.Notebook
}


func (f *myframe)evtFont(wx.Event){
	fontdlg := wx.NewFontDialog(f.frame)
	fontdlg.ShowModal()
	font := fontdlg.GetFontData().GetChosenFont()
	f.notebook.GetCurrentPage().SetFont(font)
	fontdlg.Destroy()
}

func (f *myframe)evtAbout(wx.Event){		
	wx.MessageBox("Welcome to wxWidgets!\nString test|测试|測試|試験|테스트")
}

func (f *myframe)InitFrame(){
    f.frame = wx.NewFrame()
    f.frame.Create(wx.NullWindow, -1, "GoLang wxWidgets Wrapper")
        
    f.statusbar = f.frame.CreateStatusBar()
    f.statusbar.SetStatusText("Welcome to wxWidgets")
        
    f.statusbar.SetFieldsCount(2)
    f.statusbar.SetStatusText("This is a statusbar!", 1)
        
    f.menubar = wx.NewMenuBar()
    menuFile := wx.NewMenu()
    
    menuItemFont := wx.NewMenuItem(menuFile, int(wx.ID_ANY), "Font...", "Select A Font", wx.ITEM_NORMAL)
    menuFile.Append(menuItemFont)
    menuItemAbout := wx.NewMenuItem(menuFile, int(wx.ID_ANY), "About", "About", wx.ITEM_NORMAL)
    menuFile.Append(menuItemAbout)
    
    f.menubar.Append(menuFile, "File")
    f.frame.SetMenuBar(f.menubar)

    mainSizer := wx.NewBoxSizer(int(wx.HORIZONTAL) )
    f.frame.SetSizer(mainSizer)
        
    f.notebook = wx.NewNotebook( f.frame, int(wx.ID_ANY), wx.DefaultPosition, wx.DefaultSize, 0 )
        
    mainSizer.Add(f.notebook, 1, int(wx.EXPAND), 5)
        
    

	textCtrl := wx.NewTextCtrl(f.notebook, int(wx.ID_ANY), "", wx.DefaultPosition, wx.DefaultSize, wx.TE_MULTILINE)
    textCtrl.SetMinSize(wx.NewSize(600, 400))

	f.notebook.AddPage(textCtrl, "This is a page", true)
	textCtrl.SetFocus()
        
    f.frame.Layout()
	mainSizer.Fit(f.frame)
    wx.Bind( f.frame, wx.EVT_MENU, f.evtFont, menuItemFont.GetId() )
    wx.Bind( f.frame, wx.EVT_MENU, f.evtAbout, menuItemAbout.GetId() )
    //wx.Unbind( f.frame, wx.EVT_MENU, f.evtFont, menuItemFont.GetId() )
    //wx.Unbind( f.frame, wx.EVT_MENU, f.evtAbout, menuItemAbout.GetId() )
    
    
    f.frame.Show()
}

func main(){
    wx1 := wx.NewApp()
    var f myframe
    f.InitFrame()
    wx1.MainLoop()
}
