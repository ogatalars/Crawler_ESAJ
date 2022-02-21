# Crawler_ESAJ

Projeto que visa capturar a quantidade de decisões com origem de ```2° grau``` e ```Colégios Recursais``` publicadas como ```Acórdãos```, ```Homologações de Acordo``` e ```Decisões Monocráticas```.

O prjoteto precisa receber as datas a serem pesquisadas como  **data inicial** e  **data final**, com isso serão pesquisadas as datas mês a mês do intervalo fornecido em todos os sistemas ESAJ: ```tjac,tjal,tjam,tjce, tjms e tjsp```.

Ao final são gerados um arquivo CSV para cada TJ.
 
## Dependências

Para o esse crawler somos dependentes do projeto [selenium](https://github.com/tebeka/selenium#readme), sendo necessário a pré-instalação do [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/) e do [selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)

## Multithread

O projeto já possuí o conceito de multithread implementado, porém não é recomendado utilizar mais de 1 thread (leia-se um navegador chrome executado por vez), pois o sistema ESAJ detecta e retorna o erro 403 com alta frequência. 


## Run
Basta dar o start no selenium server e iniciar o app em GO: 

```java -jar selenium-server-standalone.jar```

```go run main.go```

## Aviso legal (Disclaimer)

O projeto não mantém nem atualiza os dados do ESAJ, não altera nem transforma os dados durante a coleta. O autor do projeto não assume qualquer responsabilidade sobre o seu uso e resultados.

## Considerações éticas

A) Idealmente o sistema ESAJ deveria disponibilizar uma API ou no mínimo um web service para facilitar o acesso, como isso não existe projetos como esse são necessários. O TJSP não proíbe expressamente o uso de crawlers/scraapers, você pode conferir acessando o robots.txt dos sistemas de cada TJ. Isso porém não quer dizer que seu IP não será bloqueado caso você execute muitas requisições.

B) Esse projeto foi criado de forma voluntária, você pode contribuir de qualquer modo. Se encontrar uma falha, não hesite em criar um “issue” ou  procure corrigir você mesma(o) o erro e dar um “pull request”.


C) use os dados baixados para agregar valor, como por exemplo, para realizar análises ou publicar artigos, fazer inferências, elaborar recomendações aos poderes públicos etc. Baixar esses dados para reproduzi-los em sua página web é tirar proveito do trabalho alheio, mesmo sendo esses dados públicos.
