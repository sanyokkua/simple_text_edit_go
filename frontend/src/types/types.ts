export type FileInformation = {
    OpenTimeStamp: number; // OpenTimeStamp - Used to place tab in right order
    Path: string; // Path - Full path to the file (empty for new file)
    Name: string; // Name of the file (last item from the path without extension)
    Ext: string; // Ext - Extension of the file (last item after . in path if available) (empty for new file)
    Type: string; // Type based on the extension. Can be the same or different
    Exists: boolean; // Exists equal true if this file was opened and not just created
    IsOpenedNow: boolean; // true if the currently opened tab points to the current file
    IsChanged: boolean; //IsChanged equal true if this the actual content is not equal to original content
}

export type FileInMemory = {
    Descriptor: FileInformation; // Contains base file information
    OriginalContent: string; // Content that was read after open
    ActualContent: string; // Content that can be changed during the time
}

export type FilesInMemoryMap = {
    [key: number]: FileInMemory;
}

export type EditorApplication = {
    AppContext: any; // Application context
    FilesInMemory: FilesInMemoryMap; // Files that currently opened in the app memory
}
