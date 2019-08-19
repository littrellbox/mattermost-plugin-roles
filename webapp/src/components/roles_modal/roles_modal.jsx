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
                    <InputGroup className="mb-3">
                        <DropdownButton
                            as={InputGroup.Prepend}
                            variant="outline-secondary"
                            title="Nothings"
                            id="input-group-dropdown-1"
                        >
                            <MenuItem eventKey="1">Nothing</MenuItem>
                            <MenuItem eventKey="2">Nothing 2</MenuItem>
                            <MenuItem eventKey="3">Nothing 3</MenuItem>
                            <MenuItem divider />
                            <MenuItem eventKey="4">New Nothing</MenuItem>
                        </DropdownButton>
                        <FormControl
                            placeholder="Nothing name"
                            aria-label="Name of role"
                        />
                        <InputGroup.Append>
                            <Button variant="outline-secondary">Do Nothing</Button>
                        </InputGroup.Append>
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