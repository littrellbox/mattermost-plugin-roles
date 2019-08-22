import React from 'react';

// eslint-disable-next-line no-unused-vars
import PropTypes from 'prop-types';
import {FormattedMessage} from 'react-intl';
import {Modal} from 'react-bootstrap';
import {InputGroup} from 'react-bootstrap'
import {FormControl} from 'react-bootstrap'
import {DropdownButton} from 'react-bootstrap'
import {Panel} from 'react-bootstrap'
import {MenuItem} from 'react-bootstrap'
import {Checkbox} from 'react-bootstrap'
import {Button} from 'react-bootstrap'

var roleDropdownName = "Roles";

const SMALL_SCREEN_WIDTH = 900;

export default class RolesModal extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    changeDropdownName(e) {
        //TODO: Change the dropdown text
    }

    constructor(props) {
        super(props)

        this.changeDropdownName = this.changeDropdownName.bind(this);
    }

    isSmall = () => {
        return window.innerWidth <= SMALL_SCREEN_WIDTH;
    }

    handleChildClick(e) {
        e.stopPropagation();
    }

    render() {
        
        var menuItemStyle = {
            color: this.props.theme.buttonColor
        }
        
        var buttonStyle = {
            color: this.props.theme.buttonColor,
            backgroundColor: this.props.theme.buttonBg
        }
        
        var panelBodyStyle = {
            color: this.props.theme.centerChannelColor,
            backgroundColor: this.props.theme.centerChannelBg
        }

        var panelFooterStyle = {
            color: this.props.theme.centerChannelColor,
            backgroundColor: this.props.theme.sidebarBg
        }

        return (
            <Modal
                show={true}
                onClick={this.handleChildClick}
                onHide={this.props.doClose}
            >
                <style> {`
                    .menuItemStyle {
                        color: ` + this.props.theme.buttonColor + ` !important;
                    }
                `} </style>
                <Modal.Header
                    closeButton={true}
                >
                    <h4 className='modal-title'>
                        {'Nothing'}
                    </h4>
                </Modal.Header>
                <Modal.Body>
                    <InputGroup>
                        <DropdownButton
                            componentClass={InputGroup.Button}
                            title="Roles"
                            id="input-dropdown-addon"
                            style={buttonStyle}
                            onChange={this.changeDropdownName}
                        >
                            <MenuItem eventKey="1">Nothing</MenuItem>
                            <MenuItem eventKey="2">Nothing 2</MenuItem>
                            <MenuItem eventKey="3">Nothing 3</MenuItem>
                            <MenuItem divider />
                            <MenuItem eventKey="4">New Role</MenuItem>
                        </DropdownButton>
                        <FormControl
                            placeholder="Role Name"
                            aria-label="Name of role"
                        />
                        <InputGroup.Button>
                            <Button style={buttonStyle} variant="outline-secondary">Create</Button>
                        </InputGroup.Button>
                    </InputGroup>
                    <br/>
                    <Panel>
                        <Panel.Body style={panelBodyStyle}>
                            <Checkbox>
                                Send Messages
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to send messages.</Panel.Footer>
                    </Panel>            
                    <Panel>
                        <Panel.Body style={panelBodyStyle}>
                            <Checkbox>
                                Upload Files
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to upload files.</Panel.Footer>
                    </Panel>     
                    <Panel>
                        <Panel.Body style={panelBodyStyle}>
                            <Checkbox>
                                Mute Users
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to mute members using the /mute command.</Panel.Footer>
                    </Panel>     
                    <Panel>
                        <Panel.Body style={panelBodyStyle}>
                            <Checkbox>
                                Kick Users
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to kick members using the /kick command.</Panel.Footer>
                    </Panel>         
                    <Panel>
                        <Panel.Body style={panelBodyStyle}>
                            <Checkbox>
                                Ban Users
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to ban members using the /ban command.</Panel.Footer>
                    </Panel>       
                    <Panel>
                        <Panel.Body style={panelBodyStyle}>
                            <Checkbox>
                                Pin Messages
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to pin messages.</Panel.Footer>
                    </Panel>       
                </Modal.Body>
                <Modal.Footer>
                    <button
                        type='button'
                        className='btn btn-cancel'
                        onClick={this.props.doClose}
                        style={buttonStyle}
                    >
                        {'Close'}
                    </button>
                </Modal.Footer>
            </Modal>
        );
    }
}

/*
import {InputGroup} from 'react-bootstrap';
import {FormControl} from 'react-bootstrap';
import {DropdownButton} from 'react-bootstrap';
import {Dropdown} from 'react-bootstrap';
import {Button} from 'react-bootstrap';


{this.props.enabled ?
<span>
    <i
           className='icon fa fa-plug'
           style={iconStyle}
       />
    <FormattedMessage
           id='sidebar.demo'
           defaultMessage='Roles'
       />
   </span> :
<span/>
}*/
