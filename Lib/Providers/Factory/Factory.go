package Factory

import (
	"go-microservices/Lib/Core/Source"
	"go-microservices/Lib/Providers/Db/MySql"
	"go-microservices/Lib/Providers/Interface"
)

func CreateProvider() Interface.ProviderInterface {

	switch Source.SOURCE_CURRENT_PROVIDER {
	case "my-sql":
		return createMySqlProvider()

	}

	//Default data provider is MySql
	return createMySqlProvider()

}

func createMySqlProvider() MySql.MySqlProvider {
	//Initialize provider
	provider := MySql.MySqlProvider{}

	//Set provider connection settings
	provider.SetDatasource(Source.SOURCE_USER)
	provider.SetHost(Source.SOURCE_HOST)
	provider.SetPassword(Source.SOURCE_PASSWORD)
	provider.SetUser(Source.SOURCE_USER)

	return provider
}