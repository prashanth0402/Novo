<template>
    <v-container>
        <v-slide-y-transition mode="out-in" appear>
            <v-breadcrumbs :items="items">
                <template v-slot:divider>
                    <v-icon>mdi-chevron-right</v-icon>
                </template>
            </v-breadcrumbs>
        </v-slide-y-transition>
        <v-slide-x-transition mode="out-in" appear>
            <roleMasterMobile v-if="this.$vuetify.breakpoint.width < 600"></roleMasterMobile>
            <roleMaster v-else />
            <!-- <roleMaster /> -->

        </v-slide-x-transition>
    </v-container>
</template>
<script>
import roleMaster from "../roleMaster.vue"
import roleMasterMobile from "../roleMasterMobile.vue";
export default {

    components: {
        roleMaster,
        roleMasterMobile
    },
    data: () => ({
        items: [
            {
                text: 'RoleTask',
                disabled: true,
            },
            {
                text: 'CreateTask',
                disabled: false,
                href: '/createtask',
            },
        ],
    }),
};

</script>