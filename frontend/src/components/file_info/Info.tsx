import React from "react";
import {FileStruct} from "../../types/backend";

type InfoProps = {
    currentFile: FileStruct | null;
};

function Info(props: InfoProps) {
    const notset = "NOT SET";
    return (
        <div>
            <p>File Path: <b><i>{props.currentFile?.Path ? props.currentFile?.Path : notset}</i></b></p>
            <p>File Extension: <b><i>{props.currentFile?.Extension ? props.currentFile?.Extension : notset}</i></b></p>
            <p>File Type: <b><i>{props.currentFile?.Type ? props.currentFile?.Type : notset}</i></b></p>
            <p>File is New: <b><i>{props.currentFile?.New ? "Yes" : "No"}</i></b></p>
            <p>File is changed: <b><i>{props.currentFile?.Changed ? "Yes" : "No"}</i></b></p>
        </div>
    );
}

export default Info;