package dbhandler


import(
  "testing"
  "os"
  "database/sql"
  "fmt"
  "strings"
)

var (
  MockDBURLProto string
  MockDB string 
  MockDBType = "mysql"
)

func init() {
  MockDBURLProto = os.Getenv("MOCKDBURLPROTO")
  MockDB = os.Getenv("MOCKDB")
  MockCreateDatabase()
}

func MockCreateDatabase() {
  noDBURL := strings.Replace(MockDBURLProto, MockDB, "", 1)
  dbi, err := newMockDBI(MockDBType, noDBURL)
  if err != nil {
    fmt.Printf("DB create failed: %s\n", err)
    os.Exit(1)
  }
  defer dbi.Close()
  
  // check if db exists
  dbname, err := dbi.GetStringValue(fmt.Sprintf("SHOW DATABASES LIKE '%s'", MockDB))
  if err != nil && err != sql.ErrNoRows {
    fmt.Printf("query for exiting db failed: %s\n", err)
    os.Exit(1)
  } 
  
  if dbname == MockDB {
    return // db exists so return
  }
  
  _, err = dbi.SQLSession.Exec(fmt.Sprintf("CREATE DATABASE %s", MockDB))
  if err != nil {
    fmt.Printf("Creating database failed: %s\n", err)
    os.Exit(1)
  }
}

func TestParseEnv(t *testing.T) {
  dbi, err := NewDBI(MockDBType)
  if err != nil {
    t.Fatal(err)
  }
  defer dbi.Close()
  
  _, err = dbi.SQLSession.Exec("SHOW DATABASES")
  if err != nil {
    t.Fatal(err)
  }
  
}