This is a file extension hiding program that uses a code outside UTF-8 which makes the os can know the correct order of lettres (Left-to-right of Right-to-Left).
usage:
```
go run main.go <your_file>
```
you can edit the output by changing it in here (the default is ".txt"):
```
rtlOverride := "\u202Etxt"
```
If you have any better ideas I am all ears.
