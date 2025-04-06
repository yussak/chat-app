package infrastructure

import (
	"server/db"
	"time"
)

type Workspace struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	OwnerID   int       `json:"ownerId"`
	Theme     string    `json:"theme"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FindAll() ([]Workspace, error) {
	query := `SELECT id, name, owner_id, theme, created_at, updated_at FROM workspaces`
		rows, err := db.DB.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var workspaces []Workspace
		for rows.Next() {
			var workspace Workspace
			if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.OwnerID, &workspace.Theme, &workspace.CreatedAt, &workspace.UpdatedAt); err != nil {
				return nil, err
			}
			workspaces = append(workspaces, workspace)
		}

		return workspaces, nil
}