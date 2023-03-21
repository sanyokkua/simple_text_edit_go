import React from "react"
import {Button, DropdownItemProps, DropdownProps, Form, Modal} from "semantic-ui-react";
import {FileTypeInformation, InformationStruct} from "../types/types";
import {GetFileTypeInformation} from "../../wailsjs/go/jsapi/JsStruct";

export type DialogResult = {
    FileName: string | null | undefined;
    FileType: string | null | undefined;
    FileExt: string | null | undefined;
};
type DialogProps = {
    showDialog: boolean;
    fileInfo: InformationStruct | null;
    onAcceptBtnClicked: (res: DialogResult) => void;
    onCancelBtnClicked: () => void;
};
type DialogState = {
    fileTypeInfo: FileTypeInformation[];
    selectedName: string | null | undefined;
    selectedType: string | null | undefined;
    selectedExtension: string | null | undefined;
    isFormValid: boolean;
    isNameValid: boolean;
    dropdownFileTypes: DropdownItemProps[];
    dropdownFileExtensions: DropdownItemProps[];
};

function mapInfoToTypeList(data: FileTypeInformation[]): DropdownItemProps[] {
    return data.map(fileInfo => {
        return {
            key: fileInfo.Key,
            value: fileInfo.Key,
            text: fileInfo.Name
        }
    });
}

function mapInfoToExtensionList(selectedType: string | null | undefined, data: FileTypeInformation[]): DropdownItemProps[] {
    if (selectedType === null || selectedType === undefined || selectedType.length == 0) {
        return []
    }

    const selectedTypeKey: FileTypeInformation | null | undefined = data.find((value) => value.Key === selectedType);
    if (selectedTypeKey === null || selectedTypeKey === undefined) {
        return []
    }

    return selectedTypeKey.Extensions.map(value => {
        return {
            key: value,
            value: value,
            text: value
        }
    });
}

function isEmptyString(value: string | null | undefined) {
    return value === null || value === undefined || value.length === 0;
}

function isFileNameValid(fileName: string | null | undefined): boolean {
    if (fileName === undefined || fileName === null) {
        return true
    }
    if (isEmptyString(fileName)) {
        return true
    }
    if (fileName.trim().length === 0 && fileName.length !== fileName.trim().length) {
        return false
    }

    const sizeLimit: number = 255;
    if (fileName.length > sizeLimit) {
        return false;
    }

    const regExp: RegExp = /^((\d|\w|-|)+(\s)?)*$/;
    return regExp.test(fileName);
}

export class InfoChangeDialog extends React.Component<DialogProps, DialogState> {
    constructor(props: DialogProps) {
        super(props);
        this.state = {
            fileTypeInfo: [],
            selectedName: null,
            selectedType: null,
            selectedExtension: null,
            isFormValid: true,
            isNameValid: true,
            dropdownFileTypes: [],
            dropdownFileExtensions: [],
        }
    }

    componentDidMount() {
        GetFileTypeInformation()
            .then((data: FileTypeInformation[]) => {
                this.setState({fileTypeInfo: data}, () => this.updateTypesAndExtensions());
            }).catch((e) => console.log(e));
    }

    validateForm() {
        const fileName = this.state.selectedName;
        const fileType = this.state.selectedType;
        const fileExtension = this.state.selectedExtension;

        const isFileEmpty = isEmptyString(fileName);
        const isTypeEmpty = isEmptyString(fileType);
        const isExtensionEmpty = isEmptyString(fileExtension);

        if (isFileEmpty && isTypeEmpty && isExtensionEmpty) {
            this.setState({isFormValid: true, isNameValid: true});
            return;
        }

        const isValidName = isFileNameValid(fileName);
        const isTypeAndExtensionSelected = !isTypeEmpty && !isExtensionEmpty;
        const isWholeFormValid = isValidName && isTypeAndExtensionSelected;

        if (isWholeFormValid) {
            this.setState({isFormValid: true, isNameValid: true});
            return;
        }

        this.setState({isFormValid: isWholeFormValid, isNameValid: isValidName});
    }

    updateTypesAndExtensions() {
        const typeInformation = this.state.fileTypeInfo;
        const selectedType = this.state.selectedType;

        const fileTypes: DropdownItemProps[] = mapInfoToTypeList(typeInformation);
        const fileExtensions: DropdownItemProps[] = mapInfoToExtensionList(selectedType, typeInformation);

        if (fileExtensions.length == 1) {
            this.setState({
                dropdownFileTypes: fileTypes,
                dropdownFileExtensions: fileExtensions,
                selectedExtension: fileExtensions[0].value as string
            }, () => this.validateForm());
            return;
        }

        this.setState({
            dropdownFileTypes: fileTypes,
            dropdownFileExtensions: fileExtensions,
        }, () => this.validateForm());
    }

    onFileNameChanged(name: string) {
        this.setState({
            selectedName: name
        }, () => this.validateForm());
    }

    onTypeChanged(type: DropdownProps) {
        this.setState({
            selectedType: type.value as string,
            selectedExtension: null,
        }, () => {
            this.updateTypesAndExtensions();
        });
    }

    onExtensionChanged(extension: DropdownProps) {
        this.setState({
            selectedExtension: extension.value as string,
        }, () => {
            this.validateForm();
        });
    }

    onReturnResults() {
        const res: DialogResult = {
            FileName: this.state.selectedName,
            FileType: this.state.selectedType,
            FileExt: this.state.selectedExtension,
        }
        this.props.onAcceptBtnClicked(res)
    }

    render() {
        const fileName: string = this.props.fileInfo?.FileName || "";
        const fileType: string = this.props.fileInfo?.FileType || "";
        const fileExt: string = this.props.fileInfo?.FileExtension || "";
        const dropdownFileExtensions = this.state.dropdownFileExtensions;
        const selectedType = this.state.selectedType || "";
        const showExtForm: boolean = selectedType.length > 0 && dropdownFileExtensions.length > 1;

        return (
            <Modal dimmer={"blurring"} open={this.props.showDialog} size={'large'}>
                <Modal.Header>Current File Information</Modal.Header>
                <Modal.Content>
                    <Form>
                        <p>Current name: <b>{fileName}</b></p>
                        <p>For new file - this name will be suggested during saving</p>
                        <p>For existing file - saving will require manual save with name that was typed</p>
                        <Form.Group>
                            <Form.Input label='Change Name'
                                        type='text'
                                        error={!this.state.isNameValid ? "File name is not valid" : undefined}
                                        onChange={(event, data) => this.onFileNameChanged(data.value)}/>
                        </Form.Group>

                        <p>Current type: <b>{fileType}</b></p>
                        <p>Change of type will change syntax highlight language</p>
                        <p>Change of extension only possible after selecting type</p>
                        <Form.Group>
                            <Form.Select label={"Change File Type"}
                                         options={this.state.dropdownFileTypes}
                                         onChange={(event, data) => this.onTypeChanged(data)}
                            />
                        </Form.Group>

                        {showExtForm && <div>
                            <p>Current extension is: <b>{fileExt}</b></p>
                            <Form.Group>
                                <Form.Select label={"Change File Type"}
                                             disabled={dropdownFileExtensions.length === 0}
                                             options={dropdownFileExtensions}
                                             placeholder={this.state.selectedExtension || undefined}
                                             onChange={(event, data) => this.onExtensionChanged(data)}
                                />
                            </Form.Group>
                        </div>}
                    </Form>
                </Modal.Content>
                <Modal.Actions>
                    <Button disabled={!this.state.isFormValid} positive
                            onClick={() => this.onReturnResults()}>Ok</Button>
                    <Button negative
                            onClick={() => this.props.onCancelBtnClicked()}>Cancel</Button>
                </Modal.Actions>
            </Modal>
        )
    }
}
