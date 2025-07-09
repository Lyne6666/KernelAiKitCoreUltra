// cmd/kernelaikitcoreultra/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"kernelaikitcoreultra/internal/kernelaikitcoreultra"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := kernelaikitcoreultra.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
