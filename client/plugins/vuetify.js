// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Vuetify
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { VVirtualScroll } from 'vuetify/labs/VVirtualScroll'


export default createVuetify({
    components: {
        ...components,
        VVirtualScroll
    },
    directives,
    icons: {
        defaultSet: 'mdi',
    }
});
