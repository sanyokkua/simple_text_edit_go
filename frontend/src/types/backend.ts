export type FileStruct = {
    Id: number;
    Path: string;
    Name: string;
    Type: string;
    Extension: string;
    InitialContent: string;
    ActualContent: string;
    New: boolean;
    Opened: boolean;
    Changed: boolean;
};

export type FileInfoStruct = {
    Id: number;
    Path: string;
    Name: string;
    Type: string;
    Extension: string;
    New: boolean;
    Opened: boolean;
    Changed: boolean;
};

export type FileTypesJson = {
    Key: string;
    Name: string;
    Extensions: string[];
};

export type KeyValuePair = {
    Key: string;
    Value: string;
};
