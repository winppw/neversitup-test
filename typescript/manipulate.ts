/**
 * 
 * @param text string of value example "aabb", "abcde"
 * @returns {string[]} string array of result
 */ 
export function manipulate(text: string) : string[] {
    // TODO : Start your code here
    if (text.length <= 1) {
        return [text];
    }

    const chars = text.split('');
    const result = new Set<string>(); // Using a set to avoid duplicates

    chars.forEach((char, idx) => {
        const remainingChars = [...chars.slice(0, idx), ...chars.slice(idx + 1)].join('');
        const permutations = manipulate(remainingChars);
        permutations.forEach(permutation => {
            result.add(char + permutation);
        });
    });

    return Array.from(result); // Convert the set back to an array
    // throw new Error("Not implemented");
}