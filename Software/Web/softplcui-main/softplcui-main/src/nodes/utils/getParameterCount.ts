export const getParameterElementUsingNumber = (arr: string[]): number => {
    let count = 0;
    let countEmptyBetween = 0;

    arr.forEach((str) => {
        if (!str) {
            countEmptyBetween += 1;
        } else {
            count += 1 + countEmptyBetween;
            countEmptyBetween = 0;
        }
    });

    return count;
};
