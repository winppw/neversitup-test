import { countSmilyFace } from "./smily";

// Write your test here
function testCountSmilyFace() {
    const testCases = [
        { input: [':)', ';(', ';}', ':-D'], expectedOutput: 2 },
        { input: [';D', ':-(', ':-)', ';~)'], expectedOutput: 3 },
        { input: [';]', ':[', ';*', ':$', ';-D'], expectedOutput: 1 },
    ];

    testCases.forEach(testCase => {
        const output = countSmilyFace(testCase.input);
        console.log(`With input [${testCase.input.join(", ")}]: Your function should return ${testCase.expectedOutput}, got ${output}. ${output === testCase.expectedOutput ? "Success" : "Failure"}`);
    });
}

testCountSmilyFace();