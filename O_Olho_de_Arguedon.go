package main

import "fmt"

func main() {
	//Inicio
	Fase := 0
	Vigor := 10
	Nome := "nome"
	Escolha := 0
	fmt.Println("\n\n  Depois de dias intensos de busca e aventuras, você finalmente se deparou com a dungeon final de sua jornada ! " +
		"\nDungeon essa, que abriga o olho de Arguedon, o artefato místico, que é procurado por todos, para diversos fins, como por exemplo... " +
		"\nA destruição de toda Crioterra. Sua terra natal está em perigo, e você como o maior paladino da região, deve obter esse artefato, para que não caiam em mão erradas. " +
		"\n\nAntes de tudo, qual o seu nome, aventureiro ?")
	fmt.Scan(&Nome)
	//Primeira etapa
	Fase = 1
	for Fase == 1 {
		fmt.Printf("\n----------------------------------------------------\n\n")
		fmt.Println("Pois bem ", Nome, ", Ao entrar na primeira etapa do calabouso final, você se depara com três portas de madeiras velhas:"+
			"\n 1 - Esquerda\n 2 - Meio\n 3 - Direita\n Qual você decide entrar ?\nPs: para checar seu nome e vigor atual, digite: 4")
		fmt.Scan(&Escolha)
		if Escolha == 1 {
			fmt.Printf("Ao entrar na primeira porta você repentinamente é atacado por um rápido vulto. Vocẽ consegue desviar de seu golpe por um triz" +
				"\nO inimigo, é um zumbi, empunhando uma espada, você avança e consegue decapitá-lo facilmente com sua espada, mas você não contava com a orda" +
				"\nQue o cercou após essa distração, você os enfrenta bravamente, defendendo-se dos golpes com seu escudo e os atacando com sua espada" +
				"\nDepois de alguns árduos minutos, você os derrota, porém, sai de lá ferido, e que agora você perdeu 3 pontos de vigor")
			Vigor = Vigor - 3
			Fase = 2
		} else if Escolha == 2 {
			fmt.Println("Ao entrar na porta do meio, você simplesmente encontra um corredor vazio, mas muito, MUITO fedorento(tipo MUITO fedorento mesmo) ")
			Fase = 2
		} else if Escolha == 3 {
			fmt.Println("Ao entrar na última porta, você percebe que deu sorte, e se depara com um pequeno baú, e ao abri-lo, encontra um elmo de ouro" +
				"\nque com certeza vai lhe proteger ainda mais, você ganhou +2 pontos de vigor")
			Vigor = Vigor + 2
			Fase = 2
		} else if Escolha == 4 {
			fmt.Println("seu nome é:", Nome, ", o mais famoso paladino de toda Crioterra.\nvigor =", Vigor)
			fmt.Println("Digite qualquer coisa para voltar as opções")
			fmt.Scan(&Escolha)
		} else {
			fmt.Println("Opaaaaa, essa opção aí não existe não.")
			fmt.Println("Digite qualquer coisa para voltar as opções")
			fmt.Scan(&Escolha)
		}
	}
	//Segunda etapa
	for Fase == 2 {
		fmt.Printf("\n\n----------------------------------------------------\n\n")
		fmt.Println("Ao passar pelo corredor, você se depara com mais três portas, qual a sua escolha ?" +
			"\n Ps: o sistema de escolha é igual a antes, e vai continuar sendo assim na próxima etapa(isso se houver uma... >:D )")
		fmt.Scan(&Escolha)

		if Escolha == 1 {
			fmt.Printf("Ao entrar na primeira porta, você se depara com um lindo peitoral dourada pendurado na parede, e você tem certeza" +
				"\nque isso será útil, então você o usa,e ganha +3 de vigor")
			Vigor = Vigor + 3
			Fase = 3
		} else if Escolha == 2 {
			fmt.Println("Você põe a mão na maçaneta para abrir a porta, mas logo a mesma porta é aberta com muita força encima de vocẽ. Lhe empurrando para trás," +
				"\nvocê então vê na sua frente um enorme troll, segurando um tacape, e pronto para lhe esmagar, a luta então é ardua, você consegue desviar dos golpes" +
				"\ndele por um triz, mas segurar aqueles impactos com seu escudo machuca muito, mas isso não irá lhe impedir, então ao desviar de um impacto, voce consegue" +
				"\n acertar um golpe fatal na criatura, que cai derrotada no chão, mas essa batalha foi árdua, você continua, mas ferido. você perdeu 4 de vigor. ")
			Vigor = Vigor - 4
			Fase = 3
		} else if Escolha == 3 {
			fmt.Println("Você abre a porta, e encontra um corredor escuro e aparentemente vazio. Você não encontra nada que lhe ajude, mas também nada que atrapalhe.")
			Fase = 3
		} else if Escolha == 4 {
			fmt.Println("seu nome é:", Nome, ",o mais famoso paladino de toda Crioterra.\nvigor =", Vigor)
			fmt.Println("Digite qualquer coisa para voltar as opções")
			fmt.Scan(&Escolha)
		} else {
			fmt.Println("Opaaaaa, essa opção aí não existe não.")
			fmt.Println("Digite qualquer coisa para voltar as opções")
			fmt.Scan(&Escolha)
		}
	}
	//Terceira Etapa
	for Fase == 3 {
		fmt.Printf("\n\n----------------------------------------------------\n\n")
		fmt.Println("Ao passar pelo mais novo corredor, você se depara com mais três portas, qual a sua escolha ?")
		fmt.Scan(&Escolha)

		if Escolha == 1 {
			fmt.Printf("Ao entrar na porta da esquerda, você encontra um escudo muito belo e bem esculpido, um pouco maior que o seu, você realiza a troca" +
				"\ne ganha +1 de vigor")
			Vigor = Vigor + 1
			Fase = 4
		} else if Escolha == 2 {
			fmt.Println("Você abre a porta do meio, e percebe simplesmente um corredor escuro, que leva a uma sala iluminada.")
			Fase = 4
		} else if Escolha == 3 {
			fmt.Println("Você abre a porta, e vê um enorme e irritado minotauro, segurando seu grande machado, ele ja começa atacando com a ferocidade de um animal enjaulado." +
				"\nvocê desvia dos primeiros golpes, e os acerta alguns com espada, mas a força e resistência dele, são bem superiores, então você defende um dos golpes" +
				"\ncom seu escudo, que voa para o outro lado da sala, é ai então que o minotauro desfere um golpe com tudo em seu torso, que perfura seu peitoral...")
			Vigor = Vigor - 5
			//Desfecho da luta
			if Vigor > 0 {
				fmt.Println("Você sente seu sangue escorrento pelo seu peito, e seu corpo enfraquecer, mas você não pode deixar a Crioterra em risco, você precisa"+
					"\nse levantar e lutar. Então", Nome, "reune toda sua força e determinação, e retira o machado do seu peito, o empunhando com toda sua força de vontade."+
					"\nvendo a cena, o minotauro avança com tudo, pronto para lhe empalar com seus chifres, mas você não irá permitir, então", Nome, "Desfere um último golpe"+
					"\ndecapitando o minotauro com seu próprio machado..."+
					"\ndepois de alguns segundos de descanso, você segue a sala final, com sua espada, escudo, e amor pela sua terra.")
				Fase = 4
			} else {
				fmt.Println(Nome, "sente seu sangue escorrendo, e toda aquela dor do machado em seu corpo, ele tenta buscar forças para levantar, mas infelizmente "+
					"\nnão consegue encontra-la, sua visão começa a escurecer, enquanto o minotauro avança para realizar o golpe final. Crioterra está condenada...")
				fmt.Println("GAME OVER")
				break
			}
		} else if Escolha == 4 {
			fmt.Println("seu nome é:", Nome, ",o mais famoso paladino de toda Crioterra.\nvigor =", Vigor)
			fmt.Println("Digite qualquer coisa para voltar as opções")
			fmt.Scan(&Escolha)
		} else {
			fmt.Println("Opaaaaa, essa opção aí não existe não.")
			fmt.Println("Digite qualquer coisa para voltar as opções")
			fmt.Scan(&Escolha)
		}
	}
	//Etapa Final
	for Fase == 4 {
		fmt.Printf("\n\n----------------------------------------------------\n\n")
		fmt.Println("Ao seguir no corredor, você entra em uma sala enorme, e muito bem iluminada, com diversas tochas.",Nome,"vê então o olho de Arguedon,"+
			"\num lindo e brilhante cristal, exposto em um pedestal. É quando do teto, desce um enorme dragão, de escamas azuis brilhantes. O guardião do cristal"+
			"\n o encara com raiva, partindo diretamente para o ataque.")
		fmt.Printf("\n\n----------------------------------------------------\n\n")

		if Vigor < 10 {
			fmt.Printf("O dragão avança em sua direção com toda ferocidade que uma criatura mística poderia ter.",Nome,"consegue desviar do primeiro avanço," +
				"\ne tenta atacar o dragão com sua espada, mas o mesmo é muito rápido, e o grande paladino não consegue acerta-lo. O dragão então começa a cuspir" +
				"\nsua rajada de fogo, e",Nome,"faz de tudo para defender, mas ele não possui mais forças suficiente para isso, então o fogo do dragão, passa pela" +
				"\nsua guarda, e o desintegra. Condenando assim a crioterra a quem capturar o olho de Arguedon, e se forem mãos erradas, é o fim...")
			fmt.Println("GAME OVER")
			break
		} else if Vigor >= 10 && Vigor < 16 {
			fmt.Println("O dragão avança em sua direção com toda ferocidade que uma criatura mística poderia ter.",Nome,"consegue defender com seu escudo" +
				"\nmas cai pra trás devido ao impacto, o dragão então avança em sua direção e tenta fatia-lo, mas o paladino consegue se esquivar com dificuldade" +
				"\na luta é intensa, e acirrada, e tudo dependerá, do quão forte você é. O dragão então recua e começa a disparar uma rajada de fogo." +
				"\n",Nome,"defende com seu escudo, com muito esforço e dificuldade, mas ele já chegou muito longe, e não pode desistir de sua terra agora" +
				"\nentão ele reuni toda sua força e avança na direção da criatura, enquanto defende-se de sua rajada, e ao chegar perto do dragão, que já está ficando sem fôlego, epga sua espada e crava na cabeça da criatura, assim, derrotando-a" +
				"\na luta foi dificil para o grande paladino,",Nome,". Que agora possui o olho de Arguedon, portanto, Crioterra estará segura agora. Ou sera que não... ?")
			break
		} else {
			fmt.Println("O dragão avança em sua direção com toda ferocidade que uma criatura mística poderia ter.",Nome,"consegue desviar do primeiro avanço." +
				"\nAo desviar",Nome,"usa sua espada para rasgar a boca do dragão, o ferindo gravemente e o enfraquecendo, o dragão então cospe fogo no paladino" +
				"\nque defende com seu escudo sem nenhuma dificuldade, já que está no auge de sua força e o dragão fraco pelo ferimento.",Nome,"então, avança" +
				"\nem meio ao fogo e ao chegar perto do dragão, que já está ficando sem fôlego, epga sua espada e crava na cabeça da criatura, assim, derrotando-a" +
				"\na luta foi fácil para o grande paladino,",Nome,". Que agora possui o olho de Arguedon, portanto, Crioterra estará segura agora. Ou sera que não... ?")
			break
		}
	}
}
