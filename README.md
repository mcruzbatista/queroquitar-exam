# QueroQuitar Exam

## Proposal
 - Pegar um gif random do Giphy e mostrar numa página com um "greeter"
 
### Tech

    - Echo
    - Ginkgo/Gomega
    - Resty
    - Gjson
    

### How ?

 - Giphy tem uma api então nada mais justo do que consumir.
    
    - Service: Consumir api da Giphy retornando a url ou error
    - Delivery: Com uma rota "/giphy" de chamada "GET" e uma func no package
    Html para chamar o service e retorna o template pronto com gif
    -Template: template para o html que vai ser exposto no render
    
### Test
 - Usei ginkgo pois estou querendo usar mais o BDD pois acredito que é melhor para
 uma nova pessoa entrar no projeto em andamento e para manutenção
     