import React from 'react';

// eslint-disable-next-line no-unused-vars
import PropTypes from 'prop-types';
import {FormattedMessage} from 'react-intl';
import {Modal} from 'react-bootstrap';
import {InputGroup} from 'react-bootstrap'
import {FormControl} from 'react-bootstrap'
import {DropdownButton} from 'react-bootstrap'
import {Dropdown} from 'react-bootstrap'
import {MenuItem} from 'react-bootstrap'
import {Button} from 'react-bootstrap'


const SMALL_SCREEN_WIDTH = 900;

export default class RolesModal extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    constructor(props) {
        super(props)
    }

    isSmall = () => {
        return window.innerWidth <= SMALL_SCREEN_WIDTH;
    }

    handleChildClick(e) {
        e.stopPropagation();
        //e.preventDefault();
        //e.stopImmediatePropagation()
        console.log('handleChildClick');
    }

    render() {
        
        var menuItemStyle = {
            color: this.props.theme.buttonColor
        }
        
        var modalStyle = {
            overflowY: 'visible'
        }
        
        var buttonStyle = {
            color: this.props.theme.buttonColor,
            backgroundColor: this.props.theme.buttonBg
        }
        
        return (
            <Modal
                show={true}
                onClick={this.handleChildClick}
            >
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
                            title="Nothings"
                            id="input-dropdown-addon"
                            style={buttonStyle}
                        >
                            <MenuItem style={menuItemStyle} eventKey="1">Nothing</MenuItem>
                            <MenuItem style={menuItemStyle} eventKey="2">Nothing 2</MenuItem>
                            <MenuItem style={menuItemStyle} eventKey="3">Nothing 3</MenuItem>
                            <MenuItem divider />
                            <MenuItem style={menuItemStyle} eventKey="4">New Nothing</MenuItem>
                        </DropdownButton>
                        <FormControl
                            placeholder="Nothing name"
                            aria-label="Name of role"
                        />
                        <InputGroup.Button>
                            <Button style={buttonStyle} variant="outline-secondary">Do Nothing</Button>
                        </InputGroup.Button>
                    </InputGroup>
                </Modal.Body>
                <Modal.Footer>
                    <button
                        type='button'
                        className='btn btn-cancel'
                        onClick={this.props.doClose}
                    >
                        {'k'}
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