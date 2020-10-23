package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"os"
	"time"
)

func sendMail(receiver, name string) error {

	from := "yourEmail"
	password := "yourPassword"
	host := "smtp.gmail.com" //as per your email provider
	port := "587"

	body := "From: " + from + "\n" +
		"To: " + receiver + "\n" +
		"Subject: [NHSPC 2020] Thank You\n\n" +

		`Dear Mr. ` + name + `,

Greetings from National High School Programming Contest 2020. The prestigious programming contest for the high school aspiring students of Bangladesh has been completed successfully this year. It's our immense pleasure to have you in our team as an honorable Judge of NHSPC.
We want to say "Thank You" for your support and contribution to make a problem solving generation for upcoming days. Thank you Mr. ` + name + `.

Regards,
Humaun Kabir
http://ihumaun.com
Academic Coordinator, National High School Programming Contest 2020
`

	return smtp.SendMail(

		host+":"+port,
		smtp.PlainAuth("", from, password, host),
		from,
		[]string{receiver},
		[]byte(body))

}

func main() {

	csvfile, err := os.Open("input_judge.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	n := 0

	f, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if n > 0 { //Ignoring the column header in csv

			mailerr := sendMail(record[2], record[1])

			if mailerr != nil {
				//log.Fatalln("Error occured", mailerr)

				fmt.Println("SL "+record[0]+" Couldn't send email to :", record[2])
				timeString := time.Now().String()
				_, err1 := fmt.Fprintln(f, timeString)

				logs := "SL " + record[0] + " Couldn't send email to :" + record[2] + "\n"
				_, err2 := fmt.Fprintln(f, logs)

				if err1 != nil || err2 != nil {
					fmt.Println("Cannot write to logs.txt")
				}

			} else {
				fmt.Println("Sent to SL: ", record[0], " Email:", record[2])
			}

		}

		n++
	}
}
