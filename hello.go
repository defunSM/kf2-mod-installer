package main

import (
     "os"
     "io"
     "log"
     "fmt"
     "bufio"
     "strings"
     "net/http"
     "github.com/gen2brain/go-unarr"
)

func DownloadFiles(url string, filename string) {
	newFile, err := os.Create(filename)
     if err != nil {
          log.Fatal(err)
     }
     defer newFile.Close()

     //HTTP GET request 
     response, err := http.Get(url)
     if err != nil { panic(err) }
     defer response.Body.Close()

     // Write bytes from HTTP response to file.
     // response.Body satisfies the reader interface.
     // newFile satisfies the writer interface.
     // That allows us to use io.Copy which accepts
     // any type that implements reader and writer interface
     numBytesWritten, err := io.Copy(newFile, response.Body)
     if err != nil {
          log.Fatal(err)
     }
     log.Printf("Downloaded %d byte file.\n", numBytesWritten)
     // Extracting into current working directory
     ExtractFiles(filename)
}

func CreatePath(path string) string {
     // creates a path given the contents from unarr list function
     currentDirectory, err := os.Getwd()
     if err != nil { panic(err) }
     filePaths := strings.Split(path, "/")
     lengthOfFilePath := len(filePaths)
     filePath := strings.Join(filePaths[:lengthOfFilePath-1], "/")
     completePath := currentDirectory + "/" + filePath

     return completePath
}

func ExtractFiles(filename string) {
     //opening the 7z file
     a, err := unarr.NewArchive(filename)
     if err != nil { panic(err) }
     defer a.Close()

     data, err := a.List()
     if err != nil { panic(err) }
     for _, file := range data {
          
          fmt.Println(file)
          err := a.EntryFor(file)
          if err != nil { panic(err) }

          // returns path as string that needs to be passed to MkdirAll
          completePath := CreatePath(file)
          _ = os.MkdirAll(completePath, os.ModePerm)

          currentEntry := make([]byte, 1000000)
          currentEntry, err = a.ReadAll()
          if err != nil { panic(err) }
          
          // Create enmpty file
          newFile, err := os.Create(file)
          if err != nil { panic(err) }
          defer newFile.Close()
          
          // Opens file
          bytefile, err := os.OpenFile(file, os.O_WRONLY, 0666)
          if err != nil { log.Fatal(err) }
          defer bytefile.Close()
          
          // Creates a buffer for the bytes to be written to
          bufferedWriter := bufio.NewWriter(bytefile)
          bytesWritten, err := bufferedWriter.Write(currentEntry,)
          if err != nil { log.Fatal(err) }

          // commiting the changes and displaying
          bufferedWriter.Flush()
          fmt.Printf("Bytes: %d\n", bytesWritten)

     }
}

func readFirstLine(filepath string) ([]string){
     
     file, err := os.Open(filepath)
     if err != nil { panic(err) }
     defer file.Close()
     
     var filecontents []string

     scanner := bufio.NewScanner(file)
     for scanner.Scan() {
          filecontents = append(filecontents, scanner.Text())
     }
     fmt.Println(filecontents)
     return filecontents
}

func main() {
     // Downloads files from a link
     // DownloadFiles("https://drive.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j", "KFGame.7z")
     var contents []string = readFirstLine("settings.txt")
     DownloadFiles(contents[0], "KFGame.7z")
}