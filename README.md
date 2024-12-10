# Questionnaire System

A robust questionnaire management system built with Go, featuring secure authentication, questionnaire creation and management, response collection, and file handling capabilities.

## Table of Contents
- [Architecture Overview](#architecture-overview)
- [System Design](#system-design)
  - [Data Flow Diagrams](#data-flow-diagrams)
  - [Entity Relationship Diagram](#entity-relationship-diagram)
- [Technical Stack](#technical-stack)
- [Features](#features)
- [Setup Instructions](#setup-instructions)
- [API Documentation](#api-documentation)

## Architecture Overview

The project follows a Clean Architecture pattern with Domain-Driven Design (DDD) principles, organized in the following layers:

1. **Domain Layer** (`internal/*/domain`)
   - Contains business logic and domain models
   - Defines core business rules and entities
   - Independent of external concerns

2. **Port Layer** (`internal/*/port`)
   - Defines interfaces for external dependencies
   - Maintains separation between domain and infrastructure
   - Contains repository interfaces and service contracts

3. **Service Layer** (`internal/*/service`)
   - Implements business use cases
   - Orchestrates domain objects
   - Maintains business rules integrity

4. **Infrastructure Layer** (`pkg/adapter`)
   - Implements external interfaces (database, email, etc.)
   - Handles technical concerns
   - Contains concrete implementations of port interfaces

5. **API Layer** (`api/handler`)
   - Handles HTTP requests and responses
   - Implements REST endpoints
   - Manages request validation and response formatting

### Key Design Patterns Used

1. **Repository Pattern**
   - Abstracts data persistence operations
   - Enables swapping of data storage implementations
   - Centralizes data access logic

2. **Middleware Pattern**
   - Handles cross-cutting concerns
   - Manages authentication and authorization
   - Implements rate limiting

## System Design

### Data Flow Diagrams

#### Authentication Flow
```mermaid
sequenceDiagram
    participant User
    participant AuthService
    participant EmailService
    participant Redis
    participant Database

    User->>AuthService: Register/Login Request
    AuthService->>EmailService: Generate & Send OTP
    EmailService-->>User: Send OTP Email
    AuthService->>Redis: Store OTP
    User->>AuthService: Submit OTP
    AuthService->>Redis: Verify OTP
    AuthService->>Database: Create/Verify User
    AuthService-->>User: Return JWT Token
```

#### Questionnaire Management Flow
```mermaid
sequenceDiagram
    User->>AuthMiddleware: Create Request
    AuthMiddleware->>QuestionnaireService: Validate
    QuestionnaireService->>Database: Store
    User->>MediaService: Upload Media
    MediaService->>Database: Store Reference
    QuestionnaireService-->>User: Return ID
```

#### Media Handling Flow
```mermaid
sequenceDiagram
    participant User
    participant AuthMiddleware
    participant MediaService
    participant FileSystem
    participant Database

    User->>AuthMiddleware: Upload Media Request
    AuthMiddleware->>MediaService: Authenticated Request
    MediaService->>FileSystem: Store File
    MediaService->>Database: Store Metadata
    MediaService-->>User: Return Media UUID

    User->>AuthMiddleware: Download Media Request
    AuthMiddleware->>MediaService: Verify Access
    MediaService->>Database: Check Permissions
    MediaService->>FileSystem: Retrieve File
    MediaService-->>User: Return File
```

### Entity Relationship Diagram

```mermaid
erDiagram
    users ||--o{ questionnaires : creates
    users ||--o{ responses : submits
    users ||--o{ media : uploads
    questionnaires ||--|{ questions : contains
    questions ||--|{ options : has
    questions ||--o{ responses : receives
    questions ||--o{ media : includes

    users {
        integer id PK
        string user_id UK
        string email UK
        string password
        string nat_id UK
        timestamp created_at
        integer role
    }

    questionnaires {
        integer id PK
        string questionnaire_id UK
        integer owner_id FK
        string title
        string description
        integer duration
        boolean editable
        boolean randomable
        timestamp created_at
        timestamp valid_to
    }

    questions {
        integer id PK
        string questionnaire_id FK
        enum type
        integer number
        integer count
        string title
        string media FK
    }

    options {
        integer id PK
        integer question_id FK
        string text
        boolean is_answer
    }

    responses {
        integer id PK
        enum type
        integer user_id FK
        integer question_id FK
        string data
        integer option_id FK
        timestamp created_at
    }

    media {
        integer id PK
        string uuid UK
        integer user_id FK
        string path
        string type
        integer size
        string name
        timestamp created_at
    }
```

## Technical Stack

- **Language**: Go 1.23
- **Web Framework**: Fiber
- **Database**: MySQL
- **Cache**: Redis
- **Authentication**: JWT + OTP
- **Containerization**: Docker + Docker Compose

## Implemented Features

1. **Authentication System**
   - Two-factor authentication with email OTP
   - JWT-based session management
   - Secure password hashing

2. **Rate Limiting**
   - IP-based rate limiting for API endpoints
   - Configurable request limits

3. **Questionnaire Management**
   - Create and manage questionnaires
   - Support for multiple question types
   - Option management for questions

4. **Media Handling**
   - Secure file upload and storage
   - Access control for media files
   - Support for various file types

5. **Response Management**
   - Store and manage questionnaire responses
   - Support for different response types

## Setup Instructions

### Prerequisites
- Docker and Docker Compose
- Git

### Quick Start

1. Clone the repository:
```bash
git clone https://github.com/your-repo/questionnaire
cd questionnaire
```

2. Run the application:
```bash
docker-compose up
```

That's it! The application will automatically:
- Start MySQL and Redis services
- Wait for databases to be ready
- Run database migrations
- Start the application server

The application will be available at `http://localhost:3000`

### Environment

All necessary configurations are already set in:
- `config.json`: Database, Redis, and email configurations
- `Private_key.pem`: JWT signing key
- `Public_key.pem`: JWT verification key

## API Documentation

### Authentication Endpoints

- `POST /api/v1/auth/register/init`: Initialize registration
- `POST /api/v1/auth/register/complete`: Complete registration
- `POST /api/v1/auth/login/init`: Initialize login
- `POST /api/v1/auth/login/complete`: Complete login

### Questionnaire Endpoints

- `POST /api/v1/questionnaire`: Create questionnaire
- `POST /api/v1/questionnaire/questions`: Add questions
- `GET /api/v1/questionnaire/:id`: Get questionnaire
- `PUT /api/v1/questionnaire/:id`: Update questionnaire
- `DELETE /api/v1/questionnaire/:id`: Delete questionnaire

### Media Endpoints

- `POST /api/v1/media/upload`: Upload media
- `GET /api/v1/media/download/:uuid`: Download media

## Error Handling

The application uses standardized error responses:

```json
{
    "error": {
        "message": "Error description"
    }
}
```

Common HTTP status codes:
- `400`: Bad Request
- `401`: Unauthorized
- `403`: Forbidden
- `404`: Not Found
- `429`: Too Many Requests
- `500`: Internal Server Error