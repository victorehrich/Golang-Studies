package main

import (
	"fmt"
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
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.google.com/", "https://www.youtube.com/"}
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
		} else {
			fmt.Println("Site:", site, "esta com problemas; status: ", resp.Status)
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
