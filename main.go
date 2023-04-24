package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// App Setup
func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

		// fetchHTMLContent("https://github.com/openshift/hive/blob/v1/docs/architecture.md", textView)
		// URL https://github.com/openshift/ops-sop/blob/master/backporting.md

	fetchHTMLContent("https://github.com/openshift/ops-sop/blob/master/team_guides/Thor/Onboarding.md", textView)

	textView.Highlight("0")

	// Input Handling
	textView.SetDoneFunc(func(key tcell.Key) {
		currentSelection := textView.GetHighlights()
		if len(currentSelection) > 0 {
			index, _ := strconv.Atoi(currentSelection[0])
			if key == tcell.KeyEnter {
				// TODO: Update with the Link
				url := textView.GetRegionText(currentSelection[0])
				fetchHTMLContent(url, textView)
				fmt.Println(url)
			}
			if key == tcell.KeyTab {
				index = (index + 1) % numLinks
			} else if key == tcell.KeyBacktab {
				index = (index - 1 + numLinks) % numLinks
			} else {
				return
			}
			textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
		}
	})

	// Run the App
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}
}
