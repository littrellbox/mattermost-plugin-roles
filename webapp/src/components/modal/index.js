import {connect} from 'react-redux';

// eslint-disable-next-line no-unused-vars
import {isEnabled} from 'selectors';

import RulesModal from './modal';

// eslint-disable-next-line no-unused-vars
const mapStateToProps = (state) => ({
});

export default connect(mapStateToProps)(RulesModal);

//enabled: isEnabled(state),