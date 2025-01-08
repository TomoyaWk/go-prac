package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const CSV_PATH = "./csv/expense.csv"

// 支払い
type Expense struct {
	ID          int
	Date        time.Time
	Amount      int
	Description string
}

// Get list
func GetExpenses() {
	expenses, err := ParseCSV()
	if err != nil {
		log.Fatal(err)
	}

	header := color.New(color.FgGreen, color.Underline).SprintfFunc()
	table := table.New("ID", "Date", "Amount", "Description")
	table.WithHeaderFormatter(header)

	//表示
	for _, ex := range expenses {
		table.AddRow(ex.ID, ex.Date.Format("2006-01-02"), ex.Amount, ex.Description)
	}
	table.Print()
}

// Create Expense
func CreateExpense(description string, amount int) (Expense, error) {
	read, err := ParseCSV()
	if err != nil {
		log.Fatal(err)
		return Expense{}, err
	}
	var created = Expense{
		ID:          len(read) + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}
	expenses := append(read, created)
	if err := SaveCSV(expenses); err != nil {
		return Expense{}, err
	}
	return created, nil
}

// csvファイル読み込みでexpensesを返す
func ParseCSV() ([]Expense, error) {
	csvFile, err := os.Open(CSV_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	//shift-jis　->　utf-8変換
	r := csv.NewReader(
		transform.NewReader(csvFile, japanese.ShiftJIS.NewDecoder()),
	)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	expenses := []Expense{}
	for i, row := range rows {
		if i == 0 {
			continue //ヘッダーをスキップ
		}
		//ID
		id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal(err)
		}
		//日付
		date, err := time.Parse("2006-01-02", row[1])
		if err != nil {
			log.Fatal(err)
		}
		//金額
		amount, err := strconv.Atoi(row[2])
		if err != nil {
			log.Fatal(err)
		}
		//expense変換
		expenses = append(expenses, Expense{
			ID:          id,
			Date:        date,
			Amount:      amount,
			Description: row[3],
		})
	}
	return expenses, nil
}

// CSV書き込み
func SaveCSV(expenses []Expense) error {
	csvFile, err := os.Create(CSV_PATH)
	if err != nil {
		log.Fatal(err)
		return err
	}

	writer := csv.NewWriter(
		transform.NewWriter(csvFile, japanese.ShiftJIS.NewEncoder()),
	)
	var headerRow = []string{"ID", "Date", "Amount", "Description"}
	writer.Write(headerRow)
	for _, e := range expenses {
		err := writer.Write([]string{
			strconv.Itoa(e.ID),
			e.Date.Format("2006-01-02"),
			strconv.Itoa(e.Amount),
			e.Description,
		})
		if err != nil {
			return err
		}
	}
	writer.Flush()
	defer csvFile.Close()
	return nil
}
