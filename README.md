# neversitup-test

## Setup (Typescript)

1. Install node and ts-node <br>

<pre>
npm install -g
npm install ts-node -g
</pre>

2. Run testing

<pre>
cd typesript
ts-node manipulate.spec.ts
</pre>

## Setup (Golang)

1. Install golang <br>

<pre>
go mod init
</pre>

2. Run testing

<pre>
cd golang # And run into file
</pre>

## Project Structure

- `/golang`: All file for interview test in golang.
- `/typescript`: All file for interview test in typescript.

<pre>
neversitup-test
 ┣ golang
 ┃ ┣ go.mod
 ┃ ┣ manipulate.go
 ┃ ┣ manipulate_test.go
 ┃ ┣ odd_number.go
 ┃ ┣ odd_number_test.go
 ┃ ┣ smily.go
 ┃ ┗ smily_test.go
 ┣ typescript
 ┃ ┣ manipulate.spec.ts
 ┃ ┣ manipulate.ts
 ┃ ┣ odd-number.spec.ts
 ┃ ┣ odd-number.ts
 ┃ ┣ smily.spec.ts
 ┃ ┗ smily.ts
 ┗ README.md
</pre>
