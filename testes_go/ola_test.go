package ola

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Hello("Benjamin")
		esperado := WaitHello("Benjamin")
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Hello("")
		esperado := WaitHello("Bruna")
		verificaMensagemCorreta(t, resultado, esperado)
	})

	// t.Run("diz olá para as pessoas", func(t *testing.T) {
	// 	resultName := Hello("Bruna")
	// 	waitName := WaitHello("Bruna")

	// 	if resultName != waitName {
	// 		t.Errorf("resultado '%s', esperado '%s'", resultName, waitName)
	// 	}
	// })

	// t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
	// 	resultName := Hello("")
	// 	waitName := WaitHello("Olá, mundo")

	// 	if resultName != waitName {
	// 		t.Errorf("resultado '%s', esperado '%s'", resultName, waitName)
	// 	}
	// })
}
