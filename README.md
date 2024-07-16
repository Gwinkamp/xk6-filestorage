# xk6-filestorage

Extension for [k6](https://k6.io). Allows to manage the file system during the execution of k6 tests.

## Requirements

* [Golang 1.22.4](https://go.dev/)
* [xk6](https://k6.io/blog/extending-k6-with-xk6/)

```shell
go install go.k6.io/xk6/cmd/xk6@latest
```

## Build

From local repository:

```shell
xk6 build --with xk6-filestorage=.
```

From remote repository:

```shell
xk6 build --with github.com/Gwinkamp/xk6-filestorage
```

## Usage

In load testing scenarios:

```javascript
import { sleep } from "k6";
import fs from 'k6/x/filestorage';

export const options = {
  target: 1,
  duration: '10s',
};

export default function () {
  const storage = new fs.FileStorage("./testdata");

  const test0txt = storage.readFile("test-0.txt");
  const test00txt = storage.readFile("subdir/test-00.txt");

  console.log(`list files: ${storage.listFiles()}`);
  console.log(`is test-3.bin file exists: ${storage.hasFile("test-3.bin")}`);
  console.log(`is subdir/test-01.xml file exists: ${storage.hasFile("subdir/test-01.xml")}`);
  console.log(`is test-4.doc file exists: ${storage.hasFile("test-4.doc")}`);
  console.log(`test-0.txt file name: ${test0txt.name}, path: ${test0txt.path}, content: ${test0txt.content}`);
  console.log(`subdir/test-0.txt file name: ${test00txt.name}, path: ${test00txt.path}, content: ${test00txt.content}`);
  console.log(`random file: ${storage.readRandFile().name}`);
  console.log(`random xml file name: ${storage.readRandFileWithExt(".xml").name}`);

  sleep(1)
}
```

To run this script, you need to run the k6 executable file, which was previously built with the `xk6 build` command

```shell
./k6 run scripts/example.js
```
