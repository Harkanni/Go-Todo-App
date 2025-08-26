# 📝 Go TodoApp  

A simple Todo application built with the **Go programming language**.  
This project explores Go fundamentals such as structs, slices, functions, and especially **concurrency with Goroutines and channels**.  

---

## 🚀 Features (Current Implementation)  

- **Define a TODO item (`TODO struct`)**  
  - Fields: `title (string)`, `duration (time.Duration)`, `status (string)`, `completed (bool)`  

- **Store TODO items**  
  - `TODO_LIST` (global slice for active tasks)  

- **Store completed tasks**  
  - `Completed_Todo_List` (global slice for finished tasks)  

- **Add a TODO item (`addTodoItem`)**  
  - Creates and appends a new task to `TODO_LIST`  
  - Prints confirmation message  

- **List all TODO items (`getAllTodoItems`)**  
  - Iterates through `TODO_LIST` and prints task details  

- **Get a TODO item by ID (`getTodoItem`)**  
  - Returns the task if index is valid, else returns empty TODO  

- **Move TODO item to completed (`moveTodoItemToCompleted`)**  
  - Removes a task from `TODO_LIST` and adds it to `Completed_Todo_List`  
  - ⚠️ Known issue: bug when indexing after slicing  

---

## 🔮 Future Improvements  

- Task timers and reminders (`time.Duration`)  
- Multiple task states (e.g., pending, in-progress, completed)  
- Mark tasks as completed directly (toggle `completed = true`)  
- Delete or edit tasks  
- Filter/search tasks by state or keyword  
- Save/load tasks for persistence (file, DB, JSON)  
- Recurring tasks and deadlines  
- CLI interaction (menus, prompts, arguments)  

---

## 🛠️ Tech Stack  

- **Language:** Go  
- **Core Concepts:** Goroutines, Channels, Slices, Structs  

---

## 📂 Project Status  

✅ Functional basic Todo application  
⚠️ Known bug when moving items after slicing  
🔜 Planned enhancements for more robust task management  

---

## 🤝 Contributing  

This is primarily a learning project, but suggestions and improvements are welcome.  
Feel free to fork the repo and create pull requests.  

---

## 📜 License  

This project is open-source and available under the [MIT License](LICENSE).  
