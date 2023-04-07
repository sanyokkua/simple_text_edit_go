// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from "../models";

export function CloseFile(arg1: number): Promise<void>;

export function GetFileTypeExtension(arg1: string): Promise<types.FrontendKeyValueArrayContainerStruct>;

export function GetFileTypes(): Promise<types.FrontendKeyValueArrayContainerStruct>;

export function GetFilesShortInfo(): Promise<types.FrontendFileInfoArrayContainerStruct>;

export function GetOpenedFile(): Promise<types.FrontendFileContainerStruct>;

export function NewFile(): Promise<void>;

export function SwitchOpenedFileTo(arg1: number): Promise<void>;

export function UpdateFileContent(arg1: number, arg2: string): Promise<void>;
