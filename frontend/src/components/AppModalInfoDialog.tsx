import React from "react";
import {Button, Modal} from "semantic-ui-react";

type AppModalInfoDialogProps = {
    header: string;
    message: string;
    show: boolean;
    onClose: () => void;
};

class AppModalInfoDialog extends React.Component<AppModalInfoDialogProps, any> {
    render() {
        return (
            <Modal dimmer={"blurring"} open={this.props.show} size={"mini"}>
                <Modal.Header>{this.props.header}</Modal.Header>
                <Modal.Content>
                    {this.props.message}
                </Modal.Content>
                <Modal.Actions>
                    <Button positive onClick={() => this.props.onClose()}>Ok</Button>
                </Modal.Actions>
            </Modal>
        );
    }
}

export default AppModalInfoDialog;