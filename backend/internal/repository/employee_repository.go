package repository

import (
	"database/sql"
	"fmt"
	"time"

	"backend/internal/models"
)

// EmployeeRepository содержит методы для работы с таблицей Employees.
type EmployeeRepository struct {
	db *sql.DB
}

// NewEmployeeRepository возвращает новый репозиторий для сотрудников.
func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

// Create добавляет нового сотрудника в БД.
func (r *EmployeeRepository) Create(emp *models.Employee) error {
	query := `
		INSERT INTO Employees (first_name, last_name, position, hire_date, salary, email, phone)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		emp.FirstName,
		emp.LastName,
		emp.Position,
		emp.HireDate,
		emp.Salary,
		emp.Email,
		emp.Phone,
	).Scan(&emp.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать сотрудника: %w", err)
	}
	return nil
}

// Delete удаляет сотрудника по его ID.
func (r *EmployeeRepository) Delete(id int) error {
	query := `DELETE FROM Employees WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить сотрудника: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("сотрудник с id=%d не найден", id)
	}
	return nil
}

// GetAll возвращает список всех сотрудников.
func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	query := `
		SELECT id, first_name, last_name, position, hire_date, salary, email, phone
		FROM Employees
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список сотрудников: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		var hireDate time.Time
		if err := rows.Scan(
			&e.ID,
			&e.FirstName,
			&e.LastName,
			&e.Position,
			&hireDate,
			&e.Salary,
			&e.Email,
			&e.Phone,
		); err != nil {
			return nil, fmt.Errorf("ошибка чтения данных сотрудника: %w", err)
		}
		e.HireDate = hireDate
		employees = append(employees, e)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерировании по сотрудникам: %w", err)
	}

	return employees, nil
}

// GetByID возвращает данные одного сотрудника по его ID.
func (r *EmployeeRepository) GetByID(id int) (*models.Employee, error) {
	query := `
		SELECT id, first_name, last_name, position, hire_date, salary, email, phone
		FROM Employees
		WHERE id = $1
	`
	var e models.Employee
	var hireDate time.Time
	err := r.db.QueryRow(query, id).Scan(
		&e.ID,
		&e.FirstName,
		&e.LastName,
		&e.Position,
		&hireDate,
		&e.Salary,
		&e.Email,
		&e.Phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Если сотрудник не найден, возвращаем nil
		}
		return nil, fmt.Errorf("не удалось получить сотрудника: %w", err)
	}
	e.HireDate = hireDate
	return &e, nil
}

// Update обновляет данные сотрудника.
func (r *EmployeeRepository) Update(emp *models.Employee) error {
	query := `
		UPDATE Employees
		SET first_name = $1,
		    last_name = $2,
		    position = $3,
		    hire_date = $4,
		    salary = $5,
		    email = $6,
		    phone = $7
		WHERE id = $8
	`
	result, err := r.db.Exec(query,
		emp.FirstName,
		emp.LastName,
		emp.Position,
		emp.HireDate,
		emp.Salary,
		emp.Email,
		emp.Phone,
		emp.ID,
	)
	if err != nil {
		return fmt.Errorf("не удалось обновить данные сотрудника: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("сотрудник с id=%d не найден", emp.ID)
	}
	return nil
}
