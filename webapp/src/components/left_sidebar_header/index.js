import {connect} from 'react-redux';

// eslint-disable-next-line no-unused-vars
import {isEnabled} from 'selectors';

import LeftSidebarHeader from './left_sidebar_header';

// eslint-disable-next-line no-unused-vars
const mapStateToProps = (state) => ({
    teamid: state.entities.teams.currentTeamId
});

export default connect(mapStateToProps)(LeftSidebarHeader);

//enabled: isEnabled(state),