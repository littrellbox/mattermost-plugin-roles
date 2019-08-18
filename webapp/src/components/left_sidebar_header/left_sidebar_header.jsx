import React from 'react';

// eslint-disable-next-line no-unused-vars
import PropTypes from 'prop-types';
import {FormattedMessage} from 'react-intl';

// LeftSidebarHeader is a pure component, later connected to the Redux store so as to
// show the plugin's enabled / disabled status.
export default class LeftSidebarHeader extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    toggleHover() {
        this.setState({hover: !this.state.hover});
    }

    constructor(props) {
        super(props)
        this.state = {
            hover: false
        }
        this.toggleHover = this.toggleHover.bind(this);
    }

    render() {
        const iconStyle = {
            display: 'inline-block',
            margin: '0 7px 0 1px',
        };
        var style;

        if (this.state.hover) {
            style = {
                padding: '3.5px 12px 3.5px 15px',
                color: 'rgba(255,255,255,0.6)',
                backgroundColor: this.props.theme.sidebarTextHoverBg,   
            }
        } else {
            style = {
                padding: '3.5px 12px 3.5px 15px',
                color: 'rgba(255,255,255,0.6)',
            }
        } 

        return (
            <div style={style} onMouseEnter={this.toggleHover} onMouseLeave={this.toggleHover}>
                <i
                    className='icon fa fa-plug'
                    style={iconStyle}
                />
                <FormattedMessage
                    id='sidebar.demo'
                    defaultMessage='Nothing'
                />
            </div>
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