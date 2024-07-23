package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url string, filepath string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to download file: %s, status code: %d", url, resp.StatusCode)
    }

    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

func main() {
    csvFile, err := os.Open("your_exported_file.csv")
    if err != nil {
        fmt.Printf("Error opening CSV file: %v\n", err)
        return
    }
    defer csvFile.Close()

    reader := csv.NewReader(csvFile)
    headers, err := reader.Read() // Read the header line
    if err != nil {
        fmt.Printf("Error reading CSV header: %v\n", err)
        return
    }

    // Find the indexes of the relevant columns
    var idIndex, urlIndex int
    for i, header := range headers {
        if header == "Id" {
            idIndex = i
        } else if header == "VersionData" {
            urlIndex = i
        }
    }

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Printf("Error reading CSV record: %v\n", err)
            continue
        }

        id := record[idIndex]
        url := record[urlIndex]

        if url == "" {
            fmt.Printf("Skipping empty URL for ID: %s\n", id)
            continue
        }

        // Ensure URL is properly formatted
        if !filepath.IsAbs(url) {
            url = "https://jokr6.lightning.force.com/" + url
        }

        filepath := fmt.Sprintf("%s.pdf", id)
        err = downloadFile(url, filepath)
        if err != nil {
            fmt.Printf("Error downloading file %s: %v\n", id, err)
        } else {
            fmt.Printf("Successfully downloaded file: %s\n", filepath)
        }
    }
}
