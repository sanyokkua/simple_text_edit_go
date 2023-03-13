import React, {useState} from "react"
import Editor from '@monaco-editor/react';
import {EventsOn} from "../wailsjs/runtime";


function App() {
    // @ts-ignore
    const [text, setText] = useState<string>("");

    EventsOn("FileIsChosen", (filePath)=>{
        setText(filePath)
    })

    return (
        <Editor
            height="100vh"
            defaultLanguage="plaintext"
            defaultValue=""
            value={text}
        />
    )
}

export default App
