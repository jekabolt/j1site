package jdblayer

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type pgDataStore struct {
	*sql.DB
}

func NewPgqlDataStore(conn string) (*pgDataStore, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &pgDataStore{
		DB: db,
	}, nil

}

func (pgql *pgDataStore) AddMember(cm *CrewMember) error {
	pgql.QueryRow("INSERT INTO users(id,Name,SecurityClearance,Title) VALUES($1,$2,$3,$4)", cm.ID, cm.Name, cm.SecClearance, cm.Position)
	return nil
}

func (pgql *pgDataStore) FindMember(id int) (CrewMember, error) {
	cm := CrewMember{}
	err := pgql.QueryRow("Select * from users where id = $1", id).Scan(&cm.ID, &cm.Name, &cm.SecClearance, &cm.Position)
	if err != nil {
		return cm, err
	}
	return cm, err
}
func (pgql *pgDataStore) AllMembers() (crew, error) {
	rows, err := pgql.Query("Select * from users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := crew{}
	for rows.Next() {
		member := CrewMember{}
		err = rows.Scan(&member.ID, &member.Name, &member.SecClearance, &member.Position)
		if err == nil {
			members = append(members, member)
		}
	}
	err = rows.Err()
	return members, err
}
