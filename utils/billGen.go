package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	I "github.com/mohdjishin/GoCart/interfaces"
	"github.com/mohdjishin/GoCart/model"
)

// func init() {
// 	currentTime := time.Now()

// 	date := currentTime.Format("06-Jan-02")
// 	// fmt.Println(currentTime)

// 	m := pdf.NewMaroto(consts.Portrait, consts.A4)

// 	m.SetPageMargins(20, 10, 20)
// 	buildHeading(date, m)
// 	BuildFruiteList(m)
// 	BuildTotal("9000", m)

// 	err := m.OutputFileAndClose("pdf/sample.pdf")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("PDF saved successfully")

// }

type BillGenerator struct{}

func NewBillGenerator() I.IBillGenerator {
	return &BillGenerator{}
}

func (*BillGenerator) GenerateInvoice(bill model.Invoice) string {

	name := uuid.New().String()
	fmt.Println(bill)
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	currentTime := time.Now()

	date := currentTime.Format("06-Jan-02")

	m.SetPageMargins(20, 10, 20)
	buildHeading(date, bill.Name, bill.Phone, m)
	BuildProListeList(m, bill)

	BuildTotal(bill.Total, m)

	err := m.OutputFileAndClose("media/pdf/" + name + ".pdf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PDF saved successfully")
	return name + ".pdf"
}

func BuildTotal(total string, m pdf.Maroto) {

	m.Row(30, func() {
		m.Text("total : "+total, props.Text{
			Top:   5,
			Right: 10,
			Align: consts.Right,
			Color: color.NewBlack(),
			Style: consts.BoldItalic,
		})
	})
}

func buildHeading(date string, name string, phone string, m pdf.Maroto) {
	m.RegisterHeader(func() {

		m.Row(50, func() {
			m.Col(12, func() {

				err := m.FileImage("media/images/logo.png", props.Rect{
					Center:  true,
					Percent: 100,
				})
				if err != nil {
					fmt.Println(err)
				}
			})
		})

	})
	m.Row(10, func() {
		m.Text("Date :"+date, props.Text{
			Top:   3,
			Style: consts.Bold,
			Align: consts.Left,
			Color: color.NewBlack(),
		})
	})
	m.Row(10, func() {
		m.Text("Name :"+name, props.Text{
			Top:   3,
			Style: consts.Bold,
			Align: consts.Left,
			Color: color.NewBlack(),
		})
	})
	m.Row(10, func() {
		m.Text("Phone :"+phone, props.Text{
			Top:   3,
			Style: consts.Bold,
			Align: consts.Left,
			Color: color.NewBlack(),
		})
	})

	m.Row(10, func() {
		m.Text("Invoice", props.Text{
			Top:   3,
			Style: consts.BoldItalic,
			Align: consts.Center,
			Color: darkPurpleColor(),
		})
	})
}

func darkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}

}

func BuildProListeList(m pdf.Maroto, bill model.Invoice) {
	tableHeadings := []string{"order ID", "Product name", "Quantity", "unit price", "total amount"}

	contents := [][]string{{bill.OrderId, bill.ProductName, string(bill.Quantity), bill.Price, bill.Total}}
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {

			m.Text("Product", props.Text{
				Top:    2,
				Size:   12,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 4, 2, 2, 2},
		},
		ContentProp: props.TableListContent{

			Size:      9,
			GridSizes: []uint{2, 4, 2, 2, 2},
		},
		Align:              consts.Left,
		HeaderContentSpace: 1,
		Line:               true,
		// AlternatedBackground: &ligtPurpleColor,
	})

}

func ligtPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{

		Red:   3,
		Green: 166,
		Blue:  166,
	}

}
