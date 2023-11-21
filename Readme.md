### Desafio Luiza Labs 

:writing_hand:Para resolver o desafio decidi aplicar arquitetura limpa, arquitetura hexagonal, que na minha opinião é mais uma arquitetura de portas e adaptadores, e também apliquei fatiamento vertical.

Entender que no final das contas a arquitetura hexagonal tem muitas formas de ser aplicada e depende muito da visão do arquiteto. Minha aplicação é a seguinte:


├── core
│   ├── module (una entidad del negocio)
│   │   ├── application
│   │   ├── domain
│   │   ├── infrastructure



- [Como você executa o projeto](#como-você-executa-o-projeto)
- [Considerações](#considerações)
- [Caracteristicas](#caracteristicas)

## Como você executa o projeto

1. Execute o contêiner postgres primeiro, na raiz do projeto
2. Execute ou construa o aplicativo escrito em Golang

```bash
docker-compose up -d

go run cmd/main.go
```


## Considerações

- Criei diferentes módulos, como database, framework, etc. Minha justificativa é porque acho que módulos são algo que pode mudar e não é algo específico do negócio.

- É claro que outro mecanismo de pasta pode ser aplicado, por exemplo, você poderia criar uma implementação concreta que reúna o que está atualmente em main.go para separar ainda mais a aplicação

- Tenho utilizado inversão de dependência e injeção de dependência para manter uma abstração, um desacoplamento, ou seja, a parte mais acoplada seria a infraestrutura e a menos, o domínio.

Aqui está um exemplo "simulado" de como algo semelhante seria implementado em Kotlin:

```kotlin
interface UserRepository {
    fun getById(userId: String): User?
    fun getByDateRange(startDate: LocalDate, endDate: LocalDate): List<User>
    fun insertFile(data: List<User>)
}

class MockUserRepository(private val connectionDb: IConnectionDb) : UserRepository {
    private val users: MutableList<User> = mutableListOf()

    override fun getById(userId: String): User? {
        return users.find { it.userId == userId }
    }

    override fun getByDateRange(startDate: LocalDate, endDate: LocalDate): List<User> {
        return users.filter { user ->
            user.orders.any { order ->
                order.products.any { product ->
                    product.buyDate in startDate..endDate
                }
            }
        }
    }

    override fun insertFile(data: List<User>) {
        users.addAll(data)
    }
}

fun main() {
    val connectionDb: IConnectionDb = MockConnectionDb()
    val mockRepo: UserRepository = MockUserRepository(connectionDb)
}
```

## Caracteristicas

- Swagger -> http://localhost:8080/api/v1/docs
- Fiber (framework)
- ORM (GORM)
- Database (Postgres)
- Docker
- Teste teste de carga simulado com Artillery
- Mock github action