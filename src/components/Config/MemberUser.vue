<template>
    <v-container>
        <v-breadcrumbs :items="items">
            <template v-slot:divider>
                <v-icon>mdi-chevron-right</v-icon>
            </template>
        </v-breadcrumbs>
        <MemberTable class="mt-10" />
        <MemberDialog />
    </v-container>
</template>

<script>
import MemberDialog from './AddUser/MemberDialog.vue';
import MemberTable from './AddUser/MemberTable.vue';

export default {
    components: {
        MemberDialog,
        MemberTable
    },
    data() {
        return {
            items: [
                {
                    text: 'Config',
                    disabled: false,
                },
                {
                    text: 'ManageRole',
                    disabled: true,
                    href: '/managerole',
                },
            ],

        };
    },
}
</script>

<style scoped></style>