package utils

import "book-trade/config"

//* Tip: same path as writen in html is used in this parameter, just type there and paste here
func BuildStaticURLs(files []string) []string {
    serverAddress := config.GetEnv("SERVER_ADDRESS", "")
    baseURL := "http://" + serverAddress

    fullURLs := make([]string, len(files))
    for i, file := range files {
        fullURLs[i] = baseURL + file
    }

    return fullURLs
}

func BuildSingleStaticURL(filePath string) string {
	serverAddress := config.GetEnv("SERVER_ADDRESS", "")
	return "http://" + serverAddress + filePath
}
