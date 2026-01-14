Subscription Management Service

A backend service for managing user subscriptions and calculating their total cost for a selected period.

Method	Endpoint	Description
POST	/api/subs	Create subscription
GET	/api/subs/{id}	Get subscription by ID
GET	/api/subs/list	List subscriptions
PUT	/api/subs/{id}	Update subscription
DELETE	/api/subs/{id}	Delete subscription
GET	/api/subs/total	Calculate total cost

Database PostgreSQL

Swagger UI is available at:

http://localhost:7777/swagger/index.html

To generate Swagger docs:
swag init

Docker
Start the service
docker-compose up --build

Start PostgreSQL manually

Create .env from .env.example

Run the service:

go run main.go

Go-Echo framework-GORM-PostgreSQL