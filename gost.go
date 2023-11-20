package main

import (
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	symbols := "aapl,nflx,meta,amzn,tsla,goog"
	res, err := fetchQuery(symbols)

	if err != nil {
		panic(err)
	}

	if err := ui.Init(); err != nil {
		panic(err)
	}

	defer ui.Close()

	charts := []*widgets.BarChart{}

	for s, _ := range res {
		charts = append(charts, mkChart(s, res))
	}

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()

	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/3, charts[0]),
			ui.NewCol(1.0/3, charts[1]),
			ui.NewCol(1.0/3, charts[2]),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/3, charts[3]),
			ui.NewCol(1.0/3, charts[4]),
			ui.NewCol(1.0/3, charts[5]),
		),
	)

	ui.Render(grid)
	<-ui.PollEvents()
}

func mkChart(symbols string, res qMap) *widgets.BarChart {
	bc := widgets.NewBarChart()
	bc.Data = []float64{}
	bc.Labels = []string{}
	bc.BarWidth = 1
	bc.BarGap = 0

	vals := res[symbols]
	var min float64

	for i := len(vals) - 1; i >= 0; i-- {
		price, err := strconv.ParseFloat(vals[i].price, 64)
		if err != nil {
			panic(err)
		}

		bc.Data = append(bc.Data, price)

		if min == 0 || price < min {
			min = price
		}

		bc.Labels = append(bc.Labels, weekday(vals[i].date))
	}

	for i, _ := range bc.Data {
		bc.Data[i] -= min
	}

	bc.NumFormatter = func(f float64) string {
		return ""
	}

	bc.Title = symbols
	bc.BarWidth = 1
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}

	return bc
}

func weekday(date string) string {
	dt, _ := time.Parse("2006-01-02", date)
	return string(dt.Weekday().String()[0])
}
