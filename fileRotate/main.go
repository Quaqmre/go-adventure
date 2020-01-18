package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f := "loggerfile"
	numFmt := fmt.Sprintf(".%%0%dd", len(fmt.Sprintf("%d", 50))) //max rotte yapılacak file sayı formatıdır %02d
	_, _ = os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	list, _ := filepath.Glob(fmt.Sprintf("%s.*", f)) // fileları formata göre listeler
	sort.Sort(sort.Reverse(sort.StringSlice(list)))  //fileları harf sırasının tersine göre sıralar
	for _, file := range list {
		parts := strings.Split(file, ".")
		numPart := parts[len(parts)-1]
		num, _ := strconv.Atoi(numPart)                                               //string bir sayıyı integere parse eder
		newName := fmt.Sprintf(strings.Join(parts[:len(parts)-1], ".")+numFmt, num+1) //yeni file ismi belirlenir
		// don't check error because there's nothing we can do
		os.Rename(file, newName) //dosya ismi değiştirilir.
	}

}
