# Fluora 🚀
An extensible AI toolkit built in **Go** to effortlessly create and deploy **generative AI tools** using structured instructions.

## 🌟 Features
- 📌 **AI Models Supported:** Currently supports **Google Gemini** models.
- 🔄 **Extensible:** Easily add new AI providers and instructions.
- 📂 **Configurable:** Modify behavior via YAML configuration.
- 📝 **Swagger API Docs:** Auto-generated API documentation.


## 📦 Installation

### **1️⃣ Clone Repository**
```sh
git clone https://github.com/Ja7ad/fluora.git
cd fluora
```

### **2️⃣ Install Dependencies**
Ensure you have Go **1.24** installed.
```sh
go mod tidy
```

## 🚀 Running the Service

### **1️⃣ Prepare Configuration**
Create a config file (e.g., `config.yml`) with the following **example setup**:

```yaml
rest:
  address: "127.0.0.1:8080"
  origins:
    - http://example.com

ai:
  - provider: "gemini"
    api_key: "YOUR_GEMINI_API_KEY"
    rate_limit: 100

logger:
  colorful: true
  max_backups: 0
  rotate_log_after_days: 1
  compress: false
  targets: ['console']
  levels:
    default: 'info'
    _ai: 'warn'
    _service: 'info'
    _rest: 'error'
```

📌 **Note:** Replace `"YOUR_GEMINI_API_KEY"` with a valid **Gemini API Key**.

---

### **2️⃣ Start the Server**
Run the Fluora service:
```sh
go run main.go -config=config.yml
```

Alternatively, build and run:
```sh
go build -o fluora
./fluora -config=config.yml
```


## 📜 **API Documentation (Swagger)**
Once the service is running, you can access the **Swagger UI** at:

```
http://127.0.0.1:8080/swagger/
```

It provides:
- 📌 **Full API Documentation**
- 📩 **Interactive Requests**


## 🔧 **Usage Example**
### **Rewrite API Request**
**Endpoint:** `POST /textcraft/rewrite`

**Request Body:**
```json
{
  "provider": "gemini",
  "model": "gemini-2.0-flash",
  "language": "English",
  "sentence": "This is an example sentence that needs rewriting.",
  "num_suggestions": 5,
  "tone": "formal"
}
```

**Response Example:**
```json
{
  "original_sentence": "This is an example sentence that needs rewriting.",
  "rewritten_sentence": "This sentence requires revision.",
  "suggestions": [
    "This example sentence needs rewording.",
    "An example sentence requiring refinement.",
    "This sentence should be rewritten differently.",
    "A revision is needed for this sentence.",
    "This statement could be rephrased."
  ]
}
```

## 📎 **TODO**
- [ ] support ratelimit
- [ ] support openai
- [ ] support ollama
- [ ] client API key support instead configuration ai if not configurated
- [ ] add codecov report
- [ ] add linter
- [ ] improve code coverage test
- [ ] add subtitle translator instruction feature


## 📜 **License**
Fluora is released under the [**MIT License**](/LICENSE).
