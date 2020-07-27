
package main

import (
     "os"
     "io"
    // "log"
     "fmt"
     //"bufio"
     "runtime"
     "strings"
     "net/http"
     "archive/zip"
     "path/filepath"
     "github.com/sqweek/dialog"
     //"github.com/TheTitanrain/w32"
)

func DownloadFiles(url string, filename string) {
	newFile, err := os.Create(filename)
     if err != nil { panic(err) }
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
     if err != nil { panic(err) }
     fmt.Println("Downloaded bytes: ", numBytesWritten)
     // Extracting into current working directory
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

// func readFirstLine(filepath string) ([]string){
     
//      file, err := os.Open(filepath)
//      if err != nil { panic(err) }
//      defer file.Close()
     
//      var filecontents []string

//      scanner := bufio.NewScanner(file)
//      for scanner.Scan() {
//           filecontents = append(filecontents, scanner.Text())
//      }
//      fmt.Println(filecontents)
//      return filecontents
// }

func Unzip(src, dest string) error {
     r, err := zip.OpenReader(src)
     if err != nil {
         return err
     }
     defer func() {
         if err := r.Close(); err != nil {
             panic(err)
         }
     }()
 
     os.MkdirAll(dest, 0755)
 
     // Closure to address file descriptors issue with all the deferred .Close() methods
     extractAndWriteFile := func(f *zip.File) error {
         rc, err := f.Open()
         if err != nil {
             return err
         }
         defer func() {
             if err := rc.Close(); err != nil {
                 panic(err)
             }
         }()
 
         path := filepath.Join(dest, f.Name)
 
         if f.FileInfo().IsDir() {
             os.MkdirAll(path, f.Mode())
         } else {
             os.MkdirAll(filepath.Dir(path), f.Mode())
             f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
             if err != nil {
                 return err
             }
             defer func() {
                 if err := f.Close(); err != nil {
                     panic(err)
                 }
             }()
 
             _, err = io.Copy(f, rc)
             if err != nil {
                 return err
             }
         }
         return nil
     }
 
     for _, f := range r.File {
         err := extractAndWriteFile(f)
         if err != nil {
             return err
         }
     }
 
     return nil
 }



func WindowsCleanUp(filename string) {
    if runtime.GOOS == "windows" {


        currentDirectory, err := os.Getwd()
        if err != nil { panic(err) }

        Unzip(filename, currentDirectory)

        pathToZip := currentDirectory+"\\"+filename
        os.Remove(pathToZip)
        fmt.Println(pathToZip)

        
        //cmd := exec.Command(`pause`)
        //cmd.Run()
     }
 }

 func main() {
     // Downloads files from a link
     // DownloadFiles("https://drive.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j", "KFGame.7z")
     // var contents []string = readFirstLine("settings.txt")
    switch answer := dialog.Message("%s", "Do want to install current server mods?").Title("KF2 Mod Installer?").YesNo(); answer {
    case true:
        filename := "KFGame.zip" 
        DownloadFiles("https://drive.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j", filename)
        WindowsCleanUp(filename)
    case false:
        break
    }   
}