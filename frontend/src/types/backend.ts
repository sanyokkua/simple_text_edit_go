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

export type FileInfoUpdateStruct = {
    Id: number;
    Type: string;
    Extension: string;
}

export type KeyValuePairStruct = {
    Key: string;
    Value: string;
};
