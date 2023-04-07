export function isStringEmpty(value: string | null | undefined) {
    return value === null || value === undefined || value.length === 0;
}

export function isFileNameValid(fileName: string | null | undefined): boolean {
    if (fileName === undefined || fileName === null) {
        return true;
    }
    if (isStringEmpty(fileName)) {
        return true;
    }
    if (fileName.trim().length === 0 && fileName.length !== fileName.trim().length) {
        return false;
    }

    const sizeLimit: number = 255;
    if (fileName.length > sizeLimit) {
        return false;
    }

    const regExp: RegExp = /^((\d|\w|-|)+(\s)?)*$/;
    return regExp.test(fileName);
}