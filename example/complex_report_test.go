package example

import (
	"fmt"
	"github.com/cindoralla/gopdf"
	"github.com/cindoralla/gopdf/core"
	//gopdf "gitee.com/cindoralla/gopdf"
	"path/filepath"
	"testing"
)

const (
	ErrFile    = 1
	FONT_MY    = "微软雅黑"
	FONT_MD    = "MPBOLD"
	DateFormat = "2006-01-02 15:04:05"
)

var (
	largeFont      = core.Font{Family: FONT_MY, Size: 15}
	titleFont      = core.Font{Family: FONT_MY, Size: 13}
	headFont       = core.Font{Family: FONT_MY, Size: 11}
	textFont       = core.Font{Family: FONT_MY, Size: 8}
	textFontU      = core.Font{Family: FONT_MY, Size: 8, Style: "U"}
	cornerMarkFont = core.Font{Family: FONT_MY, Size: 6}
)

func ComplexReport() {
	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: FONT_MY,
		FileName: "ttf//microsoft.ttf",
	}
	font2 := core.FontMap{
		FontName: FONT_MD,
		FileName: "ttf//mplus-1p-bold.ttf",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2})
	r.SetPage("A4", "P")
	r.FisrtPageNeedHeader = true
	r.FisrtPageNeedFooter = true

	r.RegisterExecutor(core.Executor(ComplexReportExecutor), core.Detail)
	//r.RegisterExecutor(core.Executor(SimpleTableExecutor), core.Detail)
	//r.RegisterExecutor(core.Executor(ComplexReportFooterExecutor), core.Footer)
	r.RegisterExecutor(core.Executor(ComplexReportHeaderExecutor), core.Header)

	r.Execute(fmt.Sprintf("report.pdf"))
	r.SaveAtomicCellText("report.txt")
}

func ComplexReportExecutor(report *core.Report) {
	var (
		//data      ReportDetail
		lineSpace = 1.0
		lineHight = 16.0
	)

	dir, _ := filepath.Abs("pictures")
	qrcodeFile := fmt.Sprintf("%v/a.png", dir)

	line1 := gopdf.NewHLine(report).SetMargin(core.Scope{Top: -40, Bottom: 20}).
		SetWidth(0.08).SetColor(0.5)

	line1.GenerateAtomicCell()

	// todo: 任务详情
	div := gopdf.NewDivWithWidth(595, lineHight, lineSpace, report)
	div.SetFont(largeFont)
	div.SetContent("文章的名称").GenerateAtomicCell()

	//cornerMarkDiv := gopdf.NewTextCell(30, lineHight, lineSpace, report)
	//curX, curY := report.GetXY()
	//fmt.Println("height:",div.GetHeight())
	//report.SetXY(curX+float64(largeFont.Size)*2-lineSpace*2, curY-div.GetHeight()-lineHight/3)
	//cornerMarkDiv.SetFont(cornerMarkFont).
	//	SetFontColor("255,0,0").
	//	SetContent("=").
	//	GenerateAtomicCell(40)
	//report.SetXY(curX, curY)

	// 二维码
	im := gopdf.NewImageWithWidthAndHeight(qrcodeFile, 18, 18, report)
	im.SetMargin(core.Scope{Left: 0, Top: 10})
	im.SetAutoBreak()
	im.GenerateAtomicCell()

	// 基本信息
	report.SetMargin(30, -11)
	baseInfoDiv := gopdf.NewDivWithWidth(100, lineHight, lineSpace, report)
	baseInfoDiv.SetFont(textFont)
	baseInfoDiv.SetContent("社会主义接班人").GenerateAtomicCell()

	report.SetMargin(0, 50)
	line1.SetWidth(0.1).SetColor(0.5)
	line1.GenerateAtomicCell()

	report.SetMargin(0, 3)
	baseInfoDiv.Copy("文章信息").SetFont(titleFont).GenerateAtomicCell()

	report.SetMargin(0, 10)
	articleInfoDiv := gopdf.NewTextCell(400, 40, lineSpace, report)
	articleInfoDiv.SetFont(textFont).SetFontColor("0, 153, 255").
		SetContent("802             120").
		GenerateAtomicCell(40)

	report.SetMargin(80, 160)
	articleInfoDiv2 := gopdf.NewTextCell(400, 40, lineSpace, report)
	articleInfoDiv2.SetFont(textFont).SetContent("字数             段落数").GenerateAtomicCell(40)

	report.SetMargin(80, 220)
	line1.SetWidth(0.1).SetColor(0.5)
	line1.GenerateAtomicCell()

	baseInfoDiv.Copy("语法检查结果").SetFont(titleFont).GenerateAtomicCell()

	report.SetMargin(0, 10)
	grammarInfoDiv := gopdf.NewTextCell(400, 20, lineSpace, report)
	grammarInfoDiv.SetFont(textFont).SetFontColor("255, 0, 0").
		SetContent("24				 15				  2").GenerateAtomicCell(20)

	report.SetMargin(80, 240)
	grammarInfoDiv2 := gopdf.NewTextCell(400, 20, lineSpace, report)
	grammarInfoDiv2.SetFont(textFont).SetContent("基础错误		标点错误		高级错误").GenerateAtomicCell(20)

	report.SetMargin(80, 300)
	line1.SetWidth(0.1).SetColor(0.5)
	line1.GenerateAtomicCell()

	baseInfoDiv.Copy("详细").SetFont(titleFont).GenerateAtomicCell()
	// 第一行
	report.SetMargin(0, 10)
	grammarInfoDiv.Copy("18").SetFontColor("255,0,0").SetFont(textFont).
		SetBorder(core.Scope{Top: 5, Left: 20}).
		SetBackColor("240, 240, 245").
		GenerateAtomicCell(20)
	report.SetMargin(80, 310)
	grammarInfoDiv.Copy("缺少成分").SetFont(textFont).SetFontColor("0,0,0").
		SetBorder(core.Scope{Left: 70}).
		//SetBackColor("240, 240, 245").
		GenerateAtomicCell(20)
	// 第二行
	report.SetMargin(80, 330)
	grammarInfoDiv.Copy("13").SetFontColor("255,0,0").SetFont(textFont).
		SetBorder(core.Scope{Top: 5, Left: 20}).
		SetBackColor("255, 255, 255").
		GenerateAtomicCell(20)
	report.SetMargin(80, 330)
	grammarInfoDiv.Copy("删除单词").SetFont(textFont).SetFontColor("0,0,0").
		SetBorder(core.Scope{Top: 5, Left: 70}).
		//SetBackColor("240, 240, 245").
		GenerateAtomicCell(20)
	// 第三行
	report.SetMargin(80, 350)
	grammarInfoDiv.Copy("8").SetFontColor("255,0,0").SetFont(textFont).
		SetBorder(core.Scope{Top: 5, Left: 20}).
		SetBackColor("240, 240, 245").
		GenerateAtomicCell(20)
	report.SetMargin(80, 350)
	grammarInfoDiv.Copy("修改时间时态").SetFont(textFont).SetFontColor("0,0,0").
		SetBorder(core.Scope{Top: 5, Left: 70}).
		//SetBackColor("240, 240, 245").
		GenerateAtomicCell(20)
	// 第四行
	report.SetMargin(80, 370)
	grammarInfoDiv.Copy("5").SetFontColor("255,0,0").SetFont(textFont).
		SetBorder(core.Scope{Top: 5, Left: 20}).
		SetBackColor("255, 255, 255").
		GenerateAtomicCell(20)
	report.SetMargin(80, 370)
	grammarInfoDiv.Copy("需要替换").SetFont(textFont).SetFontColor("0,0,0").
		SetBorder(core.Scope{Top: 5, Left: 70}).
		//SetBackColor("240, 240, 245").
		GenerateAtomicCell(20)

	report.SetMargin(80, 440)
	line1.SetWidth(0.1).SetColor(0.5)
	line1.GenerateAtomicCell()

	baseInfoDiv.Copy("完整原文").SetFont(titleFont).GenerateAtomicCell()
	report.SetMargin(0, 10)
	baseInfoDiv.Copy("Madam Curie").SetFont(largeFont).GenerateAtomicCell()

	txt := `The year 1866 was marked by a bizarre development, an unexplained and downright inexplicable phenomenon that surely no one has forgotten. Without getting into those rumors that upset civilians in the seaports and deranged the public mind even far inland, it must be said that professional seamen were especially alarmed. Traders, shipowners, captains of vessels, skippers, and master mariners from Europe and America, naval officers from every country, and at their heels the various national governments on these two continents, were all extremely disturbed by the business.`
	report.SetMargin(0, 20)

	grammarInfoDiv.Copy(txt).SetFont(textFont).SetFontColor("0,0,0").GenerateAtomicCell(130)

	fmt.Println(report.GetPageEndXY())

	report.SetMargin(80, 660)
	line1.SetWidth(0.1).SetColor(0.5)
	line1.GenerateAtomicCell()

	baseInfoDiv.Copy("详细结果").SetFont(titleFont).GenerateAtomicCell()

	report.SetMargin(0, 50)
	line1.SetWidth(0.1).SetColor(0.5)
	line1.GenerateAtomicCell()

	// 第一个结果
	report.SetMargin(0, -10)
	grammarInfoDiv.Copy("1.").SetFontColor("0,0,0").SetFont(textFont).
		GenerateAtomicCell(20)

	report.SetMargin(80, 675)
	grammarInfoDiv.Copy(">").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 20}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 675)
	grammarInfoDiv.Copy(",").SetFontColor("0,255,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 30}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 675)
	grammarInfoDiv.Copy("似乎缺少[,]，请考虑添加。").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 120}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 675)
	grammarInfoDiv.Copy("缺少成分").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 320}).
		GenerateAtomicCell(20)

	// 第二个结果
	report.SetMargin(80, 705)
	grammarInfoDiv.Copy("2.").SetFontColor("0,0,0").SetFont(textFont).
		GenerateAtomicCell(20)

	report.SetMargin(80, 705)
	grammarInfoDiv.Copy("with").SetFontColor("255,0,0").SetFont(textFontU).
		SetBorder(core.Scope{Left: 20}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 705)
	grammarInfoDiv.Copy("[with]可能是冗余的，请考虑删除。").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 120}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 705)
	grammarInfoDiv.Copy("删除单词").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 320}).
		GenerateAtomicCell(20)

	// 第三个结果
	report.SetMargin(80, 735)
	grammarInfoDiv.Copy("3.").SetFontColor("0,0,0").SetFont(textFont).
		GenerateAtomicCell(20)

	report.SetMargin(80, 735)
	grammarInfoDiv.Copy("had").SetFontColor("255,0,0").SetFont(textFontU).
		SetBorder(core.Scope{Left: 20}).
		GenerateAtomicCell(20)
	report.SetMargin(80, 735)
	grammarInfoDiv.Copy("hads").SetFontColor("0,255,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 40}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 735)
	grammarInfoDiv.Copy("[had]的时态可能有误，请考虑修改。").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 120}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 735)
	grammarInfoDiv.Copy("修改动词时态").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 320}).
		GenerateAtomicCell(20)

	// 第四个结果
	report.SetMargin(80, 765)
	grammarInfoDiv.Copy("4.").SetFontColor("0,0,0").SetFont(textFont).
		GenerateAtomicCell(20)

	report.SetMargin(80, 765)
	grammarInfoDiv.Copy("word").SetFontColor("255,0,0").SetFont(textFontU).
		SetBorder(core.Scope{Left: 20}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 765)
	grammarInfoDiv.Copy(">").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 50}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 765)
	grammarInfoDiv.Copy("world").SetFontColor("0,255,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 60}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 765)
	grammarInfoDiv.Copy("请考虑替换[word]").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 120}).
		GenerateAtomicCell(20)

	report.SetMargin(80, 765)
	grammarInfoDiv.Copy("需要替换").SetFontColor("0,0,0").SetFont(textFont).
		SetBorder(core.Scope{Left: 320}).
		GenerateAtomicCell(20)
}

func ComplexReportFooterExecutor(report *core.Report) {
	content := fmt.Sprintf("第 %v / {#TotalPage#} 页", report.GetCurrentPageNo())
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(textFont).SetMarign(core.Scope{Top: 10})
	//footer.SetFontColor("60, 179, 113")
	footer.HorizontalCentered().SetContent(content).GenerateAtomicCell()
}

func ComplexReportHeaderExecutor(report *core.Report) {
	content := "报告编号：0988888888888888"
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(textFont)
	footer.SetFontColor("179,179,204")
	footer.SetBorder(core.Scope{Top: 10,Left:290})
	footer.SetContent(content).GenerateAtomicCell()

	dir, _ := filepath.Abs("pictures")
	qrcodeFile := fmt.Sprintf("%v/b.png", dir)
	im := gopdf.NewImageWithWidthAndHeight(qrcodeFile, 40, 18, report)
	im.SetMargin(core.Scope{Left: 0, Top: 10})
	im.SetAutoBreak()
	im.GenerateAtomicCell()
}

func TestJobExport(t *testing.T) {
	ComplexReport()
}
