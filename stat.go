package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/perf/benchstat"
)

func getTable(name string) [][]string {
	c := benchstat.Collection{}
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	c.AddFile(name, f)
	tmp, err := os.CreateTemp("/tmp", "*.csv")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmp.Name())
	benchstat.FormatCSV(tmp, c.Tables(), false)
	tmp.Seek(0, 0)

	csvReader := csv.NewReader(tmp)
	table, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	tmp.Close()

	return table
}

func getSerializer(str string) string {
	return str[strings.LastIndex(str, "/")+1:]
}

func main() {
	type Record struct {
		serTime   float64
		deserTime float64
		dataSize  float64
	}
	m := make(map[string]Record)
	serTable := getTable("serialization_results.txt")
	deserTable := getTable("deserialization_results.txt")
	sizeTable := getTable("data_results.txt")
	var maxSerTime, maxDeserTime, maxDataSize float64

	for _, record := range serTable[1:] {
		serializer := getSerializer(record[0])
		m[serializer] = Record{}

		if entry, ok := m[serializer]; ok {
			entry.serTime, _ = strconv.ParseFloat(record[1], 64)
			maxSerTime = math.Max(maxSerTime, entry.serTime)
			m[serializer] = entry
		}
	}
	for _, record := range deserTable[1:] {
		serializer := getSerializer(record[0])

		if entry, ok := m[serializer]; ok {
			entry.deserTime, _ = strconv.ParseFloat(record[1], 64)
			maxDeserTime = math.Max(maxDeserTime, entry.deserTime)
			m[serializer] = entry
		}
	}
	for _, record := range sizeTable[1:] {
		serializer := getSerializer(record[0])

		if entry, ok := m[serializer]; ok {
			entry.dataSize, _ = strconv.ParseFloat(record[1], 64)
			maxDataSize = math.Max(maxDataSize, entry.dataSize)
			m[serializer] = entry
		}
	}

	rows := make([][]string, 0)
	for _, record := range serTable[1:] {
		serializer := getSerializer(record[0])
		entry := m[serializer]
		rows = append(rows, []string{
			serializer,
			fmt.Sprintf("%.2fms, %3.0f%%", entry.serTime/1000000, entry.serTime/maxSerTime*100),
			fmt.Sprintf("%.2fms, %3.0f%%", entry.deserTime/1000000, entry.deserTime/maxDeserTime*100),
			fmt.Sprintf("%.2fMB, %3.0f%%", entry.dataSize/1024/1024, entry.dataSize/maxDataSize*100),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Format", "Serialization Time", "Deserialization Time", "Data Size"})
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows)
	table.Render()
}
