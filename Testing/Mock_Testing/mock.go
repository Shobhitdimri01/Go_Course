package mocktesting
// Define an interface for a dependency
type Database interface {
    GetData(id string) string
}

// Function using the interface
func FetchUserData(db Database, id string) string {
    return db.GetData(id)
}


