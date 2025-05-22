package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/Zyprush18/expense-tracker/model"
	"github.com/spf13/cobra"
)

var (
	description string
	amount      int16
)

var expenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Get a Expense Tracker",
	Long:  "this is command to running apps expense tracker",
	Run:   AddExpense,
}

func init() {
	expenseCmd.Flags().StringVarP(&description, "description", "d", "", "description for expense tracker")
	expenseCmd.Flags().Int16VarP(&amount, "amount", "a", 0, "amount for expense tracker")

	expenseCmd.MarkFlagRequired("description")
	expenseCmd.MarkFlagRequired("amount")

	rootCmd.AddCommand(expenseCmd)
}

func AddExpense(cmd *cobra.Command, args []string) {
	var expense []model.ExpenseTracker
	id := 1

	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	read, err := os.ReadFile("data.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	readfiles := string(read)
	
	// search word
	regex := regexp.MustCompile(regexp.QuoteMeta("Id"))
	matches := regex.FindAllString(readfiles, -1)	

	times := time.Now()
	dates := fmt.Sprintf("%d-%d-%d", times.Year(), times.Month(), times.Day())

	writecsv := csv.NewWriter(file)
	defer writecsv.Flush()

	expenseReq := model.ExpenseTracker{
		Id:          strconv.Itoa(id),
		Date:        dates,
		Description: description,
		Amount:      strconv.Itoa(int(amount)),
	}

	expense = append(expense, expenseReq)

	if len(matches) < 1 {
		header := []string{"Id","Date","Description","Amount"}
		if err := writecsv.Write(header);err != nil {
			log.Println(err.Error())
		}
	}

	for _, exp := range expense {
		rowsData := []string{exp.Id, exp.Date, exp.Description, exp.Amount}
		if err := writecsv.Write(rowsData); err != nil {
			log.Println(err.Error())
		}
	}
	fmt.Printf("Expense Added Successfully (ID:%d) \n",id)
}
