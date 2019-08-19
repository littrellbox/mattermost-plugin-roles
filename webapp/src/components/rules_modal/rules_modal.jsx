import React from 'react';

// eslint-disable-next-line no-unused-vars
import PropTypes from 'prop-types';
import {FormattedMessage} from 'react-intl';
import {Modal} from 'react-bootstrap';

const SMALL_SCREEN_WIDTH = 900;

export default class RulesModal extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    constructor(props) {
        super(props)
    }

    isSmall = () => {
        return window.innerWidth <= SMALL_SCREEN_WIDTH;
    }

    render() {
        return (
            <Modal
                show={true}
                onHide={this.props.doClose}
            >
                <Modal.Header
                    closeButton={true}
                >
                    <h4 className='modal-title'>
                        {'Nothing'}
                    </h4>
                </Modal.Header>
                <Modal.Body>
                    <h1>Nothing</h1>
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

/*{this.props.enabled ?
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