# Fiber Todo Application

## Endpoint Description

`localhost:8000/api/todos`

request: GET
description: To get all todos

`localhost:8000/api/todos/:id`

request: GET
description: Get todo by id

`localhost:8000/api/todos`
request: POST

```
input: {
title : String
}
```

description: Create new todo

`localhost:8000/api/todos/:id`
request: PUT

```
input: {
title : String,
completed : Boolean
}
```

description: Update todo

`localhost:8000/api/todos/:id`
request: DELETE
description: Delete todo
