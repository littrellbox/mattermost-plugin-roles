import React from 'react';

// eslint-disable-next-line no-unused-vars
import PropTypes from 'prop-types';
import {FormattedMessage} from 'react-intl';
import {makeStyleFromTheme} from 'mattermost-redux/utils/theme_utils';

// LeftSidebarHeader is a pure component, later connected to the Redux store so as to
// show the plugin's enabled / disabled status.
export default class LeftSidebarHeader extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    toggleHover() {
        this.setState({hover: !this.state.hover});
    }

    render() {
        const iconStyle = {
            display: 'inline-block',
            margin: '0 7px 0 1px',
        };
        let style = {
            margin: '.5em 0 .5em',
            padding: '0 12px 0 15px',
            color: 'rgba(255,255,255,0.6)',
        };
        if (this.state.hover) {
            style = {
                margin: '.5em 0 .5em',
                padding: '0 12px 0 15px',
                color: this.props.theme.sidebarTextHoverBg,
            };
        }

        return (
            <div style={style}>
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