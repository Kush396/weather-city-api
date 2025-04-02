# Weather API Integration & City Management Service

This project is a RESTful web API that integrates with OpenWeatherMap to fetch current weather data for cities and provides CRUD operations for managing city information locally.

## Features

- Fetch real-time weather data for any city using OpenWeatherMap API
- Create, read, update, and delete city information
- Simple in-memory storage for city data
- JSON-based API responses

## Prerequisites

- Go 1.16 or higher
- OpenWeatherMap API key (get one at [OpenWeatherMap](https://openweathermap.org/api))

## Project Structure

```
weather-city-api/
├── api/
│   ├── handlers/
│   │   ├── city_handler.go
│   │   └── weather_handler.go
│   └── routes.go
├── cmd/
│   └── main.go
├── models/
│   └── city.go
├── repository/
│   └── city_repository.go
├── services/
│   ├── city_service.go
│   └── weather_service.go
├── util/
│   └── config.go
├── app.env
└── README.md
```

## Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/Kush396/weather-city-api.git
   cd weather-city-api
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Configure your API key:
   - Open `app.env` file in the root directory
   - Replace the placeholder value for `WEATHER_API_KEY` with your OpenWeatherMap API key
   ```
   PORT=8080
   WEATHER_API_KEY=your_api_key_here
   ```

4. Run the application:
   ```
   cd cmd
   go run main.go
   ```

5. The server will start on port 8080 (or the port specified in your app.env file)

## API Endpoints

### Weather API

#### Get Weather Data for a City

```
GET /api/weather?city={cityName}
```

**Parameters:**
- `city` (required): Name of the city to get weather data for

**Example Request:**
```
GET /api/weather?city=Bharatpur
```

**Example Response:**
```json
{"base":"stations",
"clouds":{
   "all":0
},
"cod":200,
"coord":{
   "lat":27.2167,
   "lon":77.4833
},
"dt":1743553982,
"id":1276128,
"main":{
   "feels_like":19.64,
   "grnd_level":990,
   "humidity":11,
   "pressure":1011,
   "sea_level":1011,
   "temp":21.18,
   "temp_max":21.18,
   "temp_min":21.18
},
"name":"Bharatpur",
"sys":{
   "country":"IN",
   "sunrise":1743554377,
   "sunset":1743599244},
"timezone":19800,
"visibility":10000,
"weather":[
   {
      "description":"clear sky",
      "icon":"01n",
      "id":800,
      "main":"Clear"
   }
],
"wind":{
   "deg":358,
   "gust":1.89,
   "speed":1.81}
}
```

### City Management API

#### Get All Cities

```
GET /api/cities
```

**Example Response:**
```json
[
  {
    "id": 1,
    "name": "New York",
    "country": "USA",
    "description": "The Big Apple"
  },
  {
    "id": 2,
    "name": "London",
    "country": "UK",
    "description": "Capital of England"
  }
]
```

#### Get City by ID

```
GET /api/cities/{id}
```

**Example Request:**
```
GET /api/cities/1
```

**Example Response:**
```json
{
  "id": 1,
  "name": "New York",
  "country": "USA",
  "description": "The Big Apple"
}
```

#### Create a New City

```
POST /api/cities
```

**Request Body:**
```json
{
  "name": "Paris",
  "country": "France",
  "description": "City of Love"
}
```

**Example Response:**
```json
{
  "id": 3,
  "name": "Paris",
  "country": "France",
  "description": "City of Love"
}
```

#### Update a City

```
PUT /api/cities/{id}
```

**Example Request:**
```
PUT /api/cities/3
```

**Request Body:**
```json
{
  "name": "Paris",
  "country": "France",
  "description": "City of Lights"
}
```

**Example Response:**
```json
{
  "id": 3,
  "name": "Paris",
  "country": "France",
  "description": "City of Lights"
}
```

#### Delete a City

```
DELETE /api/cities/{id}
```

**Example Request:**
```
DELETE /api/cities/3
```

**Example Response:**
```json
{
  "message": "City deleted successfully"
}
```

## Error Handling

The API returns appropriate HTTP status codes along with error messages:

- 400 Bad Request: For invalid input parameters
- 404 Not Found: When a requested city is not found
- 500 Internal Server Error: For server-side errors

Example error response:
```json
{
  "error": "City not found"
}
```

## Notes

- This implementation uses an in-memory data store for city information. Data will be lost when the server restarts.
- The weather data is fetched directly from OpenWeatherMap API with no caching mechanism.
