package cli

import (
	"fmt"
	"gb_golang/internal/player"
)

func (c Cli) GetPlayers() [2]player.Player {
	thereIsAPlayers := false
	var players [2]player.Player
	var err error

	fmt.Println("Нам необходимо два пользователя. Один играет за X другой за 0.")

	for !thereIsAPlayers {
		for x := 0; x < 2; x++ {
			fmt.Println("Создадим пользователя №", x+1)
			players[x], err = player.New(CliGetPlayerName(), CliGetSide())
			if err != nil {
				fmt.Println("Ууупс! что-то пошло не так.")
				break
			}
		}

		if players[0].GetSide() == players[1].GetSide() {
			fmt.Println("Два пользователя выбрали одну и туже сторону - это не правильно.")
			fmt.Println("Попробуйте еще раз")
		} else {
			thereIsAPlayers = true
		}
	}

	fmt.Printf("Подготовка завершена.\n\n")
	return players
}
