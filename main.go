package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputFilePath := "block.txt"

	outputFilePath := "outblock.txt"

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("ファイルを開くのに失敗しました。:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("ファイルの作成に失敗しました。:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) >= 2 {

			result := fmt.Sprintf("%s\n", parts[1])

			_, err := outputFile.WriteString(result)
			if err != nil {
				fmt.Println("ファイルの書き出しに失敗しました。:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ファイルの解析に失敗しました。:", err)
		return
	}

	fmt.Println("処理が完了しました。結果は outblock.txt に保存してあります。")
}
