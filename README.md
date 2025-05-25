
# 🧾 Expense Tracker CLI (Golang)

A simple command-line **Expense Tracker** application built in **Go**, storing expenses in a **CSV file** (`data.csv`).  
Supports adding, updating, deleting, listing, and summarizing expenses by month.

---

## 🛠 Features

- Add new expenses with description and amount.
- Update existing expenses by ID.
- Delete expenses by ID.
- List all expenses in a nicely formatted table.
- Show total expenses overall or filtered by month (current year).
- Data stored in `data.csv` file.

---

## 📦 Installation

1. Clone the repo or copy the source code:

```bash
git clone https://github.com/yourusername/expense-tracker-go.git
cd expense-tracker-go
```

2. Install dependencies with Go Modules:

```bash
go mod init expense-tracker
go get github.com/spf13/cobra
go get github.com/gocarina/gocsv
go get github.com/olekukonko/tablewriter
```

3. Build the executable:

```bash
go build -o expense-tracker main.go
```

4. Run the CLI commands:

```bash
./expense-tracker <command> [flags]
```

---

## 🚀 Usage

### Add Expense

```bash
./expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
```

### Update Expense

```bash
./expense-tracker update --id 1 --description "Brunch" --amount 25
# Expense updated successfully (ID: 1)
```

### Delete Expense

```bash
./expense-tracker delete --id 2
# Expense deleted successfully
```

### List Expenses

```bash
./expense-tracker list
```

Example output:

```
┌────┬───────────┬─────────────┬────────┐
│ ID │   DATE    │ DESCRIPTION │ AMOUNT │
├────┼───────────┼─────────────┼────────┤
│ 1  │ 2025-05-25│ Lunch       │ 20     │
└────┴───────────┴─────────────┴────────┘
```

If no expenses exist:

```
Please add expense before you show list expense tracker.
```

### Summary

```bash
./expense-tracker summary
# Total expenses: $30
```

### Summary by Month

```bash
./expense-tracker summary --month 5
# Total expenses for May: $20
```

---

## 📁 Data Format (`data.csv`)

CSV file with header:

```csv
id,date,description,amount
1,2025-05-25,Lunch,20
2,2025-05-25,Dinner,10
```

---

## Dependencies

- [cobra](https://github.com/spf13/cobra) for CLI commands
- [gocsv](https://github.com/gocarina/gocsv) for CSV handling
- [tablewriter](https://github.com/olekukonko/tablewriter) for terminal tables
- Go standard library

---

## Project Idea

This project is inspired by the <a href="https://roadmap.sh/projects/expense-tracker">Expense Tracker Project on Roadmap.sh</a>, designed to practice building CLI apps using Go.
