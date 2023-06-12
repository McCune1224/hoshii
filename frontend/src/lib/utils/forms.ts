/**
Finds any missing items in a form data object and returns an array of errors messages for each missing item
*/
export function ValidateFormData(data: FormData, keys: string[]) {
    const missingItems: Array<string> = [];
    keys.forEach((key) => {
        if (!data.get(key)) {
            missingItems.push(`${key} is required`);
        }
    });
    return missingItems;
}
