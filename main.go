package gogo_dorks

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	tld "golang.org/x/net/publicsuffix"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

func parseArgs() (domain string, results int, output string) {
	flag.StringVar(&domain, "domain", "", "Domain to scan (required)")
	flag.IntVar(&results, "results", 10, "Number of results per search, default 10")
	flag.StringVar(&output, "output", "", "Output file")
	flag.Parse()

	if domain == "" {
		log.Fatal("The --domain (-d) flag is required.")
	}

	return
}

func save(file, data string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(data + "\n"); err != nil {
		log.Fatal(err)
	}
}

func GoogleSearch(query string) ([]string, error) {
	results := []string{}

	// Google search query
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(query))
	client := &http.Client{}

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return results, err
	}

	req.Header.Set("User-Agent", randomUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return results, err
	}

	sel := doc.Find("div#search .g .r a")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		href, exists := item.Attr("href")
		if exists {
			results = append(results, href)
		}
	}

	return results, nil
}

func randomUserAgent() string {
	userAgents := []string{
		// List of User Agents
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		// Add more user agents if needed
	}
	return userAgents[rand.Intn(len(userAgents))]
}

func main() {
	domain, amount, outputFile := parseArgs()
	rand.Seed(time.Now().UnixNano())

	// Extract target domain
	_, err := tld.PublicSuffix(domain)
	if err != false {
		log.Fatal(err)
	}

	// List of Google dorks (shortened for example; add the full dorks as needed)
	dorks := map[string]string{
		// "# .git folders": "inurl:\"/.git\" " + domain + " -github",
		// Add more dorks as needed
	}

	for description, dork := range dorks {
		fmt.Println("\n" + description + "\n")
		if outputFile != "" {
			save(outputFile, description)
		}

		results, err := GoogleSearch(dork)
		if err != nil {
			log.Printf("Error searching for dork '%s': %v", dork, err)
			continue
		}

		for i, result := range results {
			fmt.Println(result)
			if outputFile != "" {
				save(outputFile, result)
			}

			if (i + 1) >= amount {
				break
			}

			// Randomize sleep time
			time.Sleep(time.Duration(rand.Intn(15-1)+1) * time.Second)
		}
	}
}
