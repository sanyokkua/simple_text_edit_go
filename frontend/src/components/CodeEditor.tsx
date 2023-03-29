import React from "react";
import {FileStruct} from "../types/backend";
import CodeMirror from "@uiw/react-codemirror";
import {UpdateFileContent} from "../../wailsjs/go/frontend/frontStruct";
import {LanguageName, loadLanguage} from "@uiw/codemirror-extensions-langs";
import AppModalInfoDialog from "./AppModalInfoDialog";

type CodeEditorProps = {
    file: FileStruct | null | undefined;
};

type CodeEditorState = {
    showError: boolean;
    errorMessage: string;
};


class CodeEditor extends React.Component<CodeEditorProps, CodeEditorState> {
    constructor(props: CodeEditorProps) {
        super(props);
        this.state = {
            showError: false,
            errorMessage: "",
        };
    }

    onError(error: Error) {
        let msg = error?.message ? error?.message : JSON.stringify(error);
        this.setState({
            showError: true,
            errorMessage: msg,
        });
    }

    contentIsChanged(text: string) {
        if (!this.props.file) {
            this.onError({
                name: "Update content error",
                message: "File passed in props is null|undefined",
            });
            return;
        }
        const id: number = this.props.file.Id;
        UpdateFileContent(id, text).catch((e) => this.onError(e));
    }

    getCurrentSyntaxLang(): any | null {
        if (!this.props.file || this.props.file.Type === undefined || this.props.file.Type === null
            || this.props.file.Type === "" || this.props.file.Type === "txt") {
            return null;
        }
        return loadLanguage(this.props.file.Type as LanguageName);
    }

    getEditorExtensions(): any[] {
        const extensions: any[] = [];
        let currentSyntaxLang = this.getCurrentSyntaxLang();
        if (currentSyntaxLang) {
            extensions.push(currentSyntaxLang);
        }
        return extensions;
    }

    getEditorContent(): string {
        if (!this.props.file) {
            return "";
        }
        return this.props.file.ActualContent;
    }

    render() {
        const fileContent: string = this.getEditorContent();
        const editorExtensions: any[] = this.getEditorExtensions();

        return (
            <div>
                <CodeMirror height="100vh"
                            basicSetup={{
                                foldGutter: true,
                                allowMultipleSelections: true,
                                indentOnInput: true,
                                tabSize: 4,
                                highlightActiveLine: true,
                                highlightActiveLineGutter: true,
                                highlightSelectionMatches: true,
                                syntaxHighlighting: true,
                                bracketMatching: true,
                            }}
                            extensions={editorExtensions}
                            value={fileContent}
                            onChange={(text) => this.contentIsChanged(text)}
                />
                <AppModalInfoDialog header={"Error"} message={this.state.errorMessage} show={this.state.showError}
                                    onClose={() => this.setState({
                                        errorMessage: "",
                                        showError: false,
                                    })}
                />
            </div>
        );
    }
}

export default CodeEditor;