# 🚀 Project Zero

Welcome to **Project Zero** - cool name, right?

This project is a lightweight boilerplate that i created to kickstart new Go-Based web applications without having to rewrite the same code every time. <br>
Think of it as my __starter pack__ for spinning up HTTP servers with templating and styling baked in.

Whether you're prototyping fast or building something small and clean, this base setup will save you time and typing. ⌨️✨

---

### Tech Stack Overview

### Backend
- Built with pure **[Go](https://go.dev/doc/effective_go) (Golang)** 🦫 - no frameworks, just standard libraries for full controll and simplicity.
- Uses **[Templ](https://templ.guide/)** for HTML templating, enabling efficient and clean server-side rendering.

### Frontend
- Styled with **[SASS](https://sass-lang.com/documentation/)** for flexible and maintanable CSS.

## 📦 What's Inside?

This repo includes everything you need to get up and running:

- ✅ **Boilerplate HTTP server** — fully set up and ready to serve routes
- 📁 **Organized folder structure** — a clear project layout to keep things tidy
- 🔁 **Hot reload support** — auto-compile your Go and SASS files during development
- 🏗️ **Build command** — easily compile your app for production

## Router

In your `app`, the `router` is the component where you define and attach your routes so the application can listen and respond to incoming requests.

To keep things well-organized 🗂️, it's recomended to group related routes into logical folders. One way to do this is by creating a route group using the following structure:

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

We'll discuss the parameters for `HandleFunc` in more detail shortly.

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
> __🧩 Consistency tip:__ <br>
> Suffix all route handler function names with "Route" (e.g., `RegisterRoute`, `LoginRoute`).

As you can see the function is "tied" to the `UserRoutes` type, in another worlds, is on our routes group for user.

