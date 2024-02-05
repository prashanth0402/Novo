<template>
    <div>

    </div>
</template>

<script>
export default {
    data: () => ({
        Mobileheader: [
            { text: "", value: "domainName", editable: false, sortable: false },
        ],
        AddDomain: [],
        loading: false,
        search: '',
    }),
    methods: {
        close() {
            this.$nextTick(() => {
                this.Domaindata = Object.assign({}, this.DefaultDomain);
                this.editedIndex = -1;
            });
        },
    }
}
</script>

<style  scoped></style>