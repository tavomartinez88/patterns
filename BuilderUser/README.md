# BuilderUser

Se implementa un generador de Usuarios que representan personas del mundo real.

Un usuario esta compuesto por :
* FirstName
* LastName
* Email
* Password
* Phone

El modulo implementa el patron de diseño _Builder_ donde se interactua de manera:

                    
      GeneratorUser.BuildUser(constructor builders.BuilderUser, data map[string]string)
                            |
                            v

        BuilderUser.Build(fields map[string]string)

                            |
                            v
                          User  
                          

#### Build Module
go build ./...

#### Run Tests
go test ./...

#### Run Coverage
go test -cover ./...