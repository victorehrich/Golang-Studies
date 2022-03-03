package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	var comand int = -1

	fmt.Println("Bem vindo ao sistema, escolha a opção que deseja executar:")
	comand = menu()

	switch comand {
		case 1:
			fmt.Println("Monitorando. . .")
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

func menu() int{
	var loop bool = true;
	var comand int = -1

	for loop {
		fmt.Println("1-Monitoramento")
		fmt.Println("2-Logs")
		fmt.Println("0-Sair")
		fmt.Scan(&comand)

		if (comand == 1 || comand == 2 || comand == 0){
			fmt.Println("O comando escolhido foi:",comand)
			fmt.Println()
			loop = false
		} else{
			consoleClear()
			fmt.Println("Opção Inválida")
		}
	}
	return comand
}

func consoleClear(){
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}