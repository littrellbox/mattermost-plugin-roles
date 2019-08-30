import React from 'react';

// eslint-disable-next-line no-unused-vars
import PropTypes from 'prop-types';
import {FormattedMessage} from 'react-intl';
import RolesModal from '../roles_modal'

// LeftSidebarHeader is a pure component, later connected to the Redux store so as to
// show the plugin's enabled / disabled status.
export default class LeftSidebarHeader extends React.PureComponent {
    /*static propTypes = {
        enabled: PropTypes.bool.isRequired,
    }*/

    toggleHover() {
        this.setState({hover: !this.state.hover});
    }

    startModal() {
        var request = {}
        request.teamID = this.props.teamid
        httpRequest = new XMLHttpRequest();
        httpRequest.onreadyload = function(e) {
            if (httpRequest.readyState == 4) {
                if (httpRequest.status == 200) {
                    this.setState({
                        rolelist: JSON.parse(httpRequest.response)
                    })
                    this.setState({showingRoleBox: !this.state.showingRoleBox, });
                    this.setState({hover: false});
                }
            }
        }
        httpRequest.open('POST', "/plugins/me.william341.mattermost-plugin-roles/api/v1/roles/user/getallpermissions", true);
        httpRequest.setRequestHeader('Content-Type', 'application/json; charset=UTF-8')
        httpRequest.send(JSON.stringify(request))
    }

    constructor(props) {
        super(props)
        this.state = {
            hover: false,
            rolelist: [],
            showingRoleBox: false
        }
        this.toggleHover = this.toggleHover.bind(this);
        this.startModal = this.startModal.bind(this);
    }

    handleChildClick(e) {
        e.stopPropagation();
        //e.preventDefault();
        //e.stopImmediatePropagation()
        console.log('handleChildClick');
    }
    
    render() {
        const iconStyle = {
            display: 'inline-block',
            margin: '0 7px 0 1px',
        };
        var style;

        if (this.state.hover) {
            style = {
                padding: '8.5px 12px 8.5px 15px',
                color: 'rgba(255,255,255,0.6)',
                backgroundColor: this.props.theme.sidebarTextHoverBg,   
            }
        } else {
            style = {
                padding: '8.5px 12px 8.5px 15px',
                color: 'rgba(255,255,255,0.6)',
            }
        } 
        
        

        return (
            <div style={style} onMouseEnter={this.toggleHover} onMouseLeave={this.toggleHover} onClick={this.startModal}>
                <i
                    className='icon fa fa-plug'
                    style={iconStyle}
                />
                <FormattedMessage
                    id='sidebar.demo'
                    defaultMessage='Roles'
                />
                <div onClick={this.handleChildClick}>
                    {this.state.showingRoleBox &&
                        <RolesModal theme={this.props.theme} doClose={this.toggleModal} roleList={this.state.rolelist} teamid={this.props.teamid}/>
                    }
                </div>
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