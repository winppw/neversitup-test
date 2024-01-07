/**
 * 
 * @param texts list of string [":)", ":(", ":>"]
 * @returns amount of smily face detected
 */
export function countSmilyFace(texts: string[]) : number {
    // TODO : Start your code here
    const validSmileyRegex = /^[:;][-~]?[)D]$/;
    return texts.filter(face => validSmileyRegex.test(face)).length;
    // throw new Error("Not implemented");
}