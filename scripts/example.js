import { sleep } from "k6";
import fs from 'k6/x/filestorage';

export const options = {
  scenarios: {
    main: {
      executor: "constant-arrival-rate",
      rate: 1,
      preAllocatedVUs: 10,
      maxVUs: 100,
      duration: '10s',
    }
  }
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

