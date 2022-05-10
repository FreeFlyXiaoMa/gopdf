package example

import (
	"fmt"
	"github.com/cindoralla/gopdf"
	"github.com/cindoralla/gopdf/core"
	"github.com/sergi/go-diff/diffmatchpatch"
	"path/filepath"
	"strings"
	"testing"
	"unicode"
)

func ComplexReport2() []byte {
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

	r.RegisterExecutor(core.Executor(ComplexReportExecutor2), core.Detail)
	//r.RegisterExecutor(core.Executor(SimpleTableExecutor), core.Detail)
	r.RegisterExecutor(core.Executor(ComplexReportFooterExecutor2), core.Footer)
	r.RegisterExecutor(core.Executor(ComplexReportHeaderExecutor2), core.Header)

	r.Execute(fmt.Sprintf("complex_paraphrase_test.pdf"))
	//r.SaveAtomicCellText("complex_paraphrase_test.txt")
	return r.GetBytesPdf()
}

var (
	lineSpace        = 1.0
	lineHight        = 16.0
	blackColor       = "0,0,0"
	redColor         = "255,0,0"
	blueColor        = "0, 153, 51"
	whiteColor       = "255,255,255"
	A4Width          = 595.28
	startX           = 90.14
	startY           = 72.00
	A4Hight          = 841.89
	partDuraiton     = 20.0
	inpartDuraion    = 15.0
	smallTxtWidth    = 100.0
	smallTxtMaxHight = 40.0
	articInfoleWidth = 80.0 //文章信息的宽度
	maxHight         = 40.0
)

func ComplexReportExecutor2(report *core.Report) {
	var (
		title        = "《用户文档原文标题》改写报告"
		articleTitle = "文章信息"
		detailResult = "详细结果"
		remarks      = "绿色标识部分为改写后与原文的差异部分"

		originTxt      = "原文字数"
		afterChangeTxt = "改动后字数"
		changedTxt     = "改动字数"
		changeTxtRatio = "改动字数比例"
	)

	reportTime := fmt.Sprintf("报告时间：%s", "2022/06/21 16:41")
	paraType := fmt.Sprintf("改写类型：%s", "高级改写")
	originTxtTotalCount := "5821"
	txtTotalCount := "7865"
	changedtxtCount := "2835"
	changedTxtRatioCount := "48.6%"

	curX, curY := report.GetXY()
	curY = 50
	report.SetXY(curX, curY)
	line1 := gopdf.NewHLine(report).SetWidth(1).SetColor(0.9)

	line1.GenerateAtomicCell()
	curY += partDuraiton
	report.SetXY(curX, curY)

	publicDiv := gopdf.NewDivWithWidth(28, lineHight-4, lineSpace, report).SetFont(textFont)

	// 1. 概要
	// 1.1 标题
	txtCell := gopdf.NewTextCell(smallTxtWidth*3, lineHight, lineSpace, report)
	txtCell.SetFont(largeFont)
	txtCell.SetContent(title).GenerateAtomicCell(smallTxtMaxHight)
	curY += inpartDuraion * 2

	report.SetXY(curX, curY)
	txtCell.Copy(reportTime).SetFont(textFont).GenerateAtomicCell(smallTxtMaxHight)

	report.SetXY(curX+smallTxtWidth*2, curY)
	txtCell.Copy(paraType).SetFont(textFont).GenerateAtomicCell(smallTxtMaxHight)

	curY += partDuraiton
	report.SetXY(curX, curY)

	line1.GenerateAtomicCell()
	curY += partDuraiton
	report.SetXY(curX, curY)

	// 2.文章信息
	txtCell.Copy(articleTitle).SetFont(headFont).GenerateAtomicCell(articInfoleWidth)
	curY += inpartDuraion * 2
	// 2.1 第一行数字部分
	report.SetXY(curX, curY)
	txtCell.Copy(originTxtTotalCount).SetFont(textFont).SetFontColor(redColor).GenerateAtomicCell(maxHight)
	report.SetXY(curX+articInfoleWidth, curY)
	txtCell.Copy(txtTotalCount).SetFont(textFont).SetFontColor(blueColor).GenerateAtomicCell(maxHight)
	report.SetXY(curX+articInfoleWidth*2, curY)
	txtCell.Copy(changedtxtCount).SetFont(textFont).SetFontColor(blueColor).GenerateAtomicCell(maxHight)
	report.SetXY(curX+articInfoleWidth*3, curY-3)
	publicDiv.Copy(changedTxtRatioCount).HorizontalCentered().SetBorder(core.Scope{Top: 3}).SetFontColor(whiteColor).SetBackColor(blueColor).GenerateAtomicCell()
	// 2.2 第二行字母部分
	curY += inpartDuraion
	report.SetXY(curX, curY)
	txtCell.Copy(originTxt).SetFont(textFont).SetFontColor(blackColor).GenerateAtomicCell(maxHight)
	report.SetXY(curX+articInfoleWidth, curY)
	txtCell.Copy(afterChangeTxt).SetFont(textFont).SetFontColor(blackColor).GenerateAtomicCell(maxHight)
	report.SetXY(curX+articInfoleWidth*2, curY)
	txtCell.Copy(changedTxt).SetFont(textFont).SetFontColor(blackColor).GenerateAtomicCell(maxHight)
	report.SetXY(curX+articInfoleWidth*3, curY)
	txtCell.Copy(changeTxtRatio).SetFont(textFont).SetFontColor(blackColor).GenerateAtomicCell(maxHight)

	curY += partDuraiton
	report.SetXY(curX, curY)
	line1.GenerateAtomicCell()

	// 3.详细结果
	curY += partDuraiton
	report.SetXY(curX, curY)
	txtCell.Copy(detailResult).SetFont(headFont).SetFontColor(blackColor).GenerateAtomicCell(maxHight)
	curY += partDuraiton

	report.SetXY(curX, curY)
	txtCell.Copy(remarks).SetFont(textFont).SetFontColor(blackColor).GenerateAtomicCell(maxHight)

	curY += partDuraiton
	report.SetXY(curX, curY)

	//text1 := "(2) QIIME:2010年，美国科罗拉多大学的Rob Knight教授(现单位美国加州大学圣地亚哥分校)团队发布QIIME(发音同chime)分析流程[19]。该流程可在Linux或Mac系统中运行，相比mothur具有更多的优点，主要包括:整合了200多款相关软件和包，实现每个步骤更多软件和方法的选择;提供150多个脚本，实现各种个性化分析，并可以应对不同类型数据和实验设计;流程开放程度高，容易整合新软件和方法;增强统计和可视化，实现多样性、物种组成、差异比较和网络等众多方法和出版级图表绘制。由于QIIME允许同领域研究者较自主地开展扩增子数据的个性化分析和可视化，逐渐成为本领域最受欢迎的软件(图2)。为了满足日益增长的测序数据量和可重复计算的要求，GregoryJ. Caporaso教授于2016年起发起了基于Python 3语言从头编写的QIIME 2项目[20]。"
	//text2 := "（2） QIIME:2010年，美国科罗拉多大学的Rob Knight教授（现单位美国加州大学圣地亚哥分校）团队发布QIIME（发音同chime）分析流程[19]。该流程可在Linux或Mac系统中运行，使用简单的图形用户界面，通过输入实验参数即可快速完成结果输出及统计分析工作，为生物信息学研究提供高效可靠的工具。QIIME是目前国际上最流行的基于内容分析方法之一。相比mothur具有更多的优点，主要包括：整合了200多款相关软件和包，实现每个步骤更多软件和方法的选择；提供150多个脚本，实现各种个性化分析，并可以应对不同类型数据和实验设计；流程开放程度高，容易整合新软件和方法；增强统计和可视化，实现多样性、物种组成、差异比较和网络等众多方法和出版级图表绘制。由于QIIME允许同领域研究者较自主地开展扩增子数据的个性化分析和可视化，逐渐成为本领域最受欢迎的软件（图2）。为了满足日益增长的测序数据量和可重复计算的要求，Gregory J. Caporaso教授于2016年起发起了基于Python 3语言从头编写的QIIME 2项目[20]。"

	text1 := `The year 1866 was marked by a bizarre development, an unexplained and downright inexplicable phenomenon that surely no one has forgotten. Without getting into those rumors that upset civilians in the seaports and deranged the public mind even far inland, it must be said that professional seamen were especially alarmed. Traders, shipowners, captains of vessels, skippers, and master mariners from Europe and America, naval officers from every country, and at their heels the various national governments on these two continents, were all extremely disturbed by the business.In essence, over a period of time several ships had encountered an enormous thing at sea, a long spindle-shaped object, sometimes giving off a phosphorescent glow, infinitely bigger and faster than any whale.The relevant data on this apparition`
	text2 := `Something strange happened in 1866, a completely unexplained phenomenon that surely no one has forgotten. In the absence of rumours that upset civilians in the harbour and even unhinged the public in the interior, it must be noted that professional seafarers were particularly shocked. Traders, shipowners, captains of vessels, skippers, and master mariners from Europe and America, naval officers from every country, and at their heels the various national governments on these two continents, were all extremely disturbed by the business.In essence, over a period of time several ships had encountered an enormous thing at sea, a long spindle-shaped object, sometimes giving off a phosphorescent glow, infinitely bigger and faster than any whale.The relevant data on this apparition`

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)
	txtBodyCell := gopdf.NewTextCell(smallTxtWidth*5, lineHight, lineSpace, report).SetFont(textFont)
	fmt.Println(dmp.DiffPrettyText(diffs))

	hasChinese := IsChinese(text1)
	genColorfulTxt(hasChinese, diffs, txtBodyCell, report)

}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func diffProcess(hasChinese bool, diffs []diffmatchpatch.Diff) []diffmatchpatch.Diff {
	res := []diffmatchpatch.Diff{}
	// 中文
	if hasChinese {
		for _, diff := range diffs {
			text := diff.Text
			type_ := diff.Type
			if len(text) > 1 { // 拆出来每一个中文字母，单词写入pdf
				for _, item := range text {
					dif := diffmatchpatch.Diff{Text: string(item), Type: type_}
					res = append(res, dif)
				}
			} else {
				dif := diffmatchpatch.Diff{Text: text, Type: type_}
				res = append(res, dif)
			}
		}
		return res
	}

	// 英文
	// 1. 去掉类型为deleted
	tmp := []diffmatchpatch.Diff{}
	for _, diff := range diffs {
		if diff.Type != diffmatchpatch.DiffDelete {
			tmp = append(tmp, diff)
		}
	}
	// 2. 将剩余的字符串拼接起来，字符串中默认带有空格，起到区分单词的作用
	diffs = tmp

	// 3. 有空格分开的那种单词，把它拆分出来，一个单词一个单词的写入pdf
	diffs = tmp
	for _, diff := range diffs {
		text := diff.Text
		type_ := diff.Type
		if type_ == diffmatchpatch.DiffDelete {
			continue
		}
		if len(text) > 1 && strings.Contains(text, " ") {
			textSlice := strings.Split(text, " ")
			for _, item := range textSlice {
				dif := diffmatchpatch.Diff{Text: item + " ", Type: type_}
				res = append(res, dif)
			}
		} else {
			dif := diffmatchpatch.Diff{Text: text, Type: type_}
			res = append(res, dif)
		}
	}

	return res
}

func genColorfulTxt(hasChinese bool, diffs []diffmatchpatch.Diff, txtCell *gopdf.TextCell, report *core.Report) {
	curX, curY := report.GetXY()
	// 每一个单词一个坐标系，在每一行的结尾
	fmt.Println("diffs:", diffs)
	diffs = diffProcess(hasChinese, diffs)
	fmt.Println("after diffs:", diffs)
	wordLen := 0.0
	for _, diff := range diffs {
		text := diff.Text
		switch diff.Type {
		case diffmatchpatch.DiffEqual, diffmatchpatch.DiffInsert: // 黑色字
			color := ""
			if diff.Type == diffmatchpatch.DiffEqual {
				color = blackColor
			} else if diff.Type == diffmatchpatch.DiffInsert {
				color = blueColor
			}
			if curX >= A4Width-startX { //重起一行
				curX = startX
				curY += lineHight
			}
			report.SetXY(curX, curY)
			txtCell.Copy(text).SetFont(textFont).SetFontColor(color).GenerateAtomicCell(maxHight)
			if hasChinese {
				if len(text) == 1 { // 计算每个单词的长度
					wordLen = float64(len(text) * 4)
				} else {
					wordLen = float64(len(text) * 2)
				}
				curX += wordLen + lineSpace*3
			} else {
				wordLen = float64(len(text) * 4)
				curX += wordLen + lineSpace*3
			}
			s := fmt.Sprintf("text=%s ,curX=%.2f ,wordLen=%f", text, curX, wordLen)
			fmt.Println(s)
		case diffmatchpatch.DiffDelete: // 不插入pdf

		}

	}

}

func ComplexReportFooterExecutor2(report *core.Report) {
	content := fmt.Sprintf("第 %v / {#TotalPage#} 页", report.GetCurrentPageNo())
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(textFont).SetMarign(core.Scope{Top: 10})
	footer.SetFontColor("60, 179, 113")
	footer.HorizontalCentered().SetContent(content).GenerateAtomicCell()

}

func ComplexReportHeaderExecutor2(report *core.Report) {

	curX, curY := report.GetXY()
	report.SetXY(curX, curY+10)
	content := "报告编号：0988888888888888"
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(textFont)
	footer.SetFontColor("179,179,204")
	footer.SetBorder(core.Scope{Top: 10, Left: 290})
	footer.SetContent(content).GenerateAtomicCell()

	report.SetXY(curX, curY)
	dir, _ := filepath.Abs("pictures")
	qrcodeFile := fmt.Sprintf("%v/b.png", dir)
	im := gopdf.NewImageWithWidthAndHeight(qrcodeFile, 40, 18, report)
	im.SetMargin(core.Scope{Left: 0, Top: -10})
	im.SetAutoBreak()
	im.GenerateAtomicCell()

}

func TestPara(t *testing.T) {
	ComplexReport2()
}
