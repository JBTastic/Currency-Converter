package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type ApiResponse struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

// fetchCurrencies fetches the currency list from the API
func fetchCurrencies() (map[string]string, error) {
	resp, err := http.Get("https://api.frankfurter.dev/v1/currencies")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var currencies map[string]string
	err = json.NewDecoder(resp.Body).Decode(&currencies)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Currency Converter:\n\n")
		flag.PrintDefaults()
	}

	amount := flag.Float64("amount", 0, "Amount to convert")
	from := flag.String("from", "", "Source currency (e.g. EUR)")
	to := flag.String("to", "", "Target currency (e.g. USD)")
	listCurrencies := flag.Bool("currencies", false, "List all available currencies")
	abbr := flag.String("abbr", "", "Search for currency codes by name substring")

	flag.Parse()

	// If user wants to list all currencies or search by substring, do that and exit
	if *listCurrencies || *abbr != "" {
		currencies, err := fetchCurrencies()
		if err != nil {
			fmt.Println("❌ Error fetching currencies:", err)
			return
		}

		if *listCurrencies {
			fmt.Println("Available currencies:")
			for code, name := range currencies {
				fmt.Printf("  %s: %s\n", code, name)
			}
			return
		}

		if *abbr != "" {
			query := strings.ToLower(*abbr)
			found := false
			fmt.Printf("Searching currencies containing '%s':\n", query)
			for code, name := range currencies {
				if strings.Contains(strings.ToLower(name), query) {
					fmt.Printf("  %s: %s\n", code, name)
					found = true
				}
			}
			if !found {
				fmt.Println("  No matches found.")
			}
			return
		}
	}

	// If not listing or searching, run the currency conversion
	if *from == "" || *to == "" || *amount <= 0 {
		fmt.Println("❌ Usage: currency-converter -amount 100 -from EUR -to USD")
		return
	}

	fromUpper := strings.ToUpper(*from)
	toUpper := strings.ToUpper(*to)
	url := fmt.Sprintf("https://api.frankfurter.app/latest?amount=%f&from=%s&to=%s", *amount, fromUpper, toUpper)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Error fetching exchange rates:", err)
		return
	}
	defer resp.Body.Close()

	var result ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("❌ Error parsing response:", err)
		return
	}

	rate, ok := result.Rates[toUpper]
	if !ok {
		fmt.Printf("❌ Invalid source or target currency: %s | %s\n", fromUpper, toUpper)
		return
	}

	// Format the date
	input := result.Date
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("❌ Error parsing date:", err)
		return
	}

	output := t.Format("02.01.2006")
	fmt.Printf("💱 %.2f %s → %.2f %s (Date: %s)\n", *amount, fromUpper, rate, toUpper, output)
}
