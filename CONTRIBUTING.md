### 📄 **CONTRIBUTING.md**

# Contributing to Fluora

🚀 Welcome to **Fluora**! We appreciate your interest in contributing. Whether you're fixing a bug, improving documentation, or adding a new feature, your help is valuable. Please follow the guidelines below to ensure smooth collaboration.  

---

## 🛠️ **How to Contribute**

### **1. Fork & Clone the Repository**
1. **Fork the repository** on GitHub.
2. **Clone** your fork locally:
   ```sh
   git clone https://github.com/YOUR_USERNAME/fluora.git
   cd fluora
   ```

3. **Create a new branch** for your changes:
   ```sh
   git checkout -b feature/your-feature
   ```

---

## 📌 **Adding a New Instruction**
Fluora allows developers to create structured AI instructions dynamically. You can contribute by adding a new instruction.

### **Step 1: Create an Instruction YAML**
1. Navigate to the `instructions/` directory.
2. Create a new **`instruction_name.yml`** file inside `instructions/{instruction_name}`.
3. Use the following template:

   ```yaml
   instruction_info:
     title: "Your Instruction Title"
     info: "Short description of the instruction."
     version: "v1.0.0"

   config:
     route: "/your-endpoint"
     models:
       gemini:
         - "gemini-2.0-flash"
         - "gemini-1.5-pro"
     temperature: 0.7
     top_p: 0.9
     max_output_tokens: 256

   description: |
     Explain how the instruction works and its purpose.

   tasks:
     - Define what this instruction does.

   rules:
     - List rules that should be followed.

   input_format: |
     {
       "sentence": "Input example here.",
       "tone": "formal | informal | neutral"
     }

   output_format: |
     {
       "original_sentence": "Original input sentence.",
       "rewritten_sentence": "Updated sentence.",
       "suggestions": ["Alternative 1", "Alternative 2"]
     }

   additional_notes:
     - Any other important details.
   ```

### **Step 2: Generate Code from YAML**
1. Run the generator:
   ```sh
   make gen_instructions
   ```
2. This will generate the instruction service inside `internal/instructions/textcraft/`.

---

## 🏗 **Adding the Service and Controller**
Once the instruction is generated, you must register it in the **TextCraft service** and **controller**.

### **Step 1: Register the Instruction**
Open `internal/service/textcraft.go` and add your new instruction:
```go
services := map[string]types.Instructor{
    "rewrite":    textcraft.NewRewriteInstruction(),
    "casualize":  textcraft.NewCasualizeInstruction(),
    "expand":     textcraft.NewExpandInstruction(),
    "your_instruction": textcraft.NewYourInstruction(),
}
```

### **Step 2: Register the Route in Controller**
Modify `internal/controller/textcraft.go` and add:
```go
r.Post(t.svc.Service("your_instruction").Route(), t.YourInstructionHandler)
```

### **Step 3: Implement Handler**
Inside `internal/controller/textcraft.go`:
```go
func (t *TextCraftController) YourInstructionHandler(ctx *fiber.Ctx) error {
    r := new(dto.YourInstructionReq)
    if err := ctx.BodyParser(r); err != nil {
        return errors.New(fiber.StatusBadRequest, "Invalid request body", map[string]string{"error": err.Error()})
    }

    if err := r.Validate(); err != nil {
        return err
    }

    resp, err := t.svc.YourInstruction(ctx.Context(), r.Provider, r.Model, &dto.YourInstructionParam{
        BaseTextParam: dto.BaseTextParam{Sentence: r.Sentence, Language: r.Language},
    })
    if err != nil {
        return err
    }

    return ctx.JSON(resp)
}
```

---

## 🐞 **Fixing Issues or Adding Features**
### **Step 1: Find an Issue**
- Check the **[GitHub Issues](https://github.com/Ja7ad/fluora/issues)** for bugs or feature requests.
- If the issue is **not listed**, create a new issue and describe the bug or enhancement.

### **Step 2: Fix the Issue**
- Create a branch for your fix:
  ```sh
  git checkout -b fix/issue-name
  ```
- Make the necessary changes.

### **Step 3: Test Your Code**
- Run unit tests:
  ```sh
  go test ./...
  ```
- Check for errors and lint issues:
  ```sh
  make lint
  ```

---

## 🚀 **Submitting Your Contribution**
1. **Commit your changes**:
   ```sh
   git add .
   git commit -m "feat: Added YourInstruction functionality"
   ```
2. **Push your branch**:
   ```sh
   git push origin feature/your-feature
   ```
3. **Create a Pull Request (PR)**:
   - Go to [Fluora GitHub Repository](https://github.com/Ja7ad/fluora).
   - Click **New Pull Request**.
   - Provide a description of the change.
   - Wait for review and approval.

---

## 🎯 **Guidelines**
✅ Follow the coding style used in Fluora.  
✅ Keep PRs **small and focused** (one feature or fix per PR).  
✅ Add **comments and documentation** where necessary.  
✅ Ensure all **tests pass** before submitting a PR.  

---

## ❓ **Need Help?**
If you have any questions, open a **GitHub Discussion** or contact **@Ja7ad**.

Happy Coding! 🚀🎉
