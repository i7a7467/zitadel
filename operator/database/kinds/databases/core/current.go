package core

/*
import (
	"errors"

	"github.com/caos/zitadel/pkg/databases/db"

	"github.com/caos/orbos/pkg/tree"
)

const queriedName = "database"

var ErrNoCurrentState = errors.New("no current state for database found")

func ParseQueriedForDatabase(queried map[string]interface{}) (db.Client, error) {
	queriedDB, ok := queried[queriedName]
	if !ok {
		return nil, ErrNoCurrentState
	}
	currentDBTree, ok := queriedDB.(*tree.Tree)
	if !ok {
		return nil, errors.New("current state does not fullfil interface")
	}
	currentDB, ok := currentDBTree.Parsed.(db.Client)
	if !ok {
		return nil, errors.New("current state does not fullfil interface")
	}

	return currentDB, nil
}

func SetQueriedForDatabase(queried map[string]interface{}, databaseCurrent *tree.Tree) {
	queried[queriedName] = databaseCurrent
}


func SetQueriedForDatabaseDBList(queried map[string]interface{}, databases, users []string) {
	currentDBList := &CurrentDBList{
		Common: &tree.Common{
			Kind: "DBList",
		},
		Current: &DatabaseCurrentDBList{
			Databases: databases,
			Users:     users,
		},
	}
	currentDBList.Common.OverwriteVersion("V0")

	currentDB := &tree.Tree{
		Parsed: currentDBList,
	}

	SetQueriedForDatabase(queried, currentDB)
}
*/
