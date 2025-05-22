package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Zyprush18/expense-tracker/model"
	"github.com/spf13/cobra"
)

var (
	ids         int16
	description string
	amount      int16
)

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Expense Tracker",
	Long:  "this is command to add expense tracker",
	Run:   AddExpense,
}

var updateExpenseCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a Expense Tracker",
	Long:  "This is command to update expense by id",
	Run:   UpdateExpense,
}

func init() {
	// flags add
	addExpenseCmd.Flags().StringVarP(&description, "description", "d", "", "description for expense tracker")
	addExpenseCmd.Flags().Int16VarP(&amount, "amount", "a", 0, "amount for expense tracker")

	// require flags add
	addExpenseCmd.MarkFlagRequired("description")
	addExpenseCmd.MarkFlagRequired("amount")

	// flags update
	updateExpenseCmd.Flags().Int16VarP(&ids, "id", "i", 0, "Id Data Expense Tracker")
	updateExpenseCmd.Flags().StringVarP(&description, "description", "d", "", "description for expense tracker")
	updateExpenseCmd.Flags().Int16VarP(&amount, "amount", "a", 0, "amount for expense tracker")

	// require flags update
	updateExpenseCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(addExpenseCmd)
	rootCmd.AddCommand(updateExpenseCmd)
}

func AddExpense(cmd *cobra.Command, args []string) {
	var expense []model.ExpenseTracker
	var id int

	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	data := strings.Join(lines, ", ")

	regex := regexp.MustCompile(regexp.QuoteMeta("Id"))
	matches := regex.FindAllString(data, -1)

	times := time.Now()
	dates := fmt.Sprintf("%d-%d-%d", times.Year(), times.Month(), times.Day())

	writecsv := csv.NewWriter(file)
	defer writecsv.Flush()

	if len(lines) < 1 {
		id = 1
	} else {
		id = len(lines)
	}

	expenseReq := model.ExpenseTracker{
		Id:          strconv.Itoa(id),
		Date:        dates,
		Description: description,
		Amount:      strconv.Itoa(int(amount)),
	}

	expense = append(expense, expenseReq)

	if len(matches) < 1 {
		header := []string{"Id", "Date", "Description", "Amount"}
		if err := writecsv.Write(header); err != nil {
			log.Println(err.Error())
		}
	}

	for _, exp := range expense {
		rowsData := []string{exp.Id, exp.Date, exp.Description, exp.Amount}
		if err := writecsv.Write(rowsData); err != nil {
			log.Println(err.Error())
		}
	}
	fmt.Printf("Expense Added Successfully (ID:%d) \n", id)
}

// update
func UpdateExpense(cmd *cobra.Command, args []string) {
	fmt.Println("Ini update Expense tracker")
}
