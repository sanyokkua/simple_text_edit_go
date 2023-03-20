import React from "react"

import {EventsOn} from "../wailsjs/runtime";
import CodeMirror from '@uiw/react-codemirror';
import {LanguageName, loadLanguage} from '@uiw/codemirror-extensions-langs';
import {
    ChangeFileContent,
    ChangeFileStatusToOpened,
    FindOpenedFile,
    GetFilesInformation
} from "../wailsjs/go/jsapi/JsStruct"
import {FileStruct, InformationStruct} from "./types/types";
import {
    EventOnErrorHappened,
    EventOnFileClosed,
    EventOnFileOpened,
    EventOnFileSaved,
    EventOnNewFileCreate
} from "./types/constants";
import {Button, Menu, Modal} from "semantic-ui-react";
import {SemanticCOLORS} from "semantic-ui-react/dist/commonjs/generic";

type AppState = {
    files: InformationStruct[];
    currentFile: FileStruct;
    currentLanguage: any | null;
    errorModal: boolean;
    errorText: string;
};

const COLOR_NEW: SemanticCOLORS = "red"
const COLOR_NO_CHANGES: SemanticCOLORS = "black"
const COLOR_HAS_CHANGES: SemanticCOLORS = "blue"

class App extends React.Component<any, AppState> {
    constructor(props: any) {
        super(props);
        this.state = {
            files: [],
            currentFile: {} as FileStruct,
            currentLanguage: null,
            errorModal: false,
            errorText: "",
        };
        EventsOn(EventOnNewFileCreate, () => {
            console.log("EventOnNewFileCreate")
            this.updateState().catch((e) => this.onErrorProcessing(e))
        });
        EventsOn(EventOnFileOpened, () => {
            console.log("EventOnFileOpened")
            this.updateState().catch((e) => this.onErrorProcessing(e))
        });
        EventsOn(EventOnFileSaved, (file) => {
            console.log("EventOnFileSaved", file)
            this.updateState().catch((e) => this.onErrorProcessing(e))
        });
        EventsOn(EventOnFileClosed, (file) => {
            console.log("EventOnFileClosed", file)
            this.updateState().catch((e) => this.onErrorProcessing(e))
        });
        EventsOn(EventOnErrorHappened, (error) => {
            console.log("EventOnErrorHappened")
            this.onErrorProcessing(error)
        });
    }

    componentDidMount() {
        this.updateState().catch((e) => this.onErrorProcessing(e))
    }

    onErrorProcessing(error: any) {
        let msg = error?.message ? error?.message : JSON.stringify(error)
        console.error(msg)
        this.setState({
            errorModal: true,
            errorText: msg
        })
    }

    async updateState() {
        try {
            const files: InformationStruct[] = await GetFilesInformation()
            console.log(files)
            // Sort files by time of open/creation (internal ID of each file)
            files.sort((a, b) => a.OpenTimeStamp - b.OpenTimeStamp)

            const currentFile: FileStruct = await FindOpenedFile()
            console.log(currentFile)
            const currentFileLang = loadLanguage(currentFile.FileInfo.FileType as LanguageName)

            this.setState({
                files: files,
                currentFile: currentFile,
                currentLanguage: currentFileLang,
            });
        } catch (e) {
            this.onErrorProcessing(e)
        }
    }

    tabIsChanged(fileId: number) {
        ChangeFileStatusToOpened(fileId)
            .then(() => this.updateState())
            .catch((e) => this.onErrorProcessing(e))
    }

    contentIsChanged(text: string) {
        // @ts-ignore
        let openTimeStamp: number = this.state.currentFile?.FileInfo.OpenTimeStamp;
        ChangeFileContent(openTimeStamp, text)
            .then((hasChanges: boolean) => this.state.currentFile?.FileInfo.FileHasChanges !== hasChanges)
            .then((needsUpdate: boolean) => {
                    if (needsUpdate) {
                        return this.updateState().then()
                    }
                },
            ).catch((e) => this.onErrorProcessing(e));
    }

    render() {
        const extensions: any[] = []
        if (this.state.currentLanguage) {
            extensions.push(this.state.currentLanguage)
        }

        const menuItems = this.state.files.map(openedFile => {
            const key: string = openedFile.OpenTimeStamp.toString();
            const fileExist: boolean = openedFile.FileExists;
            const fileName: string = fileExist ? openedFile.FileName : "*New";
            const isActive: boolean = openedFile.FileIsOpened;
            const hasChanges: boolean = openedFile.FileHasChanges
            const color: SemanticCOLORS = fileExist ? hasChanges ? COLOR_HAS_CHANGES : COLOR_NO_CHANGES : COLOR_NEW;

            return <Menu.Item key={key} active={isActive} color={color}
                              onClick={() => this.tabIsChanged(openedFile.OpenTimeStamp)}>
                {fileName}
            </Menu.Item>
        });
        const fileContent = this.state.currentFile?.ActualContent;

        return (
            <div>
                <Menu tabular>{menuItems}</Menu>
                <CodeMirror
                    value={fileContent}
                    height="100vh"
                    onChange={(text) => this.contentIsChanged(text)}
                    basicSetup={{
                        foldGutter: true,
                        allowMultipleSelections: true,
                        indentOnInput: true,
                        tabSize: 4,
                        highlightActiveLine: true,
                        highlightActiveLineGutter: true,
                        highlightSelectionMatches: true,
                        syntaxHighlighting: true,
                        bracketMatching: true
                    }}
                    extensions={extensions}
                />
                <Modal dimmer={"blurring"} open={this.state.errorModal} size={'mini'}
                       onClose={() => this.setState({errorModal: false})}
                >
                    <Modal.Header>Error happened </Modal.Header>
                    <Modal.Content>
                        {this.state.errorText}
                    </Modal.Content>
                    <Modal.Actions>
                        <Button negative onClick={() => this.setState({errorModal: false})}>Ok</Button>
                    </Modal.Actions>
                </Modal>
            </div>
        )
    }
}

export default App