Quizzer
=======

Introduction
------------

*Quizzer* is a quiz server. It loads a set of tasks (questions and their respective answers)
and exposes them using a simple RESTful API.

Parameters
----------

*Quizzer* has exactly one paramter, `--config` to specify its configuration file location.
The configuration contains information about the server instance as well as the catalogue of
tasks.

Configuration
-------------

### Server configuration

The server configuration determines the servers listen address and port.

| Parameter     | Description |
|---------------|-------------|
| listenAddress | address to listen on for incoming requests |
| listenPort    | port to listen on for incoming requests |

Example:

```json
{
    "listenAddress": "0.0.0.0",
    "listenPort": 8080
}
```

### Task configuration

Tasks are a combination of the question and the answering possibilities, of which
at least one is marked as the correct solution.

Example:

```json
    "tasks": [
        {
            "question": "1+1",
            "answers": [
                { "text": "1", "isCorrect": false},
                { "text": "2", "isCorrect": true},
                { "text": "3", "isCorrect": false},
                { "text": "4", "isCorrect": false}
            ]
        }
    ]
```
