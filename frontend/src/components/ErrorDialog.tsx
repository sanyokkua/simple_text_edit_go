import React from "react"
import {Button, Modal} from "semantic-ui-react";

type HandlerFunc = () => void;

type DialogProps = {
    showDialog: boolean;
    errorText: string;
    onButtonClicked: HandlerFunc;
};

class ErrorDialog extends React.Component<DialogProps, any> {
    render() {
        return (
            <Modal dimmer={"blurring"} open={this.props.showDialog} size={'mini'}>
                <Modal.Header>Error happened</Modal.Header>
                <Modal.Content>
                    {this.props.errorText}
                </Modal.Content>
                <Modal.Actions>
                    <Button negative onClick={() => this.props.onButtonClicked()}>Ok</Button>
                </Modal.Actions>
            </Modal>
        )
    }
}

export default ErrorDialog