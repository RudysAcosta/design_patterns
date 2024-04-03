package factory

import "fmt"

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	/*
		Interfaz Database: Define un conjunto de métodos comunes (GetData y PutData) que
		deben ser implementados por todas las bases de datos. Esto permite que el código
		que utiliza las bases de datos las trate de manera uniforme sin preocuparse por
		su tipo subyacente.
	*/
	Database interface {
		GetData(string) string
		PutData(string, string)
	}
)

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDb")
	return mdb.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDb")
	return sql.database[query]
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

/*
Diseño del patrón Factory: El método DatabaseFactory actúa como una fábrica
que crea instancias de bases de datos basadas en el entorno proporcionado.
Esto permite la creación de objetos de base de datos de manera flexible y
desacoplada de la lógica de creación.
*/
func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]string),
		}
	case "development":
		return sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}
