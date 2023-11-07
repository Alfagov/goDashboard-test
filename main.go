package main

import (
	"github.com/Alfagov/goDashboard/dashboard"
	"github.com/Alfagov/goDashboard/layout"
	"github.com/Alfagov/goDashboard/models"
	"github.com/Alfagov/goDashboard/pages"
	"github.com/Alfagov/goDashboard/pkg/form"
	"github.com/Alfagov/goDashboard/pkg/graph"
	"github.com/Alfagov/goDashboard/pkg/graphContainer"
	"github.com/Alfagov/goDashboard/pkg/numeric"
	"github.com/Alfagov/goDashboard/pkg/widgets"
	"github.com/Alfagov/goDashboard/toolbox"
	"log"
	"math/rand"
	"time"
)

func main() {
	dashboard.InitDashboardGlobals()
	wd := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test1"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(1),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	wd1 := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test2"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(2),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	wd2 := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test3"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(3),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	wd3 := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test4"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(3),
				layout.SetRow(2),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	wd11 := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test5"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(2),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	wd21 := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test6"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(3),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	wd31 := numeric.NewNumeric(
		time.Second*5,
		widgets.SetName("test7"),
		widgets.SetDescription("test"),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(3),
				layout.SetRow(2),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		),
	).WithSettings(
		numeric.SetNumericInitValue(1),
		numeric.SetNumericUpdateHandler(
			func() (int, error) {
				return rand.Intn(500), nil
			},
		),
	)

	gww := graphContainer.NewGraphWidget(
		graph.NewLineGraph(
			"lg1", func() *models.LineGraphData {
				return &models.LineGraphData{
					XAxis: []string{"test", "test1", "test2", "test3"},
					YAxis: []models.LineYAxis{
						{
							Name: "label1",
							Data: []int{rand.Intn(30), rand.Intn(30),
								rand.Intn(30), rand.Intn(30)},
						},
						{
							Name: "label2",
							Data: []int{rand.Intn(30), rand.Intn(30),
								rand.Intn(30), rand.Intn(30)},
						},
					},
				}
			},
		).WithToolboxOpts(
			toolbox.NewToolbox(
				toolbox.SetZoom(false),
				toolbox.WithSaveImage("test"),
			),
		),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(1), layout.SetRow(4),
				layout.SetWidth(3), layout.SetHeight(2),
			),
		),
	)

	gbw := graphContainer.NewGraphWidget(
		graph.NewBarGraph(
			"bg1", func() *models.BarGraphData {
				return &models.BarGraphData{
					XAxis: []string{"test", "test1", "test2", "test3"},
					YAxis: []models.BarYAxis{
						{
							Name: "label1",
							Data: []int{rand.Intn(30), rand.Intn(30),
								rand.Intn(30), rand.Intn(30)},
						},
						{
							Name: "label2",
							Data: []int{rand.Intn(30), rand.Intn(30),
								rand.Intn(30), rand.Intn(30)},
						},
					},
				}
			},
			false,
		).WithToolboxOpts(
			toolbox.NewToolbox(
				toolbox.SetZoom(false),
				toolbox.WithSaveImage("test"),
			),
		),
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(1), layout.SetRow(4),
				layout.SetWidth(3), layout.SetHeight(2),
			),
		),
	)

	type formTest struct {
		Name string `label:"NameLabel" type:"select-remote" `
		Age  int    `label:"AgeLabel" type:"number" `
	}

	f := form.SubmitField("submit", "submit")

	fw1 := form.NewFormWidget[formTest]("testForm",
		widgets.SetLayout(
			layout.NewWidgetLayout(
				layout.SetColumn(1),
				layout.SetRow(1),
				layout.SetHeight(1),
				layout.SetWidth(1),
			),
		)).
		WithSettings(
			form.SetFormUpdateHandler[formTest](
				func(test formTest) *models.UpdateResponse {
					return &models.UpdateResponse{
						Success: true,
						Message: "test",
					}
				}),
			form.WithSelectHandler[formTest](
				"Name", func(s string) []string {
					return []string{"test", "test1", "test2"}
				},
			),
			form.AddFormFields[formTest](
				&f,
			),
		)

	pg := pages.NewPage(
		"test",
	).WithWidgets(gww, wd21, wd31, wd11, fw1)

	pg1 := pages.NewPage(
		"test1",
	).WithWidgets(gbw, wd, wd1, wd2, wd3)

	pgContainer := pages.NewPageContainer(
		"testContainer",
		pages.SetIndexPage("test"),
	).WithPages(pg, pg1)

	d := dashboard.NewDashboard("Dashboard", "https://www.svgrepo.com/show/217594/house-key-key.svg").
		WithPages(pgContainer)

	log.Fatal(d.Run())
}
