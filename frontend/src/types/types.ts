export type InformationStruct = {
    OpenTimeStamp: number  // openTimeStamp - Used to place tab in right order
    FilePath: string // filePath - Full path to the file (empty for new file)
    FileName: string // fileName of the file (last item from the path without extension)
    FileExtension: string // fileExtension - Extension of the file (last item after . in path if available) (empty for new file)
    FileType: string // fileType based on the extension. Can be the same or different
    FileExists: boolean   // fileExists equal true if this file was opened and not just created
    FileOpened: boolean   // fileOpened equal true if this file should be shown on the UI now
    HasChanges: boolean   // hasChanges equal true if this the actual content is not equal to original content
}
export type FileStruct = {
    FileInfo: InformationStruct // fileInfo - contains base file information
    OriginalContent: string             // originalContent - Content that was read after open
    ActualContent: string             // actualContent - Content that can be changed during the time
}