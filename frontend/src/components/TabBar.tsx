import React from "react";
import {FileInfoStruct} from "../types/backend";
import {SemanticCOLORS} from "semantic-ui-react/dist/commonjs/generic";
import {Button, Icon, Menu} from "semantic-ui-react";
import {CloseFile, SwitchOpenedFileTo} from "../../wailsjs/go/frontendapi/FrontendApiStruct";

type FilesTabBarProps = {
    files: FileInfoStruct[];
    onError: (error: Error) => void;
    onNewFile: () => void;
};
type MenuItemInfo = {
    fileId: number;
    itemKey: string
    isActive: boolean
    tabColor: SemanticCOLORS
    fileName: string

}
const COLOR_NEW: SemanticCOLORS = "red";
const COLOR_NO_CHANGES: SemanticCOLORS = "black";
const COLOR_HAS_CHANGES: SemanticCOLORS = "blue";

function fileShortInfoToMenuItem(shortInfo: FileInfoStruct): MenuItemInfo {
    const id: number = shortInfo.Id;
    const key: string = id.toString();
    const isNewFile: boolean = shortInfo.New;
    const isChanged: boolean = shortInfo.Changed;
    const isActive: boolean = shortInfo.Opened;
    const fileName: string = shortInfo.Name;
    const color: SemanticCOLORS = isNewFile ? COLOR_NEW : (isChanged ? COLOR_HAS_CHANGES : COLOR_NO_CHANGES);

    return {
        fileId: id,
        itemKey: key,
        isActive: isActive,
        tabColor: color,
        fileName: fileName,
    };
}


class TabBar extends React.Component<FilesTabBarProps, any> {
    onTabClicked(fileId: number) {
        SwitchOpenedFileTo(fileId).catch((e) => this.props.onError(e));
    }

    onTabCloseClicked(fileId: number) {
        CloseFile(fileId).catch((e) => this.props.onError(e));
    }

    render() {
        const menuItemInfos = this.props.files.map(fileShortInfoToMenuItem);
        const menuItems = menuItemInfos.map(menuItemInfo => {
            return <Menu.Item key={menuItemInfo.itemKey}
                              active={menuItemInfo.isActive}
                              color={menuItemInfo.tabColor}
                              onClick={() => this.onTabClicked(menuItemInfo.fileId)}
            >
                {menuItemInfo.fileName}
                {menuItemInfo.isActive && <Icon color={"red"} name="close" style={{float: "right"}}
                                                onClick={() => this.onTabCloseClicked(menuItemInfo.fileId)}/>}


            </Menu.Item>;
        });
        return <Menu tabular size="mini" stackable>
            {menuItems}
            <Menu.Item position="left">
                <Button icon primary onClick={() => this.props.onNewFile()}>
                    <Icon name="plus circle"/>
                </Button>
            </Menu.Item>
        </Menu>;
    }
}

export default TabBar;