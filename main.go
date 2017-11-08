package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"os/exec"
	"gopkg.in/yaml.v2"
	"github.com/PodMonitor-server/types"
	"bytes"
	"strings"
)

func main() {
	fmt.Println("************************* Server Started *****************************")
	router := mux.NewRouter()
	router.HandleFunc("/welcome",Welcome).Methods("GET")
	router.HandleFunc("/runcmd",RunCmd).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Welcome(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Logging In ...")
	json.NewEncoder(w).Encode("Welcome lets start work")

}
func RunCmd(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	clientYaml, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Logging In ...",string(clientYaml))

	var data = types.CMD
	 yaml.Unmarshal(clientYaml, &data)
	 fmt.Println(data.Cmd)
length:=len(strings.Split(data.Cmd,":"))
var output *exec.Cmd;
	switch length {
	case 0 :
		json.NewEncoder(w).Encode("No command received !!!")
		break;
	case 1 :
		json.NewEncoder(w).Encode("Invalid argument !!!")
		break;
	case 2 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1])
		break;
	case 3 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2])
		break;
	case 4 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2],strings.Split(data.Cmd," ")[3])
		break;
	case 5 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2],strings.Split(data.Cmd," ")[3],strings.Split(data.Cmd," ")[4])
		break;
	case 6 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2],strings.Split(data.Cmd," ")[3],strings.Split(data.Cmd," ")[4],strings.Split(data.Cmd," ")[5])
		break;
	case 7 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2],strings.Split(data.Cmd," ")[3],strings.Split(data.Cmd," ")[4],strings.Split(data.Cmd," ")[5],strings.Split(data.Cmd," ")[6])
		break;
	case 8 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2],strings.Split(data.Cmd," ")[3],strings.Split(data.Cmd," ")[4],strings.Split(data.Cmd," ")[5],strings.Split(data.Cmd," ")[6],strings.Split(data.Cmd," ")[7])
		break;
	case 9 :
		output=exec.Command(strings.Split(data.Cmd," ")[0],strings.Split(data.Cmd," ")[1],strings.Split(data.Cmd," ")[2],strings.Split(data.Cmd," ")[3],strings.Split(data.Cmd," ")[4],strings.Split(data.Cmd," ")[5],strings.Split(data.Cmd," ")[6],strings.Split(data.Cmd," ")[7],strings.Split(data.Cmd," ")[8])
	break;
		}


	out:=printCmd(output)
	fmt.Println(string(out))
	w.Write(out)
	json.NewEncoder(w).Encode("Welcome lets start work")


}
func printCmd(cmd *exec.Cmd) []byte{

	cmdOutput := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = stdErr
	err := cmd.Run()
	if err != nil {
		//fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		return stdErr.Bytes()
	}
	return cmdOutput.Bytes()
}

