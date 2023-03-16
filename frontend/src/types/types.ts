export type FileDescriptor = {
    FileId: number;
    FilePath: string;
    FileName: string;
    FileType: string;
    IsActive: boolean;
}
export type IntToStringMap = {
    [key: number]: string
}

export type AppFile = {
    Descriptor: FileDescriptor;
    FileContent: string;
    SelectedContent: string;
    ContentHistory: IntToStringMap;
}

export type AppFilesMap = {
    [key: number]: AppFile;
}
export type EditorApplication = {
    AppContext: any;
    AppFiles: AppFilesMap;
}