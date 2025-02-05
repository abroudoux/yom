# yom

🥕 Who owes how much to whom?

## Usage

### Taskfile

Run `task run path/to/your/grocery/ticket` to start using `yom`.

### Manual

You can create the binary:

```bash
go build -o bin/yom cmd/main.go
```

Then run it by passing your grocery ticket's path

```bash
./bin/yom path/to/your/grocery/ticket
```

## Roadmap

- [x] Improve parser to handle names and prices correctly
- [x] Handle items shared by multiple users
- [ ] Parse quantity
- [ ] Handle cases when there's more than 3 people

## License

This project is under [MIT License](LICENSE).
