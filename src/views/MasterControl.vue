<template>
    <div>
        <MasterControl v-if="allowed" class="mt-10" :IpoMasterData="IpoMasterData" :SgbMasterData="SgbMasterData" :NcbMasterData="NcbMasterData"
        @GetAllMaster="GetAllMasterData"></MasterControl>
    </div>
</template>
<script>
import EventServices from "@/services/EventServices";
import MasterControl from "../components/MasterController/MasterControl.vue";

export default {
    components: {
        MasterControl
    },
    data() {
        return {
            allowed: false,
            loading: true,
            IpoMasterData:[],
            SgbMasterData:[],
            NcbMasterData:[]
        };
    },
    methods: {
        GetAllMasterData() {
            EventServices.GetMasterControl()
                .then((response) => {
                    this.$globalData.overlay = true;
                        this.loading = true;
                    if (response.data.status == "S") {
                        this.$globalData.overlay = false;
                        this.loading = false;
                        this.IpoMasterData = response.data.ipoMasterData != null ? response.data.ipoMasterData  : []
                        this.SgbMasterData = response.data.sgbMasterData != null ? response.data.sgbMasterData  : []
                        this.NcbMasterData = response.data.ncbMasterData != null ? response.data.ncbMasterData  : []                        
                    } else {
                        //allow the access to this page
                        this.$globalData.overlay = false;
                        this.loading = false;
                        this.MessageBar("E", response.data.errMsg);
                    }
                })
                .catch((error) => {
                    this.$globalData.overlay = false;
                    this.loading = false;
                    this.MessageBar("E", error);
                });
        },
        Token() {
            this.$globalData.overlay = true;
            this.loading = true;
            EventServices.tokenValidation()
                .then((response) => {
                    if (response.data.status != "S") {
                        this.$globalData.overlay = false;
                        this.loading = false;
                        window.location = this.LoginUrl;
                    } else {
                        EventServices.RouterValidation(this.$route.path)
                            .then((response) => {
                                if (response.data.status != "S") {
                                    this.$globalData.overlay = false;
                                    this.loading = false;
                                    //   window.location = this.LoginUrl;
                                    this.$router.replace(this.$globalData.links[0].path)
                                } else {
                                    //allow the access to this page
                                    this.$globalData.overlay = false;
                                    this.loading = false;
                                    this.allowed = true;
                                }
                            })
                            .catch((error) => {
                                this.$globalData.overlay = false;
                                this.loading = false;
                                this.MessageBar("E", error);
                            });
                    }
                })
                .catch(() => {
                    this.$globalData.overlay = false;
                    this.loading = false;
                    window.location = this.LoginUrl;
                });
        },
    },
    created() {
        if (this.$globalData.logged == true) {
            if (this.$globalData.logged == true) {
                this.Token();
            } else {
                this.$router.replace("/")
            };
        } else {
            this.$router.replace("/")
        }
    },
    mounted(){
        this.GetAllMasterData()
    }

}
</script>
