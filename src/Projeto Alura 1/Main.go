package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {
	var comand int = -1

	fmt.Println("Bem vindo ao sistema, escolha a opção que deseja executar:")
	for {
		comand = menu()

		switch comand {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Lista de logs:")
			imprimeLogs()
		case 0:
			fmt.Println("Finalizando programa. . .")
			os.Exit(0)
		default:
			fmt.Println("Ops, algo deu errado.")
			os.Exit(-1)
		}
	}

}

func monitoramento() {
	fmt.Println("Monitorando. . .")
	sites := lerArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Site", i+1, "-", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}
func testaSite(site string) {
	resp, err := http.Get((site))
	if err != nil {
		fmt.Println("Ops, ocorreu um erro: ", err)
	} else {
		holderStatusResp := strings.Trim(resp.Status, "200 ")
		if strings.Split(strconv.Itoa(resp.StatusCode), "")[0] == "2" {
			fmt.Println("Site:", site, "foi carregado com o status:", holderStatusResp)
			registraLog(site, true, holderStatusResp)
		} else {
			fmt.Println("Site:", site, "esta com problemas; status: ", resp.Status)
			registraLog(site, false, resp.Status)
		}
	}
}

func menu() int {
	var loop bool = true
	var comand int = -1

	for loop {
		fmt.Println("1-Monitoramento")
		fmt.Println("2-Logs")
		fmt.Println("0-Sair")
		fmt.Scan(&comand)

		if comand == 1 || comand == 2 || comand == 0 {
			fmt.Println("O comando escolhido foi:", comand)
			fmt.Println()
			loop = false
		} else {
			consoleClear()
			fmt.Println("Opção Inválida")
		}
	}
	return comand
}

func consoleClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func lerArquivo() []string {
	var sites []string
	//arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ops, ocorreu um erro ao acessar o arquivo", err)
		os.Exit(-1)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		site, err := leitor.ReadString('\n')
		site = strings.TrimSpace(site)
		sites = append(sites, site)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ops, ocorreu um erro ao ler o arquivo", err)
			os.Exit(-1)
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool, statuscode string) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ops, ocorreu um erro ao ler o arquivo", err)
		os.Exit(-1)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + " statuscode " + statuscode + "\n")
	arquivo.Close()
}
func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ops, ocorreu um erro ao ler o arquivo", err)
		os.Exit(-1)
	}
	fmt.Println(string(arquivo))
}
