# hsleiden-Golang-Basic-REST-API

### <a name="#"></a> Todos
=======

| Method | URL                                   | Description                            | Action                       |
|--------|---------------------------------------|----------------------------------------|------------------------------|
|`GET`   | `/todos`                              | Get a list of todos                    | [find](#find)                |
|`GET`   | `/todos/{todoId}`                     | Get a todo by :id.                     | [findOne](#findOne)          |
|`POST`  | `/todos`                              | Create a new todo.                     | [create](#create)            |


#### <a name="find"></a>`GET` find
------

| Method | URL                                   | Description                            |
|--------|---------------------------------------|----------------------------------------|
|`GET`   | `/todos`                              | Get a list of todos.                   |


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

#### <a name="findOne"></a>`GET` findOne
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

#### <a name="create"></a>`POST` create
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

