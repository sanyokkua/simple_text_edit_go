// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {api} from '../models';

export function ChangeFileContent(arg1: number, arg2: string): Promise<boolean>;

export function ChangeFileStatusToOpened(arg1: number): Promise<void>;

export function FindOpenedFile(): Promise<api.OpenedFile>;

export function GetFilesInformation(): Promise<Array<api.FileInformation>>;
