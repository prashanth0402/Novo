<template>
  <div>
    <v-list v-if="!showSubMenu" :class="$vuetify.breakpoint.name != 'xs' ? 'mt-3 menu' : undefined"
      :outlined="$vuetify.breakpoint.name != 'sm' && $vuetify.breakpoint.name != 'xs'">
      <v-list-item v-for="userSubMenu in userSubMenus" :key="userSubMenu.routerId" link :to="userSubMenu.path">
        <v-list-item-title v-text="userSubMenu.router" @click="closeSubmenu"></v-list-item-title>
      </v-list-item>
    </v-list>
  </div>
</template>

<script>
import EventService from "@/services/EventServices.js";

export default {
  name: "SubMain",
  components: {},
  props: {
    parentMenuId: {
      type: Number,
    },

  },
  methods: {
    closeSubmenu() {
      this.$emit('CloseMenu')
    },
  },

  data() {
    return {
      showSubMenu: false,
      userSubMenus: [],
    };
  },
  mounted() {
    EventService.GetSubMenu(this.parentMenuId)
      .then((response) => {
        //console.log(response.data);
        if (response.data.subMenuArr == null) {
          this.userSubMenus = []
        } else {
          this.userSubMenus = response.data.subMenuArr;
          this.$globalData.subMenu = response.data.subMenuArr;
        }
        if (this.userSubMenus.length > 0) {
          this.showSubMenu = false;
        } else {
          this.showSubMenu = true;
        }
      })
      .catch((error) => {
        this.MessageBar('E', error)
      });
  },
};
</script>
<style scoped>
.menu {
  position: relative;
  border-radius: 5px;
}

.v-list-item-title:hover {
  color: #1976D2;
}

/* .menu::before {
  position: absolute;
  content: '';
  height: 20px;
  width: 20px;
  transform: rotate(45deg);
  background: #fff;
  border-top: 1px solid #dddadabd;
  border-left: 1px solid #dddadabd;
  top: -10px;
  right: 15px;
} */
</style>