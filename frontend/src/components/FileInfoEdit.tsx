import React from "react";
import {FileInfoUpdateStruct, FileStruct} from "../types/backend";
import Info from "./file_info/Info";
import Edit from "./file_info/Edit";
import {Button, Modal} from "semantic-ui-react";
import {UpdateFileInformation} from "../../wailsjs/go/frontendapi/FrontendApiStruct";

type FileInfoEditProps = {
    currentFile: FileStruct | null;
    open: boolean;
    onClose: () => void;
};

class FileInfoEdit extends React.Component<FileInfoEditProps, any> {

    onFileInfoUpdate(data: FileInfoUpdateStruct) {
        if (this.props.currentFile?.Id) {
            UpdateFileInformation(this.props.currentFile.Id, data).catch((e) => console.error(e));
        } else {
            console.warn("File Info is not updated");
        }
    }

    render() {
        return (
            <div>
                <Modal size="fullscreen" open={this.props.open}>
                    <Modal.Header>Current File Information ({this.props.currentFile?.Name})</Modal.Header>
                    <Modal.Content>
                        <Info currentFile={this.props.currentFile}/>
                        <br/>
                        <Edit currentFile={this.props.currentFile} onSubmit={(data) => this.onFileInfoUpdate(data)}/>
                    </Modal.Content>
                    <Modal.Actions>
                        <Button positive onClick={() => this.props.onClose()}>Close</Button>
                    </Modal.Actions>
                </Modal>
            </div>
        );
    }
}

export default FileInfoEdit;