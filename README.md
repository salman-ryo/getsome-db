# GetSome DB 🗄️

**GetSome DB** is a lightweight, JSON-based database server written in Go. It allows you to store and retrieve JSON data using HTTP endpoints, providing a simple and flexible solution for small-scale data storage.

### 🚀 Features:
- **JSON-Based Storage**: Store your data in JSON format for easy retrieval.
- **Simple API**: Use basic HTTP routes to interact with the database.
- **Timeouts**: Configured read, write, and idle timeouts for efficient server management.
- **Health Check**: Endpoint to check the server's status.
- **Built with Go**: Fast and efficient performance using Go.

### 📦 Installation

1. **Clone the Repository**:
   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/getsomedb.git
   ```

2. **Navigate to the Project Directory**:

   ```bash
   cd getsomedb
   ```

3. **Install Dependencies**:
   This project uses Go modules for dependency management. Run the following to install dependencies:

   ```bash
   go mod tidy
   ```

4. **Run the Server**:
   Start the server by running the `main.go` file:

   ```bash
   go run main.go
   ```

   The server will be running on `http://localhost:8080`.

### 📡 API Endpoints

- **GET `/health`**  
  Check the status of the server. It will return `"Server is healthy!"`.

  Example:
  ```bash
  curl http://localhost:8080/health
  ```

### ⚙️ Configuration

You can configure the following constants in the `server` package:

- **`port`**: Port number for the server to run (default is `8080`).
- **`readTimeOut`**: Time in seconds for the read timeout (default is `5`).
- **`writeTimeOut`**: Time in seconds for the write timeout (default is `10`).
- **`idleTimeOut`**: Time in seconds for the idle timeout (default is `120`).

### 📝 Example Usage

After running the server, you can test the health check endpoint:

```bash
curl http://localhost:8080/health
```

Response:
```json
"Server is healthy!"
```

### ⚠️ Error Handling

- If the server fails to start, the error message will be displayed in the console with details.

### 🧑‍💻 Contributing

Contributions are welcome! If you'd like to help improve GetSome DB, please fork the repository and create a pull request with your changes. You can also report bugs or suggest new features.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a pull request

### 📄 License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

### 💬 Contact

- **Author**: Your Name
- **Email**: your.email@example.com
- **Website**: [yourwebsite.com](https://yourwebsite.com)
- **GitHub**: [github.com/yourusername](https://github.com/yourusername)

---

Made with ❤️ in Go

### Breakdown of Sections:
- **Project Overview**: Describes what your project is and the key features.
- **Installation Instructions**: Step-by-step guide on how to set up the project.
- **API Endpoints**: Lists the available HTTP routes and how to interact with them.
- **Configuration**: Describes configurable constants like the server's port and timeouts.
- **Example Usage**: Shows an example of using the health check endpoint.
- **Contributing**: Explains how others can contribute to the project.
- **License**: Information about the project’s license.
- **Contact Information**: Your contact details for communication.

You can add more sections as your project grows, like database schema, additional endpoints, or more detailed usage instructions.