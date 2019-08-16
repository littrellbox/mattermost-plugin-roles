import {id as pluginId} from './manifest';

import LeftSidebarHeader from './components/left_sidebar_header';

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
        registry.registerLeftSidebarHeaderComponent(LeftSidebarHeader);
    }
}

window.registerPlugin(pluginId, new Plugin());
