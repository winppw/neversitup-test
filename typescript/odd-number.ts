/**
 * 
 * @param numbers list of numbers example [7], [0], [1,1,2], [0,1,0,0,1], [1,2,3,4,5,6,7,8,9,0]
 * @returns one number that odd number
 */ 
export function findOddNumber(numbers: number[]) : number {
    // TODO : Start your code here
    const numCounts: { [key: number]: number } = {};
    for (const num of numbers) {
        if (numCounts[num]) {
            numCounts[num]++;
        } else {
            numCounts[num] = 1;
        }
    }

    for (const num in numCounts) {
        if (numCounts[num] % 2 !== 0) {
            return parseInt(num);
        }
    }

    throw new Error("No number appears an odd number of times");
    // throw new Error("Not implemented");
}