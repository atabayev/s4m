# Tasks for s4m

Tasks 1.1-1.3 and 2.1 are located in the folder 'OtherTasks'.
Tasks 2.2 and 2.3 are implemented in the project.

## Description

The program handles HTTP requests for registering a single event, multiple events, and retrieving events based on type and date range. All events are stored in a ClickHouse database.

## Installation

Detailed installation instructions.

1. Clone the repository: `[git clone https://github.com/your_repository.git](https://github.com/atabayev/s4m.git)`
2. Navigate to the project directory: `cd s4m`
3. Install dependencies: `go mod download`.

## Usage

Make an HTTP request to the API:
- `/api/event/`, POST. with body params as:
```json
        {
    		"eventType": "login",
    		"userID": 110,
    		"evenTime": "2023-10-16T03:21:16.9379379+06:00",
    		"payload": "no"
    	}
```

- `/api/set/events`, POST, with body as:
```json
        [{
          "eventID":0,"eventType":"check","userID":105,"evenTime":"2023-10-16T03:25:01.1267593+06:00","payload":"0"
        },
        {
          "eventID":0,"eventType":"register","userID":106,"evenTime":"2023-10-16T03:25:01.1267593+06:00","payload":"{\"login\":\"user\"}"
        },
        {
          "eventID":0,"eventType":"login","userID":102,"evenTime":"2023-10-16T03:25:01.1267593+06:00","payload":""
      }]
```
    
- `/api/get/events`, GET, with params as:
```url
http://localhost:8080/api/get/events?eventType=event_type_3&finishTime=31.12.2022+23%3A59&startTime=01.01.2020+00%3A00
```



## License

This project is licensed under the MIT License - see the [LICENSE](https://mit-license.org/) file for details.

