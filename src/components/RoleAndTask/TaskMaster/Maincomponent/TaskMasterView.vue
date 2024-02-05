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
            <TaskMasterMobile v-if="this.$vuetify.breakpoint.width < 600"></TaskMasterMobile>
            <TaskMaster v-else />
        </v-slide-x-transition>
    </v-container>
</template>


<script>
import TaskMasterMobile from "../TaskMasterMobile.vue";
import TaskMaster from "../TaskMaster.vue"
export default {

    components: {
        TaskMaster,
        TaskMasterMobile
    },
    data: () => ({
        items: [
            {
                text: 'RoleTask',
                disabled: true,
            },
            {
                text: 'AssignTask',
                disabled: false,
                href: '/assigntask',
            },
        ],
    }),
};

</script>