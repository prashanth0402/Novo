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
<versionMobileTable v-if="this.$vuetify.breakpoint.width < 600"></versionMobileTable>
<versionDeskTable v-else></versionDeskTable>
        </v-slide-x-transition>
    </v-container>
</template>
<script>
import versionDeskTable from './VersionComponents/versionDeskTable.vue';
import versionMobileTable from './VersionComponents/versionMobileTable.vue'
export default {

    components: {
        versionDeskTable,versionMobileTable
    },
    data: () => ({
        items: [
            {
                text: 'Meta Data Setup',
                disabled: true,
            },
            {
                text: 'VersionControl',
                disabled: false,
                href: '/versioncontrol',
            },
        ],
    }),
};

</script>