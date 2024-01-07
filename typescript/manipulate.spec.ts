import { manipulate } from "./manipulate";

// Write your test here
function testManipulate() {
    const testCases = [
        { input: 'a', expectedOutput: ['a'] },
        { input: 'ab', expectedOutput: ['ab', 'ba'] },
        { input: 'abc', expectedOutput: ['abc', 'acb', 'bac', 'bca', 'cab', 'cba'] },
        { input: 'aabb', expectedOutput: ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa'] },
    ];

    testCases.forEach(testCase => {
        const output = manipulate(testCase.input);
        const isSuccess = arraysEqual(output.sort(), testCase.expectedOutput.sort());
        console.log(`With input '${testCase.input}':\nYour function should return ${JSON.stringify(testCase.expectedOutput)}, got ${output}. ${isSuccess ? "Success" : "Failure"}`);
    });

    function arraysEqual(a: string[], b: string[]): boolean {
        if (a.length !== b.length) return false;
        for (let i = 0; i < a.length; i++) {
            if (a[i] !== b[i]) return false;
        }
        return true;
    }
}

testManipulate();