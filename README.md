# 💸 Budget Tracker CLI

A simple and idiomatic terminal-based budget tracker written in Go, using [Cobra](https://github.com/spf13/cobra). It helps you record income and expenses, summarize your financial activity, delete transactions, and export your data to CSV.

---

## 🛠 Features

- ✅ Add income or expense entries via command line
- 📊 View a summary of your total income, expenses, and balance
- 📋 List all transactions with details
- ❌ Delete a transaction by its index
- 📤 Export all entries to CSV format
- 💾 Data stored locally as a JSON file

---

## 🧱 Project Structure

budget-tracker/
├── cmd/ # All Cobra command definitions
│ ├── add.go
│ ├── summary.go
│ ├── list.go
│ ├── delete.go
│ └── export.go
├── data/ # Data access layer
│ └── store.go # loadEntries, saveEntry, saveEntries
├── model/ # Data model
│ └── entry.go # Entry struct definition
├── main.go # App entry point
└── entries.json # Your saved transactions

yaml
Copy code

---

## 🚀 Getting Started

### **Clone the repo**

```bash
git clone https://github.com/your-username/budget-tracker.git
cd budget-tracker
2. Run the CLI
bash
Copy code
go run main.go <command> [flags]
📦 Commands
➕ Add a new entry
bash
Copy code
go run main.go add --amount 200 --type income --category freelance
Flags:

--amount (float) – required

--type (income or expense) – required

--category (string) – optional

📋 List all entries
bash
Copy code
go run main.go list
Outputs a formatted table of all entries with index, date, amount, type, and category.

❌ Delete an entry
bash
Copy code
go run main.go delete --index 1
Flags:

--index or -i – index of the entry to delete

📊 Show summary
bash
Copy code
go run main.go summary
Displays total income, total expenses, and the net balance.

📤 Export to CSV
bash
Copy code
go run main.go export
Optional flag:

bash
Copy code
--output entries.csv
Exports all entries to entries.csv or your chosen file path.

📁 Data Storage
All entries are stored in a local file called entries.json in the project root. You can back this file up or edit it manually if needed (carefully).

✍️ Example Workflow
bash
Copy code
# Add a freelance income
go run main.go add --amount 500 --type income --category freelance

# Add a grocery expense
go run main.go add --amount 75.50 --type expense --category groceries

# List entries
go run main.go list

# Show summary
go run main.go summary

# Delete an entry
go run main.go delete --index 0

# Export to CSV
go run main.go export --output report.csv
🧪 Requirements
Go 1.18+

Terminal access

Local file permissions (to read/write entries.json and .csv)

 Todo / Ideas
 Edit existing entries

 Filter by date or category

 TUI version using BubbleTea

 Cloud sync or database backend

feel free to fork, improve, or build on it.

 Contributions
Contributions welcome! Open an issue or PR if you'd like to help improve the tool.

 Author
Built with hate by Morizz00
