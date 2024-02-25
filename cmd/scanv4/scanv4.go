package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Criando o diretório scanv4
	err := os.Mkdir("scanv4", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Mudando para o diretório scanv4
	err = os.Chdir("scanv4")
	if err != nil {
		log.Fatal(err)
	}

	// Baixando o arquivo zip
	cmd := exec.Command("wget", "https://github.com/adfastltda/scanv4/releases/download/scanv4/scanv4_for_arm64_x86_x86_64.zip")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Descompactando o arquivo zip
	cmd = exec.Command("unzip", "scanv4_for_arm64_x86_x86_64.zip")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Alterando permissões do executável
	cmd = exec.Command("chmod", "+x", "scanv4")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Executando o programa com o argumento -h
	cmd = exec.Command("./scanv4", "-h")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("erro ao executar scanv4: %v\n%s", err, out)
	}

	fmt.Println("Comando scanv4 executado com sucesso.")
}
t
