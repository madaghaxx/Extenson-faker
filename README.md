# 🛠️ Extension Faker

This program hides file extensions by using a code outside UTF-8, allowing the OS to determine the correct order of letters (Left-to-Right or Right-to-Left).

## 🚀 Usage

To run the program, use the following command:

```sh
go run main.go <your_file>
```

## ✨ Customization

You can edit the output by changing the default extension (".txt") in the code:

```go
rtlOverride := "\u202Etxt"
```

## 🤝 Contributions

If you have any suggestions or improvements, feel free to share them. We welcome all ideas!

![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)
