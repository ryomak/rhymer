import Vue from 'vue'
const EventBus = new Vue();

const install = () => {
    if (install.installed) return;
    install.installed = true;
    Object.defineProperties(Vue.prototype, {
        $bus: {
            get() {
                return EventBus;
            }
        },
    });
};
const EventBusPlugin = {
    install,
};
export default EventBus;
export {
    EventBusPlugin,
};
Vue.Use(EventBusPlugin)