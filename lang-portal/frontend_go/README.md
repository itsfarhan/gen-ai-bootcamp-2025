# Language Learning Portal - Frontend

This directory contains the frontend application for the Language Learning Portal. The frontend is built using HTMX and connects to the Go backend API.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Access to the backend API (running on a specified port)

### Starting the Server

1. Navigate to the frontend_go directory:
```bash
cd path/to/frontend_go
```

2. Run the Go server:
```bash
go run main.go
```

This will start the frontend server on port 8080 by default. You can access the web interface by opening a browser and navigating to `http://localhost:8080`.

## Application Pages and Features

### Dashboard

The dashboard provides a quick overview of your language learning progress, including:
- Recent activity summary
- Statistics on words learned
- Upcoming study reminders
- Quick links to start new study sessions

### Words

The Words page allows you to manage your vocabulary:
- View all words in a paginated table
- Add new words with translations, examples, and notes
- Edit existing words
- Delete words
- Search and filter your word collection

### Groups

The Groups page helps you organize words into meaningful collections:
- Create new word groups with names and descriptions
- Add words to groups
- Remove words from groups
- View all words within a specific group
- Edit or delete groups

### Study Sessions

The Study Sessions page enables active learning:
- Start new study sessions with selected word groups
- Choose different study activities
- Track your progress during sessions
- View history of past sessions with performance statistics

### Activities

The Activities page allows you to configure different learning exercises:
- Multiple choice quizzes
- Flashcards
- Writing practice
- Customize settings for each activity type

### Settings

The Settings page allows you to personalize your learning experience:
- Set default session duration
- Configure timer settings
- Adjust notification preferences
- Customize display options

## Connecting to the Backend API

The frontend application uses HTMX to communicate with the backend API. All API calls are made to endpoints on the backend server. By default, the frontend expects the backend to be available at `http://localhost:3000/api`.

To configure a different backend API URL, you can modify the `apiBaseUrl` variable in the main.go file.

### API Endpoints Used

- `/api/words` - For word management (GET, POST, PUT, DELETE)
- `/api/groups` - For group management (GET, POST, PUT, DELETE)
- `/api/sessions` - For study session management
- `/api/activities` - For activity configuration
- `/api/dashboard` - For dashboard statistics and information
- `/api/settings` - For user settings

## Development

To modify the frontend:
1. Edit the HTML files in the frontend_go directory
2. Add or modify static assets in the assets directory
3. Restart the server to see your changes

## Troubleshooting

- If you encounter CORS issues, ensure your backend API is configured to allow requests from the frontend origin.
- If the frontend cannot connect to the backend, check that the backend server is running and accessible.
- For template loading errors, verify that all HTML files are in the correct location.

