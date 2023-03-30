import React from "react";
import {FileInfoUpdateStruct, FileStruct, KeyValuePairStruct} from "../../types/backend";
import {Button, DropdownItemProps, DropdownProps, Form} from "semantic-ui-react";
import {GetFileTypeExtension, GetFileTypes} from "../../../wailsjs/go/frontend/frontStruct";
import {isStringEmpty} from "../../utils/string_utils";

type EditProps = {
    currentFile: FileStruct | null;
    onSubmit: (data: FileInfoUpdateStruct) => void;
};

type EditState = {
    fileType: string | null | undefined;
    fileExtension: string | null | undefined;
    dropdownFileTypes: DropdownItemProps[];
    dropdownFileExtensions: DropdownItemProps[];
    isFormValid: boolean;
};

function mapKeyValueToDropdownItem(keyValuePairTypesList: KeyValuePairStruct[]): DropdownItemProps[] {
    return keyValuePairTypesList.map(fileInfo => {
        return {
            key: fileInfo.Key,
            value: fileInfo.Key,
            text: fileInfo.Value,
        };
    });
}

class Edit extends React.Component<EditProps, EditState> {
    constructor(props: EditProps) {
        super(props);
        this.state = {
            fileType: null,
            fileExtension: null,
            dropdownFileTypes: [],
            dropdownFileExtensions: [],
            isFormValid: true,
        };
    }

    componentDidMount() {
        this.updateState();
    }

    updateState() {
        GetFileTypes()
            .then(fileTypes => mapKeyValueToDropdownItem(fileTypes))
            .then(fileTypesDropdown => this.setState({dropdownFileTypes: fileTypesDropdown}))
            .then(() => GetFileTypeExtension(this.state.fileType || ""))
            .then(fileExtensions => mapKeyValueToDropdownItem(fileExtensions))
            .then(fileExtensionsDropdown => this.setState({dropdownFileExtensions: fileExtensionsDropdown}, () => {
                if (this.state.dropdownFileExtensions && this.state.dropdownFileExtensions.length == 1) {
                    this.setState({fileExtension: this.state.dropdownFileExtensions[0].value as string},
                        this.validateForm);
                }
            }))
            .then(() => this.validateForm());
    }

    onTypeChanged(type: DropdownProps) {
        this.setState({
            fileType: type.value as string,
            fileExtension: null,
        }, () => {
            this.updateState();
        });
    }

    onExtensionChanged(extension: DropdownProps) {
        this.setState({
            fileExtension: extension.value as string,
        }, () => {
            this.validateForm();
        });
    }

    onSubmit() {
        if (this.state.isFormValid) {
            this.props.onSubmit({
                Id: this.props.currentFile?.Id,
                Type: this.state.fileType,
                Extension: this.state.fileExtension,
            } as FileInfoUpdateStruct);
        }
    }

    validateForm() {
        const fileType = this.state.fileType;
        const fileExtension = this.state.fileExtension;
        const isTypeEmpty = isStringEmpty(fileType);
        const isExtensionEmpty = isStringEmpty(fileExtension);

        if (isTypeEmpty && isExtensionEmpty) {
            this.setState({isFormValid: true});
            return;
        }

        const isWholeFormValid = !isTypeEmpty && !isExtensionEmpty;
        this.setState({isFormValid: isWholeFormValid});
    }

    render() {
        const dropdownFileExtensions = this.state.dropdownFileExtensions;
        const selectedType = this.state.fileType || "";
        const showExtForm: boolean = selectedType.length > 0 && dropdownFileExtensions.length > 1;

        return (
            <Form>
                <Form.Field>
                    <Form.Select
                        label="File Type"
                        placeholder={this.state.fileType || undefined}
                        options={this.state.dropdownFileTypes}
                        onChange={(event, data) => this.onTypeChanged(data)}
                    />
                </Form.Field>
                <Form.Field>
                    {showExtForm &&
                        <Form.Select
                            label="File Extension"
                            placeholder={this.state.fileExtension || undefined}
                            options={dropdownFileExtensions}
                            disabled={dropdownFileExtensions.length === 0}
                            onChange={(event, data) => this.onExtensionChanged(data)}
                        />}
                </Form.Field>
                <Button disabled={!this.state.isFormValid} type="submit" onClick={() => this.onSubmit()}>Submit</Button>
            </Form>
        );
    }
}

export default Edit;