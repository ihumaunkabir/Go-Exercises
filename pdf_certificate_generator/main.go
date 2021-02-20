package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func generateCertificate(name, university, team string) time.Duration {

	begin := time.Now()
	m := pdf.NewMaroto(consts.Landscape, consts.A4)

	//header of the certificate
	m.Row(20, func() {
		m.Col(3, func() {
			_ = m.FileImage("images/bdosn.png", props.Rect{
				Percent: 80,
				Center:  true,
			})
		})
		m.Col(6, func() {
			m.Text("Ada Lovelace Datathon 2021", props.Text{
				Top:   6,
				Align: consts.Center,
				Size:  20,
				Style: consts.BoldItalic,
			})
		})
		m.Col(3, func() {
			_ = m.FileImage("images/herwill.jpg", props.Rect{
				Percent: 110,
				Center:  true,
			})
		})
	})

	//content of the certificate
	m.Row(130, func() {
		m.Col(12, func() {
			text := "This certificate is awarded to Ms. " + name + " from " + university + " for participating as a member of team " + team + " in Ada Lovelace Datathon 2021 a part of Ada Lovelace Celebration 2021. The 48 hours datathon was held on January 26, 2021 to January 28, 2021. We wish her all yard success."
			m.Text(text, props.Text{
				Size:            15,
				Align:           consts.Center,
				Top:             50,
				VerticalPadding: 2.0,
			})
		})
	})

	//footer of the certificate
	m.Row(1, func() {

		m.Col(4, func() {
			m.Text("Munir Hasan", props.Text{
				Top:   1,
				Align: consts.Center,
				Size:  15,
				Style: consts.BoldItalic,
			})
		})

		m.Col(4, func() {
			m.Text("", props.Text{
				Top:   1,
				Align: consts.Center,
				Size:  10,
				Style: consts.BoldItalic,
			})

		})

		m.Col(4, func() {
			m.Text("Farhana Hasan", props.Text{
				Top:   1,
				Align: consts.Center,
				Size:  15,
				Style: consts.BoldItalic,
			})

		})

	})

	m.Row(10, func() {
		m.Col(4, func() {
			m.Signature("General Secretary, Bangladesh Open Source Network - BdOSN")
		})

		m.Col(4, func() {
			m.Text("", props.Text{
				Top:   0,
				Align: consts.Center,
				Size:  20,
				Style: consts.BoldItalic,
			})
		})

		m.Col(4, func() {
			m.Signature("Founder and CEO, HerWILL")
		})
	})

	//saving the *.pdf
	err := m.OutputFileAndClose("certificates/" + name + "_" + university + "_" + team + ".pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()

	return end.Sub(begin)

}

func main() {

	//read the *.csv
	csvfile, err := os.Open("input/certdata.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	n := 0

	//traversing every row
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//Ignoring the header column in *.csv
		if n != 0 {

			gtime := generateCertificate(record[1], record[2], record[0])
			fmt.Println(gtime)
		}

		n++
	}

}
