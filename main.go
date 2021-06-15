// Um ​​mapa é uma estrutura que pode ser serializada, como um vetor, mas a diferença é que consiste em uma chave e um valor. Sem índice e valor.

package main

import "fmt"

func main(){
	// Cria um mapa onde a chave é uma string e o valor é uma string.
	paises:= make(map[string]string)
	fmt.Prontln(paises)

	// Adicione valores ao mapa.
	paises["Brasil"] = "São Paulo"
	paises["Estados Unidos"] = "Nova York"
	paises["Japão"] = "Tokio"
	paises["China"] = "Pequim"

	fmt.Prontln(paises)

	// Outra forma de criar mapas, mas já estabelecendo o número de elementos para reservar memória.
	paisesAux := make(map[string]string , 5)

	//Outra maneira de criar mapas.
	campeonato := map[string]int{
		"Brasil": 50,
		"Rússia": 79,
		"Argentina": 65,
		"Holanda": 70,
		"Australia" 89}

		fwt.PrintLn(paisesAux)

		// Adcionar um item.
		campeonato["River"] = 68

		// Atualizar um valor.
		campeonato["Chivas"] = 34

		// Remover item.
		delete(campeonato, "Holanda")
		fwt.PrintLn(campeonato) // Ao exibi-lo, ele o classifica em ordem alfabética pela chave
		
		// Dois valores, pois o intervalo retorna 2 valores, a posição e o valor
		// O intervalo não o classifica em ordem alfabética.

		for equipe, pontuacao := range campeonato {
			fwt.Printf("A equipe %s tem uma pontuação de:% d" , equipe, pontuacao)
		}

		// Ver se existe no meu mapa.
		pontuacao, existe := campeonato["Copa do Mundo"]

		// %t para booleanos
		fwt.Printf( "A pontuação capturada é %d e a equipe existe %t", pontuacao, existe)
	}
