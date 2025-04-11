package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// RoleRepository определяет методы работы с таблицей Roles.
type RoleRepository struct {
	db *sql.DB
}

// NewRoleRepository возвращает новый экземпляр репозитория для ролей.
func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// Create добавляет новую роль и возвращает сгенерированный ID.
func (r *RoleRepository) Create(role *models.Role) error {
	query := `
		INSERT INTO Roles (name, description)
		VALUES ($1, $2)
		RETURNING id
	`
	err := r.db.QueryRow(query, role.Name, role.Description).Scan(&role.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать роль: %w", err)
	}
	return nil
}

// Update обновляет данные роли по ID.
func (r *RoleRepository) Update(role *models.Role) error {
	query := `
		UPDATE Roles
		SET name = $1, description = $2
		WHERE id = $3
	`
	result, err := r.db.Exec(query, role.Name, role.Description, role.ID)
	if err != nil {
		return fmt.Errorf("не удалось обновить роль: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("роль с id=%d не найдена", role.ID)
	}
	return nil
}

// GetAll возвращает список всех ролей.
func (r *RoleRepository) GetAll() ([]models.Role, error) {
	query := `SELECT id, name, description FROM Roles`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список ролей: %w", err)
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.Description); err != nil {
			return nil, fmt.Errorf("ошибка при чтении роли: %w", err)
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// GetByID возвращает роль по её ID.
func (r *RoleRepository) GetByID(id int) (*models.Role, error) {
	query := `SELECT id, name, description FROM Roles WHERE id = $1`
	var role models.Role
	err := r.db.QueryRow(query, id).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("не удалось получить роль: %w", err)
	}
	return &role, nil
}

// Delete удаляет роль по её ID.
func (r *RoleRepository) Delete(id int) error {
	query := `DELETE FROM Roles WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить роль: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("роль с id=%d не найдена", id)
	}
	return nil
}
