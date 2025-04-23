# Project Zero
Cool name right?

This project serve as a base with code that i use in a lot of different projects, so for sake of my keys life duration i choose to create this project to not have to type that over and over again.

### Stack
This project uses a very lightweight stack.

#### Backend
It uses pure `golang` to run the http server. It alsos use `Templ` to render the pages before serving them.

#### Frontend
It only have a styling framework that is `SASS`.

### What i'm going to see here?

Here will have:
- The `boilerplate code` to have a http server running with. 
- Project structure. "AKA folders"
- A command to have all the compilers on hotReload.
- A command to build the project.

# Router

In your `app`, the `router` is the component where you define and attach your routes so the application can listen and respond to incoming requests.

To keep things well-organized, it's recomended to group related routes into logical folders. One way to do this is by creating a route group using the following structure:

```bash
routes\       # Main folder for all your routes.
│      
├── user\     # Custom folder for routes related to users.  
│   │   
│   │   # These are your defined routes.
│   │   
│   ├── registerRoute.go   
│   ├── loginRoute.go
│   │   
│   │   # This file is where you gonna define
│   │   # the type and the function to register
│   │   # the routes you defined above.
│   └── handle_routes.go
```

Inside the `user` folder, you'll define a Go file (e.g., `handle_routes.go`) that contains both a type and a function to manage the user-related routes.

### Defining a route group

First, define a struct that represents your route group. This struct should include a reference to the application, so you can register routes with it:

```go
type UserRoutes struct {
    App *app.Application
}
```

Next, define a function to register the routes for this group:

```go
func RegisterRoutes(app *app.Application){
    userRoutes := &UserRoutes{App: app}
    
    app.Router.HandleFunc("/register", userRoutes.RegisterRoute)
}
```

> [!NOTE]
> We'll discuss the parameters for `HandleFunc` in more detail shortly.

> [!NOTE]
> Grouping routes is optional. If you have a route that don't make sense the group with others, there's no need to create an entire group just for that.

### Creating routes

Inside the `user` folder, you can define your endpoints. For example, a route to handle user registration.

```go
func (user *UserRoutes) RegisterRoute(w http.ResponseWriter, r *http.Request) {
    // Implementation for the user registration route
}
```

> [!NOTE]
> For consistency, suffix all route handler function names with "Route" (e.g., `RegisterRoute`, `LoginRoute`).

As you can see the function is "tied" to the `UserRoutes` type, in another worlds, is on our routes group for user.

