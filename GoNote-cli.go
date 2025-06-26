package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

var (
	file      = "notes.csv"
	true_emo  = "\033[32m\u2705\033[0m"
	false_emo = "\033[31m\u274C\033[0m"
	x         string
	ch        string
	ls        = []string{}
	ps        = []string{}
	bs        = "========================================="
)

func main() {
	check_file()
	main_start()
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func last_number() (int, [][]string) {
	file, err := os.Open(file)
	handleError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	handleError(err)

	// Check if records is empty
	if len(records) == 0 {
		return 0, records // Return 0 if there are no records
	}

	num, err := strconv.Atoi(records[len(records)-1][0])
	handleError(err) // Handle conversion error
	return num, records
}

func contains(slice []string, value string) bool {
	// Normalize the value by trimming spaces and underscores
	normalizedValue := strings.TrimSpace(strings.ReplaceAll(value, " ", "_"))

	for _, v := range slice {
		// Normalize the existing task as well
		normalizedTask := strings.TrimSpace(v)
		if normalizedTask == normalizedValue {
			return true
		}
	}
	return false
}

func DrawIt(records [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Task", "Sts", "Timestamp", "Completed At"})

	// Set table style
	table.SetBorder(true) // Add border

	// Add colored header
	table.SetHeader([]string{
		"\033[36mID\033[0m",        // Cyan
		"\033[36mTask\033[0m",      // Cyan
		"\033[36mSts\033[0m",       // Cyan
		"\033[36mTimestamp\033[0m", // Cyan
	})

	// Add records to the table
	for _, v := range records {
		table.Append(v)
	}

	// Render the table
	table.Render()
}

func added_note() {
	filexs, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	handleError(err)
	defer filexs.Close()

	fmt.Print(
		"Enter your tasks separated by , : ",
	)

	reader := bufio.NewReader(os.Stdin)
	x, _ = reader.ReadString('\n') // Read until newline
	x = strings.ReplaceAll(x, ";", ",")
	words := strings.Split(x, ",")

	for _, i := range words {
		i = strings.TrimSpace(i)
		i = strings.ReplaceAll(i, " ", "_")
		ps = append(ps, i)
	}

	no, records := last_number()
	for _, i := range records {
		ls = append(ls, i[1])
	}
	for _, i := range ps {
		if !contains(ls, i) {
			no = no + 1
			ts := time.Now()
			tsforma := ts.Format("2006-01-02 15:04:05")
			ds := fmt.Sprintf(
				"%d,%s,%s,%s,0\n",
				no,
				i,
				false_emo,
				tsforma,
			)
			_, err = filexs.WriteString(ds)
			handleError(err)
			// fmt.Print(ds)
		}
	}
}

func check_file() {
	_, err := os.Stat(file)
	if err != nil {
		_, err := os.Create(file)
		handleError(err)
	}
}

func bar() {
	fmt.Println("╔════════════════════════╗")
	fmt.Println("║  To-Do List App In Go  ║")
	fmt.Println("╚════════════════════════╝")
}

func comple_task(rec [][]string) {
	filee, err := os.Open(file)
	handleError(err)

	defer filee.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(filee)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var ask string
	DrawIt(rec)
	fmt.Print("Enter your task id : ")
	fmt.Scanln(&ask)
	num, err := strconv.Atoi(ask)
	handleError(err)
	if len(lines) >= num {
		if string(lines[num-1][len(lines[num-1])-1]) == "0" {
			ts := time.Now()
			tsforma := ts.Format("2006-01-02 15:04:05")
			newl := lines[num-1][:len(lines[num-1])-1] + tsforma
			newl = strings.ReplaceAll(newl, false_emo, true_emo)
			lines[num-1] = newl
			fmt.Printf("Completed %s\n", true_emo)
		} else if string(lines[num-1][len(lines[num-1])-1]) != "0" {
			fmt.Println("The task is already toggled to completed")
		}
		tempFile, err := os.Create(file)
		handleError(err)
		defer tempFile.Close()

		for _, line := range lines {
			_, err := tempFile.WriteString(line + "\n")
			handleError(err)
		}
	} else if !(len(lines) >= num) {
		fmt.Printf("your id \"%d\" is out of tasks range !!\n", num)
	}
	fmt.Println(bs)
	// Check if the line number is valid
}

func main_start() {
	for {
		bar()
		fmt.Print(
			"1) added notes \n2) show notes\n3) complete a task\n4) clear Terminal\n5) clear notes db\n6) exit\n#!> ",
		)
		fmt.Scanln(&ch)

		switch ch {
		case "1":
			added_note()
			fmt.Printf("Operation Completed\n%s\n", bs)
		case "2":
			_, res := last_number()
			DrawIt(res)
			fmt.Printf("Operation Completed\n%s\n", bs)
		case "3":
			_, rec := last_number()
			comple_task(rec)
		case "4":
			fmt.Printf("\033c")
		case "5":
			var ques string
			fmt.Printf(
				"Are you sure ? %q (y)es/(n)o  : ",
				"all notes will be deleted ",
			)
			fmt.Scanln(&ques)
			ques = strings.ToLower(ques)
			if ques == "y" || ques == "yes" {
				os.Remove(file)
				fmt.Printf("notes db : %q has been cleared\n\n", file)
			} else if ques == "n" || ques == "no" {
				fmt.Println("ok :)")
			}
			check_file()
		case "6":
			fmt.Println("Exiting ...\nSee you later USER :)")
			return
		default:
			fmt.Println("Invalid Output")
		}
	}
}
