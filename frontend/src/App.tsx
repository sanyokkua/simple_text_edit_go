import React from "react"
import Editor from '@monaco-editor/react';

function App() {

    return (
        <Editor
            height="100vh"
            defaultLanguage="javascript"
            defaultValue="// some comment"
            language="text"
        />
    )
}

export default App
