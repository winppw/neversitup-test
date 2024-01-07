import { findOddNumber } from "./odd-number";

// Write your test here
function testFindOddNumber() {
    const testCases = [
        { input: [7], expectedOutput: 7 },
        { input: [0], expectedOutput: 0 },
        { input: [1, 1, 2], expectedOutput: 2 },
        { input: [0, 1, 0, 1, 0], expectedOutput: 0 },
        { input: [1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1], expectedOutput: 4 },
    ];

    testCases.forEach(testCase => {
        const output = findOddNumber(testCase.input);
        const count = testCase.input.filter(x => x === output).length;
        console.log(`[${testCase.input}] should return ${output}, because it occurs ${count} times (which is odd).`);
    });
}

testFindOddNumber();