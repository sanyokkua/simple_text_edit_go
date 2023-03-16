import React from "react"

import {EventsOn} from "../wailsjs/runtime";
import CodeMirror from '@uiw/react-codemirror';


import {GetActiveFile, GetAllOpenedFileDescriptors, MakeFileActive} from "../wailsjs/go/application/EditorApplication"
import {AppFile, FileDescriptor} from "./types/types";
import {Menu} from "semantic-ui-react";
import {MenuItemProps} from "semantic-ui-react/dist/commonjs/collections/Menu/MenuItem";

type AppState = {
    activeFile: AppFile | null,
    openedFiles: FileDescriptor[]
};

class App extends React.Component<any, AppState> {
    constructor(props: any) {
        super(props);
        this.state = {
            activeFile: null,
            openedFiles: []
        };
        EventsOn("EVENT_IS_FILE_OPENED", () => {
            this.updateState().then()
        });
    }

    async updateState() {
        const activeFile: AppFile = await GetActiveFile()
        const files: FileDescriptor[] = await GetAllOpenedFileDescriptors()
        files.sort((a, b) => a.FileId - b.FileId)
        this.setState({
            activeFile: activeFile,
            openedFiles: files
        });
    }

    async componentDidMount() {
        await this.updateState()
    }

    async tabIsClicked(event: React.MouseEvent<HTMLAnchorElement>, data: MenuItemProps, fileId: number) {
        await MakeFileActive(fileId)
        await this.updateState()
    }

    render() {
        const menuItems = this.state.openedFiles.map(openedFile => {
            return <Menu.Item
                name={openedFile.FileName}
                active={openedFile.IsActive}
                onClick={(event: React.MouseEvent<HTMLAnchorElement>, data: MenuItemProps) => this.tabIsClicked(event, data, openedFile.FileId)}
            />
        })
        console.log(menuItems);

        const activeFile = this.state.activeFile;
        console.log(activeFile);

        const fileContent = activeFile?.FileContent;
        console.log(fileContent);

        return (
            <div>
                <Menu tabular>
                    {menuItems}
                </Menu>
                <CodeMirror
                    value={fileContent}
                    height="100vh"
                />
            </div>
        )
    }
}

export default App
