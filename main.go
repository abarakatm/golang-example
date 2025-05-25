package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

// Employee represents the employee model
type Employee struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Position  string  `json:"position"`
	Salary    float64 `json:"salary"`
}

// In-memory storage for employees
var employees = make(map[string]Employee)

func main() {
	app := fiber.New()

	// Add middleware
	app.Use(logger.New())

	// Define API routes
	api := app.Group("/api")

	// Employee routes
	api.Post("/employees", createEmployee)
	api.Get("/employees", getAllEmployees)
	api.Get("/employees/:id", getEmployee)
	api.Delete("/employees/:id", deleteEmployee)

	// Start server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

// createEmployee handles the creation of a new employee
func createEmployee(c *fiber.Ctx) error {
	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Generate UUID for the employee
	employee.ID = uuid.New().String()

	// Store employee in our "database"
	employees[employee.ID] = *employee

	return c.Status(fiber.StatusCreated).JSON(employee)
}

// getAllEmployees returns all employees
func getAllEmployees(c *fiber.Ctx) error {
	// Convert map to slice for better JSON response
	employeeList := make([]Employee, 0, len(employees))
	for _, emp := range employees {
		employeeList = append(employeeList, emp)
	}

	return c.JSON(employeeList)
}

// getEmployee returns a specific employee by ID
func getEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	employee, exists := employees[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	return c.JSON(employee)
}

// deleteEmployee removes an employee by ID
func deleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	_, exists := employees[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	delete(employees, id)

	return c.Status(fiber.StatusNoContent).Send(nil)
}
