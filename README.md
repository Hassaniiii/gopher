# Gopher
This is a simple web server written in _Go_ for educational purposes.

### Requirements
- Go version 1.14 and upper

### Build
To build an executable output, run the following command in the project's directory

```
go build .
```

### Run
To run the project, you can either execute the output from the previous step by running 

```
./gopher
```

Or run the project directly by the following command

```
go run .
```

### Test
The project will start a local web server on [localhost:3000](http://localhost:3000) that is accessible through the following API

##### All Users
```
curl http://localhost:3000/users
```

##### User with ID
```
curl http://localhost:3000/users/<id>
```

##### Add a new user
To add a new user, the _ID_ parameter should always be set as 0.
```
curl -d '{"ID":0, "FirstName":"<value>", "LastName":"<value>"}' http://localhost:3000/users
```

##### Update a user by ID
```
curl -d '{"ID":<value>, "FirstName":"<value>", "LastName":"<value>"}' -X PUT http://localhost:3000/users/<id>
```

##### Remove a user by ID
```
curl -X DELETE http://localhost:3000/users/<id>
```