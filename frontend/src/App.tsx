import React from "react";

import {EventsOn} from "../wailsjs/runtime";
import {GetFilesShortInfo, GetOpenedFile, NewFile} from "../wailsjs/go/frontendapi/FrontendApiStruct";
import {
    FileInfoStruct,
    FileStruct,
    FrontendFileContainerStruct,
    FrontendFileInfoArrayContainerStruct,
} from "./types/backend";
import {
    EventOnBlockUiFalse,
    EventOnBlockUiTrue,
    EventOnFileClosed,
    EventOnFileContentIsUpdated,
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
import {Dimmer, Message} from "semantic-ui-react";

type AppState = {
    openedFiles: FileInfoStruct[];
    currentFile: FileStruct | null;
    notificationType: NotificationType;
    notificationMessage: string;
    blockUi: boolean;
};

class App extends React.Component<any, AppState> {
    constructor(props: any) {
        super(props);
        this.state = {
            openedFiles: [],
            currentFile: null,
            notificationType: NotificationType.NONE,
            notificationMessage: "",
            blockUi: false,
        };

        EventsOn(EventOnInternalWarning, (eventData) => this.onEventOnInternalWarning(eventData));
        EventsOn(EventOnInternalError, (eventData) => this.onEventOnInternalError(eventData));

        EventsOn(EventOnNewFileCreated, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileOpened, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileSaved, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileClosed, (eventData) => this.updateEditorState());

        EventsOn(EventOnFileInformationUpdated, (eventData) => this.updateEditorState());

        EventsOn(EventOnFileIsSwitched, (eventData) => this.updateEditorState());
        EventsOn(EventOnFileContentIsUpdated, (eventData) => this.updateEditorState());

        EventsOn(EventOnBlockUiTrue, (eventData) => this.onEventOnBlockUiTrue());
        EventsOn(EventOnBlockUiFalse, (eventData) => this.onEventOnBlockUiFalse());
    }

    onEventOnInternalWarning(event: string) {
        this.onEventOnBlockUiFalse();
        this.setState({
            notificationType: NotificationType.WARNING,
            notificationMessage: event,
        });
    }

    onEventOnInternalError(event: string) {
        this.onEventOnBlockUiFalse();
        this.setState({
            notificationType: NotificationType.ERROR,
            notificationMessage: event,
        });
    }

    componentDidMount() {
        this.updateEditorState();
    }

    onEventOnBlockUiTrue() {
        this.setState({blockUi: true});
    }

    onEventOnBlockUiFalse() {
        this.setState({blockUi: false});
    }

    onNewFile() {
        NewFile().catch((e) => this.onError(e));
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
            .then((files: FrontendFileInfoArrayContainerStruct) => {
                if (files.HasError) {
                    this.onError(new Error(files.Error));
                    return;
                }
                this.setState({openedFiles: files.Files});
            })
            .then(GetOpenedFile)
            .then((currentFile: FrontendFileContainerStruct) => {
                if (currentFile.HasError) {
                    this.onError(new Error(currentFile.Error));
                    return;
                }
                this.setState({currentFile: currentFile.File});
            })
            .catch((e) => this.onError(e));
    }

    render() {
        const showInfoDialog: boolean = this.state.notificationType !== NotificationType.NONE;
        const showEditor: boolean = this.state.openedFiles.length > 0;

        const editorContent = <div>
            <TabBar files={this.state.openedFiles} onError={(error: Error) => this.onError(error)}
                    onNewFile={() => this.onNewFile()}/>
            <CodeEditor file={this.state.currentFile}/>
        </div>;

        const noContent = <Message color="yellow">No Opened Files</Message>;
        return (
            <div>
                <Dimmer.Dimmable blurring dimmed={this.state.blockUi}>
                    <Dimmer active={this.state.blockUi}/>
                    {showEditor ? editorContent : noContent}
                    <AppModalInfoDialog header={this.state.notificationType}
                                        message={this.state.notificationMessage}
                                        show={showInfoDialog}
                                        onClose={() => this.setState({
                                            notificationType: NotificationType.NONE,
                                            notificationMessage: "",
                                        })}/>
                </Dimmer.Dimmable>
            </div>
        );
    }
}

export default App;