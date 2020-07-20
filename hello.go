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

func downloadFiles() {
	newFile, err := os.Create("KFGame.7z")
     if err != nil {
          log.Fatal(err)
     }
     defer newFile.Close()

     // HTTP GET request 
     url := "https://drive.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j"
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
}

func createPath(path string) string {
     // 
     currentDirectory, err := os.Getwd()
     if err != nil { panic(err) }
     filePaths := strings.Split(path, "/")
     lengthOfFilePath := len(filePaths)
     filePath := strings.Join(filePaths[:lengthOfFilePath-1], "/")
     completePath := currentDirectory + "/" + filePath

     return completePath
}

func extractFiles() {
     
     a, err := unarr.NewArchive("KFGame.7z")
     if err != nil { panic(err) }
     defer a.Close()

     data, err := a.List()
     if err != nil { panic(err) }
     for _, file := range data {
          
          fmt.Println(file)
          err := a.EntryFor(file)
          if err != nil {
               panic(err)
          }

          // returns path as string that needs to be passed to MkdirAll
          completePath := createPath(file)

          _ = os.MkdirAll(completePath, os.ModePerm)

          currentEntry := make([]byte, 1000000)
          currentEntry, err = a.ReadAll()
          if err != nil {
               panic(err)
          }
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


     //currentDirectory, err := os.Getwd()
     // result, err := zFile.Extract("/tmp")
     // if err != nil {
     //      panic(err)
     // }
     // for _, file := range result {
     //      log.Printf(file)
     // }
}



func main() {
     // Create output file
     downloadFiles()
     extractFiles()
     
}