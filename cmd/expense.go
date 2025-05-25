package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Zyprush18/expense-tracker/model"
	"github.com/gocarina/gocsv"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	ids         int16
	description string
	amount      int16
	month       int16
)

var listExpenseCmd = &cobra.Command{
	Use:   "list",
	Short: "List a Expense Tracker",
	Long:  "this is command to show list expense tracker",
	Run:   ListExpense,
}

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

var deleteExpenseCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Expense Tracker",
	Long:  "This is command to delete expense by id",
	Run:   DeleteExpense,
}

var summaryExpenseCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary a Expense Tracker",
	Long:  "This is command to summary expense ",
	Run:   SummaryExpense,
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

	// delete flags
	deleteExpenseCmd.Flags().Int16VarP(&ids, "id", "i", 0, "Id Data Expense Tracker")

	// require
	deleteExpenseCmd.MarkFlagRequired("id")

	// summary flags
	summaryExpenseCmd.Flags().Int16VarP(&month, "month", "m", 0, "Sum by month")

	rootCmd.AddCommand(listExpenseCmd)
	rootCmd.AddCommand(addExpenseCmd)
	rootCmd.AddCommand(updateExpenseCmd)
	rootCmd.AddCommand(deleteExpenseCmd)
	rootCmd.AddCommand(summaryExpenseCmd)
}

func ListExpense(cmd *cobra.Command, args []string) {
	info, err := os.Stat("data.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	if info.Size() > 0 {
		expense := []*model.ExpenseTracker{}

		file, err := os.Open("data.csv")
		if err != nil {
			log.Fatalln(err.Error())
		}

		defer file.Close()

		table := tablewriter.NewTable(os.Stdout)

		read := csv.NewReader(file)
		header, err := read.Read()
		if err != nil {
			log.Fatalln(err.Error())
		}

		file.Seek(0, 0)
		if err := gocsv.Unmarshal(file, &expense); err != nil {
			log.Fatalln(err.Error())
		}
		table.Header(header)

		table.Bulk(expense)
		table.Render()

	} else {
		fmt.Println("Please Add expense before you show list expense tracker")
	}
}

func AddExpense(cmd *cobra.Command, args []string) {
	expense := []*model.ExpenseTracker{}
	id := 1

	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {

		log.Fatalln(err.Error())
	}
	defer file.Close()

	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalln(err.Error())
	}

	info, err := os.Stat("data.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var header []string

	if info.Size() > 0 {
		read := csv.NewReader(file)
		var errorr error
		if read != nil {
			header, errorr = read.Read()
			if errorr != nil {
				log.Fatalln(errorr.Error())
			}
		}

		in, err := os.Open("data.csv")
		if err != nil {
			log.Fatalln(err.Error())
		}

		defer in.Close()

		var data []model.ExpenseTracker
		if err := gocsv.Unmarshal(in, &data); err != nil {
			log.Fatalln(err.Error())
		}

		if len(data) > 0 {
			lastdata := data[len(data)-1]
			lastid, err := strconv.Atoi(lastdata.Id)
			if err != nil {
				log.Fatalln(err.Error())
			}

			id = lastid + 1
		}

	}

	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%v-%v-%v", year, int(month), day)

	expensereq := model.ExpenseTracker{
		Id:          strconv.Itoa(id),
		Date:        date,
		Description: description,
		Amount:      strconv.Itoa(int(amount)),
	}

	expense = append(expense, &expensereq)

	if header != nil {
		if err := gocsv.MarshalWithoutHeaders(expense, file); err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		data, err := gocsv.MarshalBytes(expense)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if _, err := file.Write(data); err != nil {
			log.Fatalln(err.Error())
		}
	}

	fmt.Printf("Expense Added Successfully (ID:%d) \n", id)

}

// update
func UpdateExpense(cmd *cobra.Command, args []string) {
	in, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer in.Close()

	expense := []*model.ExpenseTracker{}

	if err := gocsv.Unmarshal(in, &expense); err != nil {
		log.Fatalln(err.Error())
	}

	found := false
	for _, exp := range expense {
		expenseId, err := strconv.Atoi(exp.Id)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if ids == int16(expenseId) {
			found = true
			if description != "" && amount != 0 {
				exp.Description = description
				exp.Amount = strconv.Itoa(int(amount))
			} else if description != "" {
				exp.Description = description
			} else if amount != 0 {
				exp.Amount = strconv.Itoa(int(amount))
			}
		}
	}

	if !found {
		log.Fatalf("expense Tracker for id %d not found \n", ids)
	}

	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	data, err := gocsv.MarshalBytes(expense)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := file.Write(data); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Expense updated Successfully (ID:%d) \n", ids)

}

func DeleteExpense(cmd *cobra.Command, args []string) {
	in, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer in.Close()

	expense := []*model.ExpenseTracker{}

	if err := gocsv.Unmarshal(in, &expense); err != nil {
		log.Fatalln(err.Error())
	}

	expenseAfterDelete := []*model.ExpenseTracker{}
	found := false
	for _, exp := range expense {
		expenseId, err := strconv.Atoi(exp.Id)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if ids != int16(expenseId) {
			found = true
			expenseAfterDelete = append(expenseAfterDelete, &model.ExpenseTracker{
				Id:          exp.Id,
				Date:        exp.Date,
				Description: exp.Description,
				Amount:      exp.Amount,
			})
		}
	}

	if !found {
		log.Fatalf("expense Tracker for id %d not found \n", ids)
	}

	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	data, err := gocsv.MarshalBytes(expenseAfterDelete)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if _, err := file.Write(data); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Expense deleted successfully")
}

func SummaryExpense(cmd *cobra.Command, args []string) {
	expense := []*model.ExpenseTracker{}
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	if err := gocsv.Unmarshal(file, &expense); err != nil {
		log.Fatalln(err.Error())
	}

	result := 0
	for _, exp := range expense {
		angka, err := strconv.Atoi(exp.Amount)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if month != 0 {
			date, err := time.Parse("2006-1-02", exp.Date)
			if err != nil {
				log.Fatalln(err.Error())
			}
			year := time.Now().Year()
			if date.Month() == time.Month(month) && date.Year() == year {
				result += angka
			}

		} else {
			result += angka
		}

	}

	fmt.Printf("Total Expenses: $%d \n", result)

}
