# Multithreading

 Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

- https://brasilapi.com.br/api/cep/v1/ + cep

- http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## arquivos:

- `brasilapi.200.json` - resposta status 200 da brasilapi

```json
{
  "cep": "39408078",
  "state": "MG",
  "city": "Montes Claros",
  "neighborhood": "Ibituruna",
  "street": "Avenida Herlindo Silveira",
  "service": "open-cep"
}
```

- `brasilapi.400.json` - resposta status 400 da brasilapi - cep com quantidade de dígitos diferente de 8

```json
{
    "message": "CEP deve conter exatamente 8 caracteres.",
    "type": "validation_error",
    "name": "CepPromiseError",
    "errors": [
        {
            "message": "CEP informado possui mais do que 8 caracteres.",
            "service": "cep_validation"
        }
    ]
}
```

- `brasilapi.404.json` - resposta status 404 da brasilapi - cep não encontrado

```json
{
    "message": "Todos os serviços de CEP retornaram erro.",
    "type": "service_error",
    "name": "CepPromiseError",
    "errors": [
        {
            "name": "ServiceError",
            "message": "A autenticacao de null falhou!",
            "service": "correios"
        },
        {
            "name": "ServiceError",
            "message": "Erro ao se conectar com o serviço ViaCEP.",
            "service": "viacep"
        },
        {
            "name": "ServiceError",
            "message": "Erro ao se conectar com o serviço WideNet.",
            "service": "widenet"
        },
        {
            "name": "ServiceError",
            "message": "CEP não encontrado na base dos Correios.",
            "service": "correios-alt"
        }
    ]
}
```

- `viacep.200.json` - resposta status 200 da viacep

```json
{
  "cep": "39408-078",
  "logradouro": "Avenida Herlindo Silveira",
  "complemento": "até 499/500",
  "unidade": "",
  "bairro": "Ibituruna",
  "localidade": "Montes Claros",
  "uf": "MG",
  "estado": "Minas Gerais",
  "regiao": "Sudeste",
  "ibge": "3143302",
  "gia": "",
  "ddd": "38",
  "siafi": "4865"
}
```

- `viacep.200.erro.json` - resposta da viacep para cep não encontrado

```json
{
    "erro": "true"
}
```

- `viacep.400.html` - resposta status 400 da viacep - cep com quantidade de dígitos diferente de 8

```html
<!DOCTYPE HTML>
<html lang="pt-br">

<head>
  <title>ViaCEP 400</title>
  <meta charset="utf-8" />
  <style type="text/css">
      h1 {
          color: #555;
          text-align: center;
          font-size: 4em;
      }
      h2, h3 {
          color: #666;
          text-align: center;
          font-size: 3em;
      }
      h3 {
          font-size: 1.5em;
      }
  </style>
</head>

<body>
    <h1>Http 400</h1>
    <h3>Verifique a URL</h3>
    <h3>{Bad Request}</h3>
</body>

</html>
```

- `saida.txt` - cópia da saída do terminal para várias execuções da rotina

```shell
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go
2024/10/09 09:37:59 Usage: go run main.go <cep>
exit status 1
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go 3940807</>
2024/10/09 09:38:12 Return from ViaCep
2024/10/09 09:38:12 cep deve conter exatamente 8 caracteres
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go 394080788
2024/10/09 09:38:17 Return from Brasilapi
2024/10/09 09:38:17 cep deve conter exatamente 8 caracteres
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go 39408079
2024/10/09 09:38:25 Return from Brasilapi
2024/10/09 09:38:25 cep não encontrado
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go 39408078
2024/10/09 09:38:30 Return from Brasilapi
2024/10/09 09:38:30 {"cep":"39408078","state":"MG","city":"Montes
Claros","neighborhood":"Ibituruna","street":"Avenida Herlindo Silveira","service":"open-cep"}
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go 39408078
2024/10/09 09:52:53 Return from ViaCep
2024/10/09 09:52:53 {
 "cep": "39408-078",
 "logradouro": "Avenida Herlindo Silveira",
 "complemento": "até 499/500",
 "unidade": "",
 "bairro": "Ibituruna",
 "localidade": "Montes Claros",
 "uf": "MG",
 "estado": "Minas Gerais",
 "regiao": "Sudeste",
 "ibge": "3143302",
 "gia": "",
 "ddd": "38",
 "siafi": "4865"
}
antonio@DG15:~/DEV/full-cycle/multithreading$ go run main.go 39408078
^C2024/10/09 09:47:26 execução cancelada
```