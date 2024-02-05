<template>
    <div class="fill_height">
        <v-container>
            <DeskDash :segmentDetail="SegmenDetail" v-if="allowed" />
        </v-container>
    </div>
</template>

<script>
import DeskDash from '../components/Dash/DeskDash.vue';
import EventServices from '../services/EventServices';
export default {
    data: () => ({
        SegmenDetail: [],
        allowed: false,
        currentTime: ""
    }),
    components: {
        DeskDash
    },
    mounted() {
        EventServices.GetDashboardDetail(this.$route.path)
            .then((response) => {
                if (response.data.status == "S") {
                    if (response.data.segmentArr != null) {
                        this.SegmenDetail = response.data.segmentArr;
                    }
                }
            })
            .catch((error) => {
                this.MessageBar("E", error)
            });
        if (this.$globalData.currentTime == "") {
            setInterval(this.GetCurrentTime, 1000);
        }
    },
    methods: {
        Token() {
            this.$globalData.overlay = true;
            this.loading = true;
            EventServices.tokenValidation()
                .then((response) => {
                    if (response.data.status != "S") {
                        this.$globalData.overlay = false;
                        window.location = this.LoginUrl;
                    } else {
                        this.allowed = true;
                        this.$globalData.overlay = false;
                    }
                })
                .catch(() => {
                    this.$globalData.overlay = false;
                    window.location = this.LoginUrl;
                });
        },
        // to find current time
        GetCurrentTime() {
            const currentTime = new Date();
            let hours = currentTime.getHours();
            let minutes = currentTime.getMinutes();
            let seconds = currentTime.getSeconds();
            hours = (hours < 10 ? "0" : "") + hours;
            minutes = (minutes < 10 ? "0" : "") + minutes;
            seconds = (seconds < 10 ? "0" : "") + seconds;
            this.currentTime = `${hours}:${minutes}:${seconds}`;

            this.$globalData.currentTime = this.currentTime;
        },
    },
    created() {
        if (this.$globalData.logged == true) {
            this.Token();
        } else {
            this.$router.replace("/")
        };
    }
}
</script>

<style  scoped>
.dashbord_bg {
    background-color: #f5f1fc;
    height: 100vh;
}

/* .fill_height {
    height: 100vh;
} */
</style>