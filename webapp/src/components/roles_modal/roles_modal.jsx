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
import {ListGroup} from 'react-bootstrap'
import {ListGroupItem} from 'react-bootstrap'
var roleDropdownName = "Roles";

const SMALL_SCREEN_WIDTH = 900;

export default class RolesModal extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    setPermission(value, permission) {
        if (this.state.currentRole == ":new") {
            return
        }
        var request = {}
        request.role_name = this.state.currentRole
        request.team_id = this.props.teamid
        request.permission = permission
        request.value = value
        httpRequest = new XMLHttpRequest();
        httpRequest.onreadyload = function(e) {

        }
        httpRequest.open('POST', "/plugins/me.william341.mattermost-plugin-roles/api/v1/roles/role/setpermission", true);
        httpRequest.setRequestHeader('Content-Type', 'application/json; charset=UTF-8')
        httpRequest.send(JSON.stringify(request))
    }

    changeSendMessages(e) {
        this.setPermission(e.target.checked, "send")
    }

    changeUpload(e) {
        this.setPermission(e.target.checked, "files")
    }
    
    changeMute(e) {
        this.setPermission(e.target.checked, "mute")
    }

    changeKick(e) {
        this.setPermission(e.target.checked, "kick")
    }

    changeBan(e) {
        this.setPermission(e.target.checked, "ban")
    }

    changePin(e) {
        this.setPermission(e.target.checked, "pin")
    }

    addUser(e) {
        if (this.state.currentRole == ":new") {
            return
        }
        var request = {}
        request.role_name = this.state.currentRole
        request.team_id = this.props.teamid
        request.target_id =
        httpRequest = new XMLHttpRequest();
        httpRequest.onreadyload = function(e) {

        }
        httpRequest.open('POST', "/plugins/me.william341.mattermost-plugin-roles/api/v1/roles/user/getallpermissions", true);
        httpRequest.setRequestHeader('Content-Type', 'application/json; charset=UTF-8')
        httpRequest.send(JSON.stringify(request))
    }

    removeUser(e) {
        if (this.state.currentRole == ":new") {
            return
        }
        var request = {}
        request.role_name = this.state.currentRole
        request.team_id = this.props.teamid
        request.permission = permission
        request.value = value
        httpRequest = new XMLHttpRequest();
        httpRequest.onreadyload = function(e) {

        }
        httpRequest.open('POST', "/plugins/me.william341.mattermost-plugin-roles/api/v1/roles/user/getallpermissions", true);
        httpRequest.setRequestHeader('Content-Type', 'application/json; charset=UTF-8')
        httpRequest.send(JSON.stringify(request))
    }

    changeDropdownName(e) {
        //change the dropdown name and fetch the permissions
        this.setState({currentRole: e})
    }

    createList() {
        var list = [];
        for (i=0; i<this.props.roleList.length; i++) {
            list.push(<MenuItem eventKey={this.props.roleList[i]}>{this.props.roleList[i]}</MenuItem>)
        }
    }

    constructor(props) {
        super(props)
        this.state = {
            currentRole: ":new"
        }

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
                            {this.createList()}
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
                        <Panel.Body style={panelBodyStyle} onChange={this.changeSendMessages}>
                            <Checkbox>
                                Send Messages
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to send messages.</Panel.Footer>
                    </Panel>            
                    <Panel>
                        <Panel.Body style={panelBodyStyle} onChange={this.changeUpload}>
                            <Checkbox>
                                Upload Files
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to upload files.</Panel.Footer>
                    </Panel>     
                    <Panel>
                        <Panel.Body style={panelBodyStyle} onChange={this.changeMute}>
                            <Checkbox>
                                Mute Users
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to mute members using the /mute command.</Panel.Footer>
                    </Panel>     
                    <Panel>
                        <Panel.Body style={panelBodyStyle} onChange={this.changeKick}>
                            <Checkbox>
                                Kick Users
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to kick members using the /kick command.</Panel.Footer>
                    </Panel>         
                    <Panel>
                        <Panel.Body style={panelBodyStyle} onChange={this.changeBan}>
                            <Checkbox>
                                Ban Users
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to ban members using the /ban command.</Panel.Footer>
                    </Panel>       
                    <Panel>
                        <Panel.Body style={panelBodyStyle} onChange={this.changePin}>
                            <Checkbox>
                                Pin Messages
                            </Checkbox>
                        </Panel.Body>
                        <Panel.Footer style={panelFooterStyle}>Allows a user to pin messages.</Panel.Footer>
                    </Panel>
                    <ListGroup>
                        <ListGroupItem>
                            <FormControl
                                type="text"
                                placeholder="Username"
                            />
                            <Button 
                                bsStyle="success"
                                onClick={this.addUser}
                            >
                                Add User
                            </Button>
                        </ListGroupItem>
                        <ListGroupItem>Member 1 
                            <Button 
                                bsStyle="danger" 
                                bsSize="xsmall"
                                onClick={this.removeUser}
                            >
                                Remove
                            </Button>
                        </ListGroupItem>
                    </ListGroup>       
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
