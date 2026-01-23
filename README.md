<div align="center">

# AbacatePay Go SDK 🥑

SDK oficial da *[AbacatePay](https://abacatepay.com/)* para integração com a API em **Golang**, de forma simples, segura e totalmente tipada.

Desenvolvido com foco em **simplicidade**, **segurança**, **controle de contexto** e **boas práticas idiomáticas de Go**.

Wrapper oficial, tipado e mantido pela equipe *Open Source da AbacatePay*.

<img src="https://res.cloudinary.com/dkok1obj5/image/upload/v1767631413/avo_clhmaf.png" width="100%" alt="AbacatePay Open Source"/>

Você pode ver documentação completa do SDK [aqui](https://docs.abacatepay.com/pages/sdks/golang).

## Instalação

Instale facilmente com o *go get* nativo do Golang

</div>

```bash
go get github.com/AbacatePay/abacatepay-go-sdk/v1
```

<div align="center">

## Uso básico
</div>

```go
package main

import (
	"context"
	"time"
	"log"

	"github.com/AbacatePay/abacatepay-go-sdk/v1"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/pix"
)

func main() {
	client, err := abacatepay.New(abacatepay.ClientConfig{
		APIKey:  "abc_dev",
		Timeout: 10 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)

	defer cancel()

	body := &coupons.CreateQrCodePIXBody{
		Amount: 100,
	}

	resp, err := client.Pix.Create(ctx, body)
	
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("PIX criado: %s\n", resp.Data.ID)
}
```

<div align="center">

## Contexto & Timeouts

Todas as requisições exigem context.Context.

Isso garante:

</div>

- Cancelamento seguro.
- Timeouts explícitos
- Melhor controle em produção

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

defer cancel()
```

<div align="center">

## Estrutura do SDK
</div>

- Cada domínio da API fica em seu próprio pacote (billing, customers, etc).
- Cliente HTTP compartilhado.
- Sem estado global.
- Compatível com aplicações concorrentes.

```go
client.Billing.Get(...)
client.Billing.Create(...)
```

<div align="center">

## Tratamento de erros
</div>

- Erros de configuração retornam imediatamente.
- Erros HTTP são propagados de forma explícita.
- Nenhum panic inesperado.

```go
resp, err := client.Billing.Create(ctx, body)

if err != nil {
	// Trate erro da API ou de rede aqui
	log.Fatal(err)
}
```

<div align="center">

Feito com 🥑 pela equipe AbacatePay<br/>
Open source, de verdade.

</div>
