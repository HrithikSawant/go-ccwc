package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Options struct {
	PrintBytes bool
	PrintLines bool
	PrintWords bool
	PrintChars bool
}

type stats struct {
	bytes    uint64
	words    uint64
	lines    uint64
	chars    uint64
	filename string
}

func CalculateStats(reader *bufio.Reader) stats {
	var prevChar rune
	var bytesCount uint64
	var linesCount uint64
	var wordsCount uint64
	var charsCount uint64

	for {
		charRead, bytesRead, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				if prevChar != rune(0) && !unicode.IsSpace(prevChar) {
					wordsCount++
				}
				break
			}
			log.Fatal(err)
		}

		bytesCount += uint64(bytesRead)
		charsCount++

		if charRead == '\n' {
			linesCount++
		}

		if !unicode.IsSpace(prevChar) && unicode.IsSpace(charRead) {
			wordsCount++
		}

		prevChar = charRead
	}

	return stats{bytes: bytesCount, words: wordsCount, lines: linesCount, chars: charsCount}
}

func CalculateStatsWithTotals(reader *bufio.Reader, filename string, options Options, totals *stats) {
	fileStats := CalculateStats(reader)
	fileStats.filename = filename

	fmt.Println(FormatStats(options, fileStats, fileStats.filename))

	totals.lines += fileStats.lines
	totals.words += fileStats.words
	totals.bytes += fileStats.bytes
}

func CalculateStatsForFile(filename string, options Options, totals *stats) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	CalculateStatsWithTotals(reader, filename, options, totals)
}

func CalculateStatsForFiles(filenames []string, options Options) {
	totals := stats{filename: "total"}

	for _, filename := range filenames {
		CalculateStatsForFile(filename, options, &totals)
	}
	if len(filenames) > 1 {
		fmt.Println(FormatStats(options, totals, totals.filename))
	}
}

func maxStatSize(fileStats stats) int {
	maxLen := 0

	lenLines := len(strconv.FormatUint(fileStats.lines, 10))
	if lenLines > maxLen {
		maxLen = lenLines
	}

	lenWords := len(strconv.FormatUint(fileStats.words, 10))
	if lenWords > maxLen {
		maxLen = lenWords
	}

	lenBytes := len(strconv.FormatUint(fileStats.bytes, 10))
	if lenBytes > maxLen {
		maxLen = lenBytes
	}

	lenChars := len(strconv.FormatUint(fileStats.chars, 10))
	if lenChars > maxLen {
		maxLen = lenChars
	}
	return maxLen + 1
}

func FormatStats(commandLineOptions Options, fileStats stats, filename string) string {
	var cols []string

	maxDigits := maxStatSize(fileStats)
	fmtString := fmt.Sprintf("%%%dd", maxDigits)

	if commandLineOptions.PrintLines {
		cols = append(cols, fmt.Sprintf(fmtString, fileStats.lines))
	}
	if commandLineOptions.PrintWords {
		cols = append(cols, fmt.Sprintf(fmtString, fileStats.words))
	}
	if commandLineOptions.PrintBytes {
		cols = append(cols, fmt.Sprintf(fmtString, fileStats.bytes))
	}
	if commandLineOptions.PrintChars {
		cols = append(cols, fmt.Sprintf(fmtString, fileStats.chars))
	}
	statsString := strings.Join(cols, " ") + " " + filename

	return statsString
}
