import React from "react";
import {createRoot} from "react-dom/client";
import App from "./App";
import "semantic-ui-css/semantic.min.css";

// @ts-ignore
const container = document.getElementById("root");

const root = createRoot(container!);

root.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>,
);
