import React from "react";

import {EventsOn} from "../wailsjs/runtime";
import {GetFilesShortInfo, GetOpenedFile} from "../wailsjs/go/frontend/frontStruct";
import {FileInfoStruct, FileStruct} from "./types/backend";
import {
    EventOnFileClosed,
    EventOnFileContentIsUpdated,
    EventOnFileInformationDisplay,
    EventOnFileInformationEdit,
    EventOnFileInformationUpdated,
    EventOnFileIsSwitched,
    EventOnFileOpened,
    EventOnFileSaved,
    EventOnInternalError,
    EventOnInternalWarning,
    EventOnNewFileCreated,
} from "./types/constants";
import CodeEditor from "./components/CodeEditor";
import TabBar from "./components/TabBar";
import {NotificationType} from "./types/frontend";
import AppModalInfoDialog from "./components/AppModalInfoDialog";
import {Message} from "semantic-ui-react";

type AppState = {
    openedFiles: FileInfoStruct[];
    currentFile: FileStruct | null;
    notificationType: NotificationType;
    notificationMessage: string;
    showFileInfoDisplayModal: boolean;
    showFileInfoEditModal: boolean;
};

class App extends React.Component<any, AppState> {
    constructor(props: any) {
        super(props);
        this.state = {
            openedFiles: [],
            currentFile: null,
            notificationType: NotificationType.NONE,
            notificationMessage: "",
            showFileInfoDisplayModal: false,
            showFileInfoEditModal: false,
        };

        EventsOn(EventOnInternalWarning, (eventData) => this.onEventOnInternalWarning(eventData));
        EventsOn(EventOnInternalError, (eventData) => this.onEventOnInternalError(eventData));

        EventsOn(EventOnNewFileCreated, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileOpened, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileSaved, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileClosed, (eventData) => this.updateEditorState());

        EventsOn(EventOnFileInformationEdit, (eventData) => this.setState({showFileInfoEditModal: true}));
        EventsOn(EventOnFileInformationDisplay, (eventData) => this.setState({showFileInfoDisplayModal: true}));
        EventsOn(EventOnFileInformationUpdated, (eventData) => this.updateEditorState());

        EventsOn(EventOnFileIsSwitched, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileContentIsUpdated, (eventData) => this.updateEditorState());
    }

    onEventOnInternalWarning(event: string) {
        this.setState({
            notificationType: NotificationType.WARNING,
            notificationMessage: event,
        });
    }

    onEventOnInternalError(event: string) {
        this.setState({
            notificationType: NotificationType.ERROR,
            notificationMessage: event,
        });
    }

    componentDidMount() {
        this.updateEditorState();
    }

    onError(error: Error) {
        const msg: string = error?.message ? error?.message : JSON.stringify(error);
        this.setState({
            notificationType: NotificationType.ERROR,
            notificationMessage: msg,
        });
    }

    updateEditorState() {
        GetFilesShortInfo()
            .then((files: FileInfoStruct[]) => this.setState({openedFiles: files}))
            .then(GetOpenedFile)
            .then((currentFile: FileStruct) => this.setState({currentFile: currentFile}))
            .catch((e) => this.onError(e));
    }

    render() {
        const showInfoDialog: boolean = this.state.notificationType !== NotificationType.NONE;
        const showEditor: boolean = this.state.openedFiles.length > 0;

        const editorContent = <div>
            <TabBar files={this.state.openedFiles} onError={(error: Error) => this.onError(error)}/>
            <CodeEditor file={this.state.currentFile}/>
        </div>;

        const noContent = <Message color="yellow">No Opened Files</Message>;
        return (
            <div>
                {showEditor ? editorContent : noContent}
                <AppModalInfoDialog header={this.state.notificationType}
                                    message={this.state.notificationMessage}
                                    show={showInfoDialog}
                                    onClose={() => this.setState({
                                        notificationType: NotificationType.NONE,
                                        notificationMessage: "",
                                    })}/>
            </div>
        );
    }
}

export default App;