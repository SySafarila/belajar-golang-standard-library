package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"syahrul", "Safarila", "SySafarila"})
	_ = writer.Write([]string{"syahrul2", "Safarila2", "SySafarila2"})
	writer.Flush()
}
