import Vue from "vue";
import Vuetify from "vuetify/lib/framework";
// import colors from "vuetify/lib/util/colors";
Vue.use(Vuetify);

export default new Vuetify({
  // theme: {
  // themes: {
  //   light: {
  //     accent: colors.orange.darken4, // #3F51B5
  //   },
  // },

  theme: {
    dark: false, // Set this to true for dark theme, and false for light theme
    themes: {
      light: {
        primary: '#1976D2',
        secondary: '#424242',
        accent: '#82B1FF',
        footer: '#f8f9fc',
        header: '#F8F9FC',
        btnColor: '#0965da',
        content: '#365048',
        contentHead: '#2A394E',
        smechip: '#1565C0'
      },
      dark: {
        primary: '#1976D2',
        secondary: '#757575',
        accent: '#64B5F6',
        footer: '#272727',
        btnColor: '#0965da',

      },
    },
  },
});
