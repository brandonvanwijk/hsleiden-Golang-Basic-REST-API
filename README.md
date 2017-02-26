# hsleiden-Golang-Basic-REST-API

### <a name="#"></a> Todos
=======

| Method | URL                                   | Description                            | Action                       |
|--------|---------------------------------------|----------------------------------------|------------------------------|
|`GET`   | `/todos`                              | Get a list of todos                    | [TodoIndex](#todoIndex)      |
|`GET`   | `/todos/{todoId}`                     | Get a todo by :id.                     | [TodoShow](#todoShow)        |
|`POST`  | `/todos`                              | Create a new todo.                     | [TodoCreate](#todoCreate)    |


#### <a name="todoIndex"></a>`GET` TodoIndex
------

| Method | URL                                   | Description                            |
|--------|---------------------------------------|----------------------------------------|
|`GET`   | `/todos`                              | Get a list of estimates.               |

**Params**

| Param    | Type  | Default         | Description                        |
|----------|-------|-----------------|------------------------------------|
|`limit`   | `int` | 25              | Amount of items returned           |
|`offset`  | `int` | 0               | Offset                             |

**Response**
```json
[
  {
    "id": 1,
    "name": "Write Presentation",
    "completed": false,
    "due": "0001-01-01T00:00:00Z"
  },
  {
    "id": 2,
    "name": "Host meetup",
    "completed": false,
    "due": "0001-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "name": "New Todo",
    "completed": false,
    "due": "0001-01-01T00:00:00Z"
  },
  {
    "id": 4,
    "name": "Brandon van Wijk",
    "completed": false,
    "due": "0001-01-01T00:00:00Z"
  }
]
```

#### <a name="todoShow"></a>`GET` TodoShow
------

| Method | URL                                   | Description                            |
|--------|---------------------------------------|----------------------------------------|
|`GET`   | `/todos/{todoId}`                     | Get a todo by it's id.                 |

**Params**

| Param    | Type  | Default         | Description                        |
|----------|-------|-----------------|------------------------------------|
|`todoId`  | `int` | -               | Id of todo that you want to return |

**Response**
```json
[
  {
    "id": 1,
    "name": "Write Presentation",
    "completed": false,
    "due": "0001-01-01T00:00:00Z"
  }
]
```

#### <a name="todoCreate"></a>`POST` TodoCreate
------

| Method | URL                                   | Description                            |
|--------|---------------------------------------|----------------------------------------|
|`POST`   | `/todos`                             | Creates new Todo                       |

**Params**

| Param   | Type   | Default         | Description                        |
|---------|--------|-----------------|------------------------------------|
|`name`   |`string`| none            | Name of new Todo                   |

**Response**
```json
[
  {
    "id": 5,
    "name": "Test User",
    "completed": false,
    "due": "0001-01-01T00:00:00Z"
  }
]
```

