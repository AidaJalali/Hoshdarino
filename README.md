# Hoshdarino

![Project Logo](frontend/public/assets/logo.jpeg)

Hoshdarino is an intelligent software system that provides real-time information about natural disasters. The system helps users stay informed about potential hazards by displaying a map of the country and identifying high-risk areas, enabling them to take necessary precautions. The platform also sends urgent alerts to users during critical situations.

## 🌟 Features

### Interactive Map
- National map display with zoom and navigation capabilities
- Highlighted disaster-affected areas and high-risk zones
- Real-time updates on disaster locations

### Natural Disaster Coverage
- **Earthquakes**: Tectonic movements and high-risk zones
- **Floods**: Heavy rainfall, snowmelt, and flood-prone areas
- **Fires**: Wildfires and urban fires
- **River Overflows**: Rising river levels and flood predictions
- **Storms and Tornadoes**: Severe weather alerts and forecasts
- **Snow and Frost**: Heavy snowfall and freezing conditions

### Alert System
- User subscription for immediate alerts
- SMS and notification system for dangerous areas
- Customizable alert preferences based on disaster type

### Data Analysis
- Real-time disaster data updates
- Historical data analysis and pattern recognition
- Statistical visualization tools

## 🚀 Project Goals

- Provide real-time information about various natural disasters
- Raise public awareness about natural hazards
- Create a fast alerting platform for people at risk
- Use analytical data to identify disaster occurrence patterns

## 🛠️ Technical Stack

### Backend (Go)
- **Framework**: Gin/Echo for high-performance HTTP routing
- **Database**: MongoDB with official Go driver
- **Authentication**: JWT-based authentication
- **API**: RESTful API with OpenAPI/Swagger documentation
- **Testing**: Go's built-in testing package + testify
- **Logging**: Zap logger for structured logging
- **Configuration**: Viper for configuration management

### Frontend (React + TypeScript)
- **Framework**: React 18+ with TypeScript
- **Build Tool**: Vite for fast development and building
- **State Management**: Zustand for global state
- **Styling**: Tailwind CSS for utility-first styling
- **API Client**: React Query for data fetching
- **Maps**: Mapbox GL JS for interactive maps
- **Testing**: Vitest + React Testing Library
- **Form Handling**: React Hook Form
- **Validation**: Zod for schema validation

## 📁 Project Structure

```
hoshdarino/
├── frontend/                    # React frontend application
│   ├── public/                  # Static files
│   │   ├── assets/             # Images, fonts, etc.
│   │   └── index.html          # Main HTML file
│   ├── src/                    # Source code
│   │   ├── components/         # React components
│   │   │   ├── map/           # Map-related components
│   │   │   ├── alerts/        # Alert system components
│   │   │   └── common/        # Shared components
│   │   ├── pages/             # Page components
│   │   ├── services/          # API services
│   │   ├── hooks/             # Custom React hooks
│   │   ├── store/             # Zustand store
│   │   ├── types/             # TypeScript types
│   │   ├── utils/             # Utility functions
│   │   ├── styles/            # Global styles
│   │   └── App.tsx            # Main application component
│   ├── package.json           # Frontend dependencies
│   └── tsconfig.json          # TypeScript configuration
│
├── backend/                    # Go backend application
│   ├── cmd/                    # Application entry points
│   │   └── api/               # Main API server
│   ├── internal/              # Private application code
│   │   ├── api/              # API handlers
│   │   ├── config/           # Configuration
│   │   ├── middleware/       # HTTP middleware
│   │   ├── models/           # Data models
│   │   ├── repository/       # Database operations
│   │   ├── service/          # Business logic
│   │   └── utils/            # Utility functions
│   ├── pkg/                  # Public library code
│   ├── api/                  # API documentation
│   ├── go.mod                # Go module file
│   └── go.sum                # Go module checksum
│
├── docs/                     # Project documentation
│   ├── api/                  # API documentation
│   ├── architecture/         # Architecture diagrams
│   └── deployment/           # Deployment guides
│
├── scripts/                  # Utility scripts
│   ├── setup.sh             # Setup script
│   └── deploy.sh            # Deployment script
│
├── .gitignore               # Git ignore file
└── README.md                # Main project README
```

## 📅 Project Timeline

| Phase | Description | Duration | Timeline |
|-------|-------------|----------|----------|
| Week 1 | Development environment setup and database creation | 1 week | 1404/1/12 - 1404/1/19 |
| Week 2 | Map-based user interface implementation | 1 week | 1404/1/20 - 1404/1/27 |
| Week 3 | Alert and notification system development | 1 week | 1404/1/28 - 1404/2/4 |
| Weeks 4-5 | Data analysis and prediction algorithms | 2 weeks | 1404/2/6 - 1404/2/19 |
| Weeks 6-7 | System testing and debugging | 2 weeks | 1404/2/20 - 1404/3/2 |
| Weeks 8-9 | Optimization and final deployment | 2 weeks | 1404/3/3 - 1404/3/16 |

## 🔧 Installation

### Prerequisites
- Go 1.21 or later
- Node.js 18 or later
- MongoDB 6.0 or later
- Git

### Backend Setup
```bash
cd backend
go mod download
go run cmd/api/main.go
```

### Frontend Setup
```bash
cd frontend
npm install
npm run dev
```

## 📝 License

*License information will be added as the project progresses*

## 🤝 Contributing

*Contribution guidelines will be added as the project progresses*

## 📞 Contact

*Contact information will be added as the project progresses*  
